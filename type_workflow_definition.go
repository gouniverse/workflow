package workflow

import "github.com/gouniverse/dataobject"

func NewWorkflowDefinition() WorkflowDefinitionInterface {
	return &workflowDefinition{}
}

type workflowDefinition struct {
	dataobject.DataObject
}

func (wf *workflowDefinition) CreatedAt() string {
	return wf.Get("created_at")
}

func (wf *workflowDefinition) SetCreatedAt(createdAt string) WorkflowDefinitionInterface {
	wf.Set("created_at", createdAt)
	return wf
}

func (wf *workflowDefinition) ID() string {
	return wf.Get("id")
}

func (wf *workflowDefinition) SetID(id string) WorkflowDefinitionInterface {
	wf.Set("id", id)
	return wf
}

func (wf *workflowDefinition) Name() string {
	return wf.Get("name")
}

func (wf *workflowDefinition) SetName(name string) WorkflowDefinitionInterface {
	wf.Set("name", name)
	return wf
}

func (wf *workflowDefinition) UpdatedAt() string {
	return wf.Get("updated_at")
}

func (wf *workflowDefinition) SetUpdatedAt(updatedAt string) WorkflowDefinitionInterface {
	wf.Set("updated_at", updatedAt)
	return wf
}
