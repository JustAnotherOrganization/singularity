package singularity

func addDefaultCommands(instance *SlackInstance) {
	instance.Commands.registerCommand("test", testCommand) //TODO Get rid of the .
	instance.Commands.registerCommand("version", versionCommand)
	instance.Commands.registerCommand("setprefix", setCommmandPrefix)
}

func testCommand(command Command) {
	message := Message{}
	message.Text = "This is a test!"
	message.Channel = command.Channel.ID
	command.Instance.SendMessage(message)
}

func versionCommand(command Command) {
	message := Message{}
	message.Text = "version is 0.0.1"
	message.Channel = command.Channel.ID
	command.Instance.SendMessage(message)
}

func setCommmandPrefix(command Command) {
	if len(command.Args) == 0 || command.Args[0] == "" {
		message := Message{}
		message.Text = "Must specify what to set the prefix to!"
		message.Channel = command.Channel.ID
		command.Instance.SendMessage(message)
	}

	command.Instance.Commands.setPrefix(command.Args[0])

	message := Message{}
	message.Text = "set the prefix to " + command.Args[0]
	message.Channel = command.Channel.ID
	command.Instance.SendMessage(message)
}
