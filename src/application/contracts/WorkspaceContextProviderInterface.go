package contracts

import "parsdevkit.net/components/template"

type WorkspaceContextProviderInterface interface {
	BaseContextProviderInterface
	Context(source template.ContextProviderSource) interface{}
}
