package singularity

//SlackInstance ...
type SlackInstance struct {
	Name        string
	singularity *Singularity
	input       chan Message
	output      chan Message
	quit        chan int
}

//Quit ...
func (instance *SlackInstance) Quit() {
	instance.singularity.RemoveTeam(instance.Name) //Remove Yo Self.
	instance.quit <- 0
}

//Please dont call this outside of Singularity::Shutdown
func (instance *SlackInstance) quitShutdown() {
	instance.quit <- 0
}
