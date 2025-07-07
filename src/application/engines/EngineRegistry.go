package engines

import (
	"sort"

	internalRegistry "parsdevkit.net/internal/registry"
)

var registry = internalRegistry.New[string, EngineInterface]()

func Register(m EngineInterface) {
	name := m.GetConfig().Name
	registry.Register(name, m)
}

func Get(name string) EngineInterface {
	return registry.Get(name)
}

func All() []EngineInterface {
	return registry.All()
}
func AllSorted() []EngineInterface {
	all := registry.All()
	sort.SliceStable(all, func(i, j int) bool {
		return all[i].GetConfig().Order < all[j].GetConfig().Order
	})
	return all
}

func AllSortedReverse() []EngineInterface {
	all := registry.All()
	sort.SliceStable(all, func(i, j int) bool {
		return all[i].GetConfig().Order > all[j].GetConfig().Order
	})
	return all
}
