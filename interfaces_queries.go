package workflow

import "context"

type StoreInterface interface {
	AutoMigrate() error
	EnableDebug(debug bool) StoreInterface

	AutoMigrateEnabled() bool
	DebugEnabled() bool

	// EdgeCount(ctx context.Context,options EdgeQueryInterface) (int64, error)
	// EdgeCreate(ctx context.Context,Edge EdgeInterface) error
	// EdgeDelete(ctx context.Context,Edge EdgeInterface) error
	// EdgeDeleteByID(ctx context.Context,id string) error
	// EdgeFindByID(ctx context.Context,EdgeID string) (EdgeInterface, error)
	// EdgeList(ctx context.Context,query EdgeQueryInterface) ([]EdgeInterface, error)
	// EdgeSoftDelete(ctx context.Context,Edge EdgeInterface) error
	// EdgeSoftDeleteByID(ctx context.Context,id string) error
	// EdgeUpdate(ctx context.Context,Edge EdgeInterface) error

	EdgeDefinitionCount(ctx context.Context, options EdgeDefinitionQueryInterface) (int64, error)
	EdgeDefinitionCreate(ctx context.Context, EdgeDefinition EdgeDefinitionInterface) error
	EdgeDefinitionDelete(ctx context.Context, EdgeDefinition EdgeDefinitionInterface) error
	EdgeDefinitionDeleteByID(ctx context.Context, id string) error
	EdgeDefinitionFindByID(ctx context.Context, EdgeDefinitionID string) (EdgeDefinitionInterface, error)
	EdgeDefinitionList(ctx context.Context, query EdgeDefinitionQueryInterface) ([]EdgeDefinitionInterface, error)
	EdgeDefinitionSoftDelete(ctx context.Context, EdgeDefinition EdgeDefinitionInterface) error
	EdgeDefinitionSoftDeleteByID(ctx context.Context, id string) error
	EdgeDefinitionUpdate(ctx context.Context, EdgeDefinition EdgeDefinitionInterface) error

	// StepCount(ctx context.Context,options StepQueryInterface) (int64, error)
	// StepCreate(ctx context.Context,Step StepInterface) error
	// StepDelete(ctx context.Context,Step StepInterface) error
	// StepDeleteByID(ctx context.Context,id string) error
	// StepFindByID(ctx context.Context,StepID string) (StepInterface, error)
	// StepList(ctx context.Context,query StepQueryInterface) ([]StepInterface, error)
	// StepSoftDelete(ctx context.Context,Step StepInterface) error
	// StepSoftDeleteByID(ctx context.Context,id string) error
	// StepUpdate(ctx context.Context,Step StepInterface) error

	StepDefinitionCount(ctx context.Context, options StepDefinitionQueryInterface) (int64, error)
	StepDefinitionCreate(ctx context.Context, StepDefinition StepDefinitionInterface) error
	StepDefinitionDelete(ctx context.Context, StepDefinition StepDefinitionInterface) error
	StepDefinitionDeleteByID(ctx context.Context, id string) error
	StepDefinitionFindByID(ctx context.Context, StepDefinitionID string) (StepDefinitionInterface, error)
	StepDefinitionList(ctx context.Context, query StepDefinitionQueryInterface) ([]StepDefinitionInterface, error)
	StepDefinitionSoftDelete(ctx context.Context, StepDefinition StepDefinitionInterface) error
	StepDefinitionSoftDeleteByID(ctx context.Context, id string) error
	StepDefinitionUpdate(ctx context.Context, StepDefinition StepDefinitionInterface) error

	// WorkflowCount(ctx context.Context,options WorkflowQueryInterface) (int64, error)
	// WorkflowCreate(ctx context.Context,Workflow WorkflowInterface) error
	// WorkflowDelete(ctx context.Context,Workflow WorkflowInterface) error
	// WorkflowDeleteByID(ctx context.Context,id string) error
	// WorkflowFindByID(ctx context.Context,WorkflowID string) (WorkflowInterface, error)
	// WorkflowList(ctx context.Context,query WorkflowQueryInterface) ([]WorkflowInterface, error)
	// WorkflowSoftDelete(ctx context.Context,Workflow WorkflowInterface) error
	// WorkflowSoftDeleteByID(ctx context.Context,id string) error
	// WorkflowUpdate(ctx context.Context,Workflow WorkflowInterface) error

	WorkflowDefinitionCount(ctx context.Context, options WorkflowDefinitionQueryInterface) (int64, error)
	WorkflowDefinitionCreate(ctx context.Context, WorkflowDefinition WorkflowDefinitionInterface) error
	WorkflowDefinitionDelete(ctx context.Context, WorkflowDefinition WorkflowDefinitionInterface) error
	WorkflowDefinitionDeleteByID(ctx context.Context, id string) error
	WorkflowDefinitionFindByID(ctx context.Context, WorkflowDefinitionID string) (WorkflowDefinitionInterface, error)
	WorkflowDefinitionList(ctx context.Context, query WorkflowDefinitionQueryInterface) ([]WorkflowDefinitionInterface, error)
	WorkflowDefinitionSoftDelete(ctx context.Context, WorkflowDefinition WorkflowDefinitionInterface) error
	WorkflowDefinitionSoftDeleteByID(ctx context.Context, id string) error
	WorkflowDefinitionUpdate(ctx context.Context, workflowDefinition WorkflowDefinitionInterface) error
}

