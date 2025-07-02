package schemas

import (
	"fmt"
)

var schemaRegistry = make(map[string]func() SchemaInterface)

func Register(m SchemaInterface) {
	name := m.GetKey()
	if _, exists := schemaRegistry[name]; exists {
		panic(fmt.Sprintf("Schema %s already registered", name))
	}
	schemaRegistry[name] = func() SchemaInterface { return m }
}

func Get(name string) (SchemaInterface, error) {
	result, ok := schemaRegistry[name]
	if !ok {
		return nil, fmt.Errorf("no schema found for Type/Kind %s", name)
	}

	return result(), nil
}

func All() []SchemaInterface {
	all := []SchemaInterface{}
	for _, m := range schemaRegistry {
		all = append(all, m())
	}
	return all
}
