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

	// Parametreleri kontrol et
	for i := 0; i < fnType.NumIn(); i++ {
		depType := fnType.In(i)
		if _, ok := c.providers[depType]; !ok {
			return fmt.Errorf("dependency %s not registered for provider %s", depType, fnType)
		}
	}

	returnType := fnType.Out(0)
	c.providers[returnType] = fn
	return nil
}

// Interface’e karşılık bir concrete register etmek
func (c *Container) RegisterInterface(interfaceType reflect.Type, provider interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	fn := reflect.ValueOf(provider)
	fnType := fn.Type()

	if fnType.Kind() != reflect.Func {
		panic(errors.New("provider must be a function"))
	}

	if fnType.NumOut() != 1 {
		panic(errors.New("provider must return exactly one value"))
	}

	returnType := fnType.Out(0)

	if !returnType.Implements(interfaceType) {
		panic(fmt.Errorf("return type %s does not implement %s", returnType, interfaceType))
	}

	c.providers[interfaceType] = fn
}

func (c *Container) Get(t any) any {
	typ := reflect.TypeOf(t)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return c.resolve(typ, make(map[reflect.Type]bool)).Interface()
}

func (c *Container) Invoke(function interface{}) reflect.Value {
	fn := reflect.ValueOf(function)
	fnType := fn.Type()

	if fnType.Kind() != reflect.Func {
		panic(fmt.Errorf("Invoke requires a function"))
	}

	args := make([]reflect.Value, fnType.NumIn())

	for i := 0; i < fnType.NumIn(); i++ {
		paramType := fnType.In(i)

		// Özel parametre (string gibi) için ayrı işlem yapılabilir
		if paramType.Kind() == reflect.String {
			// Örneğin environment string parametresini dışarıdan geçebilirsin
			// Burada örnek olarak boş string veriyoruz
			args[i] = reflect.ValueOf("")
			continue
		}

		instance := c.GetByType(paramType)

		args[i] = instance
	}

	results := fn.Call(args)
	return results[0]
}

func (c *Container) GetByType(t reflect.Type) reflect.Value {
	return c.resolve(t, make(map[reflect.Type]bool))
}
func (c *Container) resolve(t reflect.Type, visited map[reflect.Type]bool) reflect.Value {
	if visited[t] {
		panic(fmt.Errorf("circular dependency detected on type %s", t))
	}

	c.lock.RLock()
	provider, ok := c.providers[t]
	c.lock.RUnlock()

	if !ok {
		panic(fmt.Errorf("no provider registered for %s", t))
	}

	visited[t] = true
	defer func() {
		delete(visited, t)
	}()

	providerType := provider.Type()
	args := make([]reflect.Value, providerType.NumIn())

	for i := 0; i < providerType.NumIn(); i++ {
		depType := providerType.In(i)
		depValue := c.resolve(depType, visited)
		args[i] = depValue
	}

	results := provider.Call(args)
	return results[0]
}
