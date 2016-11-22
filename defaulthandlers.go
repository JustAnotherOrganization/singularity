package singularity

import "strings"

/*
Default Handlers.
These are automatically loaded to help handle the singularity instance and keep the data up to date.
*/

//Load is used to load the default handlers.
func addDefaultHandlers(instance *SlackInstance) {
	instance.RegisterHandler("message", HandleCommands)
	instance.RegisterHandler("message", messageHandlerForSubtypes)

	instance.RegisterHandler("bot_added", handleBotAddedEvent)
	instance.RegisterHandler("bot_changed", handleBotChangedEvent)

	instance.RegisterHandler("channel_created", channelCreated)
	instance.RegisterHandler("channel_deleted", channelDeleted)
	instance.RegisterHandler("channel_archived", channelArchived)
	instance.RegisterHandler("channel_unarchived", channelUnarchived)
	instance.RegisterHandler("channel_joined", channelJoined)
	instance.RegisterHandler("channel_left", channelLeft)
	// TODO handle channel_history_changed event

	instance.RegisterHandler("error", handleError)
}

// This function is needed to handle the important events that come through `message`, marked by subtype (ugh)
func messageHandlerForSubtypes(message Message, instance *SlackInstance) {
	if message.SubType != "" {
		// TODO Go find the list of subtypes: https://api.slack.com/events/message
		switch message.SubType {

		default:
			// TODO: This subtype is yet to be supported!
		}
	}
}

// HandleCommands - if this is a command and we have a handler, handle it
// Do not lock rtmResp here since commands might need to access it.
func HandleCommands(message Message, instance *SlackInstance) {
	if message.SubType == "" { //No subtypes please.
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

/*{
    "type": "error",
    "error": {
        "code": 1,
        "msg": "Socket URL has expired"
    }
}*/
func handleError(err struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	} `json:"error"`
}, instance *SlackInstance) {
	instance.log(LogError, "Slack sent an error(%v): %v", err.Error.Code, err.Error.Message)
}

/*{
    "type": "channel_created",
		"channel": {
         ...
     }
}*/
func channelCreated(channelCreated struct {
	Channel Channel `json:"channel"`
}, instance *SlackInstance) {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()
	instance.rtmResp.Channels = append(instance.rtmResp.Channels, channelCreated.Channel)
}

/*{
    "type": "channel_deleted",
    "channel": "C024BE91L"
}*/
func channelDeleted(channelDeleted struct {
	Channel string `json:"channel"`
}, instance *SlackInstance) {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()
	for i := 0; i < len(instance.rtmResp.Channels); i++ {
		if instance.rtmResp.Channels[i].ID == channelDeleted.Channel {
			instance.rtmResp.Channels = append(instance.rtmResp.Channels[:i], instance.rtmResp.Channels[1+i:]...)
			i--
		}
	}
}

/*{
    "type": "channel_archive",
    "channel": "C024BE91L",
    "user": "U024BE7LH"
}*/
func channelArchived(channelArchived struct {
	Channel string `json:"channel"`
	User    string `json:"user"`
}, instance *SlackInstance) {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()
	for i := 0; i < len(instance.rtmResp.Channels); i++ {
		if instance.rtmResp.Channels[i].ID == channelArchived.Channel {
			instance.rtmResp.Channels[i].IsArchived = true
		}
	}
}

/*{
    "type": "channel_unarchive",
    "channel": "C024BE91L",
    "user": "U024BE7LH"
}*/
func channelUnarchived(channelUnarchived struct {
	Channel string `json:"channel"`
	User    string `json:"user"`
}, instance *SlackInstance) {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()
	for i := 0; i < len(instance.rtmResp.Channels); i++ {
		if instance.rtmResp.Channels[i].ID == channelUnarchived.Channel {
			instance.rtmResp.Channels[i].IsArchived = false
		}
	}
}

/*{
    "type": "channel_left",
    "channel": "C024BE91L"
}*/
func channelLeft(channelLeft struct {
	Channel string `json:"channel"`
}, instance *SlackInstance) {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()
	for i := 0; i < len(instance.rtmResp.Channels); i++ {
		if instance.rtmResp.Channels[i].ID == channelLeft.Channel {
			instance.rtmResp.Channels[i].IsMember = false
			for ii := 0; ii < len(instance.rtmResp.Channels[ii].Members); ii++ {
				if instance.rtmResp.Channels[i].Members[ii] == instance.rtmResp.Self.ID {
					instance.rtmResp.Channels[i].Members = append(instance.rtmResp.Channels[i].Members[:ii], instance.rtmResp.Channels[i].Members[ii+1:]...)
					ii--
				}
			}
		}
	}
}

/*{
    "type": "channel_joined",
    "channel": {
        ...
    }
}*/
func channelJoined(channelJoined struct {
	Channel Channel `json:"channel"`
}, instance *SlackInstance) {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()
	exist := false
	for i := 0; i < len(instance.rtmResp.Channels); i++ {
		if instance.rtmResp.Channels[i].ID == channelJoined.Channel.ID {
			// For now, it is easier just to override the object as we can assume that they send us the full object as it should be.
			instance.rtmResp.Channels[i] = channelJoined.Channel
			exist = true
		}
	}
	// Incase we dont have the channel cached.
	if !exist {
		instance.rtmResp.Channels = append(instance.rtmResp.Channels, channelJoined.Channel)
	}
}

/*{
    "type": "channel_rename",
    "channel": {
        "id":"C02ELGNBH",
        "name":"new_name",
        "created":1360782804
    }
}*/
func channelRename(channelRename struct {
	Channel Channel `json:"channel"`
}, instance *SlackInstance) {
	instance.rtmResp.Lock()
	defer instance.rtmResp.Unlock()
	exist := false
	for i := 0; i < len(instance.rtmResp.Channels); i++ {
		if instance.rtmResp.Channels[i].ID == channelRename.Channel.ID {
			instance.rtmResp.Channels[i].Name = channelRename.Channel.Name
			exist = true
		}
	}
	if !exist {
		instance.rtmResp.Channels = append(instance.rtmResp.Channels, channelRename.Channel)
	}
}
