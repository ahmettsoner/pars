package contracts

import "parsdevkit.net/components/template"

type ProjectContextProviderInterface interface {
	BaseContextProviderInterface
	Context(source template.ContextProviderSource) interface{}
}
