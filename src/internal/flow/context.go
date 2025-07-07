package flow

import (
	"fmt"
	"sync"
)

type FlowContext struct {
	Data   map[string]any
	Logs   []string
	Errors []error
	mu     sync.Mutex
}

func NewContext() *FlowContext {
	return &FlowContext{
		Data:   map[string]any{},
		Logs:   []string{},
		Errors: []error{},
	}
}

func (fc *FlowContext) Log(format string, args ...any) {
	fc.mu.Lock()
	defer fc.mu.Unlock()
	msg := fmt.Sprintf(format, args...)
	fc.Logs = append(fc.Logs, msg)
	fmt.Println("[FLOW]", msg)
}

func (fc *FlowContext) AddError(err error) {
	fc.mu.Lock()
	defer fc.mu.Unlock()
	fc.Errors = append(fc.Errors, err)
}
