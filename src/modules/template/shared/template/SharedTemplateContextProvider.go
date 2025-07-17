package shared_task

import (
	"parsdevkit.net/components/template"
	"parsdevkit.net/models"
	"parsdevkit.net/modules/resource/object_resource"

	"parsdevkit.net/application/models/label"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	shared_template_payload_structs "parsdevkit.net/modules/template/shared_template_payload/structs"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/modules/template/shared_template_contract"
)

type SharedTemplateContextProvider struct {
	environment string
}

func NewSharedTemplateContextProvider(environment string) shared_template_contract.ContextProviderInterface {
	return &SharedTemplateContextProvider{
		environment: environment,
	}
}

func (s SharedTemplateContextProvider) GetConfig() contracts.ContextProviderConfig {
	return contracts.ContextProviderConfig{
		Name: shared_template_payload_structs.MODULE_KEY,
	}
}
func (s *SharedTemplateContextProvider) Context(source template.ContextProviderSource) interface{} {

	if model, ok := source.Template.(shared_template_payload_structs.TemplateBaseStruct); ok {
		if modelProject, ok := source.Project.(application_project_payload_structs.ProjectBaseStruct); ok {
			return TemplateComposite{
				SharedTemplate: s.structToModel(modelProject.Specifications.Platform.Type, model),
				Original:       model,
			}
		}
	}

	return nil
}
func (s *SharedTemplateContextProvider) structToModel(platform models.PlatformType, model shared_template_payload_structs.TemplateBaseStruct) SharedTemplate {

	var result SharedTemplate = SharedTemplate{
		Name: model.Header.Name,
	}
	return result
}
func (s *SharedTemplateContextProvider) LabelListToModel(labels ...label.Label) []object_resource.ObjectLabel {

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
	SharedTemplate
	Original shared_template_payload_structs.TemplateBaseStruct
}

type SharedTemplate struct {
	Name string
}
