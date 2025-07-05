package platforms

import (
	"fmt"

	"parsdevkit.net/application/schemas"
	"parsdevkit.net/models"
)

var platformRegistry = make(map[models.PlatformType]interface{})

func Register[T schemas.SchemaInterface](m PlatformInterface[T]) {
	name := m.GetKey()
	if _, exists := platformRegistry[name]; exists {
		panic(fmt.Sprintf("Platform %s already registered", name))
	}
	platformRegistry[name] = m
}

func Get[T schemas.SchemaInterface](name models.PlatformType) PlatformInterface[T] {
	result, ok := platformRegistry[name]
	if !ok {
		panic(fmt.Errorf("no platform found for %s", name))
	}

	return result.(PlatformInterface[T])
}

func All[T schemas.SchemaInterface]() []PlatformInterface[T] {
	all := []PlatformInterface[T]{}
	for _, m := range platformRegistry {
		all = append(all, m.(PlatformInterface[T]))
	}
	return all
}
