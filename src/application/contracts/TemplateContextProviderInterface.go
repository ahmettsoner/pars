package contracts

import "parsdevkit.net/components/template"

type TemplateContextProviderInterface interface {
	Context(source template.ContextProviderSource) any
}
