package example

import (
	"testing"

	"github.com/gouniverse/workflow"
)

func TestWorkflow(t *testing.T) {
	// Create a new workflow
	// wf := NewWorkflow()

	// // Set the title of the workflow
	// wf.SetTitle("Example Workflow")

	// // Add a step to the workflow
	// wf.AddStep(NewStep("Step 1", func() {
	// 	t.Log("Step 1")
	// }))

	// // Add a step to the workflow
	// wf.AddStep(NewStep("Step 2", func() {
	// 	t.Log("Step 2")
	// }))

	// // Run the workflow
	// wf.Run()

	wf := workflow.NewWorkflow()
	wf.SetName("Example Workflow")
}
