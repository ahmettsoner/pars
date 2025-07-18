package engines

import (
	"fmt"
	"strings"

	"parsdevkit.net/application/contextgenerator"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/models/class"
	"parsdevkit.net/application/models/label"
	layerPkg "parsdevkit.net/application/models/layer"
	sectionPkg "parsdevkit.net/application/models/section"
	applicationStructs "parsdevkit.net/application/structs"
	applicationResource "parsdevkit.net/application/structs/resource"
	applicationTemplate "parsdevkit.net/application/structs/template"
	templatePkg "parsdevkit.net/components/template"
	templateEngine "parsdevkit.net/components/template/engines"
	"parsdevkit.net/modules/project/application_project_contract"
	"parsdevkit.net/modules/resource/data_resource_contract"
	"parsdevkit.net/modules/resource/object_resource_contract"
	"parsdevkit.net/modules/template/code_template_contract"
	"parsdevkit.net/modules/template/file_template_contract"
	"parsdevkit.net/modules/template/shared_template_contract"
	"parsdevkit.net/pkg/utilities/encrypt"
	"parsdevkit.net/pkg/utilities/file"

	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"

	"parsdevkit.net/persistence/contexts"
	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"parsdevkit.net/application/schemas"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
)

type TemplateOperationContent struct {
	Workspace schemas.SchemaInterface
	Project   schemas.SchemaInterface
	Resource  schemas.SchemaInterface
	Template  schemas.SchemaInterface
	Layer     layerPkg.Layer
	Section   sectionPkg.SectionIdentifier
	Class     class.Class
}

type TemplateOperations struct {
	environment                 string
	generationHistoryRepository *repositories.GenerationHistoryRepository
}

func NewTemplateOperations(environment string) TemplateOperations {
	dbContext := contexts.NewDbContext(environment)
	return TemplateOperations{
		environment:                 environment,
		generationHistoryRepository: repositories.NewGenerationHistoryRepository(dbContext),
	}
}

