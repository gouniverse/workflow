package workflow

import "errors"

func StepDefinitionQuery() StepDefinitionQueryInterface {
	return &stepDefinitionQueryImpl{
		properties: make(map[string]any),
	}
}

type stepDefinitionQueryImpl struct {
	properties map[string]any
}

func (q *stepDefinitionQueryImpl) Validate() error {
	if q.HasCreatedAtGte() && q.CreatedAtGte() == "" {
		return errors.New(`step definition query. created_at_gte cannot be empty`)
	}

	if q.HasCreatedAtLte() && q.CreatedAtLte() == "" {
		return errors.New(`step definition query. created_at_lte cannot be empty`)
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

	if q.HasStatus() && q.Status() == "" {
		return errors.New(`step definition query. status cannot be empty`)
	}

	if q.HasStatusIn() && len(q.StatusIn()) < 1 {
		return errors.New(`step definition query. status_in cannot be empty array`)
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

func (q *stepDefinitionQueryImpl) IsCountOnly() bool {
	if q.hasProperty("count_only") {
		return q.properties["count_only"].(bool)
	}
	return false
}

func (q *stepDefinitionQueryImpl) SetCountOnly(countOnly bool) StepDefinitionQueryInterface {
	q.properties["count_only"] = countOnly
	return q
}

func (q *stepDefinitionQueryImpl) Columns() []string {
	if q.hasProperty("columns") {
		return q.properties["columns"].([]string)
	}
	return []string{}
}

func (q *stepDefinitionQueryImpl) SetColumns(columns []string) StepDefinitionQueryInterface {
	q.properties["columns"] = columns
	return q
}

func (q *stepDefinitionQueryImpl) SoftDeletedIncluded() bool {
	if q.hasProperty("soft_deleted_included") {
		return q.properties["soft_deleted_included"].(bool)
	}
	return false
}

func (q *stepDefinitionQueryImpl) SetSoftDeletedIncluded(softDeletedIncluded bool) StepDefinitionQueryInterface {
	q.properties["soft_deleted_included"] = softDeletedIncluded
	return q
}

func (q *stepDefinitionQueryImpl) HasCreatedAtGte() bool {
	return q.hasProperty("created_at_gte")
}

func (q *stepDefinitionQueryImpl) CreatedAtGte() string {
	if q.HasCreatedAtGte() {
		return q.properties["created_at_gte"].(string)
	}
	return ""
}

func (q *stepDefinitionQueryImpl) SetCreatedAtGte(createdAtGte string) StepDefinitionQueryInterface {
	q.properties["created_at_gte"] = createdAtGte
	return q
}

func (q *stepDefinitionQueryImpl) HasCreatedAtLte() bool {
	return q.hasProperty("created_at_lte")
}

func (q *stepDefinitionQueryImpl) CreatedAtLte() string {
	if q.HasCreatedAtLte() {
		return q.properties["created_at_lte"].(string)
	}
	return ""
}

func (q *stepDefinitionQueryImpl) SetCreatedAtLte(createdAtLte string) StepDefinitionQueryInterface {
	q.properties["created_at_lte"] = createdAtLte
	return q
}

func (q *stepDefinitionQueryImpl) HasID() bool {
	return q.hasProperty("id")
}

func (q *stepDefinitionQueryImpl) ID() string {
	if q.hasProperty("id") {
		return q.properties["id"].(string)
	}
	return ""
}

func (q *stepDefinitionQueryImpl) SetID(id string) StepDefinitionQueryInterface {
	q.properties["id"] = id
	return q
}

func (q *stepDefinitionQueryImpl) HasIDIn() bool {
	return q.hasProperty("id_in")
}

func (q *stepDefinitionQueryImpl) IDIn() []string {
	if q.HasIDIn() {
		return q.properties["id_in"].([]string)
	}
	return []string{}
}

func (q *stepDefinitionQueryImpl) SetIDIn(idIn []string) StepDefinitionQueryInterface {
	q.properties["id_in"] = idIn
	return q
}

func (q *stepDefinitionQueryImpl) HasLimit() bool {
	return q.hasProperty("limit")
}

func (q *stepDefinitionQueryImpl) Limit() int {
	if q.HasLimit() {
		return q.properties["limit"].(int)
	}
	return 0
}

func (q *stepDefinitionQueryImpl) SetLimit(limit int) StepDefinitionQueryInterface {
	q.properties["limit"] = limit
	return q
}

func (q *stepDefinitionQueryImpl) HasOffset() bool {
	return q.hasProperty("offset")
}

func (q *stepDefinitionQueryImpl) Offset() int {
	if q.HasOffset() {
		return q.properties["offset"].(int)
	}
	return 0
}

func (q *stepDefinitionQueryImpl) SetOffset(offset int) StepDefinitionQueryInterface {
	q.properties["offset"] = offset
	return q
}

func (q *stepDefinitionQueryImpl) HasOrderBy() bool {
	return q.hasProperty("order_by")
}

func (q *stepDefinitionQueryImpl) OrderBy() string {
	if q.HasOrderBy() {
		return q.properties["order_by"].(string)
	}
	return ""
}

func (q *stepDefinitionQueryImpl) SetOrderBy(orderBy string) StepDefinitionQueryInterface {
	q.properties["order_by"] = orderBy
	return q
}

func (q *stepDefinitionQueryImpl) HasSortOrder() bool {
	return q.hasProperty("sort_order")
}

func (q *stepDefinitionQueryImpl) SortOrder() string {
	if q.HasSortOrder() {
		return q.properties["sort_order"].(string)
	}
	return ""
}

func (q *stepDefinitionQueryImpl) SetSortOrder(sortOrder string) StepDefinitionQueryInterface {
	q.properties["sort_order"] = sortOrder
	return q
}

func (q *stepDefinitionQueryImpl) HasStatus() bool {
	return q.hasProperty("status")
}

func (q *stepDefinitionQueryImpl) Status() string {
	if q.HasStatus() {
		return q.properties["status"].(string)
	}
	return ""
}

func (q *stepDefinitionQueryImpl) SetStatus(status string) StepDefinitionQueryInterface {
	q.properties["status"] = status
	return q
}

func (q *stepDefinitionQueryImpl) HasStatusIn() bool {
	return q.hasProperty("status_in")
}

func (q *stepDefinitionQueryImpl) StatusIn() []string {
	if q.HasStatusIn() {
		return q.properties["status_in"].([]string)
	}
	return []string{}
}

func (q *stepDefinitionQueryImpl) SetStatusIn(statusIn []string) StepDefinitionQueryInterface {
	q.properties["status_in"] = statusIn
	return q
}

func (q *stepDefinitionQueryImpl) HasUpdatedAtGte() bool {
	return q.hasProperty("updated_at_gte")
}

func (q *stepDefinitionQueryImpl) UpdatedAtGte() string {
	if q.HasUpdatedAtGte() {
		return q.properties["updated_at_gte"].(string)
	}
	return ""
}

func (q *stepDefinitionQueryImpl) SetUpdatedAtGte(updatedAtGte string) StepDefinitionQueryInterface {
	q.properties["updated_at_gte"] = updatedAtGte
	return q
}

func (q *stepDefinitionQueryImpl) HasUpdatedAtLte() bool {
	return q.hasProperty("updated_at_lte")
}

func (q *stepDefinitionQueryImpl) UpdatedAtLte() string {
	if q.HasUpdatedAtLte() {
		return q.properties["updated_at_lte"].(string)
	}
	return ""
}

func (q *stepDefinitionQueryImpl) SetUpdatedAtLte(updatedAtLte string) StepDefinitionQueryInterface {
	q.properties["updated_at_lte"] = updatedAtLte
	return q
}

func (q *stepDefinitionQueryImpl) HasWorkflowDefinitionID() bool {
	return q.hasProperty("workflow_definition_id")
}

func (q *stepDefinitionQueryImpl) WorkflowDefinitionID() string {
	if q.HasWorkflowDefinitionID() {
		return q.properties["workflow_definition_id"].(string)
	}
	return ""
}

func (q *stepDefinitionQueryImpl) SetWorkflowDefinitionID(workflowDefinitionID string) StepDefinitionQueryInterface {
	q.properties["workflow_definition_id"] = workflowDefinitionID
	return q
}

func (q *stepDefinitionQueryImpl) hasProperty(key string) bool {
	_, ok := q.properties[key]
	return ok
}