type EdgeQueryInterface interface {
	Validate() error

	IsCountOnly() bool
	SetCountOnly(countOnly bool) EdgeQueryInterface

	Columns() []string
	SetColumns(columns []string) EdgeQueryInterface

	SoftDeletedIncluded() bool
	SetSoftDeletedIncluded(softDeletedIncluded bool) EdgeQueryInterface

	HasCreatedAtGte() bool
	CreatedAtGte() string
	SetCreatedAtGte(createdAtGte string) EdgeQueryInterface

	HasCreatedAtLte() bool
	CreatedAtLte() string
	SetCreatedAtLte(createdAtLte string) EdgeQueryInterface

	HasEdgeDefinitionID() bool
	EdgeDefinitionID() string
	SetEdgeDefinitionID(edgeDefinitionID string) EdgeQueryInterface

	HasID() bool
	ID() string
	SetID(id string) EdgeQueryInterface

	HasIDIn() bool
	IDIn() []string
	SetIDIn(idIn []string) EdgeQueryInterface

	HasFromStepID() bool
	FromStepID() string
	SetFromStepID(fromStepID string) EdgeQueryInterface

	HasToStepID() bool
	ToStepID() string
	SetToStepID(toStepID string) EdgeQueryInterface

	HasUpdatedAtGte() bool
	UpdatedAtGte() string
	SetUpdatedAtGte(updatedAtGte string) EdgeQueryInterface

	HasUpdatedAtLte() bool
	UpdatedAtLte() string
	SetUpdatedAtLte(updatedAtLte string) EdgeQueryInterface

	HasLimit() bool
	Limit() int
	SetLimit(limit int) EdgeQueryInterface

	HasOffset() bool
	Offset() int
	SetOffset(offset int) EdgeQueryInterface

	HasSortOrder() bool
	SortOrder() string
	SetSortOrder(sortOrder string) EdgeQueryInterface

	HasOrderBy() bool
	OrderBy() string
	SetOrderBy(orderBy string) EdgeQueryInterface
}

type EdgeDefinitionQueryInterface interface {
	Validate() error

	IsCountOnly() bool
	SetCountOnly(countOnly bool) EdgeDefinitionQueryInterface

	Columns() []string
	SetColumns(columns []string) EdgeDefinitionQueryInterface

	SoftDeletedIncluded() bool
	SetSoftDeletedIncluded(softDeletedIncluded bool) EdgeDefinitionQueryInterface

	HasCreatedAtGte() bool
	CreatedAtGte() string
	SetCreatedAtGte(createdAtGte string) EdgeDefinitionQueryInterface

	HasCreatedAtLte() bool
	CreatedAtLte() string
	SetCreatedAtLte(createdAtLte string) EdgeDefinitionQueryInterface

	HasID() bool
	ID() string
	SetID(id string) EdgeDefinitionQueryInterface

	HasIDIn() bool
	IDIn() []string
	SetIDIn(idIn []string) EdgeDefinitionQueryInterface

	HasFromStepID() bool
	FromStepID() string
	SetFromStepID(fromStepID string) EdgeDefinitionQueryInterface

	HasToStepID() bool
	ToStepID() string
	SetToStepID(toStepID string) EdgeDefinitionQueryInterface

	HasUpdatedAtGte() bool
	UpdatedAtGte() string
	SetUpdatedAtGte(updatedAtGte string) EdgeDefinitionQueryInterface

	HasUpdatedAtLte() bool
	UpdatedAtLte() string
	SetUpdatedAtLte(updatedAtLte string) EdgeDefinitionQueryInterface

	HasLimit() bool
	Limit() int
	SetLimit(limit int) EdgeDefinitionQueryInterface

	HasOffset() bool
	Offset() int
	SetOffset(offset int) EdgeDefinitionQueryInterface

	HasSortOrder() bool
	SortOrder() string
	SetSortOrder(sortOrder string) EdgeDefinitionQueryInterface

	HasOrderBy() bool
	OrderBy() string
	SetOrderBy(orderBy string) EdgeDefinitionQueryInterface
}

