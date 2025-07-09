package flowx

type BaseStep struct{}

func (BaseStep) IgnoreError() bool {
	return false
}
