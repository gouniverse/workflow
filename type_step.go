package workflow

import "github.com/gouniverse/dataobject"

func NewStep() StepInterface {
	return &step{}
}

type step struct {
	dataobject.DataObject
}

func (s *step) CreatedAt() string {
	return s.Get(COLUMN_CREATED_AT)
}

func (s *step) SetCreatedAt(createdAt string) StepInterface {
	s.Set(COLUMN_CREATED_AT, createdAt)
	return s
}

func (s *step) ID() string {
	return s.Get(COLUMN_ID)
}

func (s *step) SetID(id string) StepInterface {
	s.Set(COLUMN_ID, id)
	return s
}

func (s *step) Name() string {
	return s.Get(COLUMN_NAME)
}

func (s *step) SetName(name string) StepInterface {
	s.Set(COLUMN_NAME, name)
	return s
}

func (s *step) Status() string {
	return s.Get(COLUMN_STATUS)
}

func (s *step) SetStatus(status string) StepInterface {
	s.Set(COLUMN_STATUS, status)
	return s
}

func (s *step) StepDefinitionID() string {
	return s.Get(COLUMN_STEP_DEFINITION_ID)
}

func (s *step) SetStepDefinitionID(stepDefinitionID string) StepInterface {
	s.Set(COLUMN_STEP_DEFINITION_ID, stepDefinitionID)
	return s
}

func (s *step) UpdatedAt() string {
	return s.Get(COLUMN_UPDATED_AT)
}

func (s *step) SetUpdatedAt(updatedAt string) StepInterface {
	s.Set(COLUMN_UPDATED_AT, updatedAt)
	return s
}

func (s *step) WorkflowID() string {
	return s.Get(COLUMN_WORKFLOW_ID)
}

func (s *step) SetWorkflowID(workflowID string) StepInterface {
	s.Set(COLUMN_WORKFLOW_ID, workflowID)
	return s
}
