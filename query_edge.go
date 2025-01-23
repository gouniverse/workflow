package workflow

import "errors"

func EdgeQuery() EdgeQueryInterface {
	return &edgeDefinitionImpl{
		properties: make(map[string]any),
	}
}

type edgeDefinitionImpl struct {
	properties map[string]any
}

func (q *edgeDefinitionImpl) Validate() error {
	if q.HasCreatedAtGte() && q.CreatedAtGte() == "" {
		return errors.New(`step definition query. created_at_gte cannot be empty`)
	}

	if q.HasCreatedAtLte() && q.CreatedAtLte() == "" {
		return errors.New(`step definition query. created_at_lte cannot be empty`)
	}

	if q.HasEdgeDefinitionID() && q.EdgeDefinitionID() == "" {
		return errors.New(`step definition query. edge_definition_id cannot be empty`)
	}

	if q.HasFromStepID() && q.FromStepID() == "" {
		return errors.New(`step definition query. from_step_id cannot be empty`)
	}

	if q.HasID() && q.ID() == "" {
		return errors.New(`step definition query. id cannot be empty`)
	}

	if q.HasIDIn() && len(q.IDIn()) < 1 {
		return errors.New(`step definition query. id_in cannot be empty array`)
	}

	if q.HasLimit() && q.Limit() < 0 {
		return errors.New(`step definition query. limit cannot be negative`)
	}

	if q.HasOffset() && q.Offset() < 0 {
		return errors.New(`step definition query. offset cannot be negative`)
	}

	if q.HasToStepID() && q.ToStepID() == "" {
		return errors.New(`step definition query. to_step_id cannot be empty`)
	}

	if q.HasUpdatedAtGte() && q.UpdatedAtGte() == "" {
		return errors.New(`step definition query. updated_at_gte cannot be empty`)
	}

	if q.HasUpdatedAtLte() && q.UpdatedAtLte() == "" {
		return errors.New(`step definition query. updated_at_lte cannot be empty`)
	}

	return nil
}

func (q *edgeDefinitionImpl) IsCountOnly() bool {
	if q.hasProperty("count_only") {
		return q.properties["count_only"].(bool)
	}
	return false
}

func (q *edgeDefinitionImpl) SetCountOnly(countOnly bool) EdgeQueryInterface {
	q.properties["count_only"] = countOnly
	return q
}

func (q *edgeDefinitionImpl) Columns() []string {
	if q.hasProperty("columns") {
		return q.properties["columns"].([]string)
	}
	return []string{}
}

func (q *edgeDefinitionImpl) SetColumns(columns []string) EdgeQueryInterface {
	q.properties["columns"] = columns
	return q
}

func (q *edgeDefinitionImpl) SoftDeletedIncluded() bool {
	if q.hasProperty("soft_deleted_included") {
		return q.properties["soft_deleted_included"].(bool)
	}
	return false
}

func (q *edgeDefinitionImpl) SetSoftDeletedIncluded(softDeletedIncluded bool) EdgeQueryInterface {
	q.properties["soft_deleted_included"] = softDeletedIncluded
	return q
}

func (q *edgeDefinitionImpl) HasCreatedAtGte() bool {
	return q.hasProperty("created_at_gte")
}

func (q *edgeDefinitionImpl) CreatedAtGte() string {
	if q.HasCreatedAtGte() {
		return q.properties["created_at_gte"].(string)
	}
	return ""
}

func (q *edgeDefinitionImpl) SetCreatedAtGte(createdAtGte string) EdgeQueryInterface {
	q.properties["created_at_gte"] = createdAtGte
	return q
}

func (q *edgeDefinitionImpl) HasCreatedAtLte() bool {
	return q.hasProperty("created_at_lte")
}

func (q *edgeDefinitionImpl) CreatedAtLte() string {
	if q.HasCreatedAtLte() {
		return q.properties["created_at_lte"].(string)
	}
	return ""
}

func (q *edgeDefinitionImpl) SetCreatedAtLte(createdAtLte string) EdgeQueryInterface {
	q.properties["created_at_lte"] = createdAtLte
	return q
}

func (q *edgeDefinitionImpl) HasEdgeDefinitionID() bool {
	return q.hasProperty("edge_definition_id")
}

func (q *edgeDefinitionImpl) EdgeDefinitionID() string {
	if q.HasEdgeDefinitionID() {
		return q.properties["edge_definition_id"].(string)
	}
	return ""
}

func (q *edgeDefinitionImpl) SetEdgeDefinitionID(edgeDefinitionID string) EdgeQueryInterface {
	q.properties["edge_definition_id"] = edgeDefinitionID
	return q
}