type StepQueryInterface interface {
	Validate() error

	IsCountOnly() bool
	SetCountOnly(countOnly bool) StepQueryInterface

	Columns() []string
	SetColumns(columns []string) StepQueryInterface

	SoftDeletedIncluded() bool
	SetSoftDeletedIncluded(softDeletedIncluded bool) StepQueryInterface

	HasCreatedAtGte() bool
	CreatedAtGte() string
	SetCreatedAtGte(createdAtGte string) StepQueryInterface

	HasCreatedAtLte() bool
	CreatedAtLte() string
	SetCreatedAtLte(createdAtLte string) StepQueryInterface

	HasID() bool
	ID() string
	SetID(id string) StepQueryInterface

	HasIDIn() bool
	IDIn() []string
	SetIDIn(idIn []string) StepQueryInterface

	HasStatus() bool
	Status() string
	SetStatus(status string) StepQueryInterface

	HasStatusIn() bool
	StatusIn() []string
	SetStatusIn(statusIn []string) StepQueryInterface

	HasUpdatedAtGte() bool
	UpdatedAtGte() string
	SetUpdatedAtGte(updatedAtGte string) StepQueryInterface

	HasUpdatedAtLte() bool
	UpdatedAtLte() string
	SetUpdatedAtLte(updatedAtLte string) StepQueryInterface

	HasWorkflowID() bool
	WorkflowID() string
	SetWorkflowID(workflowID string) StepQueryInterface

	HasStepDefinitionID() bool
	StepDefinitionID() string
	SetStepDefinitionID(stepDefinitionID string) StepQueryInterface

	HasLimit() bool
	Limit() int
	SetLimit(limit int) StepQueryInterface

	HasOffset() bool
	Offset() int
	SetOffset(offset int) StepQueryInterface

	HasSortOrder() bool
	SortOrder() string
	SetSortOrder(sortOrder string) StepQueryInterface

	HasOrderBy() bool
	OrderBy() string
	SetOrderBy(orderBy string) StepQueryInterface
}

type StepDefinitionQueryInterface interface {
	Validate() error

	IsCountOnly() bool
	SetCountOnly(countOnly bool) StepDefinitionQueryInterface

	Columns() []string
	SetColumns(columns []string) StepDefinitionQueryInterface

	SoftDeletedIncluded() bool
	SetSoftDeletedIncluded(softDeletedIncluded bool) StepDefinitionQueryInterface

	HasCreatedAtGte() bool
	CreatedAtGte() string
	SetCreatedAtGte(createdAtGte string) StepDefinitionQueryInterface

	HasCreatedAtLte() bool
	CreatedAtLte() string
	SetCreatedAtLte(createdAtLte string) StepDefinitionQueryInterface

	HasID() bool
	ID() string
	SetID(id string) StepDefinitionQueryInterface

	HasIDIn() bool
	IDIn() []string
	SetIDIn(idIn []string) StepDefinitionQueryInterface

	HasStatus() bool
	Status() string
	SetStatus(status string) StepDefinitionQueryInterface

	HasStatusIn() bool
	StatusIn() []string
	SetStatusIn(statusIn []string) StepDefinitionQueryInterface

	HasUpdatedAtGte() bool
	UpdatedAtGte() string
	SetUpdatedAtGte(updatedAtGte string) StepDefinitionQueryInterface

	HasUpdatedAtLte() bool
	UpdatedAtLte() string
	SetUpdatedAtLte(updatedAtLte string) StepDefinitionQueryInterface

	HasWorkflowDefinitionID() bool
	WorkflowDefinitionID() string
	SetWorkflowDefinitionID(workflowDefinitionID string) StepDefinitionQueryInterface

	HasLimit() bool
	Limit() int
	SetLimit(limit int) StepDefinitionQueryInterface

	HasOffset() bool
	Offset() int
	SetOffset(offset int) StepDefinitionQueryInterface

	HasOrderBy() bool
	OrderBy() string
	SetOrderBy(orderBy string) StepDefinitionQueryInterface

	HasSortOrder() bool
	SortOrder() string
	SetSortOrder(sortOrder string) StepDefinitionQueryInterface
}

