package file

import (
	"fmt"

	"parsdevkit.net/pkg/utilities/json"
	filetemplate "parsdevkit.net/structs/template/file-template"
	filetemplateStruct "parsdevkit.net/structs/template/file-template"

	engineOperations "parsdevkit.net/engines"

	"parsdevkit.net/application"
	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/pkg/utilities/encrypt"

	"github.com/sirupsen/logrus"
)

type FileTemplateEngine struct{}

func (s FileTemplateEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*filetemplateStruct.TemplateBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s FileTemplateEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
func (s FileTemplateEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  "Template.File",
		Order: 4000,
	}
}

func (s FileTemplateEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
	templateService := ioc.Get[contracts.TemplateServiceInterface[filetemplate.TemplateBaseStruct]]()

	for _, template := range templates {
		if err := template.Validate(); err != nil {
			jsonObject, _ := json.ToJson(template)
			return fmt.Errorf("template invalid data: '%s'\n%w", jsonObject, err)
		}
	}

	for _, template := range templates {
		ok, err := templateService.IsExists(template.Header.Name, template.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: File template ('%s') kontrolünde hata oluştu\n%w", template.Header.Name, err)
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

		if _, err := s.generate(template); err != nil {
			return err
		}

		fmt.Printf("%v Template created\n", template.Header.Name)
	}

	logrus.Debugf("updating %v templates ", len(templatesForUpdate))
	for _, template := range templatesForUpdate {

		if _, err := templateService.Save(template); err != nil {
			return err
		}

		if _, err := s.generate(template); err != nil {
			return err
		}

		fmt.Printf("%v Template updated\n", template.Header.Name)
	}
	return nil
}

func (s FileTemplateEngine) removeTemplates(templates []filetemplateStruct.TemplateBaseStruct, permanent bool) error {

	templateService := ioc.Get[contracts.TemplateServiceInterface[filetemplate.TemplateBaseStruct]]()
	templatesReadyToDelete := make([]filetemplateStruct.TemplateBaseStruct, 0)
	for _, template := range templates {
		ok, err := templateService.IsExists(template.Header.Name, template.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: File template ('%s') kontrolünde hata oluştu\n%w", template.Header.Name, err)
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
func (s FileTemplateEngine) generate(model filetemplateStruct.TemplateBaseStruct) (*filetemplateStruct.TemplateBaseStruct, error) {

	templateService := ioc.Get[contracts.TemplateServiceInterface[filetemplate.TemplateBaseStruct]]()

	result, err := templateService.GetByName(model.Header.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	templateOperations := engineOperations.NewFileTemplateOperations(application.GetEnvironment())
	err = templateOperations.GenerateByTemplate(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
