package shared

import (
	"fmt"

	"parsdevkit.net/core/utilities/json"
	sharedtemplateStruct "parsdevkit.net/structs/template/shared-template"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/application"
	"parsdevkit.net/core/schemas"
	"parsdevkit.net/core/utils"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/core/utilities/encrypt"
)

type SharedTemplateEngine struct{}

func (s SharedTemplateEngine) Validate(data []schemas.Schema) bool {
	for _, item := range data {
		_, ok := item.(*sharedtemplateStruct.TemplateBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s SharedTemplateEngine) Process(ctx *application.ApplicationContext, data []schemas.Schema) error {
	sharedtemplates := make([]sharedtemplateStruct.TemplateBaseStruct, 0, len(data))

	for _, item := range data {
		sharedtemplate, ok := item.(*sharedtemplateStruct.TemplateBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Process: expected sharedtemplateStruct.TemplateBaseStruct, got %T", item)
		}

		sharedtemplates = append(sharedtemplates, *sharedtemplate)
	}

	return s.createTemplates(sharedtemplates, true)
}
func (s SharedTemplateEngine) Destroy(ctx *application.ApplicationContext, data []schemas.Schema) error {
	sharedtemplates := make([]sharedtemplateStruct.TemplateBaseStruct, 0, len(data))

	for _, item := range data {
		sharedtemplate, ok := item.(*sharedtemplateStruct.TemplateBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Destroy: expected sharedtemplateStruct.TemplateBaseStruct, got %T", item)
		}

		sharedtemplates = append(sharedtemplates, *sharedtemplate)
	}

	return s.removeTemplates(sharedtemplates, true)
}

func (s SharedTemplateEngine) createTemplates(templates []sharedtemplateStruct.TemplateBaseStruct, init bool) error {

	templatesReadyToCreate := make([]sharedtemplateStruct.TemplateBaseStruct, 0)
	templatesForUpdate := make([]sharedtemplateStruct.TemplateBaseStruct, 0)
	templateService := services.NewSharedTemplateService(utils.GetEnvironment())

	for _, template := range templates {
		if err := template.Validate(); err != nil {
			jsonObject, _ := json.ToJson(template)
			return fmt.Errorf("template invalid data: '%s'\n%w", jsonObject, err)
		}
	}

	for _, template := range templates {
		ok, err := templateService.IsExists(template.Header.Name, template.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: Shared template ('%s') kontrolünde hata oluştu\n%w", template.Header.Name, err)
		}

		if ok {
			newModelHash, err := encrypt.CalculateHashFromObject(template)
			if err != nil {
				return err
			}
			structHash, err := templateService.GetHash(template.Header.Name)
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

		fmt.Printf("Creating %v Template\n", template.Header.Name)
		if _, err := templateService.Save(template); err != nil {
			return err
		}
		fmt.Printf("%v Template created\n", template.Header.Name)
	}

	logrus.Debugf("updating %v templates ", len(templatesForUpdate))
	for _, template := range templatesForUpdate {

		if _, err := templateService.Save(template); err != nil {
			return err
		}

		fmt.Printf("%v Template updated\n", template.Header.Name)
	}
	return nil
}

func (s SharedTemplateEngine) removeTemplates(templates []sharedtemplateStruct.TemplateBaseStruct, permanent bool) error {

	templateService := services.NewSharedTemplateService(utils.GetEnvironment())
	templatesReadyToDelete := make([]sharedtemplateStruct.TemplateBaseStruct, 0)
	for _, template := range templates {
		ok, err := templateService.IsExists(template.Header.Name, template.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: Shared template ('%s') kontrolünde hata oluştu\n%w", template.Header.Name, err)
		}
		if ok {
			templatesReadyToDelete = append(templatesReadyToDelete, template)
		}
	}

	for _, template := range templatesReadyToDelete {

		if _, err := templateService.Remove(template.Header.Name, template.Specifications.Workspace, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Template deleted\n", template.Header.Name)

	}

	return nil
}
