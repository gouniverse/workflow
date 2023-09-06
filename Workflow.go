package workflow

import "log"

type Workflow struct {
	id    string
	title string
	steps []Step
}

func (wf *Workflow) ID() string {
	return wf.id
}

func (wf *Workflow) SetID(id string) *Workflow {
	wf.id = id
	return wf
}

func (wf *Workflow) Title() string {
	return wf.title
}

func (wf *Workflow) SetTitle(title string) *Workflow {
	wf.title = title
	return wf
}

func (wf *Workflow) Steps() []Step {
	return wf.steps
}

func (wf *Workflow) SetSteps(steps []Step) *Workflow {
	wf.steps = steps
	return wf
}

func (wf *Workflow) AddStep(step Step) *Workflow {
	wf.steps = append(wf.steps, step)
	return wf
}

func (wf *Workflow) Run() {
	steps := wf.Steps()
	log.Println("START:", wf.Title())
	for _, step := range steps {
		step.Run()
	}
	log.Println("END:", wf.Title())
}
