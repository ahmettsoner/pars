package contextgenerator

import (
	"parsdevkit.net/application/contextproviders"
	"parsdevkit.net/application/contracts"

	layerPkg "parsdevkit.net/application/models/layer"
	sectionPkg "parsdevkit.net/application/models/section"
	"parsdevkit.net/components/template"
)

func NewTemplateDataContext(source template.ContextProviderSource) map[string]any {

	var result map[string]any = make(map[string]any)

	if source.Workspace != nil {
		contextProvider := contextproviders.ContextProviderFactory[contracts.WorkspaceContextProviderInterface](source.Workspace)
		result["Workspace"] = contextProvider.Context(source)
	}
	if source.Group != nil {
		contextProvider := contextproviders.ContextProviderFactory[contracts.GroupContextProviderInterface](source.Group)
		result["Group"] = contextProvider.Context(source)
	}
	if source.Project != nil {
		contextProvider := contextproviders.ContextProviderFactory[contracts.ProjectContextProviderInterface](source.Project)
		result["Project"] = contextProvider.Context(source)
	}
	if source.Resource != nil {
		contextProvider := contextproviders.ContextProviderFactory[contracts.ResourceContextProviderInterface](source.Resource)
		result["Resource"] = contextProvider.Context(source)
	}
	if source.Template != nil {
		contextProvider := contextproviders.ContextProviderFactory[contracts.TemplateContextProviderInterface](source.Template)
		result["Template"] = contextProvider.Context(source)
	}

	if source.Resource != nil {
		if (source.Layer != layerPkg.LayerIdentifier{}) {
			contextProvider := contextproviders.ContextProviderFactory[contracts.ResourceContextProviderInterface](source.Resource)
			result["Layer"] = contextProvider.LayerToModelContext(source)
		}
		if (source.Section != sectionPkg.SectionIdentifier{}) {
			contextProvider := contextproviders.ContextProviderFactory[contracts.ResourceContextProviderInterface](source.Resource)
			result["Section"] = contextProvider.SectionToModelContext(source)
		}
	}

	return result
}
