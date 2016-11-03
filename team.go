package singularity

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

//SlackInstance ...
type SlackInstance struct {
	rtmResp RTMResp

	//Important Data Stuff
	singularity *Singularity
	connection  *websocket.Conn

	//Handlers
	Commands               *CommandHandler //TODO unexport
	handlers               *EventAPIHandler
	customEventHandler     func(input chan ChanMessage, output chan ChanMessage)
	customCommandHandler   func()
	customWebsocketHandler func()
	log                    func(level int, message string, i ...interface{})

	//Friendly name
	Name string

	//Token
	token string

	input  chan ChanMessage //input represents what comes from slack
	output chan ChanMessage //output goes to slack
	quit   chan int

	//Configuration
	Configuration Configuration //TODO unexport, expose funcs.
	transport     *http.Transport
}

//Quit ...
func (instance *SlackInstance) Quit() {
	instance.singularity.RemoveTeam(instance.Name) //Remove Yo Self.
	instance.quit <- 0
	//TODO Add something to wait for the go rutines to end?
}

//Please dont call this outside of Singularity::Shutdown
func (instance *SlackInstance) quitShutdown() {
	instance.quit <- 0
}

//NewTeam creates a new team, that isn't started. To start a team, you'll need to call the Start function on it.
func (singularity *Singularity) NewTeam(Token string) *SlackInstance {
	singularity.Lock()
	defer singularity.Unlock()
	instance := &SlackInstance{token: Token, singularity: singularity}
	//Configuration
	instance.Configuration = &defaultConfig{config: make(map[string]interface{})} //TODO move outside. configs should be configured before a team is started.
	//defaulthandlers //TODO Move so that
	instance.Commands = NewCommandHandler()
	instance.Commands.setPrefix(".")
	addDefaultCommands(instance)
	instance.handlers = NewHandler1()
	addDefaultHandlers(instance)

	singularity.Teams = append(singularity.Teams, *instance)
	return instance
}

//Start starts the slack instance.
func (instance *SlackInstance) Start() error {
	helper := HTTPHelper{Client: &http.Client{}, Transport: "https://"} //TODO Don't hard code this.
	_, err := helper.post("rtm.start", &instance.rtmResp, "token", instance.token)
	if err != nil {
		return err
	}

	//Connect to websocket.
	//TODO Support custom proxies?
	instance.connection, err = websocket.Dial(instance.rtmResp.URL, "", "http://api.slack.com") //TODO Don't hard code here.
	if err != nil {
		return err
	}
	var hello struct {
		Hello string `json:"type"`
	}
	err = websocket.JSON.Receive(instance.connection, &hello)
	if err != nil {
		return err
	}
	if hello.Hello != "hello" {
		return errors.New("Slack did not respond with the correct message")
	}
	instance.Name = instance.rtmResp.Team.Name
	//Set the logger to the default one if it isn't already set.
	if instance.log == nil {
		instance.log = instance.singularity.log
	}

	//Channels for this Instance.
	instance.input = make(chan ChanMessage, 5)  //TODO Configure Amount
	instance.output = make(chan ChanMessage, 5) //TODO Configure Amount
	instance.quit = make(chan int)

	//Start Go-Routines for handling the things.
	if instance.customEventHandler != nil {
		go instance.customEventHandler(instance.input, instance.output)
	} else {
		go instance.handleChans()
	}

	if instance.customWebsocketHandler != nil {
		go instance.customWebsocketHandler()
	} else {
		go instance.handleWebsocket()
	}

	return nil
}

