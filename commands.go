package singularity

import (
	"errors"
	"strings"
	"sync"
)

// CommandHandler holds command handlers
// mutex for locking when adding handlers
// handlers is a map holding all CommandHandlers
type CommandHandler struct {
	sync.Mutex
	handlers map[string]func(Command)
}

// Command type
type Command struct {
	Command  string
	Args     []string
	Instance *SlackInstance
	Channel  Channel
	User     User
	Team     Team
}

func (commandHandler *CommandHandler) execute(command Command) error {
	commandHandler.Lock()
	defer func() {
		if err := recover(); err != nil {
			//TODO handle error.
		}
	}()
	defer commandHandler.Unlock()

	function, ok := commandHandler.handlers[command.Command]
	if !ok {
		//TODO return error.
	}

	function(command)

	return nil
}

// IsCommand check if we have a handler for this command
func (commandHandler *CommandHandler) IsCommand(command string) bool {
	commandHandler.Lock()
	defer commandHandler.Unlock()

	_, ok := commandHandler.handlers[strings.ToLower(command)]
	return ok
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
