package providers

import (
	"fmt"
	"sort"
)

var providerRegistry = make(map[string]ProviderInterface)

func Register(m ProviderInterface) {
	name := m.GetKey()
	if _, exists := providerRegistry[name]; exists {
		panic(fmt.Sprintf("Provider %s already registered", name))
	}
	providerRegistry[name] = m
}

func Get(name string) (ProviderInterface, error) {
	result, ok := providerRegistry[name]
	if !ok {
		return nil, fmt.Errorf("no provider found for %s", name)
	}

	return result, nil
}

func All() []ProviderInterface {
	all := []ProviderInterface{}
	for _, m := range providerRegistry {
		all = append(all, m)
	}
	return all
}
func AllSorted() []ProviderInterface {
	all := All()
	sort.SliceStable(all, func(i, j int) bool {
		return all[i].GetConfig().Order < all[j].GetConfig().Order
	})
	return all
}

// Order’a göre büyükten küçüğe
func AllSortedReverse() []ProviderInterface {
	all := All()
	sort.SliceStable(all, func(i, j int) bool {
		return all[i].GetConfig().Order > all[j].GetConfig().Order
	})
	return all
}
