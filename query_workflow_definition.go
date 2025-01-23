package workflow

import "errors"

func WorkflowDefinitionQuery() WorkflowDefinitionQueryInterface {
	return &workflowDefinitionQuery{
		properties: make(map[string]any),
	}
}

type workflowDefinitionQuery struct {
	properties map[string]any
}

func (q *workflowDefinitionQuery) Validate() error {
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

func (q *workflowDefinitionQuery) IsCountOnly() bool {
	if q.hasProperty("countOnly") {
		return q.properties["countOnly"].(bool)
	}

	return false
}

func (q *workflowDefinitionQuery) SetCountOnly(countOnly bool) WorkflowDefinitionQueryInterface {
	q.properties["countOnly"] = countOnly
	return q
}

func (q *workflowDefinitionQuery) Columns() []string {
	if q.hasProperty("columns") {
		return q.properties["columns"].([]string)
	}

	return nil
}

func (q *workflowDefinitionQuery) SetColumns(columns []string) WorkflowDefinitionQueryInterface {
	q.properties["columns"] = columns
	return q
}

func (q *workflowDefinitionQuery) HasCreatedAtGte() bool {
	return q.hasProperty("created_at_gte")
}

func (q *workflowDefinitionQuery) CreatedAtGte() string {
	return q.properties["created_at_gte"].(string)
}

func (q *workflowDefinitionQuery) SetCreatedAtGte(createdAtGte string) WorkflowDefinitionQueryInterface {
	q.properties["created_at_gte"] = createdAtGte
	return q
}

func (q *workflowDefinitionQuery) HasCreatedAtLte() bool {
	return q.hasProperty("created_at_lte")
}

func (q *workflowDefinitionQuery) CreatedAtLte() string {
	return q.properties["created_at_lte"].(string)
}

func (q *workflowDefinitionQuery) SetCreatedAtLte(createdAtLte string) WorkflowDefinitionQueryInterface {
	q.properties["created_at_lte"] = createdAtLte
	return q
}

func (q *workflowDefinitionQuery) HasID() bool {
	return q.hasProperty("id")
}

func (q *workflowDefinitionQuery) ID() string {
	return q.properties["id"].(string)
}

func (q *workflowDefinitionQuery) SetID(id string) WorkflowDefinitionQueryInterface {
	q.properties["id"] = id
	return q
}

func (q *workflowDefinitionQuery) HasIDIn() bool {
	return q.hasProperty("id_in")
}

func (q *workflowDefinitionQuery) IDIn() []string {
	return q.properties["id_in"].([]string)
}

func (q *workflowDefinitionQuery) SetIDIn(idIn []string) WorkflowDefinitionQueryInterface {
	q.properties["id_in"] = idIn
	return q
}

func (q *workflowDefinitionQuery) HasLimit() bool {
	return q.hasProperty("limit")
}

func (q *workflowDefinitionQuery) Limit() int {
	return q.properties["limit"].(int)
}

func (q *workflowDefinitionQuery) SetLimit(limit int) WorkflowDefinitionQueryInterface {
	q.properties["limit"] = limit
	return q
}

func (q *workflowDefinitionQuery) HasOffset() bool {
	return q.hasProperty("offset")
}

func (q *workflowDefinitionQuery) Offset() int {
	return q.properties["offset"].(int)
}

func (q *workflowDefinitionQuery) SetOffset(offset int) WorkflowDefinitionQueryInterface {
	q.properties["offset"] = offset
	return q
}

func (q *workflowDefinitionQuery) HasOrderBy() bool {
	return q.hasProperty("order_by")
}

func (q *workflowDefinitionQuery) OrderBy() string {
	return q.properties["order_by"].(string)
}

func (q *workflowDefinitionQuery) SetOrderBy(orderBy string) WorkflowDefinitionQueryInterface {
	q.properties["order_by"] = orderBy
	return q
}

func (q *workflowDefinitionQuery) HasSortOrder() bool {
	return q.hasProperty("sort_order")
}

func (q *workflowDefinitionQuery) SortOrder() string {
	return q.properties["sort_order"].(string)
}

func (q *workflowDefinitionQuery) SetSortOrder(sortOrder string) WorkflowDefinitionQueryInterface {
	q.properties["sort_order"] = sortOrder
	return q
}

func (q *workflowDefinitionQuery) HasStatus() bool {
	return q.hasProperty("status")
}

func (q *workflowDefinitionQuery) Status() string {
	return q.properties["status"].(string)
}

func (q *workflowDefinitionQuery) SetStatus(status string) WorkflowDefinitionQueryInterface {
	q.properties["status"] = status
	return q
}

func (q *workflowDefinitionQuery) HasStatusIn() bool {
	return q.hasProperty("status_in")
}

func (q *workflowDefinitionQuery) StatusIn() []string {
	return q.properties["status_in"].([]string)
}

func (q *workflowDefinitionQuery) SetStatusIn(statusIn []string) WorkflowDefinitionQueryInterface {
	q.properties["status_in"] = statusIn
	return q
}

func (q *workflowDefinitionQuery) HasUpdatedAtGte() bool {
	return q.hasProperty("updated_at_gte")
}

func (q *workflowDefinitionQuery) UpdatedAtGte() string {
	return q.properties["updated_at_gte"].(string)
}

func (q *workflowDefinitionQuery) SetUpdatedAtGte(updatedAtGte string) WorkflowDefinitionQueryInterface {
	q.properties["updated_at_gte"] = updatedAtGte
	return q
}

func (q *workflowDefinitionQuery) HasUpdatedAtLte() bool {
	return q.hasProperty("updated_at_lte")
}

func (q *workflowDefinitionQuery) UpdatedAtLte() string {
	return q.properties["updated_at_lte"].(string)
}

func (q *workflowDefinitionQuery) SetUpdatedAtLte(updatedAtLte string) WorkflowDefinitionQueryInterface {
	q.properties["updated_at_lte"] = updatedAtLte
	return q
}

func (q *workflowDefinitionQuery) SoftDeletedIncluded() bool {
	if q.hasProperty("softDeletedIncluded") {
		return q.properties["softDeletedIncluded"].(bool)
	}

	return false
}

func (q *workflowDefinitionQuery) SetSoftDeletedIncluded(softDeletedIncluded bool) WorkflowDefinitionQueryInterface {
	q.properties["softDeletedIncluded"] = softDeletedIncluded
	return q
}

func (c *workflowDefinitionQuery) hasProperty(name string) bool {
	_, ok := c.properties[name]
	return ok
}
