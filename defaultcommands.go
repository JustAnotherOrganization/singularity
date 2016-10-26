package singularity

import (
	"fmt"

	"github.com/JustAnotherOrganization/singularity/slacktypes"
)

func addDefaultCommands(instance *SlackInstance) {
	instance.Commands.registerCommand("test", testCommand) //TODO Get rid of the .
	instance.Commands.registerCommand("version", versionCommand)
	instance.Commands.registerCommand("setprefix", setCommmandPrefix)
}

func testCommand(command Command) {
	message := slacktypes.Message{}
	message.Text = "This is a test!"
	message.User = command.Instance.GetSelf().ID
	message.Channel = command.Channel.ID
	message.Type = "message"
	command.Instance.output <- Message{Body: message}
}

func versionCommand(command Command) {
	message := slacktypes.Message{}
	message.Text = "version is 0.0.1"
	message.User = command.Instance.GetSelf().ID
	message.Channel = command.Channel.ID
	message.Type = "message"
	command.Instance.output <- Message{Body: message}
}

func setCommmandPrefix(command Command) {
	fmt.Println("Setting prefix")
	if len(command.Args) == 0 || command.Args[0] == "" {
		message := slacktypes.Message{}
		message.Text = "Must specify what to set the prefix to!"
		message.User = command.Instance.GetSelf().ID
		message.Channel = command.Channel.ID
		message.Type = "message"
		command.Instance.output <- Message{Body: message}
	}

	command.Instance.Commands.setPrefix(command.Args[0])

	message := slacktypes.Message{}
	message.Text = "set the prefix to " + command.Args[0]
	message.User = command.Instance.GetSelf().ID
	message.Channel = command.Channel.ID
	message.Type = "message"
	command.Instance.output <- Message{Body: message}
}
