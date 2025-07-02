package ioc

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type Container struct {
	providers map[reflect.Type]reflect.Value
	lock      sync.RWMutex
}

func NewContainer() *Container {
	return &Container{
		providers: make(map[reflect.Type]reflect.Value),
	}
}

// Concrete veya struct tipi için basit register
func (c *Container) Register(provider interface{}) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	fn := reflect.ValueOf(provider)
	fnType := fn.Type()

	if fnType.Kind() != reflect.Func {
		return errors.New("provider must be a function")
	}

	if fnType.NumOut() != 1 {
		return errors.New("provider must return exactly one value")
	}

	returnType := fnType.Out(0)
	c.providers[returnType] = fn
	return nil
}

// Interface’e karşılık bir concrete register etmek
func (c *Container) RegisterInterface(interfaceType reflect.Type, provider interface{}) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	fn := reflect.ValueOf(provider)
	fnType := fn.Type()

	if fnType.Kind() != reflect.Func {
		return errors.New("provider must be a function")
	}

	if fnType.NumOut() != 1 {
		return errors.New("provider must return exactly one value")
	}

	returnType := fnType.Out(0)

	if !returnType.Implements(interfaceType) {
		return fmt.Errorf("return type %s does not implement %s", returnType, interfaceType)
	}

	c.providers[interfaceType] = fn
	return nil
}

func (c *Container) Get(t any) (any, error) {
	typ := reflect.TypeOf(t)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	v, err := c.resolve(typ)
	if err != nil {
		return nil, err
	}
	return v.Interface(), nil
}

func (c *Container) resolve(t reflect.Type) (reflect.Value, error) {
	c.lock.RLock()
	provider, ok := c.providers[t]
	c.lock.RUnlock()

	if !ok {
		return reflect.Value{}, fmt.Errorf("no provider registered for %s", t)
	}

	providerType := provider.Type()
	args := make([]reflect.Value, providerType.NumIn())

	for i := 0; i < providerType.NumIn(); i++ {
		depType := providerType.In(i)
		depValue, err := c.resolve(depType)
		if err != nil {
			return reflect.Value{}, err
		}
		args[i] = depValue
	}

	results := provider.Call(args)
	return results[0], nil
}
