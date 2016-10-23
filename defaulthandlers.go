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

//Load is used to load the default handlers.
func addDefaultHandlers(instance *SlackInstance) {
	instance.RegisterHandler("message", MessageHandler)
	instance.RegisterHandler("message", HandleCommands)
}

func MessageHandler(message slacktypes.Message, instance *SlackInstance) {
	//message_deleted Subtype
	if message.SubType != "message_deleted" {
		if strings.HasPrefix(message.Text, ".test") {
			message.Text = "This is a test!"
			message.User = instance.Self.ID
			instance.output <- Message{Body: message}
		}
	}
}

func HandleCommands(message slacktypes.Message, instance *SlackInstance) {
	if message.SubType != "message_deleted" { //If it isn't deleted
		if message.User != instance.Self.ID { //If it isn't me
			cmds := strings.Split(message.Text, " ")
			//If is command.
			cmd := "getCommand()"
			if len(cmds) > 1 {
				cmds = cmds[1:] //sl[:len(sl)-1]
			} else {
				cmds = []string{}
			}
			c := Command{Command: cmd, Args: cmds, Instance: instance}
			fmt.Println(c)
		}
	}
}
