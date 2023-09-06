package workflow

import "log"

const (
	WORKFLOW_STATUS_ACTIVE   = "active"
	WORKFLOW_STATUS_CANCELED = "canceled"
	WORKFLOW_STATUS_STOPPED  = "stopped"
)

type Workflow struct {
	id        string
	status    string
	title     string
	steps     []Step
	createdAt string
	updatedAt string
}

// =============================================================================
// CONSTRUCTOR
// =============================================================================

// =============================================================================
// METHODS
// =============================================================================

func (wf *Workflow) AddStep(step Step) *Workflow {
	wf.steps = append(wf.steps, step)
	return wf
}

func (wf *Workflow) Run() {
	steps := wf.Steps()
	log.Println("START:", wf.Title())
	for _, step := range steps {
		step.Run()
	}
	log.Println("END:", wf.Title())
}

// =============================================================================
// SETTERS AND GETTERS
// =============================================================================

func (wf *Workflow) ID() string {
	return wf.id
}

func (wf *Workflow) SetID(id string) *Workflow {
	wf.id = id
	return wf
}

func (wf *Workflow) Title() string {
	return wf.title
}

func (wf *Workflow) SetTitle(title string) *Workflow {
	wf.title = title
	return wf
}

func (wf *Workflow) Steps() []Step {
	return wf.steps
}

func (wf *Workflow) SetSteps(steps []Step) *Workflow {
	wf.steps = steps
	return wf
}

func (wf *Workflow) CreatedAt() string {
	return wf.createdAt
}

func (wf *Workflow) SetCreatedAt(createdAt string) *Workflow {
	wf.createdAt = createdAt
	return wf
}

func (wf *Workflow) UpdatedAt() string {
	return wf.updatedAt
}

func (wf *Workflow) SetUpdatedAt(updatedAt string) *Workflow {
	wf.updatedAt = updatedAt
	return wf
}

func (wf *Workflow) Status() string {
	return wf.status
}

func (wf *Workflow) SetStatus(status string) *Workflow {
	wf.status = status
	return wf
}
