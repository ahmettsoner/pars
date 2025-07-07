package ioc

import (
	"reflect"

	internalIOC "parsdevkit.net/internal/ioc"
)

var c = internalIOC.NewContainer()

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
