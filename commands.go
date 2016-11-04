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
	caseSensitive bool
	prefix        string
	handlers      map[string]func(Command)
}

func (commandHandler *CommandHandler) setPrefix(prefix string) {
	commandHandler.Lock()
	defer commandHandler.Unlock()

	commandHandler.prefix = prefix
}

//NewCommandHandler returns a new command handler
func NewCommandHandler() *CommandHandler {
	return &CommandHandler{handlers: make(map[string]func(Command))}
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
	defer func() {
		if err := recover(); err != nil {
			//TODO handle error.
		}
	}()
	command.Command = strings.Replace(command.Command, commandHandler.prefix, "", 1)
	function := commandHandler.getCommand(command.Command)

	if function != nil {
		(*function)(command)
	}

	return nil
}

func (commandHandler *CommandHandler) getCommand(command string) *func(Command) {
	commandHandler.Lock()
	defer commandHandler.Unlock()
	if function, ok := commandHandler.handlers[commandHandler.toLower(command)]; ok {
		return &function
	}
	return nil
}

// IsCommand check if we have a handler for this command
func (commandHandler *CommandHandler) IsCommand(command string) bool {
	if command == "" {
		return false
	}
	commandHandler.Lock()
	defer commandHandler.Unlock()
	if !strings.HasPrefix(command, commandHandler.prefix) {
		return false
	}
	command = strings.Replace(command, commandHandler.prefix, "", 1)
	_, ok := commandHandler.handlers[commandHandler.toLower(command)]
	return ok
}

func (commandHandler *CommandHandler) registerCommand(command string, function func(Command)) error {
	commandHandler.Lock()
	defer commandHandler.Unlock()

	if _, ok := commandHandler.handlers[strings.ToLower(command)]; ok {
		return errors.New("Command already registered!") //TODO Maybe replace?
	}

	commandHandler.handlers[commandHandler.toLower(command)] = function
	return nil
}

func (commandHandler *CommandHandler) toLower(s string) string {
	if !commandHandler.caseSensitive {
		s = strings.ToLower(s)
	}
	return s
}
