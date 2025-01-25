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

func (store *store) EdgeDefinitionCount(ctx context.Context, options EdgeDefinitionQueryInterface) (int64, error) {
	options.SetCountOnly(true)
	q, _, err := store.edgeDefinitionSelectQuery(options)

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

func (store *store) EdgeDefinitionCreate(ctx context.Context, edgeDefinition EdgeDefinitionInterface) error {
	edgeDefinition.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	edgeDefinition.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	data := edgeDefinition.Data()

	sqlStr, sqlParams, errSql := goqu.Dialect(store.dbDriverName).
		Insert(store.edgeDefinitionTableName).
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

func (store *store) EdgeDefinitionDelete(ctx context.Context, edgeDefinition EdgeDefinitionInterface) error {
	if edgeDefinition == nil {
		return errors.New("workflow definition is nil")
	}

	return store.EdgeDefinitionDeleteByID(ctx, edgeDefinition.ID())
}

func (store *store) EdgeDefinitionDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("queue id is empty")
	}

	sqlStr, preparedArgs, err := goqu.Dialect(store.dbDriverName).
		From(store.edgeDefinitionTableName).
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

func (store *store) EdgeDefinitionFindByID(ctx context.Context, id string) (EdgeDefinitionInterface, error) {
	if id == "" {
		return nil, errors.New("workflow definition id is empty")
	}

	query := EdgeDefinitionQuery().SetID(id).SetLimit(1)

	list, err := store.EdgeDefinitionList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) EdgeDefinitionList(ctx context.Context, query EdgeDefinitionQueryInterface) ([]EdgeDefinitionInterface, error) {
	q, columns, err := store.edgeDefinitionSelectQuery(query)

	if err != nil {
		return []EdgeDefinitionInterface{}, err
	}

	sqlStr, sqlParams, errSql := q.Prepared(true).Select(columns...).ToSQL()

	store.logSql("list", sqlStr, sqlParams...)

	if errSql != nil {
		return []EdgeDefinitionInterface{}, errSql
	}

	db := sb.NewDatabase(store.db, store.dbDriverName)

	if db == nil {
		return []EdgeDefinitionInterface{}, errors.New("workflow store: database is nil")
	}

	modelMaps, err := db.SelectToMapString(sqlStr, sqlParams...)

	if err != nil {
		return []EdgeDefinitionInterface{}, err
	}

	list := []EdgeDefinitionInterface{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewEdgeDefinitionFromExistingData(modelMap)
		list = append(list, model)
	})

	return list, nil
}

func (store *store) EdgeDefinitionSoftDelete(ctx context.Context, edgeDefinition EdgeDefinitionInterface) error {
	if edgeDefinition == nil {
		return errors.New("queue is nil")
	}

	edgeDefinition.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	return store.EdgeDefinitionUpdate(ctx, edgeDefinition)
}

func (store *store) EdgeDefinitionSoftDeleteByID(ctx context.Context, id string) error {
	edgeDefinition, err := store.EdgeDefinitionFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.EdgeDefinitionSoftDelete(ctx, edgeDefinition)
}

func (store *store) EdgeDefinitionUpdate(ctx context.Context, edgeDefinition EdgeDefinitionInterface) error {
	edgeDefinition.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	dataChanged := edgeDefinition.DataChanged()

	delete(dataChanged, COLUMN_ID) // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Update(store.edgeDefinitionTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C(COLUMN_ID).Eq(edgeDefinition.ID())).
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

	edgeDefinition.MarkAsNotDirty()

	return err
}

func (store *store) edgeDefinitionSelectQuery(options EdgeDefinitionQueryInterface) (selectDataset *goqu.SelectDataset, columns []any, err error) {
	if options == nil {
		return nil, []any{}, errors.New("site options cannot be nil")
	}

	if err := options.Validate(); err != nil {
		return nil, []any{}, err
	}

	q := goqu.Dialect(store.dbDriverName).From(store.edgeDefinitionTableName)

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

	if options.HasFromStepID() {
		q = q.Where(goqu.C(COLUMN_FROM_STEP_ID).Eq(options.FromStepID()))
	}

	if options.HasToStepID() {
		q = q.Where(goqu.C(COLUMN_TO_STEP_ID).Eq(options.ToStepID()))
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
