package services

import (
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	"parsdevkit.net/application/structs"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"

	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"
	_string "parsdevkit.net/pkg/utilities/string"
	"parsdevkit.net/platforms/core"

	"parsdevkit.net/components/template/models/objectResources"
)

type ObjectResourceService struct {
	manager core.ApplicationPlatformManagerInterface
}

func NewObjectResourceService(manager core.ApplicationPlatformManagerInterface) ObjectResourceService {
	return ObjectResourceService{
		manager: manager,
	}
}

func (s *ObjectResourceService) DataTypeToImport(_type structs.DataType, importsMap map[string][]string) map[string][]string {

	if !_string.IsEmpty(_type.Package.Name) {
		if _type.Category == structs.DataTypeCategories.Reference {
			if aliases, exists := importsMap[_type.Package.Name]; exists {
				isNew := true
				for _, alias := range aliases {
					if _type.Package.Alias == alias {
						isNew = false
						break
					}
				}
				if isNew {
					importsMap[_type.Package.Name] = append(aliases, _type.Package.Alias)
				}
			} else {
				if !_string.IsEmpty(_type.Package.Alias) {
					importsMap[_type.Package.Name] = []string{_type.Package.Alias}
				} else {
					importsMap[_type.Package.Name] = []string{}
				}
			}
		}
		if _type.Generics != nil {
			for _, generic := range _type.Generics {
				importsMap = s.DataTypeToImport(generic, importsMap)
			}
		}
	}

	return importsMap
}

func (s *ObjectResourceService) ResourceToModel(resource object_resource_payload_structs.ResourceBaseStruct, project application_project_payload_structs.ProjectBaseStruct, layer object_resource_payload_structs.Layer, template code_template_payload_structs.TemplateBaseStruct) objectResources.ObjectResource {
	var importsMap map[string][]string = make(map[string][]string)

	packages := project.Specifications.GetAllPackageWithLayer(layer.Name)
	packages = append(packages, resource.Specifications.Package...)
	packages = append(packages, template.Specifications.Package...)

	var dataAttributes = make([]objectResources.ObjectResourceAttribute, 0)
	for _, attribute := range resource.Object.Attributes {
		var dataAttribute = objectResources.ObjectResourceAttribute{
			Name:         attribute.Name,
			Type:         s.manager.PrintDataType(attribute.Type),
			TypePackage:  attribute.Type.Name,
			TypeCategory: string(attribute.Type.Category),
			Visibility:   s.manager.PrintVisibility(attribute.Visibility),
			Labels:       s.LabelListToModel(attribute.Labels...),
			Options:      s.OptionListToModel(attribute.Options...),
			Common:       attribute.Common,
		}
		dataAttributes = append(dataAttributes, dataAttribute)

		importsMap = s.DataTypeToImport(attribute.Type, importsMap)
	}

	var dataMethods = make([]objectResources.ObjectResourceMethod, 0)
	for _, method := range resource.Object.Methods {
		var dataMethodParameters = make([]objectResources.ObjectResourceMethodParameter, 0)
		for _, methodParameter := range method.Parameters {
			var dataMethodParameter = objectResources.ObjectResourceMethodParameter{
				Name: methodParameter.Name,
				Type: s.manager.PrintDataType(methodParameter.Type),
			}

			dataMethodParameters = append(dataMethodParameters, dataMethodParameter)
			importsMap = s.DataTypeToImport(methodParameter.Type, importsMap)
		}

		var dataMethodReturnTypes = make([]string, 0)
		for _, methodReturnType := range method.ReturnTypes {
			var dataMethodReturnType = s.manager.PrintDataType(methodReturnType)

			dataMethodReturnTypes = append(dataMethodReturnTypes, dataMethodReturnType)
			importsMap = s.DataTypeToImport(methodReturnType, importsMap)
		}
		var dataMethod = objectResources.ObjectResourceMethod{
			Name:        method.Name,
			Visibility:  s.manager.PrintVisibility(method.Visibility),
			Parameters:  dataMethodParameters,
			ReturnTypes: dataMethodReturnTypes,
			Labels:      s.LabelListToModel(method.Labels...),
			Options:     s.OptionListToModel(method.Options...),
			Code:        method.Code,
			Common:      method.Common,
		}

		dataMethods = append(dataMethods, dataMethod)
	}

	var imports []objectResources.ObjectResourceImport = make([]objectResources.ObjectResourceImport, 0)

	for key, value := range importsMap {
		imports = append(imports, objectResources.ObjectResourceImport{
			Package: key,
			Aliases: value,
		})
	}

	return objectResources.ObjectResource{
		Name:       resource.Specifications.Name,
		Package:    s.manager.PrintDependencies(packages),
		Labels:     s.LabelListToModel(resource.Specifications.Labels...),
		Layers:     s.LayerListToModel(resource, project, template, resource.Specifications.Layers...),
		Dictionary: s.DictionaryListToModel(resource.Object.Dictionary...),
		Groups:     s.GroupListToModel(resource.Object.Groups...),
		Attributes: dataAttributes,
		Methods:    dataMethods,
		Imports:    imports,
	}
}

