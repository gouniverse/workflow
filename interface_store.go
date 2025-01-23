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
	HasID() bool
	ID() string
	SetID(id string) EdgeQueryInterface
}

type EdgeDefinitionQueryInterface interface {
	HasID() bool
	ID() string
	SetID(id string) EdgeDefinitionQueryInterface
}

type StepQueryInterface interface {
	HasID() bool
	ID() string
	SetID(id string) StepQueryInterface
}

type StepDefinitionQueryInterface interface {
	HasID() bool
	ID() string
	SetID(id string) StepDefinitionQueryInterface
}

type WorkflowQueryInterface interface {
	HasID() bool
	ID() string
	SetID(id string) WorkflowQueryInterface
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
