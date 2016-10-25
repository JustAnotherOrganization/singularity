package singularity

import (
	"fmt"
	"time"
)

//Log is used to log.
func (singularity *Singularity) Log(level int, message string, i ...interface{}) {
	singularity.log(level, message, i...)
}

//SetLogger ...
func (singularity *Singularity) SetLogger(logger func(level int, message string, i ...interface{})) {
	singularity.log = logger
}

//TODO Support passing a version of instance.
func defaultLogger(level int, message string, i ...interface{}) {
	prefix := " ["
	switch level {
	case ErrDebug:
		prefix += "D"
	case ErrInfo:
		prefix += "I"
	case ErrWarn:
		prefix += "W"
	case ErrError:
		prefix += "E"
	case ErrCrit:
		prefix += "X"
	default:
		prefix += "?"
	}
	prefix += "] "
	prefix += time.Now().Format(time.RFC3339) + " | "
	fmt.Println(prefix + fmt.Sprintf(message, i...))
}

//Error Constants
const (
	ErrDebug = iota
	ErrInfo
	ErrWarn
	ErrError
	ErrCrit
)
