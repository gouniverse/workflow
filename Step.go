package workflow

import "log"

type Step struct {
	id         string
	workflowID string
	parentID   string
	title      string
	steps      []Step
	createdAt  string
	updatedAt  string
}

// =============================================================================
// CONSTRUCTOR
// =============================================================================

// =============================================================================
// METHODS
// =============================================================================

func (s *Step) AddStep(step Step) *Step {
	s.steps = append(s.steps, step)
	return s
}

func (s *Step) Run() {
	log.Println("START:", s.Title())
	steps := s.Steps()
	for _, step := range steps {
		step.Run()
	}
	log.Println("END:", s.Title())
}

// =============================================================================
// SETTERS AND GETTERS
// =============================================================================
func (step *Step) ID() string {
	return step.id
}

func (step *Step) SetID(id string) *Step {
	step.id = id
	return step
}

func (step *Step) Title() string {
	return step.title
}

func (step *Step) SetTitle(title string) *Step {
	step.title = title
	return step
}

func (step *Step) Steps() []Step {
	return step.steps
}

func (step *Step) SetSteps(steps []Step) *Step {
	step.steps = steps
	return step
}

func (step *Step) ParentID() string {
	return step.parentID
}

func (step *Step) SetParentID(parentID string) *Step {
	step.parentID = parentID
	return step
}

func (step *Step) WorkflowID() string {
	return step.workflowID
}

func (step *Step) SetWorkflowID(workflowID string) *Step {
	step.workflowID = workflowID
	return step
}

func (step *Step) CreatedAt() string {
	return step.createdAt
}

func (step *Step) SetCreatedAt(createdAt string) *Step {
	step.createdAt = createdAt
	return step
}

func (step *Step) UpdatedAt() string {
	return step.updatedAt
}

func (step *Step) SetUpdatedAt(updatedAt string) *Step {
	step.updatedAt = updatedAt
	return step
}
