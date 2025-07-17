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
	templatePkg "parsdevkit.net/components/template"
	templateEngine "parsdevkit.net/components/template/engines"
	"parsdevkit.net/modules/project/application_project_contract"
	"parsdevkit.net/modules/resource/object_resource_contract"
	"parsdevkit.net/modules/template/code_template_contract"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
	"parsdevkit.net/pkg/utilities/encrypt"
	"parsdevkit.net/pkg/utilities/file"

	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/persistence/contexts"
	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"parsdevkit.net/application/schemas"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
)

type CodeTemplateOperationContent struct {
	Workspace basic_workspace_payload_structs.WorkspaceBaseStruct
	Project   application_project_payload_structs.ProjectBaseStruct
	Resource  object_resource_payload_structs.ResourceBaseStruct
	Template  code_template_payload_structs.TemplateBaseStruct
	Layer     layerPkg.Layer
	Section   sectionPkg.SectionIdentifier
	Class     class.Class
}

type CodeTemplateOperations struct {
	environment                 string
	generationHistoryRepository *repositories.GenerationHistoryRepository
}

func NewCodeTemplateOperations(environment string) CodeTemplateOperations {
	dbContext := contexts.NewDbContext(environment)
	return CodeTemplateOperations{
		environment:                 environment,
		generationHistoryRepository: repositories.NewGenerationHistoryRepository(dbContext),
	}
}

