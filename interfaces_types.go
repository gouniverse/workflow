package workflow

type WorkflowDefinitionInterface interface {
	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	ID() string
	SetID(id string) WorkflowDefinitionInterface

	Name() string
	SetName(name string) WorkflowDefinitionInterface

	CreatedAt() string
	SetCreatedAt(createdAt string) WorkflowDefinitionInterface

	SoftDeletedAt() string
	SetSoftDeletedAt(softDeletedAt string) WorkflowDefinitionInterface

	UpdatedAt() string
	SetUpdatedAt(updatedAt string) WorkflowDefinitionInterface
}

type WorkflowInterface interface {
	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	ID() string
	SetID(id string) WorkflowInterface

	WorkflowDefinitionID() string
	SetWorkflowDefinitionID(workflowDefinitionID string) WorkflowInterface

	Status() string
	SetStatus(status string) WorkflowInterface

	Name() string
	SetName(name string) WorkflowInterface

	CreatedAt() string
	SetCreatedAt(createdAt string) WorkflowInterface

	UpdatedAt() string
	SetUpdatedAt(updatedAt string) WorkflowInterface
}

type StepDefinitionInterface interface {
	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	ID() string
	SetID(id string) StepDefinitionInterface

	WorkflowDefinitionID() string
	SetWorkflowDefinitionID(workflowDefinitionID string) StepDefinitionInterface

	Status() string
	SetStatus(status string) StepDefinitionInterface

	Name() string
	SetName(name string) StepDefinitionInterface

	CreatedAt() string
	SetCreatedAt(createdAt string) StepDefinitionInterface

	UpdatedAt() string
	SetUpdatedAt(updatedAt string) StepDefinitionInterface

	SoftDeletedAt() string
	SetSoftDeletedAt(softDeletedAt string) StepDefinitionInterface
}

type StepInterface interface {
	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	ID() string
	SetID(id string) StepInterface

	WorkflowID() string
	SetWorkflowID(workflowID string) StepInterface

	StepDefinitionID() string
	SetStepDefinitionID(stepDefinitionID string) StepInterface

	Status() string
	SetStatus(status string) StepInterface

	Name() string
	SetName(name string) StepInterface

	CreatedAt() string
	SetCreatedAt(createdAt string) StepInterface

	UpdatedAt() string
	SetUpdatedAt(updatedAt string) StepInterface
}

type EdgeDefinitionInterface interface {
	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	ID() string
	SetID(id string) EdgeDefinitionInterface

	WorkflowDefinitionID() string
	SetWorkflowDefinitionID(workflowDefinitionID string) EdgeDefinitionInterface

	FromStepDefinitionID() string
	SetFromStepDefinitionID(fromStepDefinitionID string) EdgeDefinitionInterface

	ToStepDefinitionID() string
	SetToStepDefinitionID(toStepDefinitionID string) EdgeDefinitionInterface

	CreatedAt() string
	SetCreatedAt(createdAt string) EdgeDefinitionInterface

	UpdatedAt() string
	SetUpdatedAt(updatedAt string) EdgeDefinitionInterface

	SoftDeletedAt() string
	SetSoftDeletedAt(softDeletedAt string) EdgeDefinitionInterface
}

type EdgeInterface interface {
	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	ID() string
	SetID(id string) EdgeInterface

	WorkflowID() string
	SetWorkflowID(workflowID string) EdgeInterface

	FromStepID() string
	SetFromStepID(fromStepID string) EdgeInterface

	ToStepID() string
	SetToStepID(toStepID string) EdgeInterface

	CreatedAt() string
	SetCreatedAt(createdAt string) EdgeInterface

	UpdatedAt() string
	SetUpdatedAt(updatedAt string) EdgeInterface
}
