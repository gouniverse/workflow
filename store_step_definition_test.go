package workflow

import (
	"context"
	"testing"

	"github.com/gouniverse/sb"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

func Test_Store_StepDefinitionCreate(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("StepDefinitionCreate: Error", err.Error())
	}

	stepDefinition := NewStepDefinition().
		SetWorkflowDefinitionID("WORKFLOW_DEFINITION_O1").
		SetName("Step 1")

	err = store.StepDefinitionCreate(context.Background(), stepDefinition)
	if err != nil {
		t.Fatal("StepDefinitionCreate: Error in Creating StepDefinition: received ", err.Error())
	}
}

func Test_Store_StepDefinitionDeleteByID(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("StepDefinitionDeleteByID: Error", err.Error())
	}

	stepDefinition := NewStepDefinition().
		SetWorkflowDefinitionID("WORKFLOW_DEFINITION_O1").
		SetName("Step 1")

	err = store.StepDefinitionCreate(context.Background(), stepDefinition)
	if err != nil {
		t.Fatal("StepDefinitionDeleteByID: Error in Creating StepDefinition: received ", err.Error())
	}

	err = store.StepDefinitionDeleteByID(context.Background(), stepDefinition.ID())
	if err != nil {
		t.Fatal("StepDefinitionDeleteByID: Error in Deleting StepDefinition: received ", err.Error())
	}
}

func Test_Store_StepDefinitionFindByID(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("StepDefinitionFindByID: Error", err.Error())
	}

	stepDefinition := NewStepDefinition().
		SetWorkflowDefinitionID("WORKFLOW_DEFINITION_O1").
		SetName("Step 1")

	err = store.StepDefinitionCreate(context.Background(), stepDefinition)
	if err != nil {
		t.Fatal("StepDefinitionFindByID: Error in Creating StepDefinition: received ", err.Error())
	}

	foundStepDefinition, err := store.StepDefinitionFindByID(context.Background(), stepDefinition.ID())
	if err != nil {
		t.Fatal("StepDefinitionFindByID: Error in Finding StepDefinition: received ", err.Error())
	}

	if foundStepDefinition == nil {
		t.Fatal("StepDefinitionFindByID: Error in Finding StepDefinition: received ", foundStepDefinition)
	}

	if foundStepDefinition.ID() != stepDefinition.ID() {
		t.Fatal("StepDefinitionFindByID: Error in Finding StepDefinition: received ", foundStepDefinition)
	}
}

func Test_Store_StepDefinitionList(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal(err.Error())
	}

	ids := []string{}
	names := []string{}

	for i := 0; i < 10; i++ {
		stepDefinition := NewStepDefinition().
			SetWorkflowDefinitionID("WORKFLOW_DEFINITION_O1").
			SetName("Step " + cast.ToString(i))

		err = store.StepDefinitionCreate(context.Background(), stepDefinition)
		if err != nil {
			t.Fatal("StepDefinitionList: Error in Creating StepDefinition: received ", err.Error())
		}

		ids = append(ids, stepDefinition.ID())
		names = append(names, stepDefinition.Name())
	}

	list, err := store.StepDefinitionList(context.Background(), StepDefinitionQuery().SetLimit(1))
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(list) != 1 {
		t.Fatal(`Expected "1" but received "`, len(list), `"`)
	}

	if list[0].ID() != ids[0] {
		t.Fatal(`Expected "`, ids[0], `" but received "`, list[0].ID(), `"`)
	}

	list, err = store.StepDefinitionList(context.Background(), StepDefinitionQuery().
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

func Test_Store_StepDefinitionSoftDelete(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("QueueSoftDelete: Error", err.Error())
	}

	stepDefinition := NewStepDefinition().
		SetName("Workflow 1")

	err = store.StepDefinitionCreate(context.Background(), stepDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	stepDefinition, err = store.StepDefinitionFindByID(context.Background(), stepDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if stepDefinition == nil {
		t.Fatal(`Expected stepDefinition but received "nil"`)
	}

	err = store.StepDefinitionSoftDelete(context.Background(), stepDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	stepDefinition, err = store.StepDefinitionFindByID(context.Background(), stepDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if stepDefinition != nil {
		t.Fatal(`Expected "nil" but received stepDefinition`)
	}
}

func Test_Store_StepDefinitionSoftDeleteByID(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("QueueSoftDeleteByID: Error", err.Error())
	}

	stepDefinition := NewStepDefinition().
		SetWorkflowDefinitionID("WORKFLOW_DEFINITION_O1").
		SetName("Step 1")

	err = store.StepDefinitionCreate(context.Background(), stepDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	stepDefinition, err = store.StepDefinitionFindByID(context.Background(), stepDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if stepDefinition == nil {
		t.Fatal(`Expected stepDefinition but received "nil"`)
	}

	err = store.StepDefinitionSoftDeleteByID(context.Background(), stepDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	stepDefinition, err = store.StepDefinitionFindByID(context.Background(), stepDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if stepDefinition != nil {
		t.Fatal(`Expected "nil" but received stepDefinition`)
	}
}

func Test_Store_StepDefinitionUpdate(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("StepDefinitionUpdate: Error", err.Error())
	}

	stepDefinition := NewStepDefinition().
		SetName("Workflow 1")

	err = store.StepDefinitionCreate(context.Background(), stepDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	stepDefinition, err = store.StepDefinitionFindByID(context.Background(), stepDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if stepDefinition == nil {
		t.Fatal(`Expected stepDefinition but received "nil"`)
	}

	stepDefinition.SetName("Workflow 2")

	err = store.StepDefinitionUpdate(context.Background(), stepDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	stepDefinition, err = store.StepDefinitionFindByID(context.Background(), stepDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if stepDefinition == nil {
		t.Fatal(`Expected stepDefinition but received "nil"`)
	}

	if stepDefinition.Name() != "Workflow 2" {
		t.Fatal(`Expected "Workflow 2" but received stepDefinition`)
	}
}
