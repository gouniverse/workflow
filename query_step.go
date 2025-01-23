package workflow

import "errors"

func StepQuery() StepQueryInterface {
	return &stepQueryImpl{
		properties: make(map[string]any),
	}
}

type stepQueryImpl struct {
	properties map[string]any
}

func (q *stepQueryImpl) Validate() error {
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

	if q.HasStepDefinitionID() && q.StepDefinitionID() == "" {
		return errors.New(`step definition query. step_definition_id cannot be empty`)
	}

	if q.HasUpdatedAtGte() && q.UpdatedAtGte() == "" {
		return errors.New(`step definition query. updated_at_gte cannot be empty`)
	}

	if q.HasUpdatedAtLte() && q.UpdatedAtLte() == "" {
		return errors.New(`step definition query. updated_at_lte cannot be empty`)
	}

	if q.HasWorkflowID() && q.WorkflowID() == "" {
		return errors.New(`step definition query. workflow_id cannot be empty`)
	}

	return nil
}

func (q *stepQueryImpl) IsCountOnly() bool {
	if q.hasProperty("count_only") {
		return q.properties["count_only"].(bool)
	}
	return false
}

func (q *stepQueryImpl) SetCountOnly(countOnly bool) StepQueryInterface {
	q.properties["count_only"] = countOnly
	return q
}

func (q *stepQueryImpl) Columns() []string {
	if q.hasProperty("columns") {
		return q.properties["columns"].([]string)
	}
	return []string{}
}

func (q *stepQueryImpl) SetColumns(columns []string) StepQueryInterface {
	q.properties["columns"] = columns
	return q
}

func (q *stepQueryImpl) SoftDeletedIncluded() bool {
	if q.hasProperty("soft_deleted_included") {
		return q.properties["soft_deleted_included"].(bool)
	}
	return false
}

func (q *stepQueryImpl) SetSoftDeletedIncluded(softDeletedIncluded bool) StepQueryInterface {
	q.properties["soft_deleted_included"] = softDeletedIncluded
	return q
}

func (q *stepQueryImpl) HasCreatedAtGte() bool {
	return q.hasProperty("created_at_gte")
}

func (q *stepQueryImpl) CreatedAtGte() string {
	if q.HasCreatedAtGte() {
		return q.properties["created_at_gte"].(string)
	}
	return ""
}

func (q *stepQueryImpl) SetCreatedAtGte(createdAtGte string) StepQueryInterface {
	q.properties["created_at_gte"] = createdAtGte
	return q
}

func (q *stepQueryImpl) HasCreatedAtLte() bool {
	return q.hasProperty("created_at_lte")
}

func (q *stepQueryImpl) CreatedAtLte() string {
	if q.HasCreatedAtLte() {
		return q.properties["created_at_lte"].(string)
	}
	return ""
}

func (q *stepQueryImpl) SetCreatedAtLte(createdAtLte string) StepQueryInterface {
	q.properties["created_at_lte"] = createdAtLte
	return q
}

func (q *stepQueryImpl) HasID() bool {
	return q.hasProperty("id")
}

func (q *stepQueryImpl) ID() string {
	if q.hasProperty("id") {
		return q.properties["id"].(string)
	}
	return ""
}

func (q *stepQueryImpl) SetID(id string) StepQueryInterface {
	q.properties["id"] = id
	return q
}

func (q *stepQueryImpl) HasIDIn() bool {
	return q.hasProperty("id_in")
}

func (q *stepQueryImpl) IDIn() []string {
	if q.HasIDIn() {
		return q.properties["id_in"].([]string)
	}
	return []string{}
}

func (q *stepQueryImpl) SetIDIn(idIn []string) StepQueryInterface {
	q.properties["id_in"] = idIn
	return q
}

func (q *stepQueryImpl) HasLimit() bool {
	return q.hasProperty("limit")
}

func (q *stepQueryImpl) Limit() int {
	if q.HasLimit() {
		return q.properties["limit"].(int)
	}
	return 0
}

func (q *stepQueryImpl) SetLimit(limit int) StepQueryInterface {
	q.properties["limit"] = limit
	return q
}

