package workflow

import "github.com/gouniverse/dataobject"

func NewEdgeDefinition() EdgeDefinitionInterface {
	return &edgeDefinition{}
}

type edgeDefinition struct {
	dataobject.DataObject
}

func (ed *edgeDefinition) CreatedAt() string {
	return ed.Get(COLUMN_CREATED_AT)
}

func (ed *edgeDefinition) SetCreatedAt(createdAt string) EdgeDefinitionInterface {
	ed.Set(COLUMN_CREATED_AT, createdAt)
	return ed
}

func (ed *edgeDefinition) ID() string {
	return ed.Get(COLUMN_ID)
}

func (ed *edgeDefinition) SetID(id string) EdgeDefinitionInterface {
	ed.Set(COLUMN_ID, id)
	return ed
}

func (ed *edgeDefinition) WorkflowDefinitionID() string {
	return ed.Get(COLUMN_WORKFLOW_DEFINITION_ID)
}

func (ed *edgeDefinition) SetWorkflowDefinitionID(workflowDefinitionID string) EdgeDefinitionInterface {
	ed.Set(COLUMN_WORKFLOW_DEFINITION_ID, workflowDefinitionID)
	return ed
}

func (ed *edgeDefinition) FromStepDefinitionID() string {
	return ed.Get(COLUMN_FROM_STEP_DEFINITION_ID)
}

func (ed *edgeDefinition) SetFromStepDefinitionID(fromStepDefinitionID string) EdgeDefinitionInterface {
	ed.Set(COLUMN_FROM_STEP_DEFINITION_ID, fromStepDefinitionID)
	return ed
}

func (ed *edgeDefinition) ToStepDefinitionID() string {
	return ed.Get(COLUMN_TO_STEP_DEFINITION_ID)
}

func (ed *edgeDefinition) SetToStepDefinitionID(toStepDefinitionID string) EdgeDefinitionInterface {
	ed.Set(COLUMN_TO_STEP_DEFINITION_ID, toStepDefinitionID)
	return ed
}

func (ed *edgeDefinition) UpdatedAt() string {
	return ed.Get(COLUMN_UPDATED_AT)
}

func (ed *edgeDefinition) SetUpdatedAt(updatedAt string) EdgeDefinitionInterface {
	ed.Set(COLUMN_UPDATED_AT, updatedAt)
	return ed
}
