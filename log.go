package singularity

import "fmt"

//Log is used to log.
func (singularity *Singularity) Log(message string, i ...interface{}) {
	singularity.log(message, i...)
}

//SetLogger ...
func (singularity *Singularity) SetLogger(logger func(message string, i ...interface{})) {
	singularity.log = logger
}

func defaultLogger(message string, i ...interface{}) {
	message = fmt.Sprintf(message, i...)
	fmt.Println(message)
}
