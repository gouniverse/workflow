package workflow

// func TestW(t *testing.T) {
// 	workflow := (&Workflow{}).SetID("wf1").SetTitle("Workflow 1")

// 	step1 := (&Step{}).SetID("step1").SetTitle("Step 1")
// 	step2 := (&Step{}).SetID("step2").SetTitle("Step 2")
// 	step3 := (&Step{}).SetID("step3").SetTitle("Step 3")

// 	step21 := (&Step{}).SetID("step21").SetTitle("Step 2.1")
// 	step22 := (&Step{}).SetID("step21").SetTitle("Step 2.2")

// 	step2.AddStep(*step21)
// 	step2.AddStep(*step22)

// 	// step22 := &Step{
// 	// 	ID:         "step22",
// 	// 	WorkflowID: workflow.GetID(),
// 	// 	Title:      "Step 2.2",
// 	// }

// 	workflow.AddStep(*step1)
// 	workflow.AddStep(*step2)
// 	workflow.AddStep(*step3)

// 	// t.Error(step1.WorkflowID, step1.ParentID)

// 	workflow.Run()

// }
