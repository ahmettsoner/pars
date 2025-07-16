package data_resource

import (
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/layer"
	layer2 "parsdevkit.net/application/models/layer"
	"parsdevkit.net/application/models/section"
	"parsdevkit.net/application/platforms"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/components/template"
	templatePkg "parsdevkit.net/components/template"
	"parsdevkit.net/modules/resource/data_resource_contract"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"

	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"

	"parsdevkit.net/application/models/class"
	"parsdevkit.net/application/models/option"
)

type DataResourceContextProvider struct {
	environment string
}

func NewDataResourceContextProvider(environment string) data_resource_contract.ContextProviderInterface {
	return &DataResourceContextProvider{
		environment: environment,
	}
}

func (s *DataResourceContextProvider) Context(source template.ContextProviderSource) interface{} {
	if model, ok := source.Resource.(data_resource_payload_structs.ResourceBaseStruct); ok {
		if modelProject, ok := source.Project.(application_project_payload_structs.ProjectBaseStruct); ok {
			if modelTemplate, ok := source.Template.(file_template_payload_structs.TemplateBaseStruct); ok {
				return DataResourceComposite{
					DataResource: s.structToModel(source.Workspace, modelProject, model, modelTemplate, source.Layer),
					Original:     model,
				}
			}
		}
	}
	return nil
}
func (s *DataResourceContextProvider) structToModel(workspace schemas.SchemaInterface, project application_project_payload_structs.ProjectBaseStruct, model data_resource_payload_structs.ResourceBaseStruct, template file_template_payload_structs.TemplateBaseStruct, layer layer.LayerIdentifier) DataResource {
	manager := platforms.Get[application_project_payload_structs.ProjectBaseStruct](project.Specifications.Platform.Type)

	var importsMap map[string][]string = make(map[string][]string)

	packages := project.Specifications.GetAllPackageWithLayer(layer.Name)
	packages = append(packages, template.Specifications.Package...)

	var imports []DataResourceImport = make([]DataResourceImport, 0)

	for key, value := range importsMap {
		imports = append(imports, DataResourceImport{
			Package: key,
			Aliases: value,
		})
	}
	layers := make([]layer2.LayerIdentifier, 0)
	for _, v := range model.Specifications.Layers {
		layers = append(layers, v.LayerIdentifier)
	}

	return DataResource{
		Name:    model.Specifications.Name,
		Package: manager.PrintDependencies(packages),
		Labels:  s.LabelListToModel(model.Specifications.Labels...),
		Layers:  s.LayerListToModel(workspace, project, model, template, layers...),
		// Dictionary: s.DictionaryListToModel(model.Dictionary...),
		// Groups:     s.GroupListToModel(model.Groups...),
		Data: model.Data,
	}
}
func (s *DataResourceContextProvider) LabelListToModel(labels ...label.Label) []ObjectLabel {

	var result []ObjectLabel = make([]ObjectLabel, 0)

	for _, label := range labels {
		result = append(result, ObjectLabel{
			Key:   label.Key,
			Value: label.Value,
		})
	}

	return result
}

func (s *DataResourceContextProvider) OptionListToModelDataResource(options ...option.Option) []ObjectOption {

	var result []ObjectOption = make([]ObjectOption, 0)

	for _, option := range options {
		result = append(result, ObjectOption{
			Key:   option.Key,
			Value: option.Value,
		})
	}

	return result
}
func (s *DataResourceContextProvider) LayerListToModel(workspace, project, resource, template schemas.SchemaInterface, layers ...layer.LayerIdentifier) []DataLayer {

	var result []DataLayer = make([]DataLayer, 0)

	for _, layer := range layers {
		sections := make([]section.SectionIdentifier, 0)
		// sectionLayer, ok := layer.(object_resource_payload_structs.Layer)
		// if ok {
		// }
		// for _, v := range sectionLayer.Sections {
		// 	sections = append(sections, v.SectionIdentifier)
		// }

		result = append(result, DataLayer{
			Name:     layer.Name,
			Sections: s.ObjectSectionListToModel(workspace, project, resource, template, layer, sections...),
		})
	}

	return result

}
func (s *DataResourceContextProvider) LayerToModelContext(source template.ContextProviderSource) interface{} {

	var result DataLayer = DataLayer{
		Name: source.Layer.Name,
	}

	return result
}

func (s *DataResourceContextProvider) ObjectSectionListToModel(workspace, project, resource, template schemas.SchemaInterface, _layer layer.LayerIdentifier, sections ...section.SectionIdentifier) []DataSection {

	var result []DataSection = make([]DataSection, 0)

	for _, section := range sections {

		objectSection := s.SectionToModelContext(templatePkg.NewContextProviderSource(workspace, nil, project, resource, template, _layer, section))
		objectSectionModel, ok := objectSection.(DataSection)
		if ok {
		}
		result = append(result, objectSectionModel)
	}

	return result
}

func (s *DataResourceContextProvider) SectionToModelContext(source template.ContextProviderSource) interface{} {
	var importsMap map[string][]string = make(map[string][]string)
	projectModel, ok := source.Project.(application_project_payload_structs.ProjectBaseStruct)
	if ok {
	}
	// resourceModel, ok := source.Resource.(object_resource_payload_structs.ResourceBaseStruct)
	// if ok {
	// }
	templateModel, ok := source.Template.(file_template_payload_structs.TemplateBaseStruct)
	if ok {
	}
	var sectionModel data_resource_payload_structs.Section
	// sectionModel, ok := section.(object_resource_payload_structs.Section)
	// if ok {
	// }
	manager := platforms.Get[application_project_payload_structs.ProjectBaseStruct](projectModel.Specifications.Platform.Type)

	packages := projectModel.Specifications.GetAllPackageWithLayer(source.Layer.Name)
	packages = append(packages, templateModel.Specifications.Package...)

	var imports []DataResourceImport = make([]DataResourceImport, 0)

	for key, value := range importsMap {
		imports = append(imports, DataResourceImport{
			Package: key,
			Aliases: value,
		})
	}

	return DataSection{
		Name:    source.Section.Name,
		Package: manager.PrintDependencies(packages),
		Labels:  s.LabelListToModel(sectionModel.Labels...),
		Options: s.OptionListToModelDataResource(sectionModel.Options...),
	}
}

type DataResourceComposite struct {
	DataResource
	Original data_resource_payload_structs.ResourceBaseStruct
}

type DataResource struct {
	Name       string
	Package    string
	Labels     []ObjectLabel
	Layers     []DataLayer
	Dictionary []ObjectDictionary
	Groups     []ObjectGroup
	Data       any
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

type DataLayer struct {
	Name     string
	Sections []DataSection
}

type DataSection struct {
	Name    string
	Package string
	Classes []class.Class
	Labels  []ObjectLabel
	Options []ObjectOption
}

type DataResourceAttribute struct {
	Name         string
	TypePackage  string
	Type         string
	TypeCategory string
	Visibility   string
	Labels       []ObjectLabel
	Options      []ObjectOption
	Common       bool
}

type DataResourceMethod struct {
	Name        string
	Visibility  string
	Parameters  []DataResourceMethodParameter
	ReturnTypes []string
	Labels      []ObjectLabel
	Options     []ObjectOption
	Code        string
	Common      bool
}
type DataResourceMethodParameter struct {
	Name string
	Type string
}
type DataResourceImport struct {
	Aliases []string
	Package string
}
type DataLayerComposite struct {
	DataLayer
	Original data_resource_payload_structs.Layer
}

type DataSectionComposite struct {
	DataSection
	Original data_resource_payload_structs.Section
}
