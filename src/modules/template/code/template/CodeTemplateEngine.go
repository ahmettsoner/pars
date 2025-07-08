package code_template

import (
	"fmt"

	codetemplate "parsdevkit.net/structs/template/code-template"
	codetemplateStruct "parsdevkit.net/structs/template/code-template"

	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"

	"parsdevkit.net/application"

	"github.com/sirupsen/logrus"
	engineOperations "parsdevkit.net/engines"
	"parsdevkit.net/modules/template/code_template_contract"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	workspaceStruct "parsdevkit.net/modules/workspace/basic_workspace_payload"
	"parsdevkit.net/pkg/utilities/encrypt"
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

func (s CodeTemplateEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  "Template.Code",
		Order: 4000,
	}
}

func (s CodeTemplateEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	dataStruct, err := CastArrayToConcrate(data)
	if err != nil {
		return err
	}

	err = s.create(ctx, dataStruct, true)
	if err != nil {
		return err
	}
	err = s.update(ctx, dataStruct, true)
	if err != nil {
		return err
	}

	return nil
}
func (s CodeTemplateEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	dataStruct, err := CastArrayToConcrate(data)
	if err != nil {
		return err
	}

	err = s.remove(ctx, dataStruct, true)
	if err != nil {
		return err
	}

	return nil
}
func (s CodeTemplateEngine) prepareToCreate(ctx *application.ApplicationContext, templates []codetemplateStruct.TemplateBaseStruct) ([]codetemplateStruct.TemplateBaseStruct, error) {

	service := ioc.Get[code_template_contract.TemplateInterface]()
	readyToCreateStructs := make([]codetemplateStruct.TemplateBaseStruct, 0)

	for _, template := range templates {
		if err := s.completeInformation(ctx, &template); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(template.Header.Name, template.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if !ok {
			readyToCreateStructs = append(readyToCreateStructs, template)
		}
	}
	logrus.Debugf("'%d' template(s) detected that will create", len(readyToCreateStructs))

	return readyToCreateStructs, nil
}
func (s CodeTemplateEngine) create(ctx *application.ApplicationContext, templates []codetemplateStruct.TemplateBaseStruct, init bool) error {

	service := ioc.Get[code_template_contract.TemplateInterface]()
	readyToCreateStructs, err := s.prepareToCreate(ctx, templates)
	if err != nil {
		return err
	}

	for index, template := range readyToCreateStructs {

		logrus.Debugf("trying to create %v", template.Header.Name)
		if _, err := service.Save(template); err != nil {
			return err
		}

		if _, err := s.generate(template); err != nil {
			return err
		}
		fmt.Printf("%v (%d) Group created\n", template.Header.Name, index)

	}

	return nil
}

func (s CodeTemplateEngine) prepareToUpdate(ctx *application.ApplicationContext, templates []codetemplateStruct.TemplateBaseStruct) ([]codetemplateStruct.TemplateBaseStruct, error) {

	service := ioc.Get[code_template_contract.TemplateInterface]()
	readyToUpdateStructs := make([]codetemplateStruct.TemplateBaseStruct, 0)

	for _, template := range templates {
		if err := s.completeInformation(ctx, &template); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(template.Header.Name, template.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if ok {
			newModelHash, err := encrypt.CalculateHashFromObject(template)
			if err != nil {
				return nil, err
			}
			structHash, err := service.GetHash(template.Header.Name)
			if err != nil {
				return nil, err
			}

			if newModelHash != structHash {
				readyToUpdateStructs = append(readyToUpdateStructs, template)
			}
		}
	}
	logrus.Debugf("'%d' template(s) detected that will update", len(readyToUpdateStructs))

	return readyToUpdateStructs, nil
}
func (s CodeTemplateEngine) update(ctx *application.ApplicationContext, templates []codetemplateStruct.TemplateBaseStruct, init bool) error {

	service := ioc.Get[code_template_contract.TemplateInterface]()

	readyToUpdateStructs, err := s.prepareToUpdate(ctx, templates)
	if err != nil {
		return err
	}
	for _, template := range readyToUpdateStructs {
		if _, err := service.Save(template); err != nil {
			return err
		}
		if _, err := s.generate(template); err != nil {
			return err
		}
	}
	return nil
}
func (s CodeTemplateEngine) prepareToRemove(ctx *application.ApplicationContext, templates []codetemplateStruct.TemplateBaseStruct) ([]codetemplateStruct.TemplateBaseStruct, error) {

	service := ioc.Get[code_template_contract.TemplateInterface]()
	readyToRemoveStructs := make([]codetemplateStruct.TemplateBaseStruct, 0)

	for _, template := range templates {
		if err := s.completeInformation(ctx, &template); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(template.Header.Name, template.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if ok {
			readyToRemoveStructs = append(readyToRemoveStructs, template)
		}
	}
	logrus.Debugf("'%d' template(s) detected that will remove", len(readyToRemoveStructs))

	return readyToRemoveStructs, nil
}
func (s CodeTemplateEngine) remove(ctx *application.ApplicationContext, templates []codetemplateStruct.TemplateBaseStruct, permanent bool) error {

	service := ioc.Get[code_template_contract.TemplateInterface]()

	readyToRemoveStructs, err := s.prepareToRemove(ctx, templates)
	if err != nil {
		return err
	}

	for _, template := range readyToRemoveStructs {

		if _, err := service.Remove(template.Header.Name, template.Specifications.Workspace, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Group deleted\n", template.Header.Name)

	}

	logrus.Debugf("'%d' template(s) deleting", len(readyToRemoveStructs))

	return nil
}
func (s CodeTemplateEngine) generate(model codetemplateStruct.TemplateBaseStruct) (*codetemplateStruct.TemplateBaseStruct, error) {

	templateService := ioc.Get[code_template_contract.TemplateInterface]()

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

func CastArrayToConcrate(data []schemas.SchemaInterface) ([]codetemplateStruct.TemplateBaseStruct, error) {
	r := make([]codetemplateStruct.TemplateBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*codetemplateStruct.TemplateBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected codetemplateStruct.TemplateBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}
