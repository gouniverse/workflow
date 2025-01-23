package workflow

import "errors"

func WorkflowQuery() WorkflowQueryInterface {
	return &workflowQueryImpl{
		properties: make(map[string]any),
	}
}

type workflowQueryImpl struct {
	properties map[string]any
}

func (q *workflowQueryImpl) Validate() error {
	if q.HasCreatedAtGte() && q.CreatedAtGte() == "" {
		return errors.New("workflow definition query. created_at_gte cannot be empty")
	}

	if q.HasCreatedAtLte() && q.CreatedAtLte() == "" {
		return errors.New("workflow definition query. created_at_lte cannot be empty")
	}

	if q.HasID() && q.ID() == "" {
		return errors.New("workflow definition query. id cannot be empty")
	}

	if q.HasIDIn() && len(q.IDIn()) < 1 {
		return errors.New("workflow definition query. id_in cannot be empty array")
	}

	if q.HasLimit() && q.Limit() < 0 {
		return errors.New("workflow definition query. limit cannot be negative")
	}

	if q.HasOffset() && q.Offset() < 0 {
		return errors.New("workflow definition query. offset cannot be negative")
	}

	if q.HasStatus() && q.Status() == "" {
		return errors.New("workflow definition query. status cannot be empty")
	}

	if q.HasStatusIn() && len(q.StatusIn()) < 1 {
		return errors.New("workflow definition query. status_in cannot be empty array")
	}

	if q.HasUpdatedAtGte() && q.UpdatedAtGte() == "" {
		return errors.New("workflow definition query. updated_at_gte cannot be empty")
	}

	if q.HasUpdatedAtLte() && q.UpdatedAtLte() == "" {
		return errors.New("workflow definition query. updated_at_lte cannot be empty")
	}

	if q.HasWorkflowDefinitionID() && q.WorkflowDefinitionID() == "" {
		return errors.New("workflow definition query. workflow_definition_id cannot be empty")
	}

	return nil
}

func (q *workflowQueryImpl) IsCountOnly() bool {
	if q.hasProperty("countOnly") {
		return q.properties["countOnly"].(bool)
	}

	return false
}

func (q *workflowQueryImpl) SetCountOnly(countOnly bool) WorkflowQueryInterface {
	q.properties["countOnly"] = countOnly
	return q
}

func (q *workflowQueryImpl) Columns() []string {
	if q.hasProperty("columns") {
		return q.properties["columns"].([]string)
	}

	return nil
}

func (q *workflowQueryImpl) SetColumns(columns []string) WorkflowQueryInterface {
	q.properties["columns"] = columns
	return q
}

func (q *workflowQueryImpl) HasCreatedAtGte() bool {
	return q.hasProperty("created_at_gte")
}

func (q *workflowQueryImpl) CreatedAtGte() string {
	if q.HasCreatedAtGte() {
		return q.properties["created_at_gte"].(string)
	}
	return ""
}

func (q *workflowQueryImpl) SetCreatedAtGte(createdAtGte string) WorkflowQueryInterface {
	q.properties["created_at_gte"] = createdAtGte
	return q
}

func (q *workflowQueryImpl) HasCreatedAtLte() bool {
	return q.hasProperty("created_at_lte")
}

func (q *workflowQueryImpl) CreatedAtLte() string {
	if q.HasCreatedAtLte() {
		return q.properties["created_at_lte"].(string)
	}
	return ""
}

func (q *workflowQueryImpl) SetCreatedAtLte(createdAtLte string) WorkflowQueryInterface {
	q.properties["created_at_lte"] = createdAtLte
	return q
}

func (q *workflowQueryImpl) HasID() bool {
	return q.hasProperty("id")
}

func (q *workflowQueryImpl) ID() string {
	if q.HasID() {
		return q.properties["id"].(string)
	}
	return ""
}

func (q *workflowQueryImpl) SetID(id string) WorkflowQueryInterface {
	q.properties["id"] = id
	return q
}

func (q *workflowQueryImpl) HasIDIn() bool {
	return q.hasProperty("id_in")
}

func (q *workflowQueryImpl) IDIn() []string {
	if q.HasIDIn() {
		return q.properties["id_in"].([]string)
	}

	return nil
}

func (q *workflowQueryImpl) SetIDIn(idIn []string) WorkflowQueryInterface {
	q.properties["id_in"] = idIn
	return q
}

func (q *workflowQueryImpl) HasLimit() bool {
	return q.hasProperty("limit")
}

