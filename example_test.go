package singularity_test

import "github.com/JustAnotherOrganization/singularity"

// This example shows the basic usage for deploying a bot using singularity.
func Example_use() {
	s := singularity.NewSingularity()

	team := s.NewTeam("xoxb-slackToken")

	//Example of registering an event handler
	team.RegisterHandler("message", func(message singularity.Message, team *singularity.SlackInstance) {
		if message.User != "" {
			user := team.GetUserByID(message.User)
			if message.SubType != "message_deleted" { //If it isn't deleted.
				team.Log(singularity.LogInfo, "%v said %v", user.Name, message.Text)
			} else {
				team.Log(singularity.LogInfo, "%v deleted %v", user.Name, message.Text)
			}
		}
	})

	if err := team.Start(); err != nil {
		panic(err)
	}

	s.WaitForShutdown()
}
