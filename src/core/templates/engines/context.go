package engines

import (
	"fmt"
	"reflect"
	"strings"

	"parsdevkit.net/core/utils"
	"parsdevkit.net/operation/services"
	applicationproject "parsdevkit.net/structs/project/application-project"
	objectresource "parsdevkit.net/structs/resource/object-resource"
	codetemplate "parsdevkit.net/structs/template/code-template"
	"parsdevkit.net/context/models"
)

type ContextFuncs struct{}

func (c ContextFuncs) GetContextByBaseForArray(base models.CodeTemplateDataContext, args []string) models.CodeTemplateDataContext {
	return c.GetContextByBase(base, args...)
}
func (c ContextFuncs) GetContextByBase(base models.CodeTemplateDataContext, args ...string) models.CodeTemplateDataContext {
	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	applicationProjectService := services.NewApplicationProjectService(utils.GetEnvironment())
	objectResourceService := services.NewObjectResourceService(utils.GetEnvironment())
	codeTemplateService := services.NewCodeTemplateService(utils.GetEnvironment())

	/*
		Eğer set tanımlı değilse
			base'den al

		Eğer resource tanımlı değilse
			base'den al
		Tanımlıysa
			db'den getir

		Eğer layer tanımlı değilse
			base'den al
		Tanımlıysa
			db'den getir

		Eğer proje tanımlı değilse
			ilgili sette, belirtilen layer'a sahip proje'leri listele
		Tanımlıysa
			belirtilen proje'yi al

		Eğer section tanımlı değilse
			resource'u baz al
		Tanımlıysa
			resource'ta doğrula

		Eğer template tanımlı değilse
			ilgili sette, belirtilen layer'a sahip template'leri listele
			önceden tespit edilen resource'a uygun olmayanları ayır (selector)
			section varsa class'ına göre uygun olmayanları al
		Tanımlıysa
			belirtilen template'i al
			resource'a uygunluğunu doğrula (layer, varsa selector)
			Eğer section varsa
				ilgili template'in class'ına göre uygun olmayanları ayırla
				section'a göre uygun olanları al


			CodeTemplateOperations.PopulateContext ile context bilgisini al
	*/

	workspace := base.Workspace.Original.Name
	set := base.Project.Original.Specifications.Set
	project := ""
	resource := base.Resource.Original.Name
	layer := base.Layer.Original.Name
	section := ""
	template := ""
	// templateSelector := ""

	for _, arg := range args {
		argParts := strings.Split(arg, "::")
		if len(argParts) == 2 {
			argKey := strings.ToLower(argParts[0])
			argValue := argParts[1]

			switch argKey {
			// case "workspace":
			// 	workspace = argValue
			case "set":
				set = argValue
			case "project":
				project = argValue
				set = ""
			case "resource":
				resource = argValue
			case "template":
				template = argValue
			case "layer":
				layer = argValue
			case "section":
				section = argValue
			default:
				fmt.Printf("Unknown argument: %s\n", arg)
			}
		}
	}

	//TODO: burda splitsiz sıralı argument çözümleme hazırlanacak
	// if len(args) > 0 {
	// 	layer = args[0]
	// }

	workspaceObj, err := workspaceService.GetByName(workspace)
	if err != nil {
		return models.CodeTemplateDataContext{}
	}

	//TODO: Burda resource, template selector yapısı, project, resoruce ve template için layer ve set kontrollri daha sonra eklenecek
	resourceObj, err := objectResourceService.GetByName(resource)
	if err != nil {
		return models.CodeTemplateDataContext{}
	}

	var layerObj *objectresource.Layer = nil
	if resourceObj != nil {
		for _, resourceLayer := range resourceObj.Specifications.Layers {
			if resourceLayer.Name == layer {
				layerObj = &resourceLayer
				break
			}
		}
	}

	var projectObj *applicationproject.ProjectBaseStruct = nil
	if layerObj != nil {
		var projectList []applicationproject.ProjectBaseStruct = make([]applicationproject.ProjectBaseStruct, 0)
		if utils.IsEmpty(project) {
			projectListFromDb, err := applicationProjectService.ListBySetAndLayers(set, layer)
			if err != nil {
				return models.CodeTemplateDataContext{}
			}
			projectList = *projectListFromDb
		} else {
			projectObj, err := applicationProjectService.GetByName(project)
			if err != nil {
				return models.CodeTemplateDataContext{}
			}
			projectList = append(projectList, *projectObj)
		}

		if len(projectList) > 0 {
			for _, projectObjFromDb := range projectList {
				for _, objLayer := range projectObjFromDb.Specifications.Configuration.Layers {
					if objLayer.Name == layer {
						projectObj = &projectObjFromDb
						break
					}
				}

				if layerObj != nil {
					break
				}
			}
		}
	}

	if layerObj != nil {

		var templateObj *codetemplate.TemplateBaseStruct = nil

		var templatelist []codetemplate.TemplateBaseStruct = make([]codetemplate.TemplateBaseStruct, 0)
		if utils.IsEmpty(template) {
			templateListFromDb, err := codeTemplateService.ListBySetAndLayers(set, layer)
			if err != nil {
				return models.CodeTemplateDataContext{}
			}
			templatelist = *templateListFromDb
		} else {
			templateObjFromDb, err := codeTemplateService.GetByName(project)
			if err != nil {
				return models.CodeTemplateDataContext{}
			}
			templatelist = append(templatelist, *templateObjFromDb)
		}

		if len(templatelist) > 0 {
			for _, templateObjFromDb := range templatelist {
				for _, tmpLayer := range templateObjFromDb.Specifications.Layers {
					if tmpLayer.Name == layer {
						templateObj = &templateObjFromDb
						break
					}
				}
			}
		}

		if templateObj != nil {
			if !utils.IsEmpty(section) {
				selectedContext := models.CodeTemplateDataContext{}
				for _, objSection := range layerObj.Sections {
					if objSection.Name == section {

						selectedContext = *models.NewCodeTemplateDataContext(*workspaceObj, *projectObj, *resourceObj, *templateObj, *layerObj, objSection)

						tempPackages := templateObj.Specifications.Package
						packageStr, err := TemplateEngine(strings.Join(tempPackages, "/"), selectedContext)
						if err != nil {
							return models.CodeTemplateDataContext{}
						}
						templateObj.Specifications.Package = utils.PathToArray(packageStr)

						selectedContext = *models.NewCodeTemplateDataContext(*workspaceObj, *projectObj, *resourceObj, *templateObj, *layerObj, objSection)
						break
					}
				}

				if reflect.DeepEqual(selectedContext, models.CodeTemplateDataContext{}) {
					return models.CodeTemplateDataContext{}
				}

				return selectedContext
			} else {

				selectedContext := *models.NewCodeTemplateDataContext(*workspaceObj, *projectObj, *resourceObj, *templateObj, *layerObj, objectresource.Section{})
				tempPackages := templateObj.Specifications.Package
				packageStr, err := TemplateEngine(strings.Join(tempPackages, "/"), selectedContext)
				if err != nil {
					return models.CodeTemplateDataContext{}
				}
				templateObj.Specifications.Package = utils.PathToArray(packageStr)

				selectedContext = *models.NewCodeTemplateDataContext(*workspaceObj, *projectObj, *resourceObj, *templateObj, *layerObj, objectresource.Section{})
				if reflect.DeepEqual(selectedContext, models.CodeTemplateDataContext{}) {
					return models.CodeTemplateDataContext{}
				}

				return selectedContext
			}

		}
	}

	return models.CodeTemplateDataContext{}
}
