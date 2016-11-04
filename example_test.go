package singularity_test

import "github.com/JustAnotherOrganization/singularity"

var team singularity.SlackInstance

// This example shows the basic usage for deploying a bot using singularity.
func Example_basic() {
	s := singularity.NewSingularity()

	team := s.NewTeam("xoxb-slackToken")

	if err := team.Start(); err != nil {
		panic(err)
	}

	s.WaitForShutdown()
}

// This example shows how to register a basic handler.
func Example_registerhandler() {
	team.RegisterHandler("message", func(message singularity.Message, team *singularity.SlackInstance) {
		if message.SubType != "message_deleted" && message.User != "" {
			team.Log(singularity.LogInfo, "%v said %v", team.GetUserByID(message.User).Name, message.Text)
		}
	})
}

// This example shows how to register a basic command.
func Example_registercommand() {
	team.RegisterCommand("ping", func(command singularity.Command) {
		command.Instance.SendMessage(singularity.Message{Text: "Pong!", Channel: command.Channel.ID})
	})
}
