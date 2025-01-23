package workflow

import "errors"

func EdgeDefinitionQuery() EdgeDefinitionQueryInterface {
	return &edgeDefinitionQueryImpl{
		properties: make(map[string]any),
	}
}

type edgeDefinitionQueryImpl struct {
	properties map[string]any
}

func (q *edgeDefinitionQueryImpl) Validate() error {
	if q.HasCreatedAtGte() && q.CreatedAtGte() == "" {
		return errors.New(`step definition query. created_at_gte cannot be empty`)
	}

	if q.HasCreatedAtLte() && q.CreatedAtLte() == "" {
		return errors.New(`step definition query. created_at_lte cannot be empty`)
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

	if q.HasWorkflowDefinitionID() && q.WorkflowDefinitionID() == "" {
		return errors.New(`step definition query. workflow_definition_id cannot be empty`)
	}

	return nil
}

func (q *edgeDefinitionQueryImpl) IsCountOnly() bool {
	if q.hasProperty("count_only") {
		return q.properties["count_only"].(bool)
	}
	return false
}

func (q *edgeDefinitionQueryImpl) SetCountOnly(countOnly bool) EdgeDefinitionQueryInterface {
	q.properties["count_only"] = countOnly
	return q
}

func (q *edgeDefinitionQueryImpl) Columns() []string {
	if q.hasProperty("columns") {
		return q.properties["columns"].([]string)
	}
	return []string{}
}

func (q *edgeDefinitionQueryImpl) SetColumns(columns []string) EdgeDefinitionQueryInterface {
	q.properties["columns"] = columns
	return q
}

func (q *edgeDefinitionQueryImpl) SoftDeletedIncluded() bool {
	if q.hasProperty("soft_deleted_included") {
		return q.properties["soft_deleted_included"].(bool)
	}
	return false
}

func (q *edgeDefinitionQueryImpl) SetSoftDeletedIncluded(softDeletedIncluded bool) EdgeDefinitionQueryInterface {
	q.properties["soft_deleted_included"] = softDeletedIncluded
	return q
}

func (q *edgeDefinitionQueryImpl) HasCreatedAtGte() bool {
	return q.hasProperty("created_at_gte")
}

func (q *edgeDefinitionQueryImpl) CreatedAtGte() string {
	if q.HasCreatedAtGte() {
		return q.properties["created_at_gte"].(string)
	}
	return ""
}

func (q *edgeDefinitionQueryImpl) SetCreatedAtGte(createdAtGte string) EdgeDefinitionQueryInterface {
	q.properties["created_at_gte"] = createdAtGte
	return q
}

func (q *edgeDefinitionQueryImpl) HasCreatedAtLte() bool {
	return q.hasProperty("created_at_lte")
}

func (q *edgeDefinitionQueryImpl) CreatedAtLte() string {
	if q.HasCreatedAtLte() {
		return q.properties["created_at_lte"].(string)
	}
	return ""
}

func (q *edgeDefinitionQueryImpl) SetCreatedAtLte(createdAtLte string) EdgeDefinitionQueryInterface {
	q.properties["created_at_lte"] = createdAtLte
	return q
}

func (q *edgeDefinitionQueryImpl) HasFromStepID() bool {
	return q.hasProperty("from_step_id")
}

func (q *edgeDefinitionQueryImpl) FromStepID() string {
	if q.HasFromStepID() {
		return q.properties["from_step_id"].(string)
	}
	return ""
}

func (q *edgeDefinitionQueryImpl) SetFromStepID(fromStepID string) EdgeDefinitionQueryInterface {
	q.properties["from_step_id"] = fromStepID
	return q
}

func (q *edgeDefinitionQueryImpl) HasID() bool {
	return q.hasProperty("id")
}

func (q *edgeDefinitionQueryImpl) ID() string {
	if q.hasProperty("id") {
		return q.properties["id"].(string)
	}
	return ""
}

func (q *edgeDefinitionQueryImpl) SetID(id string) EdgeDefinitionQueryInterface {
	q.properties["id"] = id
	return q
}

func (q *edgeDefinitionQueryImpl) HasIDIn() bool {
	return q.hasProperty("id_in")
}

func (q *edgeDefinitionQueryImpl) IDIn() []string {
	if q.HasIDIn() {
		return q.properties["id_in"].([]string)
	}
	return []string{}
}

func (q *edgeDefinitionQueryImpl) SetIDIn(idIn []string) EdgeDefinitionQueryInterface {
	q.properties["id_in"] = idIn
	return q
}

func (q *edgeDefinitionQueryImpl) HasLimit() bool {
	return q.hasProperty("limit")
}

func (q *edgeDefinitionQueryImpl) Limit() int {
	if q.HasLimit() {
		return q.properties["limit"].(int)
	}
	return 0
}

func (q *edgeDefinitionQueryImpl) SetLimit(limit int) EdgeDefinitionQueryInterface {
	q.properties["limit"] = limit
	return q
}

func (q *edgeDefinitionQueryImpl) HasOffset() bool {
	return q.hasProperty("offset")
}

func (q *edgeDefinitionQueryImpl) Offset() int {
	if q.HasOffset() {
		return q.properties["offset"].(int)
	}
	return 0
}

func (q *edgeDefinitionQueryImpl) SetOffset(offset int) EdgeDefinitionQueryInterface {
	q.properties["offset"] = offset
	return q
}

func (q *edgeDefinitionQueryImpl) HasOrderBy() bool {
	return q.hasProperty("order_by")
}

func (q *edgeDefinitionQueryImpl) OrderBy() string {
	if q.HasOrderBy() {
		return q.properties["order_by"].(string)
	}
	return ""
}

func (q *edgeDefinitionQueryImpl) SetOrderBy(orderBy string) EdgeDefinitionQueryInterface {
	q.properties["order_by"] = orderBy
	return q
}

func (q *edgeDefinitionQueryImpl) HasSortOrder() bool {
	return q.hasProperty("sort_order")
}

func (q *edgeDefinitionQueryImpl) SortOrder() string {
	if q.HasSortOrder() {
		return q.properties["sort_order"].(string)
	}
	return ""
}

func (q *edgeDefinitionQueryImpl) SetSortOrder(sortOrder string) EdgeDefinitionQueryInterface {
	q.properties["sort_order"] = sortOrder
	return q
}

func (q *edgeDefinitionQueryImpl) HasToStepID() bool {
	return q.hasProperty("to_step_id")
}

func (q *edgeDefinitionQueryImpl) ToStepID() string {
	if q.HasToStepID() {
		return q.properties["to_step_id"].(string)
	}
	return ""
}

func (q *edgeDefinitionQueryImpl) SetToStepID(toStepID string) EdgeDefinitionQueryInterface {
	q.properties["to_step_id"] = toStepID
	return q
}

func (q *edgeDefinitionQueryImpl) HasUpdatedAtGte() bool {
	return q.hasProperty("updated_at_gte")
}

func (q *edgeDefinitionQueryImpl) UpdatedAtGte() string {
	if q.HasUpdatedAtGte() {
		return q.properties["updated_at_gte"].(string)
	}
	return ""
}

func (q *edgeDefinitionQueryImpl) SetUpdatedAtGte(updatedAtGte string) EdgeDefinitionQueryInterface {
	q.properties["updated_at_gte"] = updatedAtGte
	return q
}

func (q *edgeDefinitionQueryImpl) HasUpdatedAtLte() bool {
	return q.hasProperty("updated_at_lte")
}

func (q *edgeDefinitionQueryImpl) UpdatedAtLte() string {
	if q.HasUpdatedAtLte() {
		return q.properties["updated_at_lte"].(string)
	}
	return ""
}

func (q *edgeDefinitionQueryImpl) SetUpdatedAtLte(updatedAtLte string) EdgeDefinitionQueryInterface {
	q.properties["updated_at_lte"] = updatedAtLte
	return q
}

func (q *edgeDefinitionQueryImpl) HasWorkflowDefinitionID() bool {
	return q.hasProperty("workflow_definition_id")
}

func (q *edgeDefinitionQueryImpl) WorkflowDefinitionID() string {
	if q.HasWorkflowDefinitionID() {
		return q.properties["workflow_definition_id"].(string)
	}
	return ""
}

func (q *edgeDefinitionQueryImpl) SetWorkflowDefinitionID(workflowDefinitionID string) EdgeDefinitionQueryInterface {
	q.properties["workflow_definition_id"] = workflowDefinitionID
	return q
}

func (q *edgeDefinitionQueryImpl) hasProperty(key string) bool {
	_, ok := q.properties[key]
	return ok
}
