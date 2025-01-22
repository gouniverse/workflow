package workflow

import "github.com/gouniverse/dataobject"

func NewEdge() EdgeInterface {
	return &edge{}
}

type edge struct {
	dataobject.DataObject
}

func (e *edge) CreatedAt() string {
	return e.Get(COLUMN_CREATED_AT)
}

func (e *edge) SetCreatedAt(createdAt string) EdgeInterface {
	e.Set(COLUMN_CREATED_AT, createdAt)
	return e
}

func (e *edge) FromStepID() string {
	return e.Get(COLUMN_FROM_STEP_ID)
}

func (e *edge) SetFromStepID(fromStepID string) EdgeInterface {
	e.Set(COLUMN_FROM_STEP_ID, fromStepID)
	return e
}

func (e *edge) ID() string {
	return e.Get(COLUMN_ID)
}

func (e *edge) SetID(id string) EdgeInterface {
	e.Set(COLUMN_ID, id)
	return e
}

func (e *edge) Name() string {
	return e.Get(COLUMN_NAME)
}

func (e *edge) SetName(name string) EdgeInterface {
	e.Set(COLUMN_NAME, name)
	return e
}

func (e *edge) ToStepID() string {
	return e.Get(COLUMN_TO_STEP_ID)
}

func (e *edge) SetToStepID(toStepID string) EdgeInterface {
	e.Set(COLUMN_TO_STEP_ID, toStepID)
	return e
}

func (e *edge) UpdatedAt() string {
	return e.Get(COLUMN_UPDATED_AT)
}

func (e *edge) SetUpdatedAt(updatedAt string) EdgeInterface {
	e.Set(COLUMN_UPDATED_AT, updatedAt)
	return e
}

func (e *edge) WorkflowID() string {
	return e.Get(COLUMN_WORKFLOW_ID)
}

func (e *edge) SetWorkflowID(workflowID string) EdgeInterface {
	e.Set(COLUMN_WORKFLOW_ID, workflowID)
	return e
}
