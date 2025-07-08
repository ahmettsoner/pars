package flowx

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
func NewContextWithData(data map[string]any) *FlowContext {
	return &FlowContext{
		Data:   data,
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

func (fc *FlowContext) Set(key string, value any) {
	fc.mu.Lock()
	defer fc.mu.Unlock()
	fc.Data[key] = value
}

func (fc *FlowContext) Get(key string) (any, bool) {
	fc.mu.Lock()
	defer fc.mu.Unlock()
	val, ok := fc.Data[key]
	return val, ok
}

func Set[T any](fc *FlowContext, key string, value T) {
	fc.Set(key, value)
}

func Get[T any](fc *FlowContext, key string) (T, bool) {
	val, ok := fc.Get(key)
	if !ok {
		var zero T
		return zero, false
	}

	casted, ok := val.(T)
	if !ok {
		var zero T
		return zero, false
	}
	return casted, true
}
