package platforms

import (
	"fmt"
	"sort"
)

var platformRegistry = make(map[string]PlatformInterface)

func Register(m PlatformInterface) {
	name := m.GetKey()
	if _, exists := platformRegistry[name]; exists {
		panic(fmt.Sprintf("Platform %s already registered", name))
	}
	platformRegistry[name] = m
}

func Get(name string) (PlatformInterface, error) {
	result, ok := platformRegistry[name]
	if !ok {
		return nil, fmt.Errorf("no platform found for %s", name)
	}

	return result, nil
}

func All() []PlatformInterface {
	all := []PlatformInterface{}
	for _, m := range platformRegistry {
		all = append(all, m)
	}
	return all
}
func AllSorted() []PlatformInterface {
	all := All()
	sort.SliceStable(all, func(i, j int) bool {
		return all[i].GetConfig().Order < all[j].GetConfig().Order
	})
	return all
}

// Order’a göre büyükten küçüğe
func AllSortedReverse() []PlatformInterface {
	all := All()
	sort.SliceStable(all, func(i, j int) bool {
		return all[i].GetConfig().Order > all[j].GetConfig().Order
	})
	return all
}
