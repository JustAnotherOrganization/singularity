package singularity

import "sync"

//NewSingularity creates a new singularity instance.
func NewSingularity() *Singularity {
	n := &Singularity{
		shutdown: make(chan int),
		log:      defaultLogger,
	}
	return n
}

//Singularity ... //TODO Complete
type Singularity struct {
	sync.Mutex
	Teams    []SlackInstance
	shutdown chan int
	log      func(level int, message string, i ...interface{})
}

//Shutdown leaves everything. It blocks as each team has to quit.
func (singularity *Singularity) Shutdown() {
	singularity.Lock()
	defer singularity.Unlock()
	for i := 0; i < len(singularity.Teams); i++ {
		singularity.Teams[i].quitShutdown() //Warning, this blocks.
		singularity.Teams = append(singularity.Teams[:i], singularity.Teams[i+1:]...)
		i--
	}
	singularity.shutdown <- 0
}

//RemoveTeam ...
func (singularity *Singularity) RemoveTeam(token string) {
	singularity.Lock()
	defer singularity.Unlock()
	for i := 0; i < len(singularity.Teams); i++ { //a[:i], a[i+1:]...
		if singularity.Teams[i].token == token {
			singularity.Teams = append(singularity.Teams[:i], singularity.Teams[i+1:]...)
			i--
		}
	}
}

//HardShutdown doesn't go to each team's shutdown.
func (singularity *Singularity) HardShutdown() {
	singularity.shutdown <- 0
}

//WaitForShutdown blocks until shutdown.
func (singularity *Singularity) WaitForShutdown() <-chan int {
	return singularity.shutdown
}
