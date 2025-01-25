package workflow

import (
	"context"
	"testing"

	"github.com/gouniverse/sb"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

func Test_Store_EdgeDefinitionCreate(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("EdgeDefinitionCreate: Error", err.Error())
	}

	edgeDefinition := NewEdgeDefinition().
		SetFromStepDefinitionID("STEP_DEFINITION_01").
		SetToStepDefinitionID("STEP_DEFINITION_02")

	err = store.EdgeDefinitionCreate(context.Background(), edgeDefinition)
	if err != nil {
		t.Fatal("EdgeDefinitionCreate: Error in Creating EdgeDefinition: received ", err.Error())
	}
}

func Test_Store_EdgeDefinitionDeleteByID(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("EdgeDefinitionDeleteByID: Error", err.Error())
	}

	edgeDefinition := NewEdgeDefinition().
		SetFromStepDefinitionID("STEP_DEFINITION_01").
		SetToStepDefinitionID("STEP_DEFINITION_02")

	err = store.EdgeDefinitionCreate(context.Background(), edgeDefinition)
	if err != nil {
		t.Fatal("EdgeDefinitionDeleteByID: Error in Creating EdgeDefinition: received ", err.Error())
	}

	err = store.EdgeDefinitionDeleteByID(context.Background(), edgeDefinition.ID())
	if err != nil {
		t.Fatal("EdgeDefinitionDeleteByID: Error in Deleting EdgeDefinition: received ", err.Error())
	}
}

func Test_Store_EdgeDefinitionFindByID(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("EdgeDefinitionFindByID: Error", err.Error())
	}

	edgeDefinition := NewEdgeDefinition().
		SetFromStepDefinitionID("STEP_DEFINITION_01").
		SetToStepDefinitionID("STEP_DEFINITION_02")

	err = store.EdgeDefinitionCreate(context.Background(), edgeDefinition)
	if err != nil {
		t.Fatal("EdgeDefinitionFindByID: Error in Creating EdgeDefinition: received ", err.Error())
	}

	foundEdgeDefinition, err := store.EdgeDefinitionFindByID(context.Background(), edgeDefinition.ID())
	if err != nil {
		t.Fatal("EdgeDefinitionFindByID: Error in Finding EdgeDefinition: received ", err.Error())
	}

	if foundEdgeDefinition == nil {
		t.Fatal("EdgeDefinitionFindByID: Error in Finding EdgeDefinition: received ", foundEdgeDefinition)
	}

	if foundEdgeDefinition.ID() != edgeDefinition.ID() {
		t.Fatal("EdgeDefinitionFindByID: Error in Finding EdgeDefinition: received ", foundEdgeDefinition)
	}
}

func Test_Store_EdgeDefinitionList(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal(err.Error())
	}

	ids := []string{}
	names := []string{}

	for i := 0; i < 10; i++ {
		edgeDefinition := NewEdgeDefinition().
			SetFromStepDefinitionID("STEP_DEFINITION_" + cast.ToString(i+1)).
			SetToStepDefinitionID("STEP_DEFINITION_" + cast.ToString(i+2))

		err = store.EdgeDefinitionCreate(context.Background(), edgeDefinition)
		if err != nil {
			t.Fatal("EdgeDefinitionList: Error in Creating EdgeDefinition: received ", err.Error())
		}

		ids = append(ids, edgeDefinition.ID())
		names = append(names, edgeDefinition.FromStepDefinitionID())
	}

	list, err := store.EdgeDefinitionList(context.Background(), EdgeDefinitionQuery().SetLimit(1))
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(list) != 1 {
		t.Fatal(`Expected "1" but received "`, len(list), `"`)
	}

	if list[0].ID() != ids[0] {
		t.Fatal(`Expected "`, ids[0], `" but received "`, list[0].ID(), `"`)
	}

	list, err = store.EdgeDefinitionList(context.Background(), EdgeDefinitionQuery().
		SetOrderBy(COLUMN_ID).
		SetSortOrder(sb.DESC).
		SetLimit(10))

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(list) != 10 {
		t.Fatal(`Expected "10" but received "`, len(list), `"`)
	}

	for i, v := range lo.Reverse(list) {
		if v.FromStepDefinitionID() != names[i] {
			t.Fatal(`Expected "`, names[i], `" but received "`, v.FromStepDefinitionID(), `"`)
		}
	}
}

