package contracts

import "parsdevkit.net/components/template"

type WorkspaceContextProviderInterface interface {
	Context(source template.ContextProviderSource) any
}
