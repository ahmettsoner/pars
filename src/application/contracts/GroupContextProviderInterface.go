package contracts

import "parsdevkit.net/components/template"

type GroupContextProviderInterface interface {
	Context(source template.ContextProviderSource) any
}