func (s CodeTemplateOperations) Generate(source schemas.SchemaInterface) error {

	var opCnt []CodeTemplateOperationContent
	var workspace *basic_workspace_payload_structs.WorkspaceBaseStruct
	var err error

	var setName, workspaceName string
	var layers []layerPkg.Layer
	var tags []string
	var labels []label.Label
	var projects *[]application_project_payload_structs.ProjectBaseStruct
	var templates *[]code_template_payload_structs.TemplateBaseStruct
	var resources *[]object_resource_payload_structs.ResourceBaseStruct

	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	templateService := ioc.Get[code_template_contract.TemplateInterface]()
	resourceService := ioc.Get[object_resource_contract.ResourceInterface]()
	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()

	switch src := source.(type) {
	case object_resource_payload_structs.ResourceBaseStruct:
		setName = src.Specifications.Set
		workspaceName = src.Specifications.Workspace
		layers = src.Specifications.Layers
		tags = src.Header.Metadata.Tags
		labels = src.Specifications.Labels

		workspace, err = workspaceService.GetByName(workspaceName)
		if err != nil {
			return err
		}

		projects, err = projectService.ListByFilter(setName, workspaceName, layerNames(layers), tags, labels)
		if err != nil {
			return err
		}
		templates, err = templateService.ListByFilter(setName, workspaceName, layerNames(layers), tags, labels)
		if err != nil {
			return err
		}
		for _, project := range *projects {
			for _, template := range *templates {
				opCnt = append(opCnt, s.extractOperationContents(project, template, src, *workspace)...)
			}
		}

	case code_template_payload_structs.TemplateBaseStruct:
		setName = src.Specifications.Set
		workspaceName = src.Specifications.Workspace
		layers = src.Specifications.Layers
		tags = src.Header.Metadata.Tags
		labels = src.Specifications.Labels

		workspace, err = workspaceService.GetByName(workspaceName)
		if err != nil {
			return err
		}

		projects, err = projectService.ListByFilter(setName, workspaceName, layerNames(layers), tags, labels)
		if err != nil {
			return err
		}
		resources, err = resourceService.ListByFilter(setName, workspaceName, layerNames(layers), tags, labels)
		if err != nil {
			return err
		}

		for _, project := range *projects {
			for _, resource := range *resources {
				opCnt = append(opCnt, s.extractOperationContents(project, src, resource, *workspace)...)
			}
		}
	default:
		return fmt.Errorf("unsupported type: %T", source)
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
func (s CodeTemplateOperations) GenerateByResource(model object_resource_payload_structs.ResourceBaseStruct) error {
	return s.Generate(model)
}
func (s CodeTemplateOperations) GenerateContent(cnt CodeTemplateOperationContent) error {
	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	var data = contextgenerator.NewTemplateDataContext(
		templatePkg.NewContextProviderSource(
			&cnt.Workspace, nil, &cnt.Project, &cnt.Resource, &cnt.Template, cnt.Layer.LayerIdentifier, cnt.Section),
	)

	fileNameStr, err := templateEngine.RenderTemplate(cnt.Template.Specifications.Output.File, data)
	if err != nil {
		return err
	}
	pathStr, err := templateEngine.RenderTemplate(cnt.Template.Specifications.Path, data)
	if err != nil {
		return err
	}

	tempPackages := cnt.Template.Specifications.Package
	packageStr, err := templateEngine.RenderTemplate(strings.Join(tempPackages, "/"), data)
	if err != nil {
		return err
	}
	cnt.Template.Specifications.Package = file.PathToArray(packageStr)

	data = contextgenerator.NewTemplateDataContext(
		templatePkg.NewContextProviderSource(
			cnt.Workspace, nil, cnt.Project, cnt.Resource, cnt.Template, cnt.Layer.LayerIdentifier, cnt.Section),
	)
	templateContentStr, err := templateEngine.RenderTemplate(cnt.Template.Specifications.Template.Content, data)
	if err != nil {
		return err
	}
	cnt.Template.Specifications.Package = tempPackages

	templateContentStr = AddCommentToGeneratedFile(cnt.Template.Specifications.Output.File, string(cnt.Resource.Configurations.Generate), string(cnt.Template.Configurations.Generate), templateContentStr)

	_, err = projectService.AddFileToLayer(cnt.Project, cnt.Layer.Name, []string{cnt.Resource.Specifications.Path, pathStr}, fileNameStr, templateContentStr)
	if err != nil {
		return err
	}

	_, newResourceModelHash, newLayerSectionModelHash, newTemplateModelHash, err := s.CheckGeneration(cnt.Project, cnt.Resource, cnt.Template, cnt.Section, cnt.Layer)
	if err != nil {
		return err
	}
	generationHistory := entities.NewGenerationHistory(cnt.Resource.Specifications.Set, cnt.Resource.Header.Name, newResourceModelHash, cnt.Template.Header.Name, newTemplateModelHash, cnt.Section.Name, newLayerSectionModelHash, cnt.Layer.Name)
	err = s.generationHistoryRepository.Create(generationHistory)
	if err != nil {
		return err
	}

	return nil
}

func (s CodeTemplateOperations) GenerateByTemplate(model code_template_payload_structs.TemplateBaseStruct) error {
	return s.Generate(model)
}

func (s CodeTemplateOperations) CheckGeneration(project application_project_payload_structs.ProjectBaseStruct, resource object_resource_payload_structs.ResourceBaseStruct, template code_template_payload_structs.TemplateBaseStruct, section sectionPkg.SectionIdentifier, layer layerPkg.Layer) (bool, string, string, string, error) {
	var generate = true

	history, err := s.generationHistoryRepository.GetLast(template.Specifications.Set, resource.Header.Name, template.Header.Name, section.Name, layer.Name)
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

	if resource.Configurations.Generate == object_resource_payload_structs.ChangeTrackers.Never || template.Configurations.Generate == code_template_payload_structs.ChangeTrackers.Never {
		generate = false
	} else if resource.Configurations.Generate == object_resource_payload_structs.ChangeTrackers.Always && template.Configurations.Generate == code_template_payload_structs.ChangeTrackers.Always {
		generate = true
	} else if resource.Configurations.Generate == object_resource_payload_structs.ChangeTrackers.OnCreate || template.Configurations.Generate == code_template_payload_structs.ChangeTrackers.OnCreate {
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

	if generate {
		if !_string.IsEmpty(template.Configurations.Selectors.Project.Name) {
			generate = false
			if template.Configurations.Selectors.Project.Name == project.Header.Name {
				generate = true
			}
		}
	}
	if generate {
		if len(template.Configurations.Selectors.Project.Labels) > 0 {
			generate = false
			for _, selectorLabel := range template.Configurations.Selectors.Project.Labels {
				for _, projectLabel := range project.Specifications.Labels {
					if selectorLabel == projectLabel {
						generate = true
						break
					}
				}
			}
		}
	}

	if generate {
		if !_string.IsEmpty(template.Configurations.Selectors.Resource.Name) {
			generate = false
			if template.Configurations.Selectors.Resource.Name == resource.Header.Name {
				generate = true
			}
		}
	}
	if generate {
		if len(template.Configurations.Selectors.Resource.Labels) > 0 {
			generate = false
			for _, selectorLabel := range template.Configurations.Selectors.Resource.Labels {
				for _, resourceLabel := range resource.Specifications.Labels {
					if selectorLabel == resourceLabel {
						generate = true
						break
					}
				}
			}
		}
	}

	if generate {
		if !_string.IsEmpty(template.Configurations.Selectors.Resource.Section.Name) {
			generate = false
			for _, layerSection := range layer.Sections {
				if template.Configurations.Selectors.Resource.Section.Name == layerSection.Name {
					generate = true
					break
				}
			}
		}
	}
	if generate {
		if len(template.Configurations.Selectors.Resource.Section.Classes) > 0 {
			generate = false
			for _, selectorSectionClass := range template.Configurations.Selectors.Resource.Section.Classes {
				for _, layerSection := range layer.Sections {
					for _, layerSectionClass := range layerSection.Classes {
						if selectorSectionClass == layerSectionClass {
							generate = true
							break
						}
					}
				}
			}
		}
	}

	return generate, newResourceModelHash, newLayerSectionModelHash, newTemplateModelHash, nil
}
func (s CodeTemplateOperations) extractOperationContents(
	project application_project_payload_structs.ProjectBaseStruct,
	template code_template_payload_structs.TemplateBaseStruct,
	resource object_resource_payload_structs.ResourceBaseStruct,
	workspace basic_workspace_payload_structs.WorkspaceBaseStruct,
) []CodeTemplateOperationContent {
	var result []CodeTemplateOperationContent

	for _, layer := range resource.Specifications.Layers {
		if len(layer.Sections) == 0 {
			result = append(result, CodeTemplateOperationContent{
				Workspace: workspace,
				Project:   project,
				Resource:  resource,
				Template:  template,
				Layer:     layer,
			})
			continue
		}

		var templateLayer layerPkg.Layer
		for _, tl := range template.Specifications.Layers {
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
							result = append(result, CodeTemplateOperationContent{
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
