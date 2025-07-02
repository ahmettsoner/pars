package engines

import (
	"fmt"
	"sort"
)

var engineRegistry = make(map[string]EngineInterface)

func Register(m EngineInterface) {
	name := m.GetConfig().Name
	if _, exists := engineRegistry[name]; exists {
		panic(fmt.Sprintf("Engine %s already registered", name))
	}
	engineRegistry[name] = m
}

func Get(name string) (EngineInterface, error) {
	result, ok := engineRegistry[name]
	if !ok {
		return nil, fmt.Errorf("no engine found for %s", name)
	}

	return result, nil
}

func All() []EngineInterface {
	all := []EngineInterface{}
	for _, m := range engineRegistry {
		all = append(all, m)
	}
	return all
}
func AllSorted() []EngineInterface {
	all := All()
	sort.SliceStable(all, func(i, j int) bool {
		return all[i].GetConfig().Order < all[j].GetConfig().Order
	})
	return all
}

// Order’a göre büyükten küçüğe
func AllSortedReverse() []EngineInterface {
	all := All()
	sort.SliceStable(all, func(i, j int) bool {
		return all[i].GetConfig().Order > all[j].GetConfig().Order
	})
	return all
}
