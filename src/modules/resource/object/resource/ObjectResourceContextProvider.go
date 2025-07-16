package object_resource

import (
	"parsdevkit.net/application/platforms"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/layer"
	layer2 "parsdevkit.net/application/models/layer"
	"parsdevkit.net/application/models/section"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs"
	"parsdevkit.net/components/template"
	templatePkg "parsdevkit.net/components/template"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"

	"parsdevkit.net/modules/resource/object_resource_contract"

	"parsdevkit.net/application/models/option"

	"parsdevkit.net/application/models/class"
)

type ObjectResourceContextProvider struct {
	environment string
}

func NewObjectResourceContextProvider(environment string) object_resource_contract.ContextProviderInterface {
	return &ObjectResourceContextProvider{
		environment: environment,
	}
}

func (s *ObjectResourceContextProvider) Context(source template.ContextProviderSource) interface{} {

	if model, ok := source.Resource.(object_resource_payload_structs.ResourceBaseStruct); ok {
		if modelProject, ok := source.Project.(application_project_payload_structs.ProjectBaseStruct); ok {
			if modelTemplate, ok := source.Template.(code_template_payload_structs.TemplateBaseStruct); ok {
				return ObjectResourceComposite{
					ObjectResource: s.structToModel(source.Workspace, modelProject, model, modelTemplate, source.Layer),
					Original:       model,
				}
			}
		}
	}
	return nil
}
func (s *ObjectResourceContextProvider) structToModel(workspace schemas.SchemaInterface, project application_project_payload_structs.ProjectBaseStruct, model object_resource_payload_structs.ResourceBaseStruct, template code_template_payload_structs.TemplateBaseStruct, layer layer.LayerIdentifier) ObjectResource {
	manager := platforms.Get[application_project_payload_structs.ProjectBaseStruct](project.Specifications.Platform.Type)

	var importsMap map[string][]string = make(map[string][]string)

	packages := project.Specifications.GetAllPackageWithLayer(layer.Name)
	packages = append(packages, model.Specifications.Package...)
	packages = append(packages, template.Specifications.Package...)

	var dataAttributes = make([]ObjectResourceAttribute, 0)
	for _, attribute := range model.Object.Attributes {
		var dataAttribute = ObjectResourceAttribute{
			Name:         attribute.Name,
			Type:         manager.PrintDataType(attribute.Type),
			TypePackage:  attribute.Type.Name,
			TypeCategory: string(attribute.Type.Category),
			Visibility:   manager.PrintVisibility(attribute.Visibility),
			Labels:       s.LabelListToModel(attribute.Labels...),
			Options:      s.OptionListToModel(attribute.Options...),
			Common:       attribute.Common,
		}
		dataAttributes = append(dataAttributes, dataAttribute)

		importsMap = s.DataTypeToImport(attribute.Type, importsMap)
	}

	var dataMethods = make([]ObjectResourceMethod, 0)
	for _, method := range model.Object.Methods {
		var dataMethodParameters = make([]ObjectResourceMethodParameter, 0)
		for _, methodParameter := range method.Parameters {
			var dataMethodParameter = ObjectResourceMethodParameter{
				Name: methodParameter.Name,
				Type: manager.PrintDataType(methodParameter.Type),
			}

			dataMethodParameters = append(dataMethodParameters, dataMethodParameter)
			importsMap = s.DataTypeToImport(methodParameter.Type, importsMap)
		}

		var dataMethodReturnTypes = make([]string, 0)
		for _, methodReturnType := range method.ReturnTypes {
			var dataMethodReturnType = manager.PrintDataType(methodReturnType)

			dataMethodReturnTypes = append(dataMethodReturnTypes, dataMethodReturnType)
			importsMap = s.DataTypeToImport(methodReturnType, importsMap)
		}
		var dataMethod = ObjectResourceMethod{
			Name:        method.Name,
			Visibility:  manager.PrintVisibility(method.Visibility),
			Parameters:  dataMethodParameters,
			ReturnTypes: dataMethodReturnTypes,
			Labels:      s.LabelListToModel(method.Labels...),
			Options:     s.OptionListToModel(method.Options...),
			Code:        method.Code,
			Common:      method.Common,
		}

		dataMethods = append(dataMethods, dataMethod)
	}

	var imports []ObjectResourceImport = make([]ObjectResourceImport, 0)

	for key, value := range importsMap {
		imports = append(imports, ObjectResourceImport{
			Package: key,
			Aliases: value,
		})
	}
	layers := make([]layer2.LayerIdentifier, 0)
	for _, v := range model.Specifications.Layers {
		layers = append(layers, v.LayerIdentifier)
	}

	return ObjectResource{
		Name:       model.Specifications.Name,
		Package:    manager.PrintDependencies(packages),
		Labels:     s.LabelListToModel(model.Specifications.Labels...),
		Layers:     s.LayerListToModel(workspace, project, model, template, layers...),
		Dictionary: s.DictionaryListToModel(model.Object.Dictionary...),
		Groups:     s.GroupListToModel(model.Object.Groups...),
		Attributes: dataAttributes,
		Methods:    dataMethods,
		Imports:    imports,
	}
}