func (s *ObjectResourceService) DataResourceToModel(resource data_resource_payload_structs.ResourceSpecification, resourceData data_resource_payload_structs.ResourceData, project application_project_payload_structs.ProjectSpecification, layer string, template file_template_payload_structs.TemplateSpecification) objectResources.DataResource {
	var importsMap map[string][]string = make(map[string][]string)

	packages := project.GetAllPackageWithLayer(layer)
	packages = append(packages, template.Package...)

	var imports []objectResources.ObjectResourceImport = make([]objectResources.ObjectResourceImport, 0)

	for key, value := range importsMap {
		imports = append(imports, objectResources.ObjectResourceImport{
			Package: key,
			Aliases: value,
		})
	}

	return objectResources.DataResource{
		Name:    resource.Name,
		Package: s.manager.PrintDependencies(packages),
		Labels:  s.LabelListToModel(resource.Labels...),
		Layers:  s.DataLayerListToModel(resource, project, template, resource.Layers...),
		// Dictionary: s.DictionaryListToModel(resource.Dictionary...),
		// Groups:     s.GroupListToModel(resource.Groups...),
		Data: resourceData.Data,
	}
}

func (s *ObjectResourceService) ObjectSectionToModel(resource object_resource_payload_structs.ResourceBaseStruct, project application_project_payload_structs.ProjectBaseStruct, layer object_resource_payload_structs.Layer, template code_template_payload_structs.TemplateBaseStruct, section object_resource_payload_structs.Section) objectResources.ObjectSection {
	var importsMap map[string][]string = make(map[string][]string)

	packages := project.Specifications.GetAllPackageWithLayer(layer.Name)
	packages = append(packages, resource.Specifications.Package...)
	packages = append(packages, template.Specifications.Package...)

	var dataAttributes = make([]objectResources.ObjectResourceAttribute, 0)
	for _, attribute := range resource.Object.Attributes {
		for _, sectionAttribute := range section.Attributes {
			if attribute.Name == sectionAttribute {
				var dataAttribute = objectResources.ObjectResourceAttribute{
					Name:         attribute.Name,
					Type:         s.manager.PrintDataType(attribute.Type),
					TypePackage:  attribute.Type.Name,
					TypeCategory: string(attribute.Type.Category),
					Visibility:   s.manager.PrintVisibility(attribute.Visibility),
					Labels:       s.LabelListToModel(attribute.Labels...),
					Options:      s.OptionListToModel(attribute.Options...),
					Common:       attribute.Common,
				}
				dataAttributes = append(dataAttributes, dataAttribute)

				importsMap = s.DataTypeToImport(attribute.Type, importsMap)
			}
		}
	}

	var dataMethods = make([]objectResources.ObjectResourceMethod, 0)
	for _, method := range resource.Object.Methods {
		var dataMethodParameters = make([]objectResources.ObjectResourceMethodParameter, 0)
		for _, sectionMethod := range section.Methods {
			if method.Name == sectionMethod {
				for _, methodParameter := range method.Parameters {
					var dataMethodParameter = objectResources.ObjectResourceMethodParameter{
						Name: methodParameter.Name,
						Type: s.manager.PrintDataType(methodParameter.Type),
					}

					dataMethodParameters = append(dataMethodParameters, dataMethodParameter)
					importsMap = s.DataTypeToImport(methodParameter.Type, importsMap)
				}

				var dataMethodReturnTypes = make([]string, 0)
				for _, methodReturnType := range method.ReturnTypes {
					var dataMethodReturnType = s.manager.PrintDataType(methodReturnType)

					dataMethodReturnTypes = append(dataMethodReturnTypes, dataMethodReturnType)
					importsMap = s.DataTypeToImport(methodReturnType, importsMap)
				}
				var dataMethod = objectResources.ObjectResourceMethod{
					Name:        method.Name,
					Visibility:  s.manager.PrintVisibility(method.Visibility),
					Parameters:  dataMethodParameters,
					ReturnTypes: dataMethodReturnTypes,
					Labels:      s.LabelListToModel(method.Labels...),
					Options:     s.OptionListToModel(method.Options...),
					Code:        method.Code,
					Common:      method.Common,
				}

				dataMethods = append(dataMethods, dataMethod)
			}
		}

	}

	var imports []objectResources.ObjectResourceImport = make([]objectResources.ObjectResourceImport, 0)

	for key, value := range importsMap {
		imports = append(imports, objectResources.ObjectResourceImport{
			Package: key,
			Aliases: value,
		})
	}

	return objectResources.ObjectSection{
		Name:       section.Name,
		Package:    s.manager.PrintDependencies(packages),
		Classes:    section.Classes,
		Labels:     s.LabelListToModel(section.Labels...),
		Options:    s.OptionListToModel(section.Options...),
		Attributes: dataAttributes,
		Methods:    dataMethods,
		Imports:    imports,
	}
}