func Test_Store_EdgeDefinitionSoftDelete(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("QueueSoftDelete: Error", err.Error())
	}

	edgeDefinition := NewEdgeDefinition().
		SetFromStepDefinitionID("STEP_DEFINITION_01").
		SetToStepDefinitionID("STEP_DEFINITION_02")

	err = store.EdgeDefinitionCreate(context.Background(), edgeDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	edgeDefinition, err = store.EdgeDefinitionFindByID(context.Background(), edgeDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if edgeDefinition == nil {
		t.Fatal(`Expected edgeDefinition but received "nil"`)
	}

	err = store.EdgeDefinitionSoftDelete(context.Background(), edgeDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	edgeDefinition, err = store.EdgeDefinitionFindByID(context.Background(), edgeDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if edgeDefinition != nil {
		t.Fatal(`Expected "nil" but received edgeDefinition`)
	}
}

func Test_Store_EdgeDefinitionSoftDeleteByID(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("QueueSoftDeleteByID: Error", err.Error())
	}

	edgeDefinition := NewEdgeDefinition().
		SetFromStepDefinitionID("STEP_DEFINITION_01").
		SetToStepDefinitionID("STEP_DEFINITION_02")

	err = store.EdgeDefinitionCreate(context.Background(), edgeDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	edgeDefinition, err = store.EdgeDefinitionFindByID(context.Background(), edgeDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if edgeDefinition == nil {
		t.Fatal(`Expected edgeDefinition but received "nil"`)
	}

	err = store.EdgeDefinitionSoftDeleteByID(context.Background(), edgeDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	edgeDefinition, err = store.EdgeDefinitionFindByID(context.Background(), edgeDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if edgeDefinition != nil {
		t.Fatal(`Expected "nil" but received edgeDefinition`)
	}
}

func Test_Store_EdgeDefinitionUpdate(t *testing.T) {
	store, err := initStore()
	if err != nil {
		t.Fatal("EdgeDefinitionUpdate: Error", err.Error())
	}

	edgeDefinition := NewEdgeDefinition().
		SetFromStepDefinitionID("STEP_DEFINITION_01").
		SetToStepDefinitionID("STEP_DEFINITION_02")

	err = store.EdgeDefinitionCreate(context.Background(), edgeDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	edgeDefinition, err = store.EdgeDefinitionFindByID(context.Background(), edgeDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if edgeDefinition == nil {
		t.Fatal(`Expected edgeDefinition but received "nil"`)
	}

	edgeDefinition.SetFromStepDefinitionID("STEP_DEFINITION_03")
	edgeDefinition.SetToStepDefinitionID("STEP_DEFINITION_04")

	err = store.EdgeDefinitionUpdate(context.Background(), edgeDefinition)
	if err != nil {
		t.Fatal(err.Error())
	}

	edgeDefinition, err = store.EdgeDefinitionFindByID(context.Background(), edgeDefinition.ID())
	if err != nil {
		t.Fatal(err.Error())
	}

	if edgeDefinition == nil {
		t.Fatal(`Expected edgeDefinition but received "nil"`)
	}

	if edgeDefinition.FromStepDefinitionID() != "STEP_DEFINITION_03" {
		t.Fatal(`Expected "STEP_DEFINITION_03" but received "`, edgeDefinition.FromStepDefinitionID(), `"`)
	}

	if edgeDefinition.ToStepDefinitionID() != "STEP_DEFINITION_04" {
		t.Fatal(`Expected "STEP_DEFINITION_04" but received "`, edgeDefinition.ToStepDefinitionID(), `"`)
	}
}
