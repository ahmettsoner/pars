package code_template

import (
	"parsdevkit.net/application/platforms"
	"parsdevkit.net/components/template"
	"parsdevkit.net/models"
	"parsdevkit.net/modules/template/code_template_contract"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"

	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	"parsdevkit.net/modules/resource/object_resource"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/models/label"
)

type CodeTemplateContextProvider struct {
	environment string
}

func NewCodeTemplateContextProvider(environment string) code_template_contract.ContextProviderInterface {
	return &CodeTemplateContextProvider{
		environment: environment,
	}
}

func (s CodeTemplateContextProvider) GetConfig() contracts.ContextProviderConfig {
	return contracts.ContextProviderConfig{
		Name: code_template_payload_structs.MODULE_KEY,
	}
}
func (s *CodeTemplateContextProvider) Context(source template.ContextProviderSource) interface{} {

	if model, ok := source.Template.(code_template_payload_structs.TemplateBaseStruct); ok {
		if modelProject, ok := source.Project.(application_project_payload_structs.ProjectBaseStruct); ok {
			return TemplateComposite{
				CodeTemplate: s.structToModel(modelProject.Specifications.Platform.Type, model),
				Original:     model,
			}
		}
	}

	return nil
}
func (s *CodeTemplateContextProvider) structToModel(platform models.PlatformType, model code_template_payload_structs.TemplateBaseStruct) CodeTemplate {
	dependencies := model.Specifications.Package
	manager := platforms.Get[application_project_payload_structs.ProjectBaseStruct](platform)

	var result CodeTemplate = CodeTemplate{
		Package: manager.PrintDependencies(dependencies),
		Labels:  s.LabelListToModel(model.Specifications.Labels...),
	}
	return result
}
func (s *CodeTemplateContextProvider) LabelListToModel(labels ...label.Label) []object_resource.ObjectLabel {

	var result []object_resource.ObjectLabel = make([]object_resource.ObjectLabel, 0)

	for _, label := range labels {
		result = append(result, object_resource.ObjectLabel{
			Key:   label.Key,
			Value: label.Value,
		})
	}

	return result
}

type TemplateComposite struct {
	CodeTemplate
	Original code_template_payload_structs.TemplateBaseStruct
}

type CodeTemplate struct {
	Name    string
	Package string
	Labels  []object_resource.ObjectLabel
	// Options []ObjectOption
	// Layers     []ObjectLayer
}
