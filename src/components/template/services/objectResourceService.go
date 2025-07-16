package services

import (
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	"parsdevkit.net/application/structs"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"

	"parsdevkit.net/modules/resource/data_resource"
	"parsdevkit.net/modules/resource/object_resource"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"
	_string "parsdevkit.net/pkg/utilities/string"
	"parsdevkit.net/platforms/core"
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

func (s *ObjectResourceService) ResourceToModel(resource object_resource_payload_structs.ResourceBaseStruct, project application_project_payload_structs.ProjectBaseStruct, layer object_resource_payload_structs.Layer, template code_template_payload_structs.TemplateBaseStruct) object_resource.ObjectResource {
	var importsMap map[string][]string = make(map[string][]string)

	packages := project.Specifications.GetAllPackageWithLayer(layer.Name)
	packages = append(packages, resource.Specifications.Package...)
	packages = append(packages, template.Specifications.Package...)

	var dataAttributes = make([]object_resource.ObjectResourceAttribute, 0)
	for _, attribute := range resource.Object.Attributes {
		var dataAttribute = object_resource.ObjectResourceAttribute{
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

	var dataMethods = make([]object_resource.ObjectResourceMethod, 0)
	for _, method := range resource.Object.Methods {
		var dataMethodParameters = make([]object_resource.ObjectResourceMethodParameter, 0)
		for _, methodParameter := range method.Parameters {
			var dataMethodParameter = object_resource.ObjectResourceMethodParameter{
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
		var dataMethod = object_resource.ObjectResourceMethod{
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

	var imports []object_resource.ObjectResourceImport = make([]object_resource.ObjectResourceImport, 0)

	for key, value := range importsMap {
		imports = append(imports, object_resource.ObjectResourceImport{
			Package: key,
			Aliases: value,
		})
	}

	return object_resource.ObjectResource{
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

func (s *ObjectResourceService) DataResourceToModel(resource data_resource_payload_structs.ResourceSpecification, resourceData data_resource_payload_structs.ResourceData, project application_project_payload_structs.ProjectSpecification, layer string, template file_template_payload_structs.TemplateSpecification) data_resource.DataResource {
	var importsMap map[string][]string = make(map[string][]string)

	packages := project.GetAllPackageWithLayer(layer)
	packages = append(packages, template.Package...)

	var imports []data_resource.DataResourceImport = make([]data_resource.DataResourceImport, 0)

	for key, value := range importsMap {
		imports = append(imports, data_resource.DataResourceImport{
			Package: key,
			Aliases: value,
		})
	}

	return data_resource.DataResource{
		Name:    resource.Name,
		Package: s.manager.PrintDependencies(packages),
		Labels:  s.LabelListToModelDataResource(resource.Labels...),
		Layers:  s.DataLayerListToModel(resource, project, template, resource.Layers...),
		// Dictionary: s.DictionaryListToModel(resource.Dictionary...),
		// Groups:     s.GroupListToModel(resource.Groups...),
		Data: resourceData.Data,
	}
}

func (s *ObjectResourceService) ObjectSectionToModel(resource object_resource_payload_structs.ResourceBaseStruct, project application_project_payload_structs.ProjectBaseStruct, layer object_resource_payload_structs.Layer, template code_template_payload_structs.TemplateBaseStruct, section object_resource_payload_structs.Section) object_resource.ObjectSection {
	var importsMap map[string][]string = make(map[string][]string)

	packages := project.Specifications.GetAllPackageWithLayer(layer.Name)
	packages = append(packages, resource.Specifications.Package...)
	packages = append(packages, template.Specifications.Package...)

	var dataAttributes = make([]object_resource.ObjectResourceAttribute, 0)
	for _, attribute := range resource.Object.Attributes {
		for _, sectionAttribute := range section.Attributes {
			if attribute.Name == sectionAttribute {
				var dataAttribute = object_resource.ObjectResourceAttribute{
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

	var dataMethods = make([]object_resource.ObjectResourceMethod, 0)
	for _, method := range resource.Object.Methods {
		var dataMethodParameters = make([]object_resource.ObjectResourceMethodParameter, 0)
		for _, sectionMethod := range section.Methods {
			if method.Name == sectionMethod {
				for _, methodParameter := range method.Parameters {
					var dataMethodParameter = object_resource.ObjectResourceMethodParameter{
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
				var dataMethod = object_resource.ObjectResourceMethod{
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

	var imports []object_resource.ObjectResourceImport = make([]object_resource.ObjectResourceImport, 0)

	for key, value := range importsMap {
		imports = append(imports, object_resource.ObjectResourceImport{
			Package: key,
			Aliases: value,
		})
	}

	return object_resource.ObjectSection{
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

func (s *ObjectResourceService) DataSectionToModel(resource data_resource_payload_structs.ResourceSpecification, project application_project_payload_structs.ProjectSpecification, layer string, template file_template_payload_structs.TemplateSpecification, section data_resource_payload_structs.Section) data_resource.DataSection {
	var importsMap map[string][]string = make(map[string][]string)

	packages := project.GetAllPackageWithLayer(layer)
	packages = append(packages, template.Package...)

	var imports []data_resource.DataResourceImport = make([]data_resource.DataResourceImport, 0)

	for key, value := range importsMap {
		imports = append(imports, data_resource.DataResourceImport{
			Package: key,
			Aliases: value,
		})
	}

	return data_resource.DataSection{
		Name:    section.Name,
		Package: s.manager.PrintDependencies(packages),
		Labels:  s.LabelListToModelDataResource(section.Labels...),
		Options: s.OptionListToModelDataResource(section.Options...),
	}
}

func (s *ObjectResourceService) ObjectSectionListToModel(resource object_resource_payload_structs.ResourceBaseStruct, project application_project_payload_structs.ProjectBaseStruct, layer object_resource_payload_structs.Layer, template code_template_payload_structs.TemplateBaseStruct, sections ...object_resource_payload_structs.Section) []object_resource.ObjectSection {

	var result []object_resource.ObjectSection = make([]object_resource.ObjectSection, 0)

	for _, section := range sections {
		result = append(result, s.ObjectSectionToModel(resource, project, layer, template, section))
	}

	return result
}

func (s *ObjectResourceService) DataSectionListToModel(resource data_resource_payload_structs.ResourceSpecification, project application_project_payload_structs.ProjectSpecification, layer string, template file_template_payload_structs.TemplateSpecification, sections ...data_resource_payload_structs.Section) []data_resource.DataSection {

	var result []data_resource.DataSection = make([]data_resource.DataSection, 0)

	for _, section := range sections {
		result = append(result, s.DataSectionToModel(resource, project, layer, template, section))
	}

	return result
}

func (s *ObjectResourceService) LayerListToModel(resource object_resource_payload_structs.ResourceBaseStruct, project application_project_payload_structs.ProjectBaseStruct, template code_template_payload_structs.TemplateBaseStruct, layers ...object_resource_payload_structs.Layer) []object_resource.ObjectLayer {

	var result []object_resource.ObjectLayer = make([]object_resource.ObjectLayer, 0)

	for _, layer := range layers {
		result = append(result, object_resource.ObjectLayer{
			Name:     layer.Name,
			Sections: s.ObjectSectionListToModel(resource, project, layer, template, layer.Sections...),
		})
	}

	return result
}

func (s *ObjectResourceService) DataLayerListToModel(resource data_resource_payload_structs.ResourceSpecification, project application_project_payload_structs.ProjectSpecification, template file_template_payload_structs.TemplateSpecification, layers ...data_resource_payload_structs.Layer) []data_resource.DataLayer {

	var result []data_resource.DataLayer = make([]data_resource.DataLayer, 0)

	for _, layer := range layers {
		result = append(result, data_resource.DataLayer{
			Name:     layer.Name,
			Sections: s.DataSectionListToModel(resource, project, layer.Name, template, layer.Sections...),
		})
	}

	return result
}

func (s *ObjectResourceService) LabelListToModel(labels ...label.Label) []object_resource.ObjectLabel {

	var result []object_resource.ObjectLabel = make([]object_resource.ObjectLabel, 0)

	for _, label := range labels {
		result = append(result, object_resource.ObjectLabel{
			Key:   label.Key,
			Value: label.Value,
		})
	}

	return result
}
func (s *ObjectResourceService) LabelListToModelDataResource(labels ...label.Label) []data_resource.ObjectLabel {

	var result []data_resource.ObjectLabel = make([]data_resource.ObjectLabel, 0)

	for _, label := range labels {
		result = append(result, data_resource.ObjectLabel{
			Key:   label.Key,
			Value: label.Value,
		})
	}

	return result
}

func (s *ObjectResourceService) OptionListToModel(options ...option.Option) []object_resource.ObjectOption {

	var result []object_resource.ObjectOption = make([]object_resource.ObjectOption, 0)

	for _, option := range options {
		result = append(result, object_resource.ObjectOption{
			Key:   option.Key,
			Value: option.Value,
		})
	}

	return result
}

func (s *ObjectResourceService) OptionListToModelDataResource(options ...option.Option) []data_resource.ObjectOption {

	var result []data_resource.ObjectOption = make([]data_resource.ObjectOption, 0)

	for _, option := range options {
		result = append(result, data_resource.ObjectOption{
			Key:   option.Key,
			Value: option.Value,
		})
	}

	return result
}

func (s *ObjectResourceService) DictionaryListToModel(dictionaries ...object_resource_payload_structs.Dictionary) []object_resource.ObjectDictionary {

	var result []object_resource.ObjectDictionary = make([]object_resource.ObjectDictionary, 0)

	for _, dictionary := range dictionaries {
		result = append(result, object_resource.ObjectDictionary{
			Key:        dictionary.Key,
			Translates: dictionary.Translates,
		})
	}

	return result
}

func (s *ObjectResourceService) GroupListToModel(groups ...object_resource_payload_structs.Group) []object_resource.ObjectGroup {

	var result []object_resource.ObjectGroup = make([]object_resource.ObjectGroup, 0)

	for _, group := range groups {
		result = append(result, object_resource.ObjectGroup{
			Name:    group.Name,
			Title:   s.MessageToModel(group.Title),
			Options: s.OptionListToModel(group.Options...),
		})
	}

	return result
}

func (s *ObjectResourceService) MessageToModel(message object_resource_payload_structs.Message) object_resource.ObjectMessage {

	var result object_resource.ObjectMessage = object_resource.ObjectMessage{
		Text:       message.Text,
		Dictionary: message.Dictionary.Key,
	}

	return result
}

func (s *ObjectResourceService) DataLayerToModel(layer data_resource_payload_structs.Layer) data_resource.DataLayer {

	var result data_resource.DataLayer = data_resource.DataLayer{
		Name: layer.Name,
	}

	return result
}

func (s *ObjectResourceService) ObjectLayerToModel(layer object_resource_payload_structs.Layer) object_resource.ObjectLayer {

	var result object_resource.ObjectLayer = object_resource.ObjectLayer{
		Name: layer.Name,
	}

	return result
}
