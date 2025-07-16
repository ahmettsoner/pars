package contracts

import "parsdevkit.net/components/template"

type ProjectContextProviderInterface interface {
	Context(source template.ContextProviderSource) any
}
