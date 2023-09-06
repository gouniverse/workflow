package workflow

const (
	QUEUE_STEP_STATUS_QUEUED    = "queued"
	QUEUE_STEP_STATUS_RUNNING   = "running"
	QUEUE_STEP_STATUS_SUCCESS   = "success"
	QUEUE_STEP_STATUS_ERROR     = "error"
	QUEUE_STEP_STATUS_CANCELED  = "canceled"
	QUEUE_STEP_STATUS_SCHEDULED = "scheduled"
)

type QueueStep struct {
	id              string
	status          string
	queueWorkflowID string
	parentID        string
	stepID          string
	createdAt       string
	updatedAt       string
}

// =============================================================================
// CONSTRUCTOR
// =============================================================================

// =============================================================================
// METHODS
// =============================================================================

// =============================================================================
// SETTERS AND GETTERS
// =============================================================================

func (step *QueueStep) ID() string {
	return step.id
}

func (step *QueueStep) SetID(id string) *QueueStep {
	step.id = id
	return step
}

func (step *QueueStep) Status() string {
	return step.status
}

func (step *QueueStep) SetStatus(status string) *QueueStep {
	step.status = status
	return step
}

func (step *QueueStep) QueueWorkflowID() string {
	return step.queueWorkflowID
}

func (step *QueueStep) SetQueueWorkflowID(queueWorkflowID string) *QueueStep {
	step.queueWorkflowID = queueWorkflowID
	return step
}

func (step *QueueStep) ParentID() string {
	return step.parentID
}

func (step *QueueStep) SetParentID(parentID string) *QueueStep {
	step.parentID = parentID
	return step
}

func (step *QueueStep) StepID() string {
	return step.stepID
}

func (step *QueueStep) SetStepID(stepID string) *QueueStep {
	step.stepID = stepID
	return step
}

func (step *QueueStep) CreatedAt() string {
	return step.createdAt
}

func (step *QueueStep) SetCreatedAt(createdAt string) *QueueStep {
	step.createdAt = createdAt
	return step
}

func (step *QueueStep) UpdatedAt() string {
	return step.updatedAt
}

func (step *QueueStep) SetUpdatedAt(updatedAt string) *QueueStep {
	step.updatedAt = updatedAt
	return step
}
