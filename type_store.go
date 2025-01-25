package workflow

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"log/slog"

	"github.com/gouniverse/base/database"
	"github.com/gouniverse/sb"
)

type NewStoreOptions struct {
	WorkflowTableName           string
	WorkflowDefinitionTableName string
	StepTableName               string
	StepDefinitionTableName     string
	EdgeTableName               string
	EdgeDefinitionTableName     string
	AutomigrateEnabled          bool
	DB                          *sql.DB
	DbDriverName                string
	DebugEnabled                bool
	SqlLogger                   *slog.Logger
}

func NewStore(opts NewStoreOptions) (StoreInterface, error) {
	store := &store{
		workflowTableName:           opts.WorkflowTableName,
		workflowDefinitionTableName: opts.WorkflowDefinitionTableName,
		stepTableName:               opts.StepTableName,
		stepDefinitionTableName:     opts.StepDefinitionTableName,
		edgeTableName:               opts.EdgeTableName,
		edgeDefinitionTableName:     opts.EdgeDefinitionTableName,
		automigrateEnabled:          opts.AutomigrateEnabled,
		db:                          opts.DB,
		dbDriverName:                opts.DbDriverName,
		debugEnabled:                opts.DebugEnabled,
		sqlLogger:                   opts.SqlLogger,
	}

	if store.workflowTableName == "" {
		return nil, errors.New("workflow store: WorkflowTableName is required")
	}

	if store.workflowDefinitionTableName == "" {
		return nil, errors.New("workflow store: WorkflowDefiniationTableName is required")
	}

	if store.stepTableName == "" {
		return nil, errors.New("workflow store: StepTableName is required")
	}

	if store.stepDefinitionTableName == "" {
		return nil, errors.New("workflow store: StepDefiniationTableName is required")
	}

	if store.edgeTableName == "" {
		return nil, errors.New("workflow store: EdgeTableName is required")
	}

	if store.edgeDefinitionTableName == "" {
		return nil, errors.New("workflow store: EdgeDefiniationTableName is required")
	}

	if store.db == nil {
		return nil, errors.New("workflow store: DB is required")
	}

	if store.dbDriverName == "" {
		store.dbDriverName = database.DatabaseType(store.db)
	}

	if store.automigrateEnabled {
		store.AutoMigrate()
	}

	return store, nil
}

type store struct {
	workflowTableName           string
	workflowDefinitionTableName string
	stepTableName               string
	stepDefinitionTableName     string
	edgeTableName               string
	edgeDefinitionTableName     string
	db                          *sql.DB
	dbDriverName                string
	automigrateEnabled          bool
	debugEnabled                bool
	sqlLogger                   *slog.Logger
}

