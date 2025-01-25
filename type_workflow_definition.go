package workflow

import (
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/uid"
)

// == CONSTRUCTORS =============================================================

func NewWorkflowDefinition() WorkflowDefinitionInterface {
	wd := &workflowDefinition{}

	wd.SetID(uid.HumanUid()).
		SetName("").
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetSoftDeletedAt(sb.MAX_DATETIME)

	return wd
}

func NewWorkflowDefinitionFromExistingData(data map[string]string) WorkflowDefinitionInterface {
	o := &workflowDefinition{}
	o.Hydrate(data)
	return o
}

// == CLASS ====================================================================

type workflowDefinition struct {
	dataobject.DataObject
}

// == SETTERS & GETTERS ========================================================

func (wf *workflowDefinition) CreatedAt() string {
	return wf.Get(COLUMN_CREATED_AT)
}

func (wf *workflowDefinition) SetCreatedAt(createdAt string) WorkflowDefinitionInterface {
	wf.Set(COLUMN_CREATED_AT, createdAt)
	return wf
}

func (wf *workflowDefinition) ID() string {
	return wf.Get(COLUMN_ID)
}

func (wf *workflowDefinition) SetID(id string) WorkflowDefinitionInterface {
	wf.Set(COLUMN_ID, id)
	return wf
}

func (wf *workflowDefinition) Name() string {
	return wf.Get(COLUMN_NAME)
}

func (wf *workflowDefinition) SetName(name string) WorkflowDefinitionInterface {
	wf.Set(COLUMN_NAME, name)
	return wf
}

func (wf *workflowDefinition) SoftDeletedAt() string {
	return wf.Get(COLUMN_SOFT_DELETED_AT)
}

func (wf *workflowDefinition) SetSoftDeletedAt(softDeletedAt string) WorkflowDefinitionInterface {
	wf.Set(COLUMN_SOFT_DELETED_AT, softDeletedAt)
	return wf
}

func (wf *workflowDefinition) UpdatedAt() string {
	return wf.Get(COLUMN_UPDATED_AT)
}

func (wf *workflowDefinition) SetUpdatedAt(updatedAt string) WorkflowDefinitionInterface {
	wf.Set(COLUMN_UPDATED_AT, updatedAt)
	return wf
}
