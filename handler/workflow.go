package handler

import (
	"github.com/Amirghrb/octopus/state"
)

type WorkflowHandler struct {
	store *state.Store
	// Workflow state.Workflow
}

func (wh *WorkflowHandler) Add(wf state.Workflow) int {
	wh.store.Workflow = append(wh.store.Workflow, wf)
	return 1
}

// func (w *WorkflowHandler) update(s state.State) state.Workflow {}

// func (w *WorkflowHandler) delete(s state.State) state.Workflow {}

// func (w *WorkflowHandler) get(s state.State) state.Workflow {}
