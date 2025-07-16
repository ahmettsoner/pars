package template

import (
	"parsdevkit.net/application/schemas"

	"parsdevkit.net/application/models/section"

	"parsdevkit.net/application/models/layer"
)

type ContextProviderSource struct {
	Workspace schemas.SchemaInterface
	Group     schemas.SchemaInterface
	Project   schemas.SchemaInterface
	Resource  schemas.SchemaInterface
	Template  schemas.SchemaInterface
	Layer     layer.LayerIdentifier
	Section   section.SectionIdentifier
}

func NewContextProviderSource(workspace, group, project, resource, template schemas.SchemaInterface, _layer layer.LayerIdentifier, _section section.SectionIdentifier) ContextProviderSource {
	return ContextProviderSource{
		Workspace: workspace,
		Group:     group,
		Project:   project,
		Resource:  resource,
		Template:  template,
		Layer:     _layer,
		Section:   _section,
	}
}
