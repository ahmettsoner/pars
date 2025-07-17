package contracts

import "parsdevkit.net/components/template"

type TemplateContextProviderInterface interface {
	BaseContextProviderInterface
	Context(source template.ContextProviderSource) interface{}
}
