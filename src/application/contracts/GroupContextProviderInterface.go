package contracts

import "parsdevkit.net/components/template"

type GroupContextProviderInterface interface {
	BaseContextProviderInterface
	Context(source template.ContextProviderSource) interface{}
}
