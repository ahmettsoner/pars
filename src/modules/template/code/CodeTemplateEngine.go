package code

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/core/utilities/json"
	codetemplate "parsdevkit.net/structs/template/code-template"
	codetemplateStruct "parsdevkit.net/structs/template/code-template"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"

	"parsdevkit.net/core/utilities/encrypt"
	"parsdevkit.net/core/utils"

	"github.com/sirupsen/logrus"
	engineOperations "parsdevkit.net/engines"
)

type CodeTemplateEngine struct{}

func (s CodeTemplateEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*codetemplateStruct.TemplateBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s CodeTemplateEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	codetemplates := make([]codetemplateStruct.TemplateBaseStruct, 0, len(data))

	for _, item := range data {
		codetemplate, ok := item.(*codetemplateStruct.TemplateBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Process: expected codetemplateStruct.TemplateBaseStruct, got %T", item)
		}

		codetemplates = append(codetemplates, *codetemplate)
	}

	return s.createTemplates(codetemplates, true)
}
func (s CodeTemplateEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	codetemplates := make([]codetemplateStruct.TemplateBaseStruct, 0, len(data))

	for _, item := range data {
		codetemplate, ok := item.(*codetemplateStruct.TemplateBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Destroy: expected codetemplateStruct.TemplateBaseStruct, got %T", item)
		}

		codetemplates = append(codetemplates, *codetemplate)
	}

	return s.removeTemplates(codetemplates, true)
}

func (s CodeTemplateEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  "Template.Code",
		Order: 4000,
	}
}
func (s CodeTemplateEngine) createTemplates(templates []codetemplateStruct.TemplateBaseStruct, init bool) error {

	templatesReadyToCreate := make([]codetemplateStruct.TemplateBaseStruct, 0)
	templatesForUpdate := make([]codetemplateStruct.TemplateBaseStruct, 0)
	templateService := ioc.Get[contracts.TemplateServiceInterface[codetemplate.TemplateBaseStruct]]()

	for _, template := range templates {
		if err := template.Validate(); err != nil {
			jsonObject, _ := json.ToJson(template)
			return fmt.Errorf("template invalid data: '%s'\n%w", jsonObject, err)
		}
	}

	for _, template := range templates {
		ok, err := templateService.IsExists(template.Header.Name, template.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: Code template ('%s') kontrolünde hata oluştu\n%w", template.Header.Name, err)
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
func (s CodeTemplateEngine) removeTemplates(templates []codetemplateStruct.TemplateBaseStruct, permanent bool) error {

	templateService := ioc.Get[contracts.TemplateServiceInterface[codetemplate.TemplateBaseStruct]]()
	templatesReadyToDelete := make([]codetemplateStruct.TemplateBaseStruct, 0)
	for _, template := range templates {
		ok, err := templateService.IsExists(template.Header.Name, template.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: Code template ('%s') kontrolünde hata oluştu\n%w", template.Header.Name, err)
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
func (s CodeTemplateEngine) generate(model codetemplateStruct.TemplateBaseStruct) (*codetemplateStruct.TemplateBaseStruct, error) {

	templateService := ioc.Get[contracts.TemplateServiceInterface[codetemplate.TemplateBaseStruct]]()

	result, err := templateService.GetByName(model.Header.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	templateEngine := engineOperations.NewCodeTemplateOperations(utils.GetEnvironment())
	err = templateEngine.GenerateByTemplate(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