func (s *ObjectResourceContextProvider) LabelListToModel(labels ...label.Label) []ObjectLabel {

	var result []ObjectLabel = make([]ObjectLabel, 0)

	for _, label := range labels {
		result = append(result, ObjectLabel{
			Key:   label.Key,
			Value: label.Value,
		})
	}

	return result
}

func (s *ObjectResourceContextProvider) LayerListToModel(workspace, project, resource, template schemas.SchemaInterface, layers ...layer.LayerIdentifier) []ObjectLayer {

	var result []ObjectLayer = make([]ObjectLayer, 0)

	for _, layer := range layers {
		sections := make([]section.SectionIdentifier, 0)
		// sectionLayer, ok := layer.(object_resource_payload_structs.Layer)
		// if ok {
		// }
		// for _, v := range sectionLayer.Sections {
		// 	sections = append(sections, v.SectionIdentifier)
		// }
		result = append(result, ObjectLayer{
			Name:     layer.Name,
			Sections: s.ObjectSectionListToModel(workspace, project, resource, template, sections...),
		})
	}

	return result
}
func (s *ObjectResourceContextProvider) ObjectSectionListToModel(workspace, project, resource, template schemas.SchemaInterface, sections ...section.SectionIdentifier) []ObjectSection {

	var result []ObjectSection = make([]ObjectSection, 0)

	for _, section := range sections {

		objectSection := s.SectionToModelContext(templatePkg.NewContextProviderSource(workspace, nil, project, resource, template, layer.LayerIdentifier{}, section))
		objectSectionModel, ok := objectSection.(ObjectSectionComposite)
		if ok {
		}
		result = append(result, objectSectionModel.ObjectSection)
	}

	return result
}
func (s *ObjectResourceContextProvider) LayerToModelContext(source template.ContextProviderSource) interface{} {

	var result ObjectLayer = ObjectLayer{
		Name: source.Layer.Name,
	}

	return ObjectLayerComposite{
		ObjectLayer: result,
		Original:    source.Layer,
	}
}

