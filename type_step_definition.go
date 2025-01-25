package workflow

import (
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/uid"
)

// == CONSTRUCTORS =============================================================

func NewStepDefinition() StepDefinitionInterface {
	o := &stepDefinition{}

	o.SetID(uid.HumanUid()).
		// SetWorkflowDefinitionID("").
		SetStatus(STEP_DEFINITION_STATUS_DRAFT).
		// SetName("").
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetSoftDeletedAt(sb.MAX_DATETIME)

	return o
}

func NewStepDefinitionFromExistingData(data map[string]string) StepDefinitionInterface {
	o := &stepDefinition{}
	o.Hydrate(data)
	return o
}

// == TYPE ====================================================================

type stepDefinition struct {
	dataobject.DataObject
}

// == SETTERS & GETTERS ========================================================

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

func (sd *stepDefinition) SoftDeletedAt() string {
	return sd.Get(COLUMN_SOFT_DELETED_AT)
}

func (sd *stepDefinition) SetSoftDeletedAt(softDeletedAt string) StepDefinitionInterface {
	sd.Set(COLUMN_SOFT_DELETED_AT, softDeletedAt)
	return sd
}

func (sd *stepDefinition) Status() string {
	return sd.Get(COLUMN_STATUS)
}

func (sd *stepDefinition) SetStatus(status string) StepDefinitionInterface {
	sd.Set(COLUMN_STATUS, status)
	return sd
}

func (sd *stepDefinition) UpdatedAt() string {
	return sd.Get(COLUMN_UPDATED_AT)
}

func (sd *stepDefinition) SetUpdatedAt(updatedAt string) StepDefinitionInterface {
	sd.Set(COLUMN_UPDATED_AT, updatedAt)
	return sd
}
