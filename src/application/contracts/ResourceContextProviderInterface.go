package contracts

import (
	"parsdevkit.net/components/template"
)

type ResourceContextProviderInterface interface {
	Context(source template.ContextProviderSource) any
	SectionToModelContext(source template.ContextProviderSource) any
	LayerToModelContext(source template.ContextProviderSource) any
}
