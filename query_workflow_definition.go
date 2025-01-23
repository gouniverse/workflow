package workflow

import "errors"

func WorkflowDefinitionQuery() WorkflowDefinitionQueryInterface {
	return &workflowDefinitionQueryImpl{
		properties: make(map[string]any),
	}
}

type workflowDefinitionQueryImpl struct {
	properties map[string]any
}

func (q *workflowDefinitionQueryImpl) Validate() error {
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

	return nil
}

func (q *workflowDefinitionQueryImpl) IsCountOnly() bool {
	if q.hasProperty("countOnly") {
		return q.properties["countOnly"].(bool)
	}

	return false
}

func (q *workflowDefinitionQueryImpl) SetCountOnly(countOnly bool) WorkflowDefinitionQueryInterface {
	q.properties["countOnly"] = countOnly
	return q
}

func (q *workflowDefinitionQueryImpl) Columns() []string {
	if q.hasProperty("columns") {
		return q.properties["columns"].([]string)
	}

	return nil
}

func (q *workflowDefinitionQueryImpl) SetColumns(columns []string) WorkflowDefinitionQueryInterface {
	q.properties["columns"] = columns
	return q
}

func (q *workflowDefinitionQueryImpl) HasCreatedAtGte() bool {
	return q.hasProperty("created_at_gte")
}

func (q *workflowDefinitionQueryImpl) CreatedAtGte() string {
	if q.HasCreatedAtGte() {
		return q.properties["created_at_gte"].(string)
	}
	return ""
}

func (q *workflowDefinitionQueryImpl) SetCreatedAtGte(createdAtGte string) WorkflowDefinitionQueryInterface {
	q.properties["created_at_gte"] = createdAtGte
	return q
}

func (q *workflowDefinitionQueryImpl) HasCreatedAtLte() bool {
	return q.hasProperty("created_at_lte")
}

func (q *workflowDefinitionQueryImpl) CreatedAtLte() string {
	if q.HasCreatedAtLte() {
		return q.properties["created_at_lte"].(string)
	}
	return ""
}

func (q *workflowDefinitionQueryImpl) SetCreatedAtLte(createdAtLte string) WorkflowDefinitionQueryInterface {
	q.properties["created_at_lte"] = createdAtLte
	return q
}

func (q *workflowDefinitionQueryImpl) HasID() bool {
	return q.hasProperty("id")
}

func (q *workflowDefinitionQueryImpl) ID() string {
	if q.HasID() {
		return q.properties["id"].(string)
	}
	return ""
}

func (q *workflowDefinitionQueryImpl) SetID(id string) WorkflowDefinitionQueryInterface {
	q.properties["id"] = id
	return q
}

func (q *workflowDefinitionQueryImpl) HasIDIn() bool {
	return q.hasProperty("id_in")
}

func (q *workflowDefinitionQueryImpl) IDIn() []string {
	if q.HasIDIn() {
		return q.properties["id_in"].([]string)
	}

	return nil
}

func (q *workflowDefinitionQueryImpl) SetIDIn(idIn []string) WorkflowDefinitionQueryInterface {
	q.properties["id_in"] = idIn
	return q
}

func (q *workflowDefinitionQueryImpl) HasLimit() bool {
	return q.hasProperty("limit")
}

func (q *workflowDefinitionQueryImpl) Limit() int {
	if q.HasLimit() {
		return q.properties["limit"].(int)
	}

	return 0
}

func (q *workflowDefinitionQueryImpl) SetLimit(limit int) WorkflowDefinitionQueryInterface {
	q.properties["limit"] = limit
	return q
}

func (q *workflowDefinitionQueryImpl) HasOffset() bool {
	return q.hasProperty("offset")
}

func (q *workflowDefinitionQueryImpl) Offset() int {
	if q.HasOffset() {
		return q.properties["offset"].(int)
	}

	return 0
}

func (q *workflowDefinitionQueryImpl) SetOffset(offset int) WorkflowDefinitionQueryInterface {
	q.properties["offset"] = offset
	return q
}

func (q *workflowDefinitionQueryImpl) HasOrderBy() bool {
	return q.hasProperty("order_by")
}

func (q *workflowDefinitionQueryImpl) OrderBy() string {
	if q.HasOrderBy() {
		return q.properties["order_by"].(string)
	}
	return ""
}

func (q *workflowDefinitionQueryImpl) SetOrderBy(orderBy string) WorkflowDefinitionQueryInterface {
	q.properties["order_by"] = orderBy
	return q
}

func (q *workflowDefinitionQueryImpl) HasSortOrder() bool {
	return q.hasProperty("sort_order")
}

func (q *workflowDefinitionQueryImpl) SortOrder() string {
	if q.HasSortOrder() {
		return q.properties["sort_order"].(string)
	}
	return ""
}

func (q *workflowDefinitionQueryImpl) SetSortOrder(sortOrder string) WorkflowDefinitionQueryInterface {
	q.properties["sort_order"] = sortOrder
	return q
}

func (q *workflowDefinitionQueryImpl) HasStatus() bool {
	return q.hasProperty("status")
}

func (q *workflowDefinitionQueryImpl) Status() string {
	if q.HasStatus() {
		return q.properties["status"].(string)
	}
	return ""
}

func (q *workflowDefinitionQueryImpl) SetStatus(status string) WorkflowDefinitionQueryInterface {
	q.properties["status"] = status
	return q
}

func (q *workflowDefinitionQueryImpl) HasStatusIn() bool {
	return q.hasProperty("status_in")
}

func (q *workflowDefinitionQueryImpl) StatusIn() []string {
	if q.HasStatusIn() {
		return q.properties["status_in"].([]string)
	}

	return []string{}
}

func (q *workflowDefinitionQueryImpl) SetStatusIn(statusIn []string) WorkflowDefinitionQueryInterface {
	q.properties["status_in"] = statusIn
	return q
}

func (q *workflowDefinitionQueryImpl) HasUpdatedAtGte() bool {
	return q.hasProperty("updated_at_gte")
}

func (q *workflowDefinitionQueryImpl) UpdatedAtGte() string {
	if q.HasUpdatedAtGte() {
		return q.properties["updated_at_gte"].(string)
	}
	return ""
}

func (q *workflowDefinitionQueryImpl) SetUpdatedAtGte(updatedAtGte string) WorkflowDefinitionQueryInterface {
	q.properties["updated_at_gte"] = updatedAtGte
	return q
}

func (q *workflowDefinitionQueryImpl) HasUpdatedAtLte() bool {
	return q.hasProperty("updated_at_lte")
}

func (q *workflowDefinitionQueryImpl) UpdatedAtLte() string {
	if q.HasUpdatedAtLte() {
		return q.properties["updated_at_lte"].(string)
	}
	return ""
}

func (q *workflowDefinitionQueryImpl) SetUpdatedAtLte(updatedAtLte string) WorkflowDefinitionQueryInterface {
	q.properties["updated_at_lte"] = updatedAtLte
	return q
}

func (q *workflowDefinitionQueryImpl) SoftDeletedIncluded() bool {
	if q.hasProperty("softDeletedIncluded") {
		return q.properties["softDeletedIncluded"].(bool)
	}

	return false
}

func (q *workflowDefinitionQueryImpl) SetSoftDeletedIncluded(softDeletedIncluded bool) WorkflowDefinitionQueryInterface {
	q.properties["softDeletedIncluded"] = softDeletedIncluded
	return q
}

func (c *workflowDefinitionQueryImpl) hasProperty(name string) bool {
	_, ok := c.properties[name]
	return ok
}