func (q *workflowQueryImpl) Limit() int {
	if q.HasLimit() {
		return q.properties["limit"].(int)
	}

	return 0
}

func (q *workflowQueryImpl) SetLimit(limit int) WorkflowQueryInterface {
	q.properties["limit"] = limit
	return q
}

func (q *workflowQueryImpl) HasOffset() bool {
	return q.hasProperty("offset")
}

func (q *workflowQueryImpl) Offset() int {
	if q.HasOffset() {
		return q.properties["offset"].(int)
	}

	return 0
}

func (q *workflowQueryImpl) SetOffset(offset int) WorkflowQueryInterface {
	q.properties["offset"] = offset
	return q
}

func (q *workflowQueryImpl) HasOrderBy() bool {
	return q.hasProperty("order_by")
}

func (q *workflowQueryImpl) OrderBy() string {
	if q.HasOrderBy() {
		return q.properties["order_by"].(string)
	}
	return ""
}

func (q *workflowQueryImpl) SetOrderBy(orderBy string) WorkflowQueryInterface {
	q.properties["order_by"] = orderBy
	return q
}

func (q *workflowQueryImpl) HasSortOrder() bool {
	return q.hasProperty("sort_order")
}

func (q *workflowQueryImpl) SortOrder() string {
	if q.HasSortOrder() {
		return q.properties["sort_order"].(string)
	}
	return ""
}

func (q *workflowQueryImpl) SetSortOrder(sortOrder string) WorkflowQueryInterface {
	q.properties["sort_order"] = sortOrder
	return q
}

func (q *workflowQueryImpl) HasStatus() bool {
	return q.hasProperty("status")
}

func (q *workflowQueryImpl) Status() string {
	if q.HasStatus() {
		return q.properties["status"].(string)
	}
	return ""
}

func (q *workflowQueryImpl) SetStatus(status string) WorkflowQueryInterface {
	q.properties["status"] = status
	return q
}

func (q *workflowQueryImpl) HasStatusIn() bool {
	return q.hasProperty("status_in")
}

func (q *workflowQueryImpl) StatusIn() []string {
	if q.HasStatusIn() {
		return q.properties["status_in"].([]string)
	}

	return []string{}
}

func (q *workflowQueryImpl) SetStatusIn(statusIn []string) WorkflowQueryInterface {
	q.properties["status_in"] = statusIn
	return q
}

func (q *workflowQueryImpl) HasUpdatedAtGte() bool {
	return q.hasProperty("updated_at_gte")
}

func (q *workflowQueryImpl) UpdatedAtGte() string {
	if q.HasUpdatedAtGte() {
		return q.properties["updated_at_gte"].(string)
	}
	return ""
}

func (q *workflowQueryImpl) SetUpdatedAtGte(updatedAtGte string) WorkflowQueryInterface {
	q.properties["updated_at_gte"] = updatedAtGte
	return q
}

func (q *workflowQueryImpl) HasUpdatedAtLte() bool {
	return q.hasProperty("updated_at_lte")
}

func (q *workflowQueryImpl) UpdatedAtLte() string {
	if q.HasUpdatedAtLte() {
		return q.properties["updated_at_lte"].(string)
	}
	return ""
}

func (q *workflowQueryImpl) SetUpdatedAtLte(updatedAtLte string) WorkflowQueryInterface {
	q.properties["updated_at_lte"] = updatedAtLte
	return q
}

func (q *workflowQueryImpl) HasWorkflowDefinitionID() bool {
	return q.hasProperty("workflow_definition_id")
}

func (q *workflowQueryImpl) WorkflowDefinitionID() string {
	if q.HasWorkflowDefinitionID() {
		return q.properties["workflow_definition_id"].(string)
	}
	return ""
}

func (q *workflowQueryImpl) SetWorkflowDefinitionID(workflowDefinitionID string) WorkflowQueryInterface {
	q.properties["workflow_definition_id"] = workflowDefinitionID
	return q
}

func (q *workflowQueryImpl) SoftDeletedIncluded() bool {
	if q.hasProperty("softDeletedIncluded") {
		return q.properties["softDeletedIncluded"].(bool)
	}

	return false
}

func (q *workflowQueryImpl) SetSoftDeletedIncluded(softDeletedIncluded bool) WorkflowQueryInterface {
	q.properties["softDeletedIncluded"] = softDeletedIncluded
	return q
}

func (c *workflowQueryImpl) hasProperty(name string) bool {
	_, ok := c.properties[name]
	return ok
}
