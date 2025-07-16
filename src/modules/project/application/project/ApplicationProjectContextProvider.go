package application_project

import (
	"parsdevkit.net/application/models/label"

	"parsdevkit.net/application/platforms"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"

	"parsdevkit.net/components/template"
	"parsdevkit.net/modules/project/application_project_contract"
)

type ApplicationProjectContextProvider struct {
	environment string
}

func NewApplicationProjectContextProvider(environment string) application_project_contract.ContextProviderInterface {

	return &ApplicationProjectContextProvider{
		environment: environment}
}

func (s *ApplicationProjectContextProvider) Context(source template.ContextProviderSource) interface{} {

	if model, ok := source.Project.(application_project_payload_structs.ProjectBaseStruct); ok {
		return ApplicationProjectComposite{
			ApplicationProject: s.structToModel(model),
			Original:           model,
		}
	}

	return nil
}
func (s *ApplicationProjectContextProvider) structToModel(model application_project_payload_structs.ProjectBaseStruct) ApplicationProject {
	dependencies := model.Specifications.GetAllPackage()
	manager := platforms.Get[application_project_payload_structs.ProjectBaseStruct](model.Specifications.Platform.Type)

	var result ApplicationProject = ApplicationProject{
		Name:    model.Header.Name,
		Package: manager.PrintDependencies(dependencies),
		Labels:  s.LabelListToModel(model.Specifications.Labels...),
	}

	return result
}

func (s *ApplicationProjectContextProvider) LabelListToModel(labels ...label.Label) []Label {

	var result []Label = make([]Label, 0)

	for _, label := range labels {
		result = append(result, Label{
			Key:   label.Key,
			Value: label.Value,
		})
	}

	return result
}

type ApplicationProjectComposite struct {
	ApplicationProject
	Original application_project_payload_structs.ProjectBaseStruct
}

type ApplicationProject struct {
	Name    string
	Package string
	Labels  []Label
}

type Label struct {
	Key   string
	Value string
}
