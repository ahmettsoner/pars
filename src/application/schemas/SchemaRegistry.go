package schemas

import (
	"fmt"
	"reflect"
	"sync"
)

var (
	schemaRegistry = make(map[string]func() SchemaInterface)
	mu             sync.RWMutex
)

// Register bir schema implementasyonunu register eder
func Register(schema SchemaInterface) {
	key := schema.GetKey()

	mu.Lock()
	defer mu.Unlock()

	if _, exists := schemaRegistry[key]; exists {
		panic(fmt.Sprintf("schema with key %q already registered", key))
	}

	// Tipin pointer versiyonunun sıfır değerini oluşturur
	schemaRegistry[key] = func() SchemaInterface {
		// schema -> *T
		typ := reflect.TypeOf(schema)
		if typ.Kind() != reflect.Ptr {
			panic(fmt.Sprintf("schema %T must be a pointer type", schema))
		}
		return reflect.New(typ.Elem()).Interface().(SchemaInterface)
	}
}

// Get, belirtilen key'e karşılık gelen yeni bir schema instance'ı döner
func Get(key string) SchemaInterface {
	mu.RLock()
	defer mu.RUnlock()

	constructor, exists := schemaRegistry[key]
	if !exists {
		panic(fmt.Sprintf("no schema registered for key: %q", key))
	}
	return constructor()
}

// All, kayıtlı tüm schema türlerinin yeni örneklerini döner
func All() []SchemaInterface {
	mu.RLock()
	defer mu.RUnlock()

	list := make([]SchemaInterface, 0, len(schemaRegistry))
	for _, constructor := range schemaRegistry {
		list = append(list, constructor())
	}
	return list
}
