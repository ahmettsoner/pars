package file

import (
	"fmt"

	"parsdevkit.net/core/utils/json"
	filetemplateStruct "parsdevkit.net/structs/template/file-template"

	engines "parsdevkit.net/engines/v2/engines"
	"parsdevkit.net/operation/services"

	"parsdevkit.net/core"
	"parsdevkit.net/core/utils"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/core/schemas"
)

type FileTemplateEngine struct{}

func (s FileTemplateEngine) Validate(data []schemas.Schema) bool {
	for _, item := range data {
		_, ok := item.(*filetemplateStruct.TemplateBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s FileTemplateEngine) Process(ctx *core.Context, data []schemas.Schema) error {
	filetemplates := make([]filetemplateStruct.TemplateBaseStruct, 0, len(data))

	for _, item := range data {
		filetemplate, ok := item.(*filetemplateStruct.TemplateBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Process: expected filetemplateStruct.TemplateBaseStruct, got %T", item)
		}

		filetemplates = append(filetemplates, *filetemplate)
	}

	return s.createTemplates(filetemplates, true)
}
func (s FileTemplateEngine) Destroy(ctx *core.Context, data []schemas.Schema) error {
	filetemplates := make([]filetemplateStruct.TemplateBaseStruct, 0, len(data))

	for _, item := range data {
		filetemplate, ok := item.(*filetemplateStruct.TemplateBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Destroy: expected filetemplateStruct.TemplateBaseStruct, got %T", item)
		}

		filetemplates = append(filetemplates, *filetemplate)
	}

	return s.removeTemplates(filetemplates, true)
}

func (s FileTemplateEngine) createTemplates(templates []filetemplateStruct.TemplateBaseStruct, init bool) error {

	templatesReadyToCreate := make([]filetemplateStruct.TemplateBaseStruct, 0)
	templatesForUpdate := make([]filetemplateStruct.TemplateBaseStruct, 0)
	templateService := services.NewFileTemplateService(utils.GetEnvironment())

	for _, template := range templates {
		if err := template.Validate(); err != nil {
			jsonObject, _ := json.ToJson(template)
			return fmt.Errorf("template invalid data: '%s'\n%w", jsonObject, err)
		}
	}

	for _, template := range templates {
		ok, err := templateService.IsExists(template.Name, template.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: File template ('%s') kontrolünde hata oluştu\n%w", template.Name, err)
		}
		if ok {
			newModelHash, err := utils.CalculateHashFromObject(template)
			if err != nil {
				return err
			}
			structHash, err := templateService.GetHash(template.Name)
			if err != nil {
				return err
			}

			if newModelHash != structHash {
				templatesForUpdate = append(templatesForUpdate, template)
			}
		} else {
			templatesReadyToCreate = append(templatesReadyToCreate, template)
		}
	}
	logrus.Debugf("'%d' template(s) detected that will create", len(templatesReadyToCreate))
	logrus.Debugf("'%d' template(s) detected that will update", len(templatesForUpdate))

	logrus.Debugf("creating %v new templates ", len(templatesReadyToCreate))
	logrus.Debugf("updating %v templates ", len(templatesForUpdate))
	for _, template := range templatesReadyToCreate {

		fmt.Printf("Creating %v Template\n", template.Name)
		if _, err := templateService.Save(template); err != nil {
			return err
		}

		if _, err := s.generate(template); err != nil {
			return err
		}

		fmt.Printf("%v Template created\n", template.Name)
	}

	logrus.Debugf("updating %v templates ", len(templatesForUpdate))
	for _, template := range templatesForUpdate {

		if _, err := templateService.Save(template); err != nil {
			return err
		}

		if _, err := s.generate(template); err != nil {
			return err
		}

		fmt.Printf("%v Template updated\n", template.Name)
	}
	return nil
}

func (s FileTemplateEngine) removeTemplates(templates []filetemplateStruct.TemplateBaseStruct, permanent bool) error {

	templateService := services.NewFileTemplateService(utils.GetEnvironment())
	templatesReadyToDelete := make([]filetemplateStruct.TemplateBaseStruct, 0)
	for _, template := range templates {
		ok, err := templateService.IsExists(template.Name, template.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: File template ('%s') kontrolünde hata oluştu\n%w", template.Name, err)
		}
		if ok {
			templatesReadyToDelete = append(templatesReadyToDelete, template)
		}
	}

	for _, template := range templatesReadyToDelete {

		if _, err := templateService.Remove(template.Name, template.Specifications.Workspace, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Template deleted\n", template.Name)

	}

	return nil
}
func (s FileTemplateEngine) generate(model filetemplateStruct.TemplateBaseStruct) (*filetemplateStruct.TemplateBaseStruct, error) {

	templateService := services.NewFileTemplateService(utils.GetEnvironment())

	result, err := templateService.GetByName(model.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	templateOperations := engines.NewFileTemplateOperations(utils.GetEnvironment())
	err = templateOperations.GenerateByTemplate(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