func (s *ObjectResourceService) DataSectionToModel(resource data_resource_payload_structs.ResourceSpecification, project application_project_payload_structs.ProjectSpecification, layer string, template file_template_payload_structs.TemplateSpecification, section data_resource_payload_structs.Section) objectResources.DataSection {
	var importsMap map[string][]string = make(map[string][]string)

	packages := project.GetAllPackageWithLayer(layer)
	packages = append(packages, template.Package...)

	var imports []objectResources.ObjectResourceImport = make([]objectResources.ObjectResourceImport, 0)

	for key, value := range importsMap {
		imports = append(imports, objectResources.ObjectResourceImport{
			Package: key,
			Aliases: value,
		})
	}

	return objectResources.DataSection{
		Name:    section.Name,
		Package: s.manager.PrintDependencies(packages),
		Labels:  s.LabelListToModel(section.Labels...),
		Options: s.OptionListToModel(section.Options...),
	}
}

func (s *ObjectResourceService) ObjectSectionListToModel(resource object_resource_payload_structs.ResourceBaseStruct, project application_project_payload_structs.ProjectBaseStruct, layer object_resource_payload_structs.Layer, template code_template_payload_structs.TemplateBaseStruct, sections ...object_resource_payload_structs.Section) []objectResources.ObjectSection {

	var result []objectResources.ObjectSection = make([]objectResources.ObjectSection, 0)

	for _, section := range sections {
		result = append(result, s.ObjectSectionToModel(resource, project, layer, template, section))
	}

	return result
}

func (s *ObjectResourceService) DataSectionListToModel(resource data_resource_payload_structs.ResourceSpecification, project application_project_payload_structs.ProjectSpecification, layer string, template file_template_payload_structs.TemplateSpecification, sections ...data_resource_payload_structs.Section) []objectResources.DataSection {

	var result []objectResources.DataSection = make([]objectResources.DataSection, 0)

	for _, section := range sections {
		result = append(result, s.DataSectionToModel(resource, project, layer, template, section))
	}

	return result
}

func (s *ObjectResourceService) LayerListToModel(resource object_resource_payload_structs.ResourceBaseStruct, project application_project_payload_structs.ProjectBaseStruct, template code_template_payload_structs.TemplateBaseStruct, layers ...object_resource_payload_structs.Layer) []objectResources.ObjectLayer {

	var result []objectResources.ObjectLayer = make([]objectResources.ObjectLayer, 0)

	for _, layer := range layers {
		result = append(result, objectResources.ObjectLayer{
			Name:     layer.Name,
			Sections: s.ObjectSectionListToModel(resource, project, layer, template, layer.Sections...),
		})
	}

	return result
}

func (s *ObjectResourceService) DataLayerListToModel(resource data_resource_payload_structs.ResourceSpecification, project application_project_payload_structs.ProjectSpecification, template file_template_payload_structs.TemplateSpecification, layers ...data_resource_payload_structs.Layer) []objectResources.DataLayer {

	var result []objectResources.DataLayer = make([]objectResources.DataLayer, 0)

	for _, layer := range layers {
		result = append(result, objectResources.DataLayer{
			Name:     layer.Name,
			Sections: s.DataSectionListToModel(resource, project, layer.Name, template, layer.Sections...),
		})
	}

	return result
}

