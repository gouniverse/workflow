package workflow

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/base/database"
	"github.com/gouniverse/sb"
	"github.com/samber/lo"
)

func (store *store) WorkflowDefinitionCount(ctx context.Context, options WorkflowDefinitionQueryInterface) (int64, error) {
	options.SetCountOnly(true)
	q, _, err := store.workflowDefinitionSelectQuery(options)

	sqlStr, sqlParams, errSql := q.Limit(1).
		Select(goqu.COUNT(goqu.Star()).As("count")).
		Prepared(true).
		ToSQL()

	if errSql != nil {
		return -1, nil
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	mapped, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return -1, err
	}

	if len(mapped) < 1 {
		return -1, nil
	}

	countStr := mapped[0]["count"]

	i, err := strconv.ParseInt(countStr, 10, 64)

	if err != nil {
		return -1, err

	}

	return i, nil
}

func (store *store) WorkflowDefinitionCreate(ctx context.Context, workflowDefinition WorkflowDefinitionInterface) error {
	workflowDefinition.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	workflowDefinition.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	data := workflowDefinition.Data()

	sqlStr, sqlParams, errSql := goqu.Dialect(store.dbDriverName).
		Insert(store.workflowDefinitionTableName).
		Prepared(true).
		Rows(data).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("create", sqlStr, sqlParams...)

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return err
	}

	return nil
}

func (store *store) WorkflowDefinitionDelete(ctx context.Context, workflowDefinition WorkflowDefinitionInterface) error {
	if workflowDefinition == nil {
		return errors.New("workflow definition is nil")
	}

	return store.WorkflowDefinitionDeleteByID(ctx, workflowDefinition.ID())
}

func (store *store) WorkflowDefinitionDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("queue id is empty")
	}

	sqlStr, preparedArgs, err := goqu.Dialect(store.dbDriverName).
		From(store.workflowDefinitionTableName).
		Prepared(true).
		Where(goqu.C(COLUMN_ID).Eq(id)).
		Delete().
		ToSQL()

	if err != nil {
		return err
	}

	store.logSql("delete", sqlStr, preparedArgs...)

	_, err = store.db.Exec(sqlStr, preparedArgs...)

	return err
}

func (store *store) WorkflowDefinitionFindByID(ctx context.Context, id string) (WorkflowDefinitionInterface, error) {
	if id == "" {
		return nil, errors.New("workflow definition id is empty")
	}

	query := WorkflowDefinitionQuery().SetID(id).SetLimit(1)

	list, err := store.WorkflowDefinitionList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) WorkflowDefinitionList(ctx context.Context, query WorkflowDefinitionQueryInterface) ([]WorkflowDefinitionInterface, error) {
	q, columns, err := store.workflowDefinitionSelectQuery(query)

	if err != nil {
		return []WorkflowDefinitionInterface{}, err
	}

	sqlStr, sqlParams, errSql := q.Prepared(true).Select(columns...).ToSQL()

	store.logSql("list", sqlStr, sqlParams...)

	if errSql != nil {
		return []WorkflowDefinitionInterface{}, errSql
	}

	db := sb.NewDatabase(store.db, store.dbDriverName)

	if db == nil {
		return []WorkflowDefinitionInterface{}, errors.New("workflow store: database is nil")
	}

	modelMaps, err := db.SelectToMapString(sqlStr, sqlParams...)

	if err != nil {
		return []WorkflowDefinitionInterface{}, err
	}

	list := []WorkflowDefinitionInterface{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewWorkflowDefinitionFromExistingData(modelMap)
		list = append(list, model)
	})

	return list, nil
}

func (store *store) WorkflowDefinitionSoftDelete(ctx context.Context, workflowDefinition WorkflowDefinitionInterface) error {
	if workflowDefinition == nil {
		return errors.New("queue is nil")
	}

	workflowDefinition.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	return store.WorkflowDefinitionUpdate(ctx, workflowDefinition)
}

func (store *store) WorkflowDefinitionSoftDeleteByID(ctx context.Context, id string) error {
	workflowDefinition, err := store.WorkflowDefinitionFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.WorkflowDefinitionSoftDelete(ctx, workflowDefinition)
}

func (store *store) WorkflowDefinitionUpdate(ctx context.Context, workflowDefinition WorkflowDefinitionInterface) error {
	workflowDefinition.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	dataChanged := workflowDefinition.DataChanged()

	delete(dataChanged, COLUMN_ID) // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Update(store.workflowDefinitionTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C(COLUMN_ID).Eq(workflowDefinition.ID())).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	if store.db == nil {
		return errors.New("workflow store: database is nil")
	}

	_, err := store.db.Exec(sqlStr, params...)

	workflowDefinition.MarkAsNotDirty()

	return err
}

func (store *store) workflowDefinitionSelectQuery(options WorkflowDefinitionQueryInterface) (selectDataset *goqu.SelectDataset, columns []any, err error) {
	if options == nil {
		return nil, []any{}, errors.New("site options cannot be nil")
	}

	if err := options.Validate(); err != nil {
		return nil, []any{}, err
	}

	q := goqu.Dialect(store.dbDriverName).From(store.workflowDefinitionTableName)

	if options.HasCreatedAtGte() && options.HasCreatedAtLte() {
		q = q.Where(
			goqu.C(COLUMN_CREATED_AT).Gte(options.CreatedAtGte()),
			goqu.C(COLUMN_CREATED_AT).Lte(options.CreatedAtLte()),
		)
	} else if options.HasCreatedAtGte() {
		q = q.Where(goqu.C(COLUMN_CREATED_AT).Gte(options.CreatedAtGte()))
	} else if options.HasCreatedAtLte() {
		q = q.Where(goqu.C(COLUMN_CREATED_AT).Lte(options.CreatedAtLte()))
	}

	if options.HasID() {
		q = q.Where(goqu.C(COLUMN_ID).Eq(options.ID()))
	}

	if options.HasIDIn() {
		q = q.Where(goqu.C(COLUMN_ID).In(options.IDIn()))
	}

	if options.HasStatus() {
		q = q.Where(goqu.C(COLUMN_STATUS).Eq(options.Status()))
	}

	if options.HasStatusIn() {
		q = q.Where(goqu.C(COLUMN_STATUS).In(options.StatusIn()))
	}

	if !options.IsCountOnly() {
		if options.HasLimit() {
			q = q.Limit(uint(options.Limit()))
		}

		if options.HasOffset() {
			q = q.Offset(uint(options.Offset()))
		}
	}

	sortOrder := sb.DESC
	if options.HasSortOrder() {
		sortOrder = options.SortOrder()
	}

	if options.HasOrderBy() {
		if strings.EqualFold(sortOrder, sb.ASC) {
			q = q.Order(goqu.I(options.OrderBy()).Asc())
		} else {
			q = q.Order(goqu.I(options.OrderBy()).Desc())
		}
	}

	columns = []any{}

	for _, column := range options.Columns() {
		columns = append(columns, column)
	}

	if options.SoftDeletedIncluded() {
		return q, columns, nil // soft deleted sites requested specifically
	}

	softDeleted := goqu.C(COLUMN_SOFT_DELETED_AT).
		Gt(carbon.Now(carbon.UTC).ToDateTimeString())

	return q.Where(softDeleted), columns, nil
}
