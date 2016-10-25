package singularity

import (
	"fmt"
	"strings"

	"github.com/JustAnotherOrganization/singularity/slacktypes"
)

/*
Default Handlers.
These are automatically loaded to help handle the singularity instance and keep the data up to date.
*/

// TODO List
// - Pass instance?
//

func testCommandHandler(command Command) {
	fmt.Println(command)
	message := slacktypes.Message{}
	message.Text = "This is a test!"
	message.User = command.Instance.Self.ID
	message.Channel = command.Channel.Name
	message.Type = "message"
	command.Instance.output <- Message{Body: message}
}

//Load is used to load the default handlers.
func addDefaultHandlers(instance *SlackInstance) {
	instance.RegisterHandler("message", MessageHandler)
	instance.RegisterHandler("message", HandleCommands)
	instance.Commands.registerCommand(".test", testCommandHandler)
}

// MessageHandler - if this is a message and we have a handler, handle it
func MessageHandler(message slacktypes.Message, instance *SlackInstance) {
	//message_deleted Subtype
	if message.SubType != "message_deleted" {
	}
}

// HandleCommands - if this is a command and we have a handler, handle it
func HandleCommands(message slacktypes.Message, instance *SlackInstance) {
	if message.SubType != "message_deleted" { //If it isn't deleted
		if message.User != instance.Self.ID { //If it isn't me
			cmds := strings.Split(message.Text, " ")
			if len(cmds) > 0 {
				cmd := cmds[0]
				if instance.Commands.IsCommand(cmd) {
					cmds = cmds[1:]
					c := Command{Command: cmd, Args: cmds, Instance: instance, Channel: Channel{Name: message.Channel}}
					instance.Commands.execute(c)
				}
			}
		}
	}
}
