package state

type State struct {
	Job      []job
	Workflow []workflow
	Server   []server
	Script   []script
}

func (s *State) add() {

}
