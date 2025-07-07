package flow

import (
	"context"
)

type Flow struct {
	Name  string
	Steps []Step
}

func NewFlow(name string) *Flow {
	return &Flow{
		Name:  name,
		Steps: []Step{},
	}
}

func (f *Flow) Step(s Step) *Flow {
	f.Steps = append(f.Steps, s)
	return f
}

func (f *Flow) Run(ctx context.Context, fc *FlowContext) error {
	fc.Log("Starting flow: %s", f.Name)

	for i, step := range f.Steps {
		fc.Log("Running step: %s", step.Name())
		if err := step.Run(ctx, fc); err != nil {
			fc.AddError(err)
			fc.Log("Step failed: %s - %v", step.Name(), err)

			for j := i - 1; j >= 0; j-- {
				fc.Log("Compensating: %s", f.Steps[j].Name())
				_ = f.Steps[j].Compensate(ctx, fc)
			}
			return err
		}
		fc.Log("Step completed: %s", step.Name())
	}

	fc.Log("Flow completed: %s", f.Name)
	return nil
}
