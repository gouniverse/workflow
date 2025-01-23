package workflow

import (
	"os"
	"testing"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func initDB(filepath string) *sql.DB {
	if filepath != ":memory:" && fileExists(filepath) {
		err := os.Remove(filepath) // remove database

		if err != nil {
			panic(err)
		}
	}

	dsn := filepath + "?parseTime=true"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}

	return db
}

func initStore() (StoreInterface, error) {
	db := initDB(":memory:")

	store, err := NewStore(NewStoreOptions{
		WorkflowDefinitionTableName: "workflow_definition",
		StepDefinitionTableName:     "step_definition",
		EdgeDefinitionTableName:     "edge_definition",
		WorkflowTableName:           "workflow",
		StepTableName:               "step",
		EdgeTableName:               "edge",
		DB:                          db,
		AutomigrateEnabled:          true,
	})

	if err != nil {
		return nil, err
	}

	return store, nil
}

func TestWithAutoMigrate(t *testing.T) {
	db := initDB(":memory:")

	storeAutomigrateFalse, errAutomigrateFalse := NewStore(NewStoreOptions{
		WorkflowDefinitionTableName: "workflow_definition_with_automigrate_false",
		StepDefinitionTableName:     "step_definition_with_automigrate_false",
		EdgeDefinitionTableName:     "edge_definition_with_automigrate_false",
		WorkflowTableName:           "workflow_with_automigrate_false",
		StepTableName:               "step_with_automigrate_false",
		EdgeTableName:               "edge_with_automigrate_false",
		DB:                          db,
		AutomigrateEnabled:          false,
	})

	if errAutomigrateFalse != nil {
		t.Fatalf("automigrateEnabled: Expected [err] to be nill received [%v]", errAutomigrateFalse.Error())
	}

	if storeAutomigrateFalse.AutoMigrateEnabled() != false {
		t.Fatalf("automigrateEnabled: Expected [false] received [%v]", storeAutomigrateFalse.AutoMigrateEnabled())
	}

	storeAutomigrateTrue, errAutomigrateTrue := NewStore(NewStoreOptions{
		WorkflowDefinitionTableName: "workflow_definition_with_automigrate_true",
		StepDefinitionTableName:     "step_definition_with_automigrate_true",
		EdgeDefinitionTableName:     "edge_definition_with_automigrate_true",
		WorkflowTableName:           "workflow_with_automigrate_true",
		StepTableName:               "step_with_automigrate_true",
		EdgeTableName:               "edge_with_automigrate_true",
		DB:                          db,
		AutomigrateEnabled:          true,
	})

	if errAutomigrateTrue != nil {
		t.Fatalf("automigrateEnabled: Expected [err] to be nill received [%v]", errAutomigrateTrue.Error())
	}

	if storeAutomigrateTrue.AutoMigrateEnabled() != true {
		t.Fatalf("automigrateEnabled: Expected [true] received [%v]", storeAutomigrateTrue.AutoMigrateEnabled())
	}
}

// func Test_Store_AutoMigrate(t *testing.T) {
// 	db := InitDB("test_vaultstore_automigrate")

// 	store, err := NewStore(NewStoreOptions{
// 		VaultTableName:     "vault_automigrate",
// 		DB:                 db,
// 		AutomigrateEnabled: false,
// 	})

// 	if err != nil {
// 		t.Fatalf("automigrateEnabled: Expected [err] to be nill received [%v]", err.Error())
// 	}

// 	if store.automigrateEnabled != false {
// 		t.Fatalf("automigrateEnabled: Expected [false] received [%v]", store.automigrateEnabled)
// 	}

// 	err = store.AutoMigrate()

// 	if err != nil {
// 		t.Fatalf("AutoMigrate Failure [%v]", err.Error())
// 	}

// 	if store.vaultTableName != "vault_automigrate" {
// 		t.Fatalf("Expected vaultTableName [vault_automigrate] received [%v]", store.vaultTableName)
// 	}
// 	if store.db == nil {
// 		t.Fatalf("DB Init Failure")
// 	}
// 	if store.automigrateEnabled != false {
// 		t.Fatalf("Failure:  AutoMigrate")
// 	}
// }
