package providers

import (
	"sort"

	internalRegistry "parsdevkit.net/internal/registry"
)

var registry = internalRegistry.New[string, ProviderInterface]()

func Register(m ProviderInterface) {
	name := m.GetKey()
	registry.Register(name, m)
}

func Get(name string) ProviderInterface {
	return registry.Get(name)
}

func All() []ProviderInterface {
	return registry.All()
}
func AllSorted() []ProviderInterface {
	all := registry.All()
	sort.SliceStable(all, func(i, j int) bool {
		return all[i].GetConfig().Order < all[j].GetConfig().Order
	})
	return all
}

// Order’a göre büyükten küçüğe
func AllSortedReverse() []ProviderInterface {
	all := registry.All()
	sort.SliceStable(all, func(i, j int) bool {
		return all[i].GetConfig().Order > all[j].GetConfig().Order
	})
	return all
}
