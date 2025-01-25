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

func (store *store) StepDefinitionCount(ctx context.Context, options StepDefinitionQueryInterface) (int64, error) {
	options.SetCountOnly(true)
	q, _, err := store.stepDefinitionSelectQuery(options)

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

func (store *store) StepDefinitionCreate(ctx context.Context, stepDefinition StepDefinitionInterface) error {
	stepDefinition.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	stepDefinition.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	data := stepDefinition.Data()

	sqlStr, sqlParams, errSql := goqu.Dialect(store.dbDriverName).
		Insert(store.stepDefinitionTableName).
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

func (store *store) StepDefinitionDelete(ctx context.Context, stepDefinition StepDefinitionInterface) error {
	if stepDefinition == nil {
		return errors.New("workflow definition is nil")
	}

	return store.StepDefinitionDeleteByID(ctx, stepDefinition.ID())
}

func (store *store) StepDefinitionDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("queue id is empty")
	}

	sqlStr, preparedArgs, err := goqu.Dialect(store.dbDriverName).
		From(store.stepDefinitionTableName).
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

func (store *store) StepDefinitionFindByID(ctx context.Context, id string) (StepDefinitionInterface, error) {
	if id == "" {
		return nil, errors.New("workflow definition id is empty")
	}

	query := StepDefinitionQuery().SetID(id).SetLimit(1)

	list, err := store.StepDefinitionList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) StepDefinitionList(ctx context.Context, query StepDefinitionQueryInterface) ([]StepDefinitionInterface, error) {
	q, columns, err := store.stepDefinitionSelectQuery(query)

	if err != nil {
		return []StepDefinitionInterface{}, err
	}

	sqlStr, sqlParams, errSql := q.Prepared(true).Select(columns...).ToSQL()

	store.logSql("list", sqlStr, sqlParams...)

	if errSql != nil {
		return []StepDefinitionInterface{}, errSql
	}

	db := sb.NewDatabase(store.db, store.dbDriverName)

	if db == nil {
		return []StepDefinitionInterface{}, errors.New("workflow store: database is nil")
	}

	modelMaps, err := db.SelectToMapString(sqlStr, sqlParams...)

	if err != nil {
		return []StepDefinitionInterface{}, err
	}

	list := []StepDefinitionInterface{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewStepDefinitionFromExistingData(modelMap)
		list = append(list, model)
	})

	return list, nil
}

func (store *store) StepDefinitionSoftDelete(ctx context.Context, stepDefinition StepDefinitionInterface) error {
	if stepDefinition == nil {
		return errors.New("queue is nil")
	}

	stepDefinition.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	return store.StepDefinitionUpdate(ctx, stepDefinition)
}

func (store *store) StepDefinitionSoftDeleteByID(ctx context.Context, id string) error {
	stepDefinition, err := store.StepDefinitionFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.StepDefinitionSoftDelete(ctx, stepDefinition)
}

func (store *store) StepDefinitionUpdate(ctx context.Context, stepDefinition StepDefinitionInterface) error {
	stepDefinition.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	dataChanged := stepDefinition.DataChanged()

	delete(dataChanged, COLUMN_ID) // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Update(store.stepDefinitionTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C(COLUMN_ID).Eq(stepDefinition.ID())).
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

	stepDefinition.MarkAsNotDirty()

	return err
}

func (store *store) stepDefinitionSelectQuery(options StepDefinitionQueryInterface) (selectDataset *goqu.SelectDataset, columns []any, err error) {
	if options == nil {
		return nil, []any{}, errors.New("site options cannot be nil")
	}

	if err := options.Validate(); err != nil {
		return nil, []any{}, err
	}

	q := goqu.Dialect(store.dbDriverName).From(store.stepDefinitionTableName)

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