// AutoMigrate migrates the tables
func (store *store) AutoMigrate() error {
	sqls := []string{
		store.sqlCreateWorkflowDefinitionTable(),
		store.sqlCreateStepDefinitionTable(),
		store.sqlCreateEdgeDefinitionTable(),
		store.sqlCreateWorkflowTable(),
		store.sqlCreateStepTable(),
		store.sqlCreateEdgeTable(),
	}

	for _, sql := range sqls {
		store.logSql("migrate", sql)
		_, err := store.db.Exec(sql)

		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

// AutoMigrateEnabled returns if the automigrate is enabled
func (store *store) AutoMigrateEnabled() bool {
	return store.automigrateEnabled
}

// EnableDebug - enables the debug option
func (st *store) EnableDebug(debugEnabled bool) StoreInterface {
	st.debugEnabled = debugEnabled
	return st
}

// DebugEnabled returns if the debug is enabled
func (store *store) DebugEnabled() bool {
	return store.debugEnabled
}

// logSql logs sql to the sql logger
func (store *store) logSql(sqlOperationType string, sql string, params ...interface{}) {
	if !store.debugEnabled {
		return
	}

	if store.sqlLogger != nil {
		store.sqlLogger.Debug("sql: "+sqlOperationType, slog.String("sql", sql), slog.Any("params", params))
	}
}

func (store *store) sqlCreateWorkflowTable() string {
	sql := sb.NewBuilder(sb.DatabaseDriverName(store.db)).
		Table(store.workflowTableName).
		Column(sb.Column{
			Name:       COLUMN_ID,
			Type:       sb.COLUMN_TYPE_STRING,
			Length:     40,
			PrimaryKey: true,
		}).
		Column(sb.Column{
			Name:   COLUMN_NAME,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 255,
		}).
		Column(sb.Column{
			Name: COLUMN_CREATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: COLUMN_UPDATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).CreateIfNotExists()

	return sql
}

func (store *store) sqlCreateWorkflowDefinitionTable() string {
	sql := sb.NewBuilder(sb.DatabaseDriverName(store.db)).
		Table(store.workflowDefinitionTableName).
		Column(sb.Column{
			Name:       COLUMN_ID,
			Type:       sb.COLUMN_TYPE_STRING,
			Length:     40,
			PrimaryKey: true,
		}).
		Column(sb.Column{
			Name:   COLUMN_NAME,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 255,
		}).
		Column(sb.Column{
			Name: COLUMN_CREATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: COLUMN_UPDATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: COLUMN_SOFT_DELETED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		CreateIfNotExists()

	return sql
}

func (store *store) sqlCreateStepTable() string {
	sql := sb.NewBuilder(sb.DatabaseDriverName(store.db)).
		Table(store.stepTableName).
		Column(sb.Column{
			Name:       COLUMN_ID,
			Type:       sb.COLUMN_TYPE_STRING,
			Length:     40,
			PrimaryKey: true,
		}).
		Column(sb.Column{
			Name: COLUMN_WORKFLOW_ID,
			Type: sb.COLUMN_TYPE_STRING,
		}).
		Column(sb.Column{
			Name:   COLUMN_NAME,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 255,
		}).
		Column(sb.Column{
			Name: COLUMN_CREATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: COLUMN_UPDATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		CreateIfNotExists()

	return sql
}

func (store *store) sqlCreateStepDefinitionTable() string {
	sql := sb.NewBuilder(sb.DatabaseDriverName(store.db)).
		Table(store.stepDefinitionTableName).
		Column(sb.Column{
			Name:       COLUMN_ID,
			Type:       sb.COLUMN_TYPE_STRING,
			Length:     40,
			PrimaryKey: true,
		}).
		Column(sb.Column{
			Name:   COLUMN_WORKFLOW_DEFINITION_ID,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 40,
		}).
		Column(sb.Column{
			Name:   COLUMN_STATUS,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 20,
		}).
		Column(sb.Column{
			Name:   COLUMN_NAME,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 255,
		}).
		Column(sb.Column{
			Name: COLUMN_CREATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: COLUMN_UPDATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: COLUMN_SOFT_DELETED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		CreateIfNotExists()

	return sql
}

func (store *store) sqlCreateEdgeTable() string {
	sql := sb.NewBuilder(sb.DatabaseDriverName(store.db)).
		Table(store.edgeTableName).
		Column(sb.Column{
			Name:       COLUMN_ID,
			Type:       sb.COLUMN_TYPE_STRING,
			Length:     40,
			PrimaryKey: true,
		}).
		Column(sb.Column{
			Name: COLUMN_WORKFLOW_ID,
			Type: sb.COLUMN_TYPE_STRING,
		}).
		Column(sb.Column{
			Name: COLUMN_FROM_STEP_ID,
			Type: sb.COLUMN_TYPE_STRING,
		}).
		Column(sb.Column{
			Name: COLUMN_TO_STEP_ID,
			Type: sb.COLUMN_TYPE_STRING,
		}).
		Column(sb.Column{
			Name: COLUMN_CREATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: COLUMN_UPDATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		CreateIfNotExists()

	return sql
}

func (store *store) sqlCreateEdgeDefinitionTable() string {
	sql := sb.NewBuilder(sb.DatabaseDriverName(store.db)).
		Table(store.edgeDefinitionTableName).
		Column(sb.Column{
			Name:       COLUMN_ID,
			Type:       sb.COLUMN_TYPE_STRING,
			Length:     40,
			PrimaryKey: true,
		}).
		Column(sb.Column{
			Name:   COLUMN_FROM_STEP_DEFINITION_ID,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 40,
		}).
		Column(sb.Column{
			Name:   COLUMN_TO_STEP_DEFINITION_ID,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 40,
		}).
		Column(sb.Column{
			Name: COLUMN_CREATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: COLUMN_UPDATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: COLUMN_SOFT_DELETED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		CreateIfNotExists()

	return sql
}

func (store *store) toQuerableContext(context context.Context) database.QueryableContext {
	if database.IsQueryableContext(context) {
		return context.(database.QueryableContext)
	}

	return database.Context(context, store.db)
}
