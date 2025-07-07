package platforms

import (
	"parsdevkit.net/application/schemas"
	internalRegistry "parsdevkit.net/internal/registry"
	"parsdevkit.net/models"
)

var registry = internalRegistry.New[models.PlatformType, interface{}]()

func Register[T schemas.SchemaInterface](m PlatformInterface[T]) {
	name := m.GetKey()
	registry.Register(name, m)
}

func Get[T schemas.SchemaInterface](name models.PlatformType) PlatformInterface[T] {
	return registry.Get(name).(PlatformInterface[T])
}

func All[T schemas.SchemaInterface]() []PlatformInterface[T] {
	var result []PlatformInterface[T]
	for _, val := range registry.All() {
		result = append(result, val.(PlatformInterface[T]))
	}
	return result
}