//TODO Panic Recovery
func (instance *SlackInstance) handleChans() {
	for {
	outOfSelect:
		select {

		case val := <-instance.input: //<-instance.input reads from (what is assumed to be) slack.
			event, err := val.GetBytes()
			if err != nil {
				//TODO Shit.
				break outOfSelect
			}

			slackType := preParseString("type", event)
			fmt.Printf("Type: %v\n", slackType)
			if slackType != "" {
				instance.handlers.execute(slackType, event, instance)
			} // TODO handle empty types.

		case val := <-instance.output: //<-instance.output sends the thing to slack.
			thingToSend, err := val.GetInterface()
			if err != nil {
				//TODO Fuck.
				break outOfSelect
			}
			err = websocket.JSON.Send(instance.connection, thingToSend) //SendBodies  :D
			if err != nil {
				//TODO fuck
			}
		case <-instance.quit:
			return
		}
	}
}

// handleWebsocket()
// This is used for grabbing incomming messages and putting them on the queue
// without having to wait for whatever to finish.
func (instance *SlackInstance) handleWebsocket() {
	var i interface{}
	for {
		func() {
			//defer Recover() //TODO implement panic recovery.
			err := websocket.JSON.Receive(instance.connection, &i)
			if err != nil {
				//TODO ERROR HANDLING
				return
			}
			instance.input <- ChanMessage{i} //Buffered.
		}()
	}
}

//GetChannelByID returns a channel matching the supplied id, or nil.
func (instance *SlackInstance) GetChannelByID(id string) *Channel {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()

	for _, channel := range instance.rtmResp.Channels {
		if channel.ID == id {
			return &channel
		}
	}
	return nil
}

//GetChannelByName returns a channel matching the supplied name, or nil.
func (instance *SlackInstance) GetChannelByName(name string) *Channel {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()

	for _, channel := range instance.rtmResp.Channels {
		if channel.Name == name {
			return &channel
		}
	}
	return nil
}

//GetSelf returns self
func (instance *SlackInstance) GetSelf() *Self {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()

	return &instance.rtmResp.Self
}

//GetTeam returns the team object stored in the RTMResp.
func (instance *SlackInstance) GetTeam() *Team {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()

	return &instance.rtmResp.Team
}

//GetUserByID returns a user matching the supplied id, or nil.
func (instance *SlackInstance) GetUserByID(id string) *User {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()

	for _, user := range instance.rtmResp.Users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

//GetUserByName returns a user matching the supplied name, or nil.
func (instance *SlackInstance) GetUserByName(name string) *User {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()

	for _, user := range instance.rtmResp.Users {
		if user.Name == name {
			return &user
		}
	}
	return nil
}

//RegisterCommand ...
func (instance *SlackInstance) RegisterCommand(command string, commandHandler func(Command)) error {
	return instance.Commands.registerCommand(command, commandHandler)
}

//RegisterHandler ...
func (instance *SlackInstance) RegisterHandler(key string, handler interface{}) error {
	instance.handlers.Lock()
	defer instance.handlers.Unlock()
	return instance.handlers.registerHandler(key, handler)
}

//RegisterHandlers will allow you to handle the events.
func (instance *SlackInstance) RegisterHandlers(handlers map[string][]interface{}) error {
	instance.handlers.Lock()
	defer instance.handlers.Unlock()
	for key, handler := range handlers {
		for _, function := range handler {
			if err := instance.handlers.registerHandler(key, function); err != nil {
				return err //No errors allowed.
			}
		}
	}
	return nil
}

//SendMessage sends a slack message
func (instance *SlackInstance) SendMessage(m Message) {
	m.Type = "message"
	m.User = instance.GetSelf().ID
	instance.output <- ChanMessage{Body: m}
}

//SetLogger sets the function that the instance should use for logging.
func (instance *SlackInstance) SetLogger(logger func(level int, message string, i ...interface{})) {
	instance.log = logger
}

//SetHTTPTransport allows you to set a transport to use when communicate with Slack.
func (instance *SlackInstance) SetHTTPTransport(transport *http.Transport) {
	instance.transport = transport
}

//WaitForEnd will wait until the team is killed or stopped.
func (instance *SlackInstance) WaitForEnd() int {
	return <-instance.quit
}
