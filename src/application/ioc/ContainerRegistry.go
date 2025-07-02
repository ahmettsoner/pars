package ioc

import (
	"reflect"
)

var c = NewContainer()

func Register(provider interface{}) {
	c.Register(provider)
}
func RegisterInterface[T any](provider interface{}) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	c.RegisterInterface(t, provider)
}

func Get[T any]() T {
	v := c.Get((*T)(nil))
	return v.(T)
}
