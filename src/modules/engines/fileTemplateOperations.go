package engines

import (
	"parsdevkit.net/application/contextgenerator"
	layerPkg "parsdevkit.net/application/models/layer"
	sectionPkg "parsdevkit.net/application/models/section"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	"parsdevkit.net/pkg/utilities/encrypt"

	"parsdevkit.net/application/ioc"
	templatePkg "parsdevkit.net/components/template"
	templateEngine "parsdevkit.net/components/template/engines"
	"parsdevkit.net/modules/project/application_project_contract"
	"parsdevkit.net/modules/resource/data_resource_contract"
	"parsdevkit.net/modules/template/file_template_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"

	"parsdevkit.net/persistence/contexts"

	"parsdevkit.net/persistence/repositories"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type FileTemplateOperations struct {
	environment                 string
	generationHistoryRepository *repositories.GenerationHistoryRepository
}

func NewFileTemplateOperations(environment string) FileTemplateOperations {
	dbContext := contexts.NewDbContext(environment)
	return FileTemplateOperations{
		environment:                 environment,
		generationHistoryRepository: repositories.NewGenerationHistoryRepository(dbContext),
	}
}

func (s FileTemplateOperations) GenerateByResource(model data_resource_payload_structs.ResourceBaseStruct) error {
	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	layers := make([]string, 0)
	for _, modelLayer := range model.Specifications.Layers {
		layers = append(layers, modelLayer.Name)
	}

	templateService := ioc.Get[file_template_contract.TemplateInterface]()
	templates, err := templateService.ListByFilter(model.Specifications.Set, model.Specifications.Workspace, layers, model.Header.Metadata.Tags, model.Specifications.Labels)
	if err != nil {
		return err
	}
	logrus.Debugf("%d Template(s) found for layer '%v' \n", len(*templates), model.Header.Name)

	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	projects, err := projectService.ListByFilter(model.Specifications.Set, model.Specifications.Workspace, layers, model.Header.Metadata.Tags, model.Specifications.Labels)
	if err != nil {
		return err
	}
	logrus.Debugf("%d Projcet(s) found for layer '%v' \n", len(*projects), model.Header.Name)

	for _, layer := range model.Specifications.Layers {

		for _, setProject := range *projects {
			projectWorkspace, err := workspaceService.GetByName(setProject.Specifications.Workspace)
			if err != nil {
				return err
			}

			for _, template := range *templates {
				err := s.GenerateContent(*projectWorkspace, setProject, model, template, layer.LayerIdentifier)
				if err != nil {
					return err
				}
			}
		}
	}

	for _, layer := range model.Specifications.Layers {
		templateService := ioc.Get[file_template_contract.TemplateInterface]()
		setTemplates, err := templateService.ListBySetAndLayers(model.Specifications.Set, layer.Name)
		if err != nil {
			return err
		}
		logrus.Debugf("%d Template(s) found for layer '%v' on Resource %v\n", len(*setTemplates), layer.Name, model.Header.Name)

		projectService := ioc.Get[application_project_contract.ProjectInterface]()
		setProjects, err := projectService.ListBySetAndLayers(model.Specifications.Set, layer.Name)
		if err != nil {
			return err
		}

		for _, setProject := range *setProjects {
			projectWorkspace, err := workspaceService.GetByName(setProject.Specifications.Workspace)
			if err != nil {
				return err
			}

			for _, setTemplate := range *setTemplates {
				err := s.GenerateContent(*projectWorkspace, setProject, model, setTemplate, layer.LayerIdentifier)
				if err != nil {
					return err
				}
			}
		}

		logrus.Debugf("%d Project(s) found for layer '%v' on Resource %v\n", len(*setProjects), layer.Name, model.Header.Name)
	}
	return nil
}

func (s FileTemplateOperations) GenerateByTemplate(model file_template_payload_structs.TemplateBaseStruct) error {
	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	layers := make([]string, 0)
	for _, modelLayer := range model.Specifications.Layers {
		layers = append(layers, modelLayer.Name)
	}

	resourceService := ioc.Get[data_resource_contract.ResourceInterface]()
	resources, err := resourceService.ListByFilter(model.Specifications.Set, model.Specifications.Workspace, layers, model.Header.Metadata.Tags, model.Specifications.Labels)
	if err != nil {
		return err
	}
	logrus.Debugf("%d Resource(s) found for layer '%v' \n", len(*resources), model.Header.Name)

	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	projects, err := projectService.ListByFilter(model.Specifications.Set, model.Specifications.Workspace, layers, model.Header.Metadata.Tags, model.Specifications.Labels)
	if err != nil {
		return err
	}
	logrus.Debugf("%d Projcet(s) found for layer '%v' \n", len(*projects), model.Header.Name)

	for _, modelLayer := range model.Specifications.Layers {
		for _, project := range *projects {
			projectWorkspace, err := workspaceService.GetByName(project.Specifications.Workspace)
			if err != nil {
				return err
			}

			for _, resource := range *resources {
				err := s.GenerateContent(*projectWorkspace, project, resource, model, modelLayer.LayerIdentifier)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s FileTemplateOperations) GenerateContent(workspace basic_workspace_payload_structs.WorkspaceBaseStruct, project application_project_payload_structs.ProjectBaseStruct, resource data_resource_payload_structs.ResourceBaseStruct, template file_template_payload_structs.TemplateBaseStruct, layer layerPkg.LayerIdentifier) error {
	projectService := ioc.Get[application_project_contract.ProjectInterface]()

	resourceLayer := layerPkg.Layer{}

	for _, selectedResourceLayer := range resource.Specifications.Layers {
		if selectedResourceLayer.LayerIdentifier == layer {
			resourceLayer = selectedResourceLayer
		}
	}

	generate, newResourceModelHash, newResourceSectionModelHash, newTemplateModelHash, err := s.CheckGeneration(project, resource, template, sectionPkg.SectionIdentifier{}, resourceLayer)
	if err != nil {
		return err
	}

	if generate {

		var data = contextgenerator.NewTemplateDataContext(
			templatePkg.NewContextProviderSource(workspace, nil, project, resource, template, resourceLayer.LayerIdentifier, sectionPkg.SectionIdentifier{}),
		)

		fileNameStr, err := templateEngine.RenderTemplate(template.Specifications.Output.File, data)
		if err != nil {
			return err
		}
		pathStr, err := templateEngine.RenderTemplate(template.Specifications.Path, data)
		if err != nil {
			return err
		}

		// data = contextgenerator.NewTemplateDataContext(
		// 	templatePkg.NewContextProviderSource(workspace, nil, project, resource, template, resourceLayer.LayerIdentifier, sectionPkg.SectionIdentifier{}),
		// )
		templateContentStr, err := templateEngine.RenderTemplate(template.Specifications.Template.Content, data)
		if err != nil {
			return err
		}

		templateContentStr = AddCommentToGeneratedFile(template.Specifications.Output.File, string(resource.Configurations.Generate), string(template.Configurations.Generate), templateContentStr)

		_, err = projectService.AddFileToLayer(project, layer.Name, []string{resource.Specifications.Path, pathStr}, fileNameStr, templateContentStr)
		if err != nil {
			return err
		}

		generationHistory := entities.NewGenerationHistory(resource.Specifications.Set, resource.Header.Name, newResourceModelHash, template.Header.Name, newTemplateModelHash, "", newResourceSectionModelHash, layer.Name)
		err = s.generationHistoryRepository.Create(generationHistory)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s FileTemplateOperations) CheckGeneration(project application_project_payload_structs.ProjectBaseStruct, resource data_resource_payload_structs.ResourceBaseStruct, template file_template_payload_structs.TemplateBaseStruct, section sectionPkg.SectionIdentifier, layer layerPkg.Layer) (bool, string, string, string, error) {
	var generate = true

	history, err := s.generationHistoryRepository.GetLast(template.Specifications.Set, resource.Header.Name, template.Header.Name, section.Name, layer.Name)
	if err != nil {
		return false, "", "", "", err
	}

	newResourceModelHash, err := encrypt.CalculateHashFromObject(resource)
	if err != nil {
		return false, "", "", "", err
	}

	newResourceSectionModelHash, err := encrypt.CalculateHashFromObject(section)
	if err != nil {
		return false, "", "", "", err
	}

	newTemplateModelHash, err := encrypt.CalculateHashFromObject(template)
	if err != nil {
		return false, "", "", "", err
	}

	if resource.Configurations.Generate == data_resource_payload_structs.ChangeTrackers.Never || template.Configurations.Generate == file_template_payload_structs.ChangeTrackers.Never {
		generate = false
	} else if resource.Configurations.Generate == data_resource_payload_structs.ChangeTrackers.Always && template.Configurations.Generate == file_template_payload_structs.ChangeTrackers.Always {
		generate = true
	} else if resource.Configurations.Generate == data_resource_payload_structs.ChangeTrackers.OnCreate || template.Configurations.Generate == file_template_payload_structs.ChangeTrackers.OnCreate {
		if history != nil {
			generate = false
		}
	} else {
		if history != nil {
			if history.ResourceHash != newResourceModelHash || history.TemplateHash != newTemplateModelHash || history.SectionHash != newResourceSectionModelHash {
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

	return generate, newResourceModelHash, newResourceSectionModelHash, newTemplateModelHash, nil
}
