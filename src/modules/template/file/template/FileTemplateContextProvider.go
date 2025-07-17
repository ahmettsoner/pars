package file_template

import (
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/platforms"
	"parsdevkit.net/components/template"
	"parsdevkit.net/models"
	"parsdevkit.net/modules/resource/object_resource"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"

	"parsdevkit.net/application/contracts"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	"parsdevkit.net/modules/template/file_template_contract"
)

type FileTemplateContextProvider struct {
	environment string
}

func NewFileTemplateContextProvider(environment string) file_template_contract.ContextProviderInterface {
	return &FileTemplateContextProvider{
		environment: environment,
	}
}
func (s FileTemplateContextProvider) GetConfig() contracts.ContextProviderConfig {
	return contracts.ContextProviderConfig{
		Name: file_template_payload_structs.MODULE_KEY,
	}
}

func (s *FileTemplateContextProvider) Context(source template.ContextProviderSource) interface{} {

	if model, ok := source.Template.(file_template_payload_structs.TemplateBaseStruct); ok {
		if modelProject, ok := source.Project.(application_project_payload_structs.ProjectBaseStruct); ok {
			return TemplateComposite{
				FileTemplate: s.structToModel(modelProject.Specifications.Platform.Type, model),
				Original:     model,
			}
		}
	}

	return nil
}
func (s *FileTemplateContextProvider) structToModel(platform models.PlatformType, model file_template_payload_structs.TemplateBaseStruct) FileTemplate {
	dependencies := model.Specifications.Package
	manager := platforms.Get[application_project_payload_structs.ProjectBaseStruct](platform)

	var result FileTemplate = FileTemplate{
		Package: manager.PrintDependencies(dependencies),
		Labels:  s.LabelListToModel(model.Specifications.Labels...),
	}
	return result
}
func (s *FileTemplateContextProvider) LabelListToModel(labels ...label.Label) []object_resource.ObjectLabel {

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
	FileTemplate
	Original file_template_payload_structs.TemplateBaseStruct
}

type FileTemplate struct {
	Name    string
	Package string
	Labels  []object_resource.ObjectLabel
	// Options []ObjectOption
	// Layers     []ObjectLayer
}
