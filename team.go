package singularity

import (
	"fmt"

	"golang.org/x/net/websocket"
)

//SlackInstance ...
type SlackInstance struct {
	RTMResp
	handlers    *EventAPIHandler
	Name        string
	singularity *Singularity
	input       chan Message
	output      chan Message
	quit        chan int
	connection  *websocket.Conn
	Commands    CommandHandler
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

func (singularity *Singularity) addTeam(connection *websocket.Conn, response RTMResp) (*SlackInstance, error) {
	singularity.Lock()
	defer singularity.Unlock()

	var instance SlackInstance
	instance.singularity = singularity //Set reference.
	instance.RTMResp = response
	instance.Name = response.Team.Name

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

func (instance *SlackInstance) RegisterHandler(key string, handler interface{}) error {
	instance.handlers.Lock()
	defer instance.handlers.Unlock()
	return instance.handlers.registerHandler(key, handler)
}