func (q *edgeDefinitionImpl) HasFromStepID() bool {
	return q.hasProperty("from_step_id")
}

func (q *edgeDefinitionImpl) FromStepID() string {
	if q.HasFromStepID() {
		return q.properties["from_step_id"].(string)
	}
	return ""
}

func (q *edgeDefinitionImpl) SetFromStepID(fromStepID string) EdgeQueryInterface {
	q.properties["from_step_id"] = fromStepID
	return q
}

func (q *edgeDefinitionImpl) HasID() bool {
	return q.hasProperty("id")
}

func (q *edgeDefinitionImpl) ID() string {
	if q.hasProperty("id") {
		return q.properties["id"].(string)
	}
	return ""
}

func (q *edgeDefinitionImpl) SetID(id string) EdgeQueryInterface {
	q.properties["id"] = id
	return q
}

func (q *edgeDefinitionImpl) HasIDIn() bool {
	return q.hasProperty("id_in")
}

func (q *edgeDefinitionImpl) IDIn() []string {
	if q.HasIDIn() {
		return q.properties["id_in"].([]string)
	}
	return []string{}
}

func (q *edgeDefinitionImpl) SetIDIn(idIn []string) EdgeQueryInterface {
	q.properties["id_in"] = idIn
	return q
}

func (q *edgeDefinitionImpl) HasLimit() bool {
	return q.hasProperty("limit")
}

func (q *edgeDefinitionImpl) Limit() int {
	if q.HasLimit() {
		return q.properties["limit"].(int)
	}
	return 0
}

func (q *edgeDefinitionImpl) SetLimit(limit int) EdgeQueryInterface {
	q.properties["limit"] = limit
	return q
}

func (q *edgeDefinitionImpl) HasOffset() bool {
	return q.hasProperty("offset")
}

func (q *edgeDefinitionImpl) Offset() int {
	if q.HasOffset() {
		return q.properties["offset"].(int)
	}
	return 0
}

func (q *edgeDefinitionImpl) SetOffset(offset int) EdgeQueryInterface {
	q.properties["offset"] = offset
	return q
}

func (q *edgeDefinitionImpl) HasOrderBy() bool {
	return q.hasProperty("order_by")
}

func (q *edgeDefinitionImpl) OrderBy() string {
	if q.HasOrderBy() {
		return q.properties["order_by"].(string)
	}
	return ""
}

func (q *edgeDefinitionImpl) SetOrderBy(orderBy string) EdgeQueryInterface {
	q.properties["order_by"] = orderBy
	return q
}

func (q *edgeDefinitionImpl) HasSortOrder() bool {
	return q.hasProperty("sort_order")
}

func (q *edgeDefinitionImpl) SortOrder() string {
	if q.HasSortOrder() {
		return q.properties["sort_order"].(string)
	}
	return ""
}

func (q *edgeDefinitionImpl) SetSortOrder(sortOrder string) EdgeQueryInterface {
	q.properties["sort_order"] = sortOrder
	return q
}

func (q *edgeDefinitionImpl) HasToStepID() bool {
	return q.hasProperty("to_step_id")
}

func (q *edgeDefinitionImpl) ToStepID() string {
	if q.HasToStepID() {
		return q.properties["to_step_id"].(string)
	}
	return ""
}

func (q *edgeDefinitionImpl) SetToStepID(toStepID string) EdgeQueryInterface {
	q.properties["to_step_id"] = toStepID
	return q
}

func (q *edgeDefinitionImpl) HasUpdatedAtGte() bool {
	return q.hasProperty("updated_at_gte")
}

func (q *edgeDefinitionImpl) UpdatedAtGte() string {
	if q.HasUpdatedAtGte() {
		return q.properties["updated_at_gte"].(string)
	}
	return ""
}

func (q *edgeDefinitionImpl) SetUpdatedAtGte(updatedAtGte string) EdgeQueryInterface {
	q.properties["updated_at_gte"] = updatedAtGte
	return q
}

func (q *edgeDefinitionImpl) HasUpdatedAtLte() bool {
	return q.hasProperty("updated_at_lte")
}

func (q *edgeDefinitionImpl) UpdatedAtLte() string {
	if q.HasUpdatedAtLte() {
		return q.properties["updated_at_lte"].(string)
	}
	return ""
}

func (q *edgeDefinitionImpl) SetUpdatedAtLte(updatedAtLte string) EdgeQueryInterface {
	q.properties["updated_at_lte"] = updatedAtLte
	return q
}

func (q *edgeDefinitionImpl) hasProperty(key string) bool {
	_, ok := q.properties[key]
	return ok
}
