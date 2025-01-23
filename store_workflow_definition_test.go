package workflow

import (
	"context"
	"testing"

	"github.com/gouniverse/sb"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

func Test_Store_WorkflowDefinitionCreate(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("WorkflowDefinitionCreate: Error", err.Error())
	}

	workflowDefinition := NewWorkflowDefinition().
		SetName("Workflow 1")

	err = store.WorkflowDefinitionCreate(context.Background(), workflowDefinition)
	if err != nil {
		t.Fatal("WorkflowDefinitionCreate: Error in Creating WorkflowDefinition: received ", err.Error())
	}
}

func Test_Store_WorkflowDefinitionDeleteByID(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("WorkflowDefinitionDeleteByID: Error", err.Error())
	}

	workflowDefinition := NewWorkflowDefinition().
		SetName("Workflow 1")

	err = store.WorkflowDefinitionCreate(context.Background(), workflowDefinition)
	if err != nil {
		t.Fatal("WorkflowDefinitionDeleteByID: Error in Creating WorkflowDefinition: received ", err.Error())
	}

	err = store.WorkflowDefinitionDeleteByID(context.Background(), workflowDefinition.ID())
	if err != nil {
		t.Fatal("WorkflowDefinitionDeleteByID: Error in Deleting WorkflowDefinition: received ", err.Error())
	}
}

func Test_Store_WorkflowDefinitionFindByID(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("WorkflowDefinitionFindByID: Error", err.Error())
	}

	workflowDefinition := NewWorkflowDefinition().
		SetName("Workflow 1")

	err = store.WorkflowDefinitionCreate(context.Background(), workflowDefinition)
	if err != nil {
		t.Fatal("WorkflowDefinitionFindByID: Error in Creating WorkflowDefinition: received ", err.Error())
	}

	foundWorkflowDefinition, err := store.WorkflowDefinitionFindByID(context.Background(), workflowDefinition.ID())
	if err != nil {
		t.Fatal("WorkflowDefinitionFindByID: Error in Finding WorkflowDefinition: received ", err.Error())
	}

	if foundWorkflowDefinition == nil {
		t.Fatal("WorkflowDefinitionFindByID: Error in Finding WorkflowDefinition: received ", foundWorkflowDefinition)
	}

	if foundWorkflowDefinition.ID() != workflowDefinition.ID() {
		t.Fatal("WorkflowDefinitionFindByID: Error in Finding WorkflowDefinition: received ", foundWorkflowDefinition)
	}
}

func Test_Store_WorkflowDefinitionList(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal(err.Error())
	}

	ids := []string{}
	names := []string{}

	for i := 0; i < 10; i++ {
		workflowDefinition := NewWorkflowDefinition().
			SetName("Workflow " + cast.ToString(i))

		err = store.WorkflowDefinitionCreate(context.Background(), workflowDefinition)
		if err != nil {
			t.Fatal("WorkflowDefinitionList: Error in Creating WorkflowDefinition: received ", err.Error())
		}

		ids = append(ids, workflowDefinition.ID())
		names = append(names, workflowDefinition.Name())
	}

	list, err := store.WorkflowDefinitionList(context.Background(), WorkflowDefinitionQuery().SetLimit(1))
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(list) != 1 {
		t.Fatal(`Expected "1" but received "`, len(list), `"`)
	}

	if list[0].ID() != ids[0] {
		t.Fatal(`Expected "`, ids[0], `" but received "`, list[0].ID(), `"`)
	}

	list, err = store.WorkflowDefinitionList(context.Background(), WorkflowDefinitionQuery().
		SetOrderBy(COLUMN_NAME).
		SetSortOrder(sb.DESC).
		SetLimit(10))

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(list) != 10 {
		t.Fatal(`Expected "10" but received "`, len(list), `"`)
	}

	t.Log(list)

	for i, v := range lo.Reverse(list) {
		if v.Name() != names[i] {
			t.Fatal(`Expected "`, names[i], `" but received "`, v.Name(), `"`)
		}
	}
}