func (s *ObjectResourceService) LabelListToModel(labels ...label.Label) []objectResources.ObjectLabel {

	var result []objectResources.ObjectLabel = make([]objectResources.ObjectLabel, 0)

	for _, label := range labels {
		result = append(result, objectResources.ObjectLabel{
			Key:   label.Key,
			Value: label.Value,
		})
	}

	return result
}

func (s *ObjectResourceService) OptionListToModel(options ...option.Option) []objectResources.ObjectOption {

	var result []objectResources.ObjectOption = make([]objectResources.ObjectOption, 0)

	for _, option := range options {
		result = append(result, objectResources.ObjectOption{
			Key:   option.Key,
			Value: option.Value,
		})
	}

	return result
}

func (s *ObjectResourceService) DictionaryListToModel(dictionaries ...object_resource_payload_structs.Dictionary) []objectResources.ObjectDictionary {

	var result []objectResources.ObjectDictionary = make([]objectResources.ObjectDictionary, 0)

	for _, dictionary := range dictionaries {
		result = append(result, objectResources.ObjectDictionary{
			Key:        dictionary.Key,
			Translates: dictionary.Translates,
		})
	}

	return result
}

func (s *ObjectResourceService) GroupListToModel(groups ...object_resource_payload_structs.Group) []objectResources.ObjectGroup {

	var result []objectResources.ObjectGroup = make([]objectResources.ObjectGroup, 0)

	for _, group := range groups {
		result = append(result, objectResources.ObjectGroup{
			Name:    group.Name,
			Title:   s.MessageToModel(group.Title),
			Options: s.OptionListToModel(group.Options...),
		})
	}

	return result
}

func (s *ObjectResourceService) MessageToModel(message object_resource_payload_structs.Message) objectResources.ObjectMessage {

	var result objectResources.ObjectMessage = objectResources.ObjectMessage{
		Text:       message.Text,
		Dictionary: message.Dictionary.Key,
	}

	return result
}

func (s *ObjectResourceService) WorkspaceToModel(workspace basic_workspace_payload_structs.WorkspaceBaseStruct) objectResources.Workspace {

	var result objectResources.Workspace = objectResources.Workspace{
		Name: workspace.Header.Name,
	}

	return result
}

func (s *ObjectResourceService) ApplicationProjectToModel(project application_project_payload_structs.ProjectBaseStruct) objectResources.ApplicationProject {
	dependencies := project.Specifications.GetAllPackage()

	var result objectResources.ApplicationProject = objectResources.ApplicationProject{
		Name:    project.Header.Name,
		Package: s.manager.PrintDependencies(dependencies),
		Labels:  s.LabelListToModel(project.Specifications.Labels...),
	}

	return result
}

func (s *ObjectResourceService) FileTemplateToModel(template file_template_payload_structs.TemplateBaseStruct) objectResources.FileTemplate {
	dependencies := template.Specifications.Package

	var result objectResources.FileTemplate = objectResources.FileTemplate{
		Name:    template.Header.Name,
		Package: s.manager.PrintDependencies(dependencies),
		Labels:  s.LabelListToModel(template.Specifications.Labels...),
	}

	return result
}

func (s *ObjectResourceService) CodeTemplateToModel(template code_template_payload_structs.TemplateBaseStruct) objectResources.CodeTemplate {
	dependencies := template.Specifications.Package

	var result objectResources.CodeTemplate = objectResources.CodeTemplate{
		Name:    template.Header.Name,
		Package: s.manager.PrintDependencies(dependencies),
		Labels:  s.LabelListToModel(template.Specifications.Labels...),
	}

	return result
}

func (s *ObjectResourceService) DataLayerToModel(layer data_resource_payload_structs.Layer) objectResources.DataLayer {

	var result objectResources.DataLayer = objectResources.DataLayer{
		Name: layer.Name,
	}

	return result
}

func (s *ObjectResourceService) ObjectLayerToModel(layer object_resource_payload_structs.Layer) objectResources.ObjectLayer {

	var result objectResources.ObjectLayer = objectResources.ObjectLayer{
		Name: layer.Name,
	}

	return result
}