func (s *ObjectResourceContextProvider) SectionToModelContext(source template.ContextProviderSource) interface{} {
	var importsMap map[string][]string = make(map[string][]string)
	projectModel, ok := source.Project.(application_project_payload_structs.ProjectBaseStruct)
	if ok {
	}
	resourceModel, ok := source.Resource.(object_resource_payload_structs.ResourceBaseStruct)
	if ok {
	}
	templateModel, ok := source.Template.(code_template_payload_structs.TemplateBaseStruct)
	if ok {
	}
	var sectionModel object_resource_payload_structs.Section
	// sectionModel, ok := section.(object_resource_payload_structs.Section)
	// if ok {
	// }
	manager := platforms.Get[application_project_payload_structs.ProjectBaseStruct](projectModel.Specifications.Platform.Type)

	packages := projectModel.Specifications.GetAllPackageWithLayer(source.Layer.Name)
	packages = append(packages, resourceModel.Specifications.Package...)
	packages = append(packages, templateModel.Specifications.Package...)

	var dataAttributes = make([]ObjectResourceAttribute, 0)
	for _, attribute := range resourceModel.Object.Attributes {
		for _, sectionAttribute := range sectionModel.Attributes {
			if attribute.Name == sectionAttribute {
				var dataAttribute = ObjectResourceAttribute{
					Name:         attribute.Name,
					Type:         manager.PrintDataType(attribute.Type),
					TypePackage:  attribute.Type.Name,
					TypeCategory: string(attribute.Type.Category),
					Visibility:   manager.PrintVisibility(attribute.Visibility),
					Labels:       s.LabelListToModel(attribute.Labels...),
					Options:      s.OptionListToModel(attribute.Options...),
					Common:       attribute.Common,
				}
				dataAttributes = append(dataAttributes, dataAttribute)

				importsMap = s.DataTypeToImport(attribute.Type, importsMap)
			}
		}
	}

	var dataMethods = make([]ObjectResourceMethod, 0)
	for _, method := range resourceModel.Object.Methods {
		var dataMethodParameters = make([]ObjectResourceMethodParameter, 0)
		for _, sectionMethod := range sectionModel.Methods {
			if method.Name == sectionMethod {
				for _, methodParameter := range method.Parameters {
					var dataMethodParameter = ObjectResourceMethodParameter{
						Name: methodParameter.Name,
						Type: manager.PrintDataType(methodParameter.Type),
					}

					dataMethodParameters = append(dataMethodParameters, dataMethodParameter)
					importsMap = s.DataTypeToImport(methodParameter.Type, importsMap)
				}

				var dataMethodReturnTypes = make([]string, 0)
				for _, methodReturnType := range method.ReturnTypes {
					var dataMethodReturnType = manager.PrintDataType(methodReturnType)

					dataMethodReturnTypes = append(dataMethodReturnTypes, dataMethodReturnType)
					importsMap = s.DataTypeToImport(methodReturnType, importsMap)
				}
				var dataMethod = ObjectResourceMethod{
					Name:        method.Name,
					Visibility:  manager.PrintVisibility(method.Visibility),
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

	var imports []ObjectResourceImport = make([]ObjectResourceImport, 0)

	for key, value := range importsMap {
		imports = append(imports, ObjectResourceImport{
			Package: key,
			Aliases: value,
		})
	}

	return ObjectSectionComposite{
		ObjectSection: ObjectSection{
			Name:       sectionModel.Name,
			Package:    manager.PrintDependencies(packages),
			Classes:    sectionModel.Classes,
			Labels:     s.LabelListToModel(sectionModel.Labels...),
			Options:    s.OptionListToModel(sectionModel.Options...),
			Attributes: dataAttributes,
			Methods:    dataMethods,
			Imports:    imports,
		},
		Original: source.Section,
	}
}
func (s *ObjectResourceContextProvider) OptionListToModel(options ...option.Option) []ObjectOption {

	var result []ObjectOption = make([]ObjectOption, 0)

	for _, option := range options {
		result = append(result, ObjectOption{
			Key:   option.Key,
			Value: option.Value,
		})
	}

	return result
}
func (s *ObjectResourceContextProvider) DataTypeToImport(_type structs.DataType, importsMap map[string][]string) map[string][]string {

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
func (s *ObjectResourceContextProvider) DictionaryListToModel(dictionaries ...object_resource_payload_structs.Dictionary) []ObjectDictionary {

	var result []ObjectDictionary = make([]ObjectDictionary, 0)

	for _, dictionary := range dictionaries {
		result = append(result, ObjectDictionary{
			Key:        dictionary.Key,
			Translates: dictionary.Translates,
		})
	}

	return result
}
func (s *ObjectResourceContextProvider) GroupListToModel(groups ...object_resource_payload_structs.Group) []ObjectGroup {

	var result []ObjectGroup = make([]ObjectGroup, 0)

	for _, group := range groups {
		result = append(result, ObjectGroup{
			Name:    group.Name,
			Title:   s.MessageToModel(group.Title),
			Options: s.OptionListToModel(group.Options...),
		})
	}

	return result
}
func (s *ObjectResourceContextProvider) MessageToModel(message object_resource_payload_structs.Message) ObjectMessage {

	var result ObjectMessage = ObjectMessage{
		Text:       message.Text,
		Dictionary: message.Dictionary.Key,
	}

	return result
}

type ObjectResourceComposite struct {
	ObjectResource
	Original object_resource_payload_structs.ResourceBaseStruct
}

type ObjectResource struct {
	Name       string
	Package    string
	Labels     []ObjectLabel
	Layers     []ObjectLayer
	Dictionary []ObjectDictionary
	Groups     []ObjectGroup
	Attributes []ObjectResourceAttribute
	Methods    []ObjectResourceMethod
	Imports    []ObjectResourceImport
}

type ObjectLabel struct {
	Key   string
	Value string
}

type ObjectDictionary struct {
	Key        string
	Translates map[string]string
}

type ObjectGroup struct {
	Name    string
	Title   ObjectMessage
	Options []ObjectOption
}

type ObjectMessage struct {
	Text       string
	Dictionary string
}
type ObjectOption struct {
	Key   string
	Value interface{}
}

type ObjectLayer struct {
	Name     string
	Sections []ObjectSection
}

type ObjectSection struct {
	Name       string
	Package    string
	Classes    []class.Class
	Labels     []ObjectLabel
	Options    []ObjectOption
	Attributes []ObjectResourceAttribute
	Methods    []ObjectResourceMethod
	Imports    []ObjectResourceImport
}

type ObjectResourceAttribute struct {
	Name         string
	TypePackage  string
	Type         string
	TypeCategory string
	Visibility   string
	Labels       []ObjectLabel
	Options      []ObjectOption
	Common       bool
}

type ObjectResourceMethod struct {
	Name        string
	Visibility  string
	Parameters  []ObjectResourceMethodParameter
	ReturnTypes []string
	Labels      []ObjectLabel
	Options     []ObjectOption
	Code        string
	Common      bool
}
type ObjectResourceMethodParameter struct {
	Name string
	Type string
}
type ObjectResourceImport struct {
	Aliases []string
	Package string
}

type ObjectLayerComposite struct {
	ObjectLayer
	Original layer.LayerIdentifier
}

type ObjectSectionComposite struct {
	ObjectSection
	Original section.SectionIdentifier
}
