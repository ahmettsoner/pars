package codeTemplate

import (
	"strings"

	"parsdevkit.net/core/utils"
	"parsdevkit.net/operation/services"
	layerPkg "parsdevkit.net/structs/layer"
	applicationproject "parsdevkit.net/structs/project/application-project"
	objectresource "parsdevkit.net/structs/resource/object-resource"
	codetemplate "parsdevkit.net/structs/template/code-template"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/engines"
	templateEngine "parsdevkit.net/templates/engines"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"parsdevkit.net/context/models"

	"github.com/sirupsen/logrus"
)

type CodeTemplateOperations struct {
	environment                 string
	generationHistoryRepository *repositories.GenerationHistoryRepository
}

func NewCodeTemplateOperations(environment string) CodeTemplateOperations {
	return CodeTemplateOperations{
		environment:                 environment,
		generationHistoryRepository: repositories.NewGenerationHistoryRepository(environment),
	}
}

func (s CodeTemplateOperations) GenerateByResource(model objectresource.ResourceBaseStruct) error {
	workspaceService := services.NewWorkspaceService(s.environment)

	for _, layer := range model.Specifications.Layers {
		templateService := services.NewCodeTemplateService(s.environment)
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

func (s CodeTemplateOperations) GenerateByTemplate(model codetemplate.TemplateBaseStruct) error {
	workspaceService := services.NewWorkspaceService(s.environment)

	for _, modelLayer := range model.Specifications.Layers {
		resourceService := services.NewObjectResourceService(s.environment)
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

		//TODO: selector işlemleri bu noktada gerçekleştirilebilir?
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

func (s CodeTemplateOperations) GenerateContent(workspace workspace.WorkspaceBaseStruct, project applicationproject.ProjectBaseStruct, resource objectresource.ResourceBaseStruct, template codetemplate.TemplateBaseStruct, layer layerPkg.LayerIdentifier) error {
	projectService := services.NewApplicationProjectService(s.environment)

	resourceLayer := objectresource.Layer{}

	for _, selectedResourceLayer := range resource.Specifications.Layers {
		if selectedResourceLayer.LayerIdentifier == layer {
			resourceLayer = selectedResourceLayer
		}
	}

	if len(resourceLayer.Sections) > 0 {

		for _, resourceLayerSection := range resourceLayer.Sections {
			for _, resourceLayerSectionClass := range resourceLayerSection.Classes {

				for _, templateLayer := range template.Specifications.Layers {
					if templateLayer.Name == resourceLayer.Name {
						for _, templateLayerSection := range templateLayer.Sections {
							for _, templateLayerSectionClass := range templateLayerSection.Classes {
								if resourceLayerSectionClass == templateLayerSectionClass {
									generate, newResourceModelHash, newLayerSectionModelHash, newTemplateModelHash, err := s.CheckGeneration(project, resource, template, resourceLayerSection, resourceLayer)
									if err != nil {
										return err
									}
									if generate {

										var data = models.NewCodeTemplateDataContext(workspace, project, resource, template, resourceLayer, resourceLayerSection)

										fileNameStr, err := templateEngine.TemplateEngine(template.Specifications.Output.File, data)
										if err != nil {
											return err
										}
										pathStr, err := templateEngine.TemplateEngine(template.Specifications.Path, data)
										if err != nil {
											return err
										}

										tempPackages := template.Specifications.Package
										packageStr, err := templateEngine.TemplateEngine(strings.Join(tempPackages, "/"), data)
										if err != nil {
											return err
										}
										template.Specifications.Package = utils.PathToArray(packageStr)

										data = models.NewCodeTemplateDataContext(workspace, project, resource, template, resourceLayer, resourceLayerSection)
										templateContentStr, err := templateEngine.TemplateEngine(template.Specifications.Template.Content, data)
										if err != nil {
											return err
										}
										template.Specifications.Package = tempPackages

										templateContentStr = engines.AddCommentToGeneratedFile(template.Specifications.Output.File, string(resource.Configurations.Generate), string(template.Configurations.Generate), templateContentStr)

										_, err = projectService.AddFileToLayer(project, layer.Name, []string{resource.Specifications.Path, pathStr}, fileNameStr, templateContentStr)
										if err != nil {
											return err
										}

										generationHistory := entities.NewGenerationHistory(resource.Specifications.Set, resource.Name, newResourceModelHash, template.Name, newTemplateModelHash, resourceLayerSection.Name, newLayerSectionModelHash, layer.Name)
										err = s.generationHistoryRepository.Create(generationHistory)
										if err != nil {
											return err
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {

		generate, newResourceModelHash, newLayerSectionModelHash, newTemplateModelHash, err := s.CheckGeneration(project, resource, template, objectresource.Section{}, resourceLayer)
		if err != nil {
			return err
		}
		if generate {
			var data = models.NewCodeTemplateDataContext(workspace, project, resource, template, resourceLayer, objectresource.Section{})

			fileNameStr, err := templateEngine.TemplateEngine(template.Specifications.Output.File, data)
			if err != nil {
				return err
			}
			pathStr, err := templateEngine.TemplateEngine(template.Specifications.Path, data)
			if err != nil {
				return err
			}

			tempPackages := template.Specifications.Package
			packageStr, err := templateEngine.TemplateEngine(strings.Join(tempPackages, "/"), data)
			if err != nil {
				return err
			}
			template.Specifications.Package = utils.PathToArray(packageStr)

			data = models.NewCodeTemplateDataContext(workspace, project, resource, template, resourceLayer, objectresource.Section{})
			templateContentStr, err := templateEngine.TemplateEngine(template.Specifications.Template.Content, data)
			if err != nil {
				return err
			}
			template.Specifications.Package = tempPackages

			templateContentStr = engines.AddCommentToGeneratedFile(template.Specifications.Output.File, string(resource.Configurations.Generate), string(template.Configurations.Generate), templateContentStr)

			_, err = projectService.AddFileToLayer(project, layer.Name, []string{resource.Specifications.Path, pathStr}, fileNameStr, templateContentStr)
			if err != nil {
				return err
			}

			generationHistory := entities.NewGenerationHistory(resource.Specifications.Set, resource.Name, newResourceModelHash, template.Name, newTemplateModelHash, "", newLayerSectionModelHash, layer.Name)
			err = s.generationHistoryRepository.Create(generationHistory)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s CodeTemplateOperations) CheckGeneration(project applicationproject.ProjectBaseStruct, resource objectresource.ResourceBaseStruct, template codetemplate.TemplateBaseStruct, section objectresource.Section, layer objectresource.Layer) (bool, string, string, string, error) {
	var generate = true

	history, err := s.generationHistoryRepository.GetLast(template.Specifications.Set, resource.Name, template.Name, section.Name, layer.Name)
	if err != nil {
		return false, "", "", "", err
	}

	newResourceModelHash, err := utils.CalculateHashFromObject(resource)
	if err != nil {
		return false, "", "", "", err
	}

	newLayerSectionModelHash, err := utils.CalculateHashFromObject(section)
	if err != nil {
		return false, "", "", "", err
	}

	newTemplateModelHash, err := utils.CalculateHashFromObject(template)
	if err != nil {
		return false, "", "", "", err
	}

	if resource.Configurations.Generate == objectresource.ChangeTrackers.Never || template.Configurations.Generate == codetemplate.ChangeTrackers.Never {
		generate = false
	} else if resource.Configurations.Generate == objectresource.ChangeTrackers.Always && template.Configurations.Generate == codetemplate.ChangeTrackers.Always {
		generate = true
	} else if resource.Configurations.Generate == objectresource.ChangeTrackers.OnCreate || template.Configurations.Generate == codetemplate.ChangeTrackers.OnCreate {
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

	return generate, newResourceModelHash, newLayerSectionModelHash, newTemplateModelHash, nil
}
