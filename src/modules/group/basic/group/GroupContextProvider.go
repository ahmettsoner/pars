package basic_group

import (
	"parsdevkit.net/application/platforms"
	applicationGroup "parsdevkit.net/application/structs/group"
	"parsdevkit.net/components/template"
	"parsdevkit.net/models"
	"parsdevkit.net/modules/group/basic_group_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type GroupContextProvider struct {
	environment string
}

func NewGroupContextProvider(environment string) basic_group_contract.ContextProviderInterface {

	return &GroupContextProvider{
		environment: environment}
}

func (s *GroupContextProvider) Context(source template.ContextProviderSource) interface{} {

	if model, ok := source.Project.(application_project_payload_structs.ProjectBaseStruct); ok {
		return GroupComposite{
			Group:    s.structToModel(model.Specifications.Platform.Type, model.Specifications.GroupObject),
			Original: model.Specifications.GroupObject,
		}
	}

	return nil
}
func (s *GroupContextProvider) structToModel(platform models.PlatformType, model applicationGroup.GroupIdentifier /*basic_group_payload_structs.GroupBaseStruct*/) Group {
	dependencies := model.Package
	manager := platforms.Get[application_project_payload_structs.ProjectBaseStruct](platform)

	var result Group = Group{
		Package: manager.PrintDependencies(dependencies),
		Name:    model.Name,
	}
	return result
}

type GroupComposite struct {
	Group
	Original applicationGroup.GroupIdentifier /*basic_group_payload_structs.GroupBaseStruct*/
}

type Group struct {
	Name    string
	Package string
}
