package flowx

import "context"

type Step interface {
	Name() string
	Run(ctx context.Context, fc *FlowContext) error
	Compensate(ctx context.Context, fc *FlowContext) error
	IgnoreError() bool
}
