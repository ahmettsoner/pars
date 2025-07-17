package contextproviders

import (
	"parsdevkit.net/application/contracts"

	"parsdevkit.net/application/schemas"
)

func ContextProviderFactory[T contracts.BaseContextProviderInterface](t schemas.SchemaInterface) T {
	var zero T
	header := t.GetHeader()
	header_key := header.GetKey()
	for _, contextProvider := range registry.All() {
		if contextProvider == nil {
			continue
		}
		key := contextProvider.GetConfig().Name
		if key == header_key {
			typed, ok := contextProvider.(T)
			if ok {
				return typed
			}
		}
	}
	return zero
}
