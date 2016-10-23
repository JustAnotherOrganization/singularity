package singularity

import (
	"errors"
	"strings"
	"sync"
)

type CommandHandler struct {
	sync.Mutex
	handlers map[string]func(Command)
}

type Command struct {
	Command  string
	Args     []string
	Instance *SlackInstance
	Channel  Channel
	User     User
	Team     Team
}

func execute(command Command) error {

	return nil
}

func (commandHandler *CommandHandler) registerCommand(command string, function func(Command)) error {
	commandHandler.Lock()
	defer commandHandler.Unlock()

	if _, ok := commandHandler.handlers[strings.ToLower(command)]; ok {
		return errors.New("Command already registered!") //TODO Maybe replace?
	}

	commandHandler.handlers[strings.ToLower(command)] = function
	return nil
}
