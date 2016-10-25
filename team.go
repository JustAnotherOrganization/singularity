package singularity

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

//SlackInstance ...
type SlackInstance struct {
	RTMResp

	//Important Data Stuff
	singularity *Singularity
	connection  *websocket.Conn

	//Handlers
	Commands               *CommandHandler
	handlers               *EventAPIHandler
	customEventHandler     func()
	customCommandHandler   func()
	customWebsocketHandler func()

	//Friendly name
	Name string

	//Token
	token string

	input  chan Message //input represents what comes from slack
	output chan Message //output goes to slack
	quit   chan int

	//Configuration
	Configuration Configuration
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
	instance.Configuration = defaultConfig{config: make(map[string]interface{})} //TODO move outside. configs should be configured before a team is started.
	//defaulthandlers
	instance.handlers = NewHandler1()
	addDefaultHandlers(instance)

	instance.Commands.handlers = make(map[string]func(Command))

	singularity.Teams = append(singularity.Teams, *instance)
	return instance
}

//Start starts the slack instance.
func (instance *SlackInstance) Start() error {
	helper := HTTPHelper{Client: &http.Client{Transport: instance.transport}, Transport: "https://"} //TODO Don't hard code this.
	_, err := helper.post("rtm.start", &instance.RTMResp, "token", instance.token)
	if err != nil {
		return err
	}

	//Connect to websocket.
	//TODO Support custom proxies?
	instance.connection, err = websocket.Dial(instance.RTMResp.URL, "", "http://api.slack.com") //TODO Don't hard code here.
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
	instance.Name = instance.RTMResp.Team.Name

	//Channels for this Instance.
	instance.input = make(chan Message, 5)  //TODO Configure Amount
	instance.output = make(chan Message, 5) //TODO Configure Amount
	instance.quit = make(chan int)
	//WebsocketChannel
	// instance.connection = connection

	//Start Go-Routines for handling the things.
	if instance.customEventHandler != nil {
		go instance.customEventHandler()
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

func (singularity *Singularity) addTeam(connection *websocket.Conn, response RTMResp) (*SlackInstance, error) {
	singularity.Lock()
	defer singularity.Unlock()

	var instance SlackInstance
	instance.singularity = singularity //Set reference.
	instance.RTMResp = response
	instance.Name = response.Team.Name
	instance.Configuration = defaultConfig{config: make(map[string]interface{})} //TODO move outside. configs should be configured before a team is started.
	instance.Commands.handlers = make(map[string]func(Command))

	instance.handlers = NewHandler1()
	addDefaultHandlers(&instance)
	//TODO Handlers are implemented here.
	//TODO Add default handlers for handling Change/Alter type request.

	//Channels for this Instance.
	instance.input = make(chan Message, 5)  //TODO Configure Amount
	instance.output = make(chan Message, 5) //TODO Configure Amount
	instance.quit = make(chan int)
	//WebsocketChannel
	instance.connection = connection

	singularity.Teams = append(singularity.Teams, instance)

	//Start Go-Routines for handling the things.
	go instance.handleChans()
	go instance.handleWebsocket()

	return nil, nil
}

//TODO Panic Recovery
func (instance *SlackInstance) handleChans() {
	for {
		select {
		//<-instance.input reads from (what is assumed to be) slack.
		case val := <-instance.input:
			event, err := val.GetBytes()
			if err != nil {
				//TODO Shit.
			}

			fmt.Println(string(event))
			slackType := preParseString("type", event)
			fmt.Printf("Type: %v\n", slackType)
			if slackType != "" {
				instance.handlers.execute(slackType, event, instance)
			}

			//<-instance.output sends the thing to slack.
		case val := <-instance.output:
			thingToSend, err := val.GetInterface()
			bytes, _ := val.GetBytes()
			fmt.Printf("Sending %+v\n", string(bytes))
			if err != nil {
				//TODO Fuck.
			}
			err = websocket.JSON.Send(instance.connection, thingToSend) //SendBodies  :D
			if err != nil {
				//TODO fuck
			}
		case <-instance.quit:
			return //Die, commie.
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
			//defer Recover()
			err := websocket.JSON.Receive(instance.connection, &i)
			if err != nil {
				//TODO ERROR HANDLING
				return
			}
			instance.input <- Message{i} //Buffered.
		}()
	}
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

//SetHTTPTransport allows you to set a transport to use when communicate with Slack.
func (instance *SlackInstance) SetHTTPTransport(transport *http.Transport) {
	instance.transport = transport
}

//WaitForEnd will wait until the team is killed or stopped.
func (instance *SlackInstance) WaitForEnd() <-chan int {
	return instance.quit
}