func (q *stepQueryImpl) HasOffset() bool {
	return q.hasProperty("offset")
}

func (q *stepQueryImpl) Offset() int {
	if q.HasOffset() {
		return q.properties["offset"].(int)
	}
	return 0
}

func (q *stepQueryImpl) SetOffset(offset int) StepQueryInterface {
	q.properties["offset"] = offset
	return q
}

func (q *stepQueryImpl) HasOrderBy() bool {
	return q.hasProperty("order_by")
}

func (q *stepQueryImpl) OrderBy() string {
	if q.HasOrderBy() {
		return q.properties["order_by"].(string)
	}
	return ""
}

func (q *stepQueryImpl) SetOrderBy(orderBy string) StepQueryInterface {
	q.properties["order_by"] = orderBy
	return q
}

func (q *stepQueryImpl) HasSortOrder() bool {
	return q.hasProperty("sort_order")
}

func (q *stepQueryImpl) SortOrder() string {
	if q.HasSortOrder() {
		return q.properties["sort_order"].(string)
	}
	return ""
}

func (q *stepQueryImpl) SetSortOrder(sortOrder string) StepQueryInterface {
	q.properties["sort_order"] = sortOrder
	return q
}

func (q *stepQueryImpl) HasStatus() bool {
	return q.hasProperty("status")
}

func (q *stepQueryImpl) Status() string {
	if q.HasStatus() {
		return q.properties["status"].(string)
	}
	return ""
}

func (q *stepQueryImpl) SetStatus(status string) StepQueryInterface {
	q.properties["status"] = status
	return q
}

func (q *stepQueryImpl) HasStatusIn() bool {
	return q.hasProperty("status_in")
}

func (q *stepQueryImpl) StatusIn() []string {
	if q.HasStatusIn() {
		return q.properties["status_in"].([]string)
	}
	return []string{}
}

func (q *stepQueryImpl) SetStatusIn(statusIn []string) StepQueryInterface {
	q.properties["status_in"] = statusIn
	return q
}

func (q *stepQueryImpl) HasStepDefinitionID() bool {
	return q.hasProperty("step_definition_id")
}

func (q *stepQueryImpl) StepDefinitionID() string {
	if q.HasStepDefinitionID() {
		return q.properties["step_definition_id"].(string)
	}
	return ""
}

func (q *stepQueryImpl) SetStepDefinitionID(stepDefinitionID string) StepQueryInterface {
	q.properties["step_definition_id"] = stepDefinitionID
	return q
}

func (q *stepQueryImpl) HasUpdatedAtGte() bool {
	return q.hasProperty("updated_at_gte")
}

func (q *stepQueryImpl) UpdatedAtGte() string {
	if q.HasUpdatedAtGte() {
		return q.properties["updated_at_gte"].(string)
	}
	return ""
}

func (q *stepQueryImpl) SetUpdatedAtGte(updatedAtGte string) StepQueryInterface {
	q.properties["updated_at_gte"] = updatedAtGte
	return q
}

func (q *stepQueryImpl) HasUpdatedAtLte() bool {
	return q.hasProperty("updated_at_lte")
}

func (q *stepQueryImpl) UpdatedAtLte() string {
	if q.HasUpdatedAtLte() {
		return q.properties["updated_at_lte"].(string)
	}
	return ""
}

func (q *stepQueryImpl) SetUpdatedAtLte(updatedAtLte string) StepQueryInterface {
	q.properties["updated_at_lte"] = updatedAtLte
	return q
}

func (q *stepQueryImpl) HasWorkflowID() bool {
	return q.hasProperty("workflow_id")
}

func (q *stepQueryImpl) WorkflowID() string {
	if q.HasWorkflowID() {
		return q.properties["workflow_id"].(string)
	}
	return ""
}

func (q *stepQueryImpl) SetWorkflowID(workflowID string) StepQueryInterface {
	q.properties["workflow_id"] = workflowID
	return q
}

func (q *stepQueryImpl) hasProperty(key string) bool {
	_, ok := q.properties[key]
	return ok
}