// func Test_Store_QueueSoftDeleteByID(t *testing.T) {
// 	store, err := InitStore("test_queue_soft_delete_by_id.db")
// 	if err != nil {
// 		t.Fatalf("QueueSoftDeleteByID: Error[%v]", err)
// 	}

// 	queuedTask := NewQueue().
// 		SetTaskID("TASK_01").
// 		SetAttempts(1)

// 	query := store.SqlCreateQueueTable()
// 	if strings.Contains(query, "unsupported driver") {
// 		t.Fatalf("QueueSoftDeleteByID: UnExpected Query, received [%v]", query)
// 	}

// 	_, err = store.db.Exec(query)
// 	if err != nil {
// 		t.Fatalf("QueueSoftDeleteByID: Table creation error: [%v]", err)
// 	}

// 	err = store.QueueCreate(queuedTask)
// 	if err != nil {
// 		t.Fatalf("QueueSoftDeleteByID: Error in Creating Queue: received [%v]", err)
// 	}

// 	err = store.QueueSoftDeleteByID(queuedTask.ID())
// 	if err != nil {
// 		t.Fatalf("QueueSoftDeleteByID: Error in Fail Queue: received [%v]", err)
// 	}

// 	queueFound, err := store.QueueFindByID(queuedTask.ID())

// 	if err != nil {
// 		t.Fatal("QueueSoftDeleteByID: Error in QueueFindByID: received:", err)
// 	}

// 	if queueFound != nil {
// 		t.Fatal("QueueSoftDeleteByID: QueueFindByID should be nil, received:", queueFound)
// 	}
// }

func Test_Store_QueueSoftDelete(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("QueueSoftDelete: Error", err.Error())
	}

	workflowDefinition := NewWorkflowDefinition().
		SetName("Workflow 1")

	err = store.WorkflowDefinitionCreate(context.Background(), workflowDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	workflowDefinition, err = store.WorkflowDefinitionFindByID(context.Background(), workflowDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if workflowDefinition == nil {
		t.Fatal(`Expected workflowDefinition but received "nil"`)
	}

	err = store.WorkflowDefinitionSoftDelete(context.Background(), workflowDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	workflowDefinition, err = store.WorkflowDefinitionFindByID(context.Background(), workflowDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if workflowDefinition != nil {
		t.Fatal(`Expected "nil" but received workflowDefinition`)
	}
}

func Test_Store_QueueSoftDeleteByID(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("QueueSoftDeleteByID: Error", err.Error())
	}

	workflowDefinition := NewWorkflowDefinition().
		SetName("Workflow 1")

	err = store.WorkflowDefinitionCreate(context.Background(), workflowDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	workflowDefinition, err = store.WorkflowDefinitionFindByID(context.Background(), workflowDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if workflowDefinition == nil {
		t.Fatal(`Expected workflowDefinition but received "nil"`)
	}

	err = store.WorkflowDefinitionSoftDeleteByID(context.Background(), workflowDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	workflowDefinition, err = store.WorkflowDefinitionFindByID(context.Background(), workflowDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if workflowDefinition != nil {
		t.Fatal(`Expected "nil" but received workflowDefinition`)
	}
}

func Test_Store_WorkflowDefinitionUpdate(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("WorkflowDefinitionUpdate: Error", err.Error())
	}

	workflowDefinition := NewWorkflowDefinition().
		SetName("Workflow 1")

	err = store.WorkflowDefinitionCreate(context.Background(), workflowDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	workflowDefinition, err = store.WorkflowDefinitionFindByID(context.Background(), workflowDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if workflowDefinition == nil {
		t.Fatal(`Expected workflowDefinition but received "nil"`)
	}

	workflowDefinition.SetName("Workflow 2")

	err = store.WorkflowDefinitionUpdate(context.Background(), workflowDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	workflowDefinition, err = store.WorkflowDefinitionFindByID(context.Background(), workflowDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if workflowDefinition == nil {
		t.Fatal(`Expected workflowDefinition but received "nil"`)
	}

	if workflowDefinition.Name() != "Workflow 2" {
		t.Fatal(`Expected "Workflow 2" but received workflowDefinition`)
	}
}
