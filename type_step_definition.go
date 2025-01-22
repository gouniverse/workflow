package workflow

import "github.com/gouniverse/dataobject"

func NewStepDefinition() StepDefinitionInterface {
	return &stepDefinition{}
}

type stepDefinition struct {
	dataobject.DataObject
}

func (sd *stepDefinition) CreatedAt() string {
	return sd.Get(COLUMN_CREATED_AT)
}

func (sd *stepDefinition) SetCreatedAt(createdAt string) StepDefinitionInterface {
	sd.Set(COLUMN_CREATED_AT, createdAt)
	return sd
}

func (sd *stepDefinition) ID() string {
	return sd.Get(COLUMN_ID)
}

func (sd *stepDefinition) SetID(id string) StepDefinitionInterface {
	sd.Set(COLUMN_ID, id)
	return sd
}

func (sd *stepDefinition) WorkflowDefinitionID() string {
	return sd.Get(COLUMN_WORKFLOW_DEFINITION_ID)
}

func (sd *stepDefinition) SetWorkflowDefinitionID(workflowDefinitionID string) StepDefinitionInterface {
	sd.Set(COLUMN_WORKFLOW_DEFINITION_ID, workflowDefinitionID)
	return sd
}

func (sd *stepDefinition) Name() string {
	return sd.Get(COLUMN_NAME)
}

func (sd *stepDefinition) SetName(name string) StepDefinitionInterface {
	sd.Set(COLUMN_NAME, name)
	return sd
}

func (sd *stepDefinition) UpdatedAt() string {
	return sd.Get(COLUMN_UPDATED_AT)
}

func (sd *stepDefinition) SetUpdatedAt(updatedAt string) StepDefinitionInterface {
	sd.Set(COLUMN_UPDATED_AT, updatedAt)
	return sd
}
