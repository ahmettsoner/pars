package ioc

import (
	"reflect"
)

var c = NewContainer()

func Register(provider interface{}) error {
	return c.Register(provider)
}
func RegisterInterface[T any](provider interface{}) error {
	t := reflect.TypeOf((*T)(nil)).Elem()
	return c.RegisterInterface(t, provider)
}

func Get[T any]() (T, error) {
	var zero T
	v, err := c.Get((*T)(nil))
	if err != nil {
		return zero, err
	}
	return v.(T), nil
}
