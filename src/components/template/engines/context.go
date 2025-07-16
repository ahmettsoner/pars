package engines

import (
	"fmt"
	"reflect"
	"strings"

	sectionPkg "parsdevkit.net/application/models/section"
	templatePkg "parsdevkit.net/components/template"
	"parsdevkit.net/modules/resource/object_resource_contract"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/context/models"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"
	"parsdevkit.net/modules/template/code_template_contract"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	"parsdevkit.net/pkg/utilities/file"
	_string "parsdevkit.net/pkg/utilities/string"
)

type ContextFuncs struct{}

func (c ContextFuncs) GetContextByBaseForArray(base models.CodeTemplateDataContext, args []string) models.CodeTemplateDataContext {
	return c.GetContextByBase(base, args...)
}
func (c ContextFuncs) GetContextByBase(base models.CodeTemplateDataContext, args ...string) models.CodeTemplateDataContext {
	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	applicationProjectService := ioc.Get[application_project_contract.ProjectInterface]()
	objectResourceService := ioc.Get[object_resource_contract.ResourceInterface]()
	codeTemplateService := ioc.Get[code_template_contract.TemplateInterface]()

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

	workspace := base.Workspace.Original.Header.Name
	set := base.Project.Original.Specifications.Set
	project := ""
	resource := base.Resource.Original.Header.Name
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

	var layerObj *object_resource_payload_structs.Layer = nil
	if resourceObj != nil {
		for _, resourceLayer := range resourceObj.Specifications.Layers {
			if resourceLayer.Name == layer {
				layerObj = &resourceLayer
				break
			}
		}
	}

	var projectObj *application_project_payload_structs.ProjectBaseStruct = nil
	if layerObj != nil {
		var projectList []application_project_payload_structs.ProjectBaseStruct = make([]application_project_payload_structs.ProjectBaseStruct, 0)
		if _string.IsEmpty(project) {
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
				for _, objLayer := range projectObjFromDb.Specifications.Layers {
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

		var templateObj *code_template_payload_structs.TemplateBaseStruct = nil

		var templatelist []code_template_payload_structs.TemplateBaseStruct = make([]code_template_payload_structs.TemplateBaseStruct, 0)
		if _string.IsEmpty(template) {
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
			if !_string.IsEmpty(section) {
				selectedContext := models.CodeTemplateDataContext{}
				for _, objSection := range layerObj.Sections {
					if objSection.Name == section {

						selectedContext = *models.NewCodeTemplateDataContext(
							templatePkg.NewContextProviderSource(
								*workspaceObj,
								nil,
								*projectObj,
								*resourceObj,
								*templateObj,
								layerObj.LayerIdentifier,
								objSection.SectionIdentifier,
							),
						)

						tempPackages := templateObj.Specifications.Package
						packageStr, err := RenderTemplate(strings.Join(tempPackages, "/"), selectedContext)
						if err != nil {
							return models.CodeTemplateDataContext{}
						}
						templateObj.Specifications.Package = file.PathToArray(packageStr)

						selectedContext = *models.NewCodeTemplateDataContext(
							templatePkg.NewContextProviderSource(
								*workspaceObj,
								nil,
								*projectObj,
								*resourceObj,
								*templateObj,
								layerObj.LayerIdentifier,
								objSection.SectionIdentifier,
							),
						)
						break
					}
				}

				if reflect.DeepEqual(selectedContext, models.CodeTemplateDataContext{}) {
					return models.CodeTemplateDataContext{}
				}

				return selectedContext
			} else {

				selectedContext := *models.NewCodeTemplateDataContext(
					templatePkg.NewContextProviderSource(
						*workspaceObj,
						nil,
						*projectObj,
						*resourceObj,
						*templateObj,
						layerObj.LayerIdentifier,
						sectionPkg.SectionIdentifier{},
					),
				)
				tempPackages := templateObj.Specifications.Package
				packageStr, err := RenderTemplate(strings.Join(tempPackages, "/"), selectedContext)
				if err != nil {
					return models.CodeTemplateDataContext{}
				}
				templateObj.Specifications.Package = file.PathToArray(packageStr)

				selectedContext = *models.NewCodeTemplateDataContext(
					templatePkg.NewContextProviderSource(
						*workspaceObj,
						nil,
						*projectObj,
						*resourceObj,
						*templateObj,
						layerObj.LayerIdentifier,
						sectionPkg.SectionIdentifier{},
					),
				)
				if reflect.DeepEqual(selectedContext, models.CodeTemplateDataContext{}) {
					return models.CodeTemplateDataContext{}
				}

				return selectedContext
			}

		}
	}

	return models.CodeTemplateDataContext{}
}
