package singularity_test

import (
	"fmt"

	"github.com/JustAnotherOrganization/singularity"
)

// This example shows the basic usage for deploying a bot using singularity.
func Example_main() {
	s := singularity.NewSingularity()

	team := s.NewTeam("xoxb-slackToken")

	//Example of registering an event handler
	team.RegisterHandler("message", func(message singularity.Message, team *singularity.SlackInstance) {
		if message.User != "" {
			user := team.GetUserByID(message.User)
			if message.SubType != "message_deleted" { //If it isn't deleted.
				fmt.Printf("%v said %v\n", user.Name, message.Text)
			} else {
				fmt.Printf("%v deleted %v\n", user.Name, message.Text)
			}
		}
	})

	if err := team.Start(); err != nil {
		panic(err)
	}

	s.WaitForShutdown()
}
