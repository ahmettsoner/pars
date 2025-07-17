package contracts

import (
	"parsdevkit.net/components/template"
)

type ResourceContextProviderInterface interface {
	BaseContextProviderInterface
	Context(source template.ContextProviderSource) interface{}
	SectionToModelContext(source template.ContextProviderSource) interface{}
	LayerToModelContext(source template.ContextProviderSource) interface{}
}
