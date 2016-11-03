# Singularity
[![GoDoc](https://godoc.org/github.com/JustAnotherOrganization/singularity?status.png)](https://godoc.org/github.com/JustAnotherOrganization/singularity)    
---
A Simple Slack Client/Framework in golang. Allows you to easily hook up commands and event listeners, or will at some point. Right now it is in progress! Nothing really here but the writings of a mad man.

## Building and using   
As of writing this...   
`go get github.com/JustAnotherOrganization/singularity`   
``` Go
package main

import (
	"fmt"

	"github.com/JustAnotherOrganization/singularity"
)

func main() {
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
```

As we add more features, more documentation will appear.

## Contributing  
If you'd like to contribute, feel free and welcome to fork this project and open a PullRequest. 

## License   

Copyright 2016 [Just Another Organization](https://github.com/JustAnotherOrganization)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use these files except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
