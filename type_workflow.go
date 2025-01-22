package workflow

import (
	"github.com/gouniverse/dataobject"
)

func NewWorkflow() WorkflowInterface {
	return &workflow{}
}

type workflow struct {
	dataobject.DataObject
}

func (wf *workflow) CreatedAt() string {
	return wf.Get(COLUMN_CREATED_AT)
}

func (wf *workflow) SetCreatedAt(createdAt string) WorkflowInterface {
	wf.Set(COLUMN_CREATED_AT, createdAt)
	return wf
}

func (wf *workflow) ID() string {
	return wf.Get(COLUMN_ID)
}

func (wf *workflow) SetID(id string) WorkflowInterface {
	wf.Set(COLUMN_ID, id)
	return wf
}

func (wf *workflow) Name() string {
	return wf.Get(COLUMN_NAME)
}

func (wf *workflow) SetName(name string) WorkflowInterface {
	wf.Set(COLUMN_NAME, name)
	return wf
}

func (wf *workflow) Status() string {
	return wf.Get(COLUMN_STATUS)
}

func (wf *workflow) SetStatus(status string) WorkflowInterface {
	wf.Set(COLUMN_STATUS, status)
	return wf
}

func (wf *workflow) UpdatedAt() string {
	return wf.Get(COLUMN_UPDATED_AT)
}

func (wf *workflow) SetUpdatedAt(updatedAt string) WorkflowInterface {
	wf.Set(COLUMN_UPDATED_AT, updatedAt)
	return wf
}

func (wf *workflow) WorkflowDefinitionID() string {
	return wf.Get(COLUMN_WORKFLOW_DEFINITION_ID)
}

func (wf *workflow) SetWorkflowDefinitionID(workflowDefinitionID string) WorkflowInterface {
	wf.Set(COLUMN_WORKFLOW_DEFINITION_ID, workflowDefinitionID)
	return wf
}

// // =============================================================================
// // CONSTRUCTOR
// // =============================================================================

// func NewWorkflow() *Workflow {
// 	return &Workflow{}
// }

// // =============================================================================
// // METHODS
// // =============================================================================

// func (wf *Workflow) AddStep(step Step) *Workflow {
// 	wf.steps = append(wf.steps, step)
// 	return wf
// }

// func (wf *Workflow) Run() {
// 	steps := wf.Steps()
// 	log.Println("START:", wf.Title())
// 	for _, step := range steps {
// 		step.Run()
// 	}
// 	log.Println("END:", wf.Title())
// }

// // =============================================================================
// // SETTERS AND GETTERS
// // =============================================================================

// func (wf *Workflow) ID() string {
// 	return wf.id
// }

// func (wf *Workflow) SetID(id string) *Workflow {
// 	wf.id = id
// 	return wf
// }

// func (wf *Workflow) Title() string {
// 	return wf.title
// }

// func (wf *Workflow) SetTitle(title string) *Workflow {
// 	wf.title = title
// 	return wf
// }

// func (wf *Workflow) Steps() []Step {
// 	return wf.steps
// }

// func (wf *Workflow) SetSteps(steps []Step) *Workflow {
// 	wf.steps = steps
// 	return wf
// }

// func (wf *Workflow) CreatedAt() string {
// 	return wf.createdAt
// }

// func (wf *Workflow) SetCreatedAt(createdAt string) *Workflow {
// 	wf.createdAt = createdAt
// 	return wf
// }

// func (wf *Workflow) UpdatedAt() string {
// 	return wf.updatedAt
// }

// func (wf *Workflow) SetUpdatedAt(updatedAt string) *Workflow {
// 	wf.updatedAt = updatedAt
// 	return wf
// }

// func (wf *Workflow) Status() string {
// 	return wf.status
// }

// func (wf *Workflow) SetStatus(status string) *Workflow {
// 	wf.status = status
// 	return wf
// }