type WorkflowQueryInterface interface {
	Validate() error

	IsCountOnly() bool
	SetCountOnly(countOnly bool) WorkflowQueryInterface

	Columns() []string
	SetColumns(columns []string) WorkflowQueryInterface

	SoftDeletedIncluded() bool
	SetSoftDeletedIncluded(softDeletedIncluded bool) WorkflowQueryInterface

	HasCreatedAtGte() bool
	CreatedAtGte() string
	SetCreatedAtGte(createdAtGte string) WorkflowQueryInterface

	HasCreatedAtLte() bool
	CreatedAtLte() string
	SetCreatedAtLte(createdAtLte string) WorkflowQueryInterface

	HasID() bool
	ID() string
	SetID(id string) WorkflowQueryInterface

	HasIDIn() bool
	IDIn() []string
	SetIDIn(idIn []string) WorkflowQueryInterface

	HasStatus() bool
	Status() string
	SetStatus(status string) WorkflowQueryInterface

	HasStatusIn() bool
	StatusIn() []string
	SetStatusIn(statusIn []string) WorkflowQueryInterface

	HasUpdatedAtGte() bool
	UpdatedAtGte() string
	SetUpdatedAtGte(updatedAtGte string) WorkflowQueryInterface

	HasUpdatedAtLte() bool
	UpdatedAtLte() string
	SetUpdatedAtLte(updatedAtLte string) WorkflowQueryInterface

	HasWorkflowDefinitionID() bool
	WorkflowDefinitionID() string
	SetWorkflowDefinitionID(workflowDefinitionID string) WorkflowQueryInterface

	HasLimit() bool
	Limit() int
	SetLimit(limit int) WorkflowQueryInterface

	HasOffset() bool
	Offset() int
	SetOffset(offset int) WorkflowQueryInterface

	HasSortOrder() bool
	SortOrder() string
	SetSortOrder(sortOrder string) WorkflowQueryInterface

	HasOrderBy() bool
	OrderBy() string
	SetOrderBy(orderBy string) WorkflowQueryInterface
}

type WorkflowDefinitionQueryInterface interface {
	Validate() error

	IsCountOnly() bool
	SetCountOnly(countOnly bool) WorkflowDefinitionQueryInterface

	Columns() []string
	SetColumns(columns []string) WorkflowDefinitionQueryInterface

	SoftDeletedIncluded() bool
	SetSoftDeletedIncluded(softDeletedIncluded bool) WorkflowDefinitionQueryInterface

	HasCreatedAtGte() bool
	CreatedAtGte() string
	SetCreatedAtGte(createdAtGte string) WorkflowDefinitionQueryInterface

	HasCreatedAtLte() bool
	CreatedAtLte() string
	SetCreatedAtLte(createdAtLte string) WorkflowDefinitionQueryInterface

	HasID() bool
	ID() string
	SetID(id string) WorkflowDefinitionQueryInterface

	HasIDIn() bool
	IDIn() []string
	SetIDIn(idIn []string) WorkflowDefinitionQueryInterface

	HasStatus() bool
	Status() string
	SetStatus(status string) WorkflowDefinitionQueryInterface

	HasStatusIn() bool
	StatusIn() []string
	SetStatusIn(statusIn []string) WorkflowDefinitionQueryInterface

	HasUpdatedAtGte() bool
	UpdatedAtGte() string
	SetUpdatedAtGte(updatedAtGte string) WorkflowDefinitionQueryInterface

	HasUpdatedAtLte() bool
	UpdatedAtLte() string
	SetUpdatedAtLte(updatedAtLte string) WorkflowDefinitionQueryInterface

	HasLimit() bool
	Limit() int
	SetLimit(limit int) WorkflowDefinitionQueryInterface

	HasOffset() bool
	Offset() int
	SetOffset(offset int) WorkflowDefinitionQueryInterface

	HasSortOrder() bool
	SortOrder() string
	SetSortOrder(sortOrder string) WorkflowDefinitionQueryInterface

	HasOrderBy() bool
	OrderBy() string
	SetOrderBy(orderBy string) WorkflowDefinitionQueryInterface
}
