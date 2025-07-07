```go

type ValidateStep struct{}

func (s *ValidateStep) Name() string { return "ValidateStep" }

func (s *ValidateStep) Run(ctx context.Context, fc *flow.FlowContext) error {
	email, ok := fc.Data["email"].(string)
	if !ok || email == "" {
		return errors.New("invalid email")
	}
	return nil
}

func (s *ValidateStep) Compensate(ctx context.Context, fc *flow.FlowContext) error {
	fc.Log("Nothing to rollback for validation")
	return nil
}

type SaveStep struct{}

func (s *SaveStep) Name() string { return "SaveStep" }

func (s *SaveStep) Run(ctx context.Context, fc *flow.FlowContext) error {
	email := fc.Data["email"].(string)
	fc.Log("User %s saved to DB", email)
	return nil
}

func (s *SaveStep) Compensate(ctx context.Context, fc *flow.FlowContext) error {
	email := fc.Data["email"].(string)
	fc.Log("Rollback: deleting user %s", email)
	return nil
}

projectFlow := flow.NewFlow("UserSignup").
	Step(&ValidateStep{}).
	Step(&SaveStep{})

ctx2 := context.Background()
fc := flow.NewContext()

	if err := projectFlow.Run(ctx2, fc); err != nil {
		fc.Log("Flow failed: %v", err)
	}

```