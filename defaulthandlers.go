package singularity

import "strings"

/*
Default Handlers.
These are automatically loaded to help handle the singularity instance and keep the data up to date.
*/

//Load is used to load the default handlers.
func addDefaultHandlers(instance *SlackInstance) {
	instance.RegisterHandler("message", MessageHandler)
	instance.RegisterHandler("message", HandleCommands)

	instance.RegisterHandler("bot_added", handleBotAddedEvent)
	instance.RegisterHandler("bot_changed", handleBotChangedEvent)
}

// MessageHandler - if this is a message and we have a handler, handle it
func MessageHandler(message Message, instance *SlackInstance) {
	//message_deleted Subtype
	if message.SubType != "message_deleted" {
	}
}

// HandleCommands - if this is a command and we have a handler, handle it
func HandleCommands(message Message, instance *SlackInstance) {
	if message.SubType != "message_deleted" { //If it isn't deleted
		if message.User != instance.GetSelf().ID { //If it isn't me
			cmds := strings.Split(message.Text, " ")
			if len(cmds) > 0 {
				cmd := cmds[0]
				if instance.Commands.IsCommand(cmd) {
					cmds = cmds[1:]
					c := Command{Command: cmd, Args: cmds, Instance: instance, User: *instance.GetUserByID(message.User), Team: *instance.GetTeam(), Channel: *instance.GetChannelByID(message.Channel)}
					instance.Commands.execute(c)
				}
			}
		}
	}
}

// handleBotAddedEvent handles when a bot_added event happens. TODO discuss exportation
/*{
    "type": "bot_added",
    "bot": {
        "id": "B024BE7LH",
        "name": "hugbot",
        "icons": {
            "image_48": "https:\/\/slack.com\/path\/to\/hugbot_48.png"
        }
    }
}*/
func handleBotAddedEvent(botAdded struct {
	Bot Bot `json:"bot"`
}, instance *SlackInstance) {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()
	instance.rtmResp.Bots = append(instance.rtmResp.Bots, botAdded.Bot)
}

/*{
    "type": "bot_changed",
    "bot": {
        "id": "B024BE7LH",
        "name": "hugbot",
        "icons": {
            "image_48": "https:\/\/slack.com\/path\/to\/hugbot_48.png"
        }
    }
}*/
func handleBotChangedEvent(botChanged struct {
	Bot Bot `json:"bot"`
}, instance *SlackInstance) {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()
	for i := 0; i < len(instance.rtmResp.Bots); i++ {
		if instance.rtmResp.Bots[i].ID == botChanged.Bot.ID {
			instance.rtmResp.Bots[i] = botChanged.Bot
		}
	}
}
