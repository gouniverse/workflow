package workflow

import "context"

type StoreInterface interface {
	AutoMigrate() error
	EnableDebug(debug bool) StoreInterface

	AutoMigrateEnabled() bool
	DebugEnabled() bool

	// EdgeCount(options EdgeQueryInterface) (int64, error)
	// EdgeCreate(Edge EdgeInterface) error
	// EdgeDelete(Edge EdgeInterface) error
	// EdgeDeleteByID(id string) error
	// EdgeFindByID(EdgeID string) (EdgeInterface, error)
	// EdgeList(query EdgeQueryInterface) ([]EdgeInterface, error)
	// EdgeSoftDelete(Edge EdgeInterface) error
	// EdgeSoftDeleteByID(id string) error
	// EdgeUpdate(Edge EdgeInterface) error

	// EdgeDefinitionCount(options EdgeDefinitionQueryInterface) (int64, error)
	// EdgeDefinitionCreate(EdgeDefinition EdgeDefinitionInterface) error
	// EdgeDefinitionDelete(EdgeDefinition EdgeDefinitionInterface) error
	// EdgeDefinitionDeleteByID(id string) error
	// EdgeDefinitionFindByID(EdgeDefinitionID string) (EdgeDefinitionInterface, error)
	// EdgeDefinitionList(query EdgeDefinitionQueryInterface) ([]EdgeDefinitionInterface, error)
	// EdgeDefinitionSoftDelete(EdgeDefinition EdgeDefinitionInterface) error
	// EdgeDefinitionSoftDeleteByID(id string) error
	// EdgeDefinitionUpdate(EdgeDefinition EdgeDefinitionInterface) error

	// StepCount(options StepQueryInterface) (int64, error)
	// StepCreate(Step StepInterface) error
	// StepDelete(Step StepInterface) error
	// StepDeleteByID(id string) error
	// StepFindByID(StepID string) (StepInterface, error)
	// StepList(query StepQueryInterface) ([]StepInterface, error)
	// StepSoftDelete(Step StepInterface) error
	// StepSoftDeleteByID(id string) error
	// StepUpdate(Step StepInterface) error

	// StepDefinitionCount(options StepDefinitionQueryInterface) (int64, error)
	// StepDefinitionCreate(StepDefinition StepDefinitionInterface) error
	// StepDefinitionDelete(StepDefinition StepDefinitionInterface) error
	// StepDefinitionDeleteByID(id string) error
	// StepDefinitionFindByID(StepDefinitionID string) (StepDefinitionInterface, error)
	// StepDefinitionList(query StepDefinitionQueryInterface) ([]StepDefinitionInterface, error)
	// StepDefinitionSoftDelete(StepDefinition StepDefinitionInterface) error
	// StepDefinitionSoftDeleteByID(id string) error
	// StepDefinitionUpdate(StepDefinition StepDefinitionInterface) error

	// WorkflowCount(options WorkflowQueryInterface) (int64, error)
	// WorkflowCreate(Workflow WorkflowInterface) error
	// WorkflowDelete(Workflow WorkflowInterface) error
	// WorkflowDeleteByID(id string) error
	// WorkflowFindByID(WorkflowID string) (WorkflowInterface, error)
	// WorkflowList(query WorkflowQueryInterface) ([]WorkflowInterface, error)
	// WorkflowSoftDelete(Workflow WorkflowInterface) error
	// WorkflowSoftDeleteByID(id string) error
	// WorkflowUpdate(Workflow WorkflowInterface) error

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
