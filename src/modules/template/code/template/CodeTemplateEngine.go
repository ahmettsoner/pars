package code_template

import (
	"fmt"

	"parsdevkit.net/pkg/utilities/json"
	codetemplate "parsdevkit.net/structs/template/code-template"
	codetemplateStruct "parsdevkit.net/structs/template/code-template"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"

	"parsdevkit.net/application"
	"parsdevkit.net/pkg/utilities/encrypt"

	"github.com/sirupsen/logrus"
	engineOperations "parsdevkit.net/engines"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	workspaceStruct "parsdevkit.net/modules/workspace/basic_workspace_payload"
	_string "parsdevkit.net/pkg/utilities/string"
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

		if err := s.completeInformation(ctx, codetemplate); err != nil {
			return err
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

		if err := s.completeInformation(ctx, codetemplate); err != nil {
			return err
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

	templateEngine := engineOperations.NewCodeTemplateOperations(application.GetEnvironment())
	err = templateEngine.GenerateByTemplate(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s CodeTemplateEngine) completeInformation(ctx *application.ApplicationContext, model *codetemplateStruct.TemplateBaseStruct) error {

	logrus.Debugf("filling model (%v) information", model.Header.Name)

	if _string.IsEmpty(model.Specifications.Name) {
		model.Specifications.Name = model.Header.Name
	}

	activeWorkspace, err := s.getWorkspace(ctx, *model)
	if err != nil {
		return err
	}

	//WARN: Doğru mu oldu?
	model.Specifications.Workspace = activeWorkspace.Header.Name
	model.Specifications.WorkspaceObject = activeWorkspace.Specifications.WorkspaceIdentifier
	logrus.Debugf("workspace (%v) detected for (%v)", activeWorkspace.Header.Name, model.Header.Name)

	if _string.IsEmpty(model.Specifications.Output.File) {
		model.Specifications.Output.File = model.Header.Name
	}

	if len(model.Specifications.Layers) == 0 {
		model.Specifications.Layers = append(model.Specifications.Layers, codetemplate.Layer{})
	}

	return nil
}

func (s CodeTemplateEngine) getWorkspace(ctx *application.ApplicationContext, model codetemplateStruct.TemplateBaseStruct) (*workspaceStruct.WorkspaceBaseStruct, error) {

	workspaceName := model.Specifications.Workspace
	if _string.IsEmpty(workspaceName) {
		workspaceName = ctx.CurrentWorkspace.Name
	}

	var result *workspaceStruct.WorkspaceBaseStruct = nil

	if !_string.IsEmpty(workspaceName) {
		workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
		workspace, err := workspaceService.GetByName(workspaceName)
		if err != nil {
			return nil, err
		}
		if workspace == nil {
			return nil, fmt.Errorf("workspace name (%v) is not correct", workspaceName)
		}
		result = workspace
	} else {
	}

	return result, nil
}
