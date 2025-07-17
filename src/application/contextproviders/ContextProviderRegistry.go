package contextproviders

import (
	internalRegistry "parsdevkit.net/internal/registry"

	"parsdevkit.net/application/contracts"
)

var registry = internalRegistry.New[string, contracts.BaseContextProviderInterface]()

func Register(m contracts.BaseContextProviderInterface) {
	name := m.GetConfig().Name
	registry.Register(name, m)
}

func Get(name string) contracts.BaseContextProviderInterface {
	return registry.Get(name)
}

func All() []contracts.BaseContextProviderInterface {
	return registry.All()
}
