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

func Get(name string) SchemaInterface {
	result, ok := schemaRegistry[name]
	if !ok {
		panic(fmt.Errorf("no schema found for Type/Kind %s", name))
	}

	return result()
}

func All() []SchemaInterface {
	all := []SchemaInterface{}
	for _, m := range schemaRegistry {
		all = append(all, m())
	}
	return all
}