func (s TemplateOperations) Generate(source schemas.SchemaInterface) error {

	var setName, workspaceName string
	var layers []layerPkg.Layer
	var tags []string
	var labels []label.Label

	sourceHeader := source.GetHeader()

	switch source.GetHeader().Type {
	case schemas.StructTypes.Resource:
		sourceSpecification := source.GetSpecification().(applicationResource.ResourceSpecification)
		setName = sourceSpecification.Set
		workspaceName = sourceSpecification.Workspace
		layers = sourceSpecification.Layers
		tags = sourceHeader.Metadata.Tags
		labels = sourceSpecification.Labels

	case schemas.StructTypes.Template:
		sourceSpecification := source.GetSpecification().(applicationTemplate.TemplateSpecification)
		setName = sourceSpecification.Set
		workspaceName = sourceSpecification.Workspace
		layers = sourceSpecification.Layers
		tags = sourceHeader.Metadata.Tags
		labels = sourceSpecification.Labels

	default:
		return fmt.Errorf("unsupported type: %T", source)
	}

	var opCnt []TemplateOperationContent
	var workspace schemas.SchemaInterface
	var projects []schemas.SchemaInterface
	var templates []schemas.SchemaInterface
	var resources []schemas.SchemaInterface
	var err error

	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	workspace, err = workspaceService.GetByName(workspaceName)
	if err != nil {
		return err
	}

	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	projectsList, err := projectService.ListByFilter(setName, workspaceName, layerNames(layers), tags, labels)
	if err != nil {
		return err
	}
	for _, v := range *projectsList {
		projects = append(projects, v)
	}

	if sourceHeader.Type != schemas.StructTypes.Resource {
		if sourceHeader.Kind == "Object" {
			resourceService := ioc.Get[object_resource_contract.ResourceInterface]()
			resourcesList, err := resourceService.ListByFilter(setName, workspaceName, layerNames(layers), tags, labels)
			if err != nil {
				return err
			}
			for _, v := range *resourcesList {
				resources = append(resources, v)
			}
		} else if sourceHeader.Kind == "Data" {
			resourceService := ioc.Get[data_resource_contract.ResourceInterface]()
			resourcesList, err := resourceService.ListByFilter(setName, workspaceName, layerNames(layers), tags, labels)
			if err != nil {
				return err
			}
			for _, v := range *resourcesList {
				resources = append(resources, v)
			}
		}
	} else {
		resources = append(resources, source)
	}

	if sourceHeader.Type != schemas.StructTypes.Template {
		if sourceHeader.Kind == "Code" {
			templateService := ioc.Get[code_template_contract.TemplateInterface]()
			templatesList, err := templateService.ListByFilter(setName, workspaceName, layerNames(layers), tags, labels)
			if err != nil {
				return err
			}
			for _, v := range *templatesList {
				templates = append(templates, v)
			}
		} else if sourceHeader.Kind == "File" {
			templateService := ioc.Get[file_template_contract.TemplateInterface]()
			templatesList, err := templateService.ListByFilter(setName, workspaceName, layerNames(layers), tags, labels)
			if err != nil {
				return err
			}
			for _, v := range *templatesList {
				templates = append(templates, v)
			}
		} else if sourceHeader.Kind == "Shared" {
			templateService := ioc.Get[shared_template_contract.TemplateInterface]()
			templatesList, err := templateService.ListByFilter(setName, workspaceName, layerNames(layers), tags, labels)
			if err != nil {
				return err
			}
			for _, v := range *templatesList {
				templates = append(templates, v)
			}
		}
	} else {
		resources = append(resources, source)
	}

	for _, project := range projects {
		for _, resource := range resources {
			for _, template := range templates {
				opCnt = append(opCnt, s.extractOperationContents(project, template, resource, workspace)...)
			}
		}
	}

	for _, cnt := range opCnt {
		generate, _, _, _, err := s.CheckGeneration(cnt.Project, cnt.Resource, cnt.Template, cnt.Section, cnt.Layer)
		if err != nil {
			return err
		}

		if generate {
			err := s.GenerateContent(cnt)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
func (s TemplateOperations) GenerateContent(cnt TemplateOperationContent) error {
	resourceSpecifications := cnt.Resource.GetSpecification().(applicationResource.ResourceSpecification)
	templateSpecifications := cnt.Template.GetSpecification().(applicationTemplate.TemplateSpecification)

	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	var data = contextgenerator.NewTemplateDataContext(
		templatePkg.NewContextProviderSource(
			cnt.Workspace, nil, cnt.Project, cnt.Resource, cnt.Template, cnt.Layer.LayerIdentifier, cnt.Section),
	)

	fileNameStr, err := templateEngine.RenderTemplate(templateSpecifications.Output.File, data)
	if err != nil {
		return err
	}
	pathStr, err := templateEngine.RenderTemplate(templateSpecifications.Path, data)
	if err != nil {
		return err
	}

	tempPackages := templateSpecifications.Package
	packageStr, err := templateEngine.RenderTemplate(strings.Join(tempPackages, "/"), data)
	if err != nil {
		return err
	}
	templateSpecifications.Package = file.PathToArray(packageStr)

	data = contextgenerator.NewTemplateDataContext(
		templatePkg.NewContextProviderSource(
			cnt.Workspace, nil, cnt.Project, cnt.Resource, cnt.Template, cnt.Layer.LayerIdentifier, cnt.Section),
	)
	templateContentStr, err := templateEngine.RenderTemplate(templateSpecifications.Template.Content, data)
	if err != nil {
		return err
	}
	templateSpecifications.Package = tempPackages

	templateContentStr = AddCommentToGeneratedFile(templateSpecifications.Output.File, string(resourceSpecifications.Generate), string(templateSpecifications.Generate), templateContentStr)

	if cnt.Project.GetHeader().Kind == "Application" {
		projectObj := cnt.Project.(application_project_payload_structs.ProjectBaseStruct)
		_, err = projectService.AddFileToLayer(projectObj, cnt.Layer.Name, []string{resourceSpecifications.Path, pathStr}, fileNameStr, templateContentStr)
		if err != nil {
			return err
		}
	}

	_, newResourceModelHash, newLayerSectionModelHash, newTemplateModelHash, err := s.CheckGeneration(cnt.Project, cnt.Resource, cnt.Template, cnt.Section, cnt.Layer)
	if err != nil {
		return err
	}

	generationHistory := entities.NewGenerationHistory(resourceSpecifications.Set, cnt.Resource.GetHeader().Name, newResourceModelHash, cnt.Template.GetHeader().Name, newTemplateModelHash, cnt.Section.Name, newLayerSectionModelHash, cnt.Layer.Name)
	err = s.generationHistoryRepository.Create(generationHistory)
	if err != nil {
		return err
	}

	return nil
}

func (s TemplateOperations) CheckGeneration(project schemas.SchemaInterface, resource schemas.SchemaInterface, template schemas.SchemaInterface, section sectionPkg.SectionIdentifier, layer layerPkg.Layer) (bool, string, string, string, error) {
	var generate = true
	resourceSpecifications := resource.GetSpecification().(applicationResource.ResourceSpecification)
	templateSpecifications := template.GetSpecification().(applicationTemplate.TemplateSpecification)

	history, err := s.generationHistoryRepository.GetLast(templateSpecifications.Set, resource.GetHeader().Name, template.GetHeader().Name, section.Name, layer.Name)
	if err != nil {
		return false, "", "", "", err
	}

	newResourceModelHash, err := encrypt.CalculateHashFromObject(resource)
	if err != nil {
		return false, "", "", "", err
	}

	newLayerSectionModelHash, err := encrypt.CalculateHashFromObject(section)
	if err != nil {
		return false, "", "", "", err
	}

	newTemplateModelHash, err := encrypt.CalculateHashFromObject(template)
	if err != nil {
		return false, "", "", "", err
	}
	if resourceSpecifications.Generate == applicationStructs.ChangeTrackers.Never || templateSpecifications.Generate == applicationStructs.ChangeTrackers.Never {
		generate = false
	} else if resourceSpecifications.Generate == applicationStructs.ChangeTrackers.Always && templateSpecifications.Generate == applicationStructs.ChangeTrackers.Always {
		generate = true
	} else if resourceSpecifications.Generate == applicationStructs.ChangeTrackers.OnCreate || templateSpecifications.Generate == applicationStructs.ChangeTrackers.OnCreate {
		if history != nil {
			generate = false
		}
	} else {
		if history != nil {
			if history.ResourceHash != newResourceModelHash || history.TemplateHash != newTemplateModelHash || history.SectionHash != newLayerSectionModelHash {
				generate = true
			} else {
				generate = false
			}
		} else {
			generate = true
		}
	}

	// if generate {
	// 	if !_string.IsEmpty(templateObj.Configurations.Selectors.Project.Name) {
	// 		generate = false
	// 		if templateObj.Configurations.Selectors.Project.Name == project.GetHeader().Name {
	// 			generate = true
	// 		}
	// 	}
	// }
	// if generate {
	// 	if len(templateObj.Configurations.Selectors.Project.Labels) > 0 {
	// 		generate = false
	// 		for _, selectorLabel := range templateObj.Configurations.Selectors.Project.Labels {
	// 			projectSpecifications := project.GetSpecification().(applicationProject.ProjectSpecification)
	// 			for _, projectLabel := range projectSpecifications.Labels {
	// 				if selectorLabel == projectLabel {
	// 					generate = true
	// 					break
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// if generate {
	// 	if !_string.IsEmpty(templateObj.Configurations.Selectors.Resource.Name) {
	// 		generate = false
	// 		if templateObj.Configurations.Selectors.Resource.Name == resource.GetHeader().Name {
	// 			generate = true
	// 		}
	// 	}
	// }
	// if generate {
	// 	if len(templateObj.Configurations.Selectors.Resource.Labels) > 0 {
	// 		generate = false
	// 		for _, selectorLabel := range templateObj.Configurations.Selectors.Resource.Labels {
	// 			resourceSpecifications := resource.GetSpecification().(applicationResource.ResourceSpecification)
	// 			for _, resourceLabel := range resourceSpecifications.Labels {
	// 				if selectorLabel == resourceLabel {
	// 					generate = true
	// 					break
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// if generate {
	// 	if !_string.IsEmpty(templateObj.Configurations.Selectors.Resource.Section.Name) {
	// 		generate = false
	// 		for _, layerSection := range layer.Sections {
	// 			if templateObj.Configurations.Selectors.Resource.Section.Name == layerSection.Name {
	// 				generate = true
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	// if generate {
	// 	if len(templateObj.Configurations.Selectors.Resource.Section.Classes) > 0 {
	// 		generate = false
	// 		for _, selectorSectionClass := range templateObj.Configurations.Selectors.Resource.Section.Classes {
	// 			for _, layerSection := range layer.Sections {
	// 				for _, layerSectionClass := range layerSection.Classes {
	// 					if selectorSectionClass == layerSectionClass {
	// 						generate = true
	// 						break
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	return generate, newResourceModelHash, newLayerSectionModelHash, newTemplateModelHash, nil
}
func (s TemplateOperations) extractOperationContents(
	project schemas.SchemaInterface,
	template schemas.SchemaInterface,
	resource schemas.SchemaInterface,
	workspace schemas.SchemaInterface,
) []TemplateOperationContent {
	var result []TemplateOperationContent

	resourceSpecifications := resource.GetSpecification().(applicationResource.ResourceSpecification)
	for _, layer := range resourceSpecifications.Layers {
		if len(layer.Sections) == 0 {
			result = append(result, TemplateOperationContent{
				Workspace: workspace,
				Project:   project,
				Resource:  resource,
				Template:  template,
				Layer:     layer,
			})
			continue
		}

		var templateLayer layerPkg.Layer
		templateSpecifications := template.GetSpecification().(applicationTemplate.TemplateSpecification)
		for _, tl := range templateSpecifications.Layers {
			if tl.Name == layer.Name {
				templateLayer = tl
				break
			}
		}

		for _, rSec := range layer.Sections {
			for _, rClass := range rSec.Classes {
				for _, tSec := range templateLayer.Sections {
					for _, tClass := range tSec.Classes {
						if rClass == tClass {
							result = append(result, TemplateOperationContent{
								Workspace: workspace,
								Project:   project,
								Resource:  resource,
								Template:  template,
								Layer:     layer,
								Section:   rSec.SectionIdentifier,
								Class:     rClass,
							})
						}
					}
				}
			}
		}
	}

	return result
}
func layerNames(layers []layerPkg.Layer) []string {
	names := make([]string, len(layers))
	for i, l := range layers {
		names[i] = l.Name
	}
	return names
}
