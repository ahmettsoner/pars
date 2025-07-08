package flowx

import (
	"context"
	"fmt"
)

type Flow struct {
	Name  string
	Steps []Step
	ctx   context.Context
}

func NewFlow(name string) *Flow {
	return &Flow{
		Name:  name,
		Steps: []Step{},
		ctx:   context.Background(),
	}
}

func (f *Flow) WithContext(ctx context.Context) *Flow {
	f.ctx = ctx
	return f
}

func (f *Flow) Step(s Step) *Flow {
	f.Steps = append(f.Steps, s)
	return f
}

func (f *Flow) Run(fc *FlowContext) error {
	ctx := f.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	return f.runInternal(ctx, fc)
}

func (f *Flow) RunWithContext(ctx context.Context, fc *FlowContext) error {
	return f.runInternal(ctx, fc)
}

func (f *Flow) runInternal(ctx context.Context, fc *FlowContext) (err error) {
	fc.Log("Starting flow: %s", f.Name)

	var executedSteps []Step // <-- çalıştırılmış adımlar burada tutulacak

	defer func() {
		if r := recover(); r != nil {
			pErr := fmt.Errorf("panic in flow '%s': %v", f.Name, r)
			fc.AddError(pErr)
			fc.Log("⚠️ Panic occurred: %v", r)

			// Sadece çalışmış adımları tersten compensate et
			for j := len(executedSteps) - 1; j >= 0; j-- {
				step := executedSteps[j]
				fc.Log("🧨 Compensating due to panic: %s", step.Name())
				_ = step.Compensate(ctx, fc)
			}

			err = pErr
		}
	}()

	for _, step := range f.Steps {
		fc.Log("Running step: %s", step.Name())

		if err := step.Run(ctx, fc); err != nil {
			fc.AddError(err)
			fc.Log("❌ Step failed: %s - %v", step.Name(), err)

			// rollback
			for j := len(executedSteps) - 1; j >= 0; j-- {
				rollbackStep := executedSteps[j]
				fc.Log("↩️ Compensating: %s", rollbackStep.Name())
				_ = rollbackStep.Compensate(ctx, fc)
			}
			return err
		}

		fc.Log("✅ Step completed: %s", step.Name())
		executedSteps = append(executedSteps, step) // sadece başarılı step’leri track et
	}

	fc.Log("🎉 Flow completed successfully: %s", f.Name)
	return nil
}
