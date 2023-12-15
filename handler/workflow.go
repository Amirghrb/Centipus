package handler

import (
	"github.com/Amirghrb/octopus/state"
)

type WorkflowHandler struct {
	State state.St
	// Workflow state.Workflow
}

func (wh *WorkflowHandler) Add(wf state.Workflow) int {
	lastState := state.CloneStore((*wh.State)[state.GetLstId()])
	lastState.Workflow = append(lastState.Workflow, wf)
	state.Mutate(wh.State, lastState)
	return 1
}

// func (w *WorkflowHandler) update(id int) int{

// }

// func (w *WorkflowHandler) delete(s ) state.Workflow {}

// func (w *WorkflowHandler) Get(id int) state.Workflow {
// 	return w.store.Workflow[id]
// }

// func (w *WorkflowHandler) GetAll() []state.Workflow {
// 	return w.store.Workflow
// }
