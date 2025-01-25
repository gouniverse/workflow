package workflow

import (
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/uid"
)

// == CONSTRUCTORS =============================================================

func NewEdgeDefinition() EdgeDefinitionInterface {
	o := &edgeDefinition{}

	o.SetID(uid.HumanUid()).
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetSoftDeletedAt(sb.MAX_DATETIME)

	return o
}

func NewEdgeDefinitionFromExistingData(data map[string]string) EdgeDefinitionInterface {
	o := &edgeDefinition{}
	o.Hydrate(data)
	return o
}

// == TYPE ====================================================================

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

func (ed *edgeDefinition) SoftDeletedAt() string {
	return ed.Get(COLUMN_SOFT_DELETED_AT)
}

func (ed *edgeDefinition) SetSoftDeletedAt(softDeletedAt string) EdgeDefinitionInterface {
	ed.Set(COLUMN_SOFT_DELETED_AT, softDeletedAt)
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
