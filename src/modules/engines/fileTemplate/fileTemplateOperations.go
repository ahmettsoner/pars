package fileTemplate

import (
	"parsdevkit.net/context/models"
	"parsdevkit.net/core/utils"
	"parsdevkit.net/operation/services"
	layerPkg "parsdevkit.net/structs/layer"
	applicationproject "parsdevkit.net/structs/project/application-project"
	dataresource "parsdevkit.net/structs/resource/data-resource"
	filetemplate "parsdevkit.net/structs/template/file-template"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/engines"
	templateEngine "parsdevkit.net/templates/engines"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type FileTemplateOperations struct {
	environment                 string
	generationHistoryRepository *repositories.GenerationHistoryRepository
}

func NewFileTemplateOperations(environment string) FileTemplateOperations {
	return FileTemplateOperations{
		environment:                 environment,
		generationHistoryRepository: repositories.NewGenerationHistoryRepository(environment),
	}
}

func (s FileTemplateOperations) GenerateByResource(model dataresource.ResourceBaseStruct) error {
	workspaceService := services.NewWorkspaceService(s.environment)

	for _, layer := range model.Specifications.Layers {
		templateService := services.NewFileTemplateService(s.environment)
		setTemplates, err := templateService.ListBySetAndLayers(model.Specifications.Set, layer.Name)
		if err != nil {
			return err
		}
		logrus.Debugf("%d Template(s) found for layer '%v' on Resource %v\n", len(*setTemplates), layer.Name, model.Name)

		projectService := services.NewApplicationProjectService(s.environment)
		setProjects, err := projectService.ListBySetAndLayers(model.Specifications.Set, layer.Name)
		if err != nil {
			return err
		}

		for _, setProject := range *setProjects {
			projectWorkspace, err := workspaceService.GetByName(setProject.Specifications.Workspace)
			if err != nil {
				return err
			}

			// templateService := dataResourceService.NewDataResourceService(setProject.Specifications.Platform.Type)
			for _, setTemplate := range *setTemplates {
				err := s.GenerateContent(*projectWorkspace, setProject, model, setTemplate, layer.LayerIdentifier)
				if err != nil {
					return err
				}
			}
		}

		logrus.Debugf("%d Project(s) found for layer '%v' on Resource %v\n", len(*setProjects), layer.Name, model.Name)
	}
	return nil
}

func (s FileTemplateOperations) GenerateByTemplate(model filetemplate.TemplateBaseStruct) error {
	workspaceService := services.NewWorkspaceService(s.environment)

	for _, modelLayer := range model.Specifications.Layers {
		resourceService := services.NewDataResourceService(s.environment)
		setResources, err := resourceService.ListBySetAndLayers(model.Specifications.Set, modelLayer.Name)
		if err != nil {
			return err
		}
		logrus.Debugf("%d Resource(s) found for layer '%v' on Template %v\n", len(*setResources), modelLayer.Name, model.Name)

		projectService := services.NewApplicationProjectService(s.environment)
		setProjects, err := projectService.ListBySetAndLayers(model.Specifications.Set, modelLayer.Name)
		if err != nil {
			return err
		}

		for _, setProject := range *setProjects {
			projectWorkspace, err := workspaceService.GetByName(setProject.Specifications.Workspace)
			if err != nil {
				return err
			}

			for _, setResource := range *setResources {

				err := s.GenerateContent(*projectWorkspace, setProject, setResource, model, modelLayer.LayerIdentifier)
				if err != nil {
					return err
				}
			}
		}

		logrus.Debugf("%d Project(s) found for layer '%v' on Template %v\n", len(*setProjects), modelLayer.Name, model.Name)
	}
	return nil
}

func (s FileTemplateOperations) GenerateContent(workspace workspace.WorkspaceBaseStruct, project applicationproject.ProjectBaseStruct, resource dataresource.ResourceBaseStruct, template filetemplate.TemplateBaseStruct, layer layerPkg.LayerIdentifier) error {
	projectService := services.NewApplicationProjectService(s.environment)

	resourceLayer := dataresource.Layer{}

	for _, selectedResourceLayer := range resource.Specifications.Layers {
		if selectedResourceLayer.LayerIdentifier == layer {
			resourceLayer = selectedResourceLayer
		}
	}

	generate, newResourceModelHash, newResourceSectionModelHash, newTemplateModelHash, err := s.CheckGeneration(project, resource, template, dataresource.Section{}, resourceLayer)
	if err != nil {
		return err
	}

	if generate {

		var data = models.NewFileTemplateDataContext(workspace, project, resource, template, resourceLayer)

		fileNameStr, err := templateEngine.TemplateEngine(template.Specifications.Output.File, data)
		if err != nil {
			return err
		}
		pathStr, err := templateEngine.TemplateEngine(template.Specifications.Path, data)
		if err != nil {
			return err
		}

		data = models.NewFileTemplateDataContext(workspace, project, resource, template, resourceLayer)
		templateContentStr, err := templateEngine.TemplateEngine(template.Specifications.Template.Content, data)
		if err != nil {
			return err
		}

		templateContentStr = engines.AddCommentToGeneratedFile(template.Specifications.Output.File, string(resource.Configurations.Generate), string(template.Configurations.Generate), templateContentStr)

		_, err = projectService.AddFileToLayer(project, layer.Name, []string{resource.Specifications.Path, pathStr}, fileNameStr, templateContentStr)
		if err != nil {
			return err
		}

		generationHistory := entities.NewGenerationHistory(resource.Specifications.Set, resource.Name, newResourceModelHash, template.Name, newTemplateModelHash, "", newResourceSectionModelHash, layer.Name)
		err = s.generationHistoryRepository.Create(generationHistory)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s FileTemplateOperations) CheckGeneration(project applicationproject.ProjectBaseStruct, resource dataresource.ResourceBaseStruct, template filetemplate.TemplateBaseStruct, section dataresource.Section, layer dataresource.Layer) (bool, string, string, string, error) {
	var generate = true

	history, err := s.generationHistoryRepository.GetLast(template.Specifications.Set, resource.Name, template.Name, section.Name, layer.Name)
	if err != nil {
		return false, "", "", "", err
	}

	newResourceModelHash, err := utils.CalculateHashFromObject(resource)
	if err != nil {
		return false, "", "", "", err
	}

	newResourceSectionModelHash, err := utils.CalculateHashFromObject(section)
	if err != nil {
		return false, "", "", "", err
	}

	newTemplateModelHash, err := utils.CalculateHashFromObject(template)
	if err != nil {
		return false, "", "", "", err
	}

	if resource.Configurations.Generate == dataresource.ChangeTrackers.Never || template.Configurations.Generate == filetemplate.ChangeTrackers.Never {
		generate = false
	} else if resource.Configurations.Generate == dataresource.ChangeTrackers.Always && template.Configurations.Generate == filetemplate.ChangeTrackers.Always {
		generate = true
	} else if resource.Configurations.Generate == dataresource.ChangeTrackers.OnCreate || template.Configurations.Generate == filetemplate.ChangeTrackers.OnCreate {
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
		if !utils.IsEmpty(template.Configurations.Selectors.Project.Name) {
			generate = false
			if template.Configurations.Selectors.Project.Name == project.Name {
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
		if !utils.IsEmpty(template.Configurations.Selectors.Resource.Name) {
			generate = false
			if template.Configurations.Selectors.Resource.Name == resource.Name {
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
		if !utils.IsEmpty(template.Configurations.Selectors.Resource.Section.Name) {
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
