package file_template

import (
	"fmt"

	filetemplate "parsdevkit.net/structs/template/file-template"
	filetemplateStruct "parsdevkit.net/structs/template/file-template"

	engineOperations "parsdevkit.net/engines"
	"parsdevkit.net/modules/template/file_template_contract"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	workspaceStruct "parsdevkit.net/modules/workspace/basic_workspace_payload"
	_string "parsdevkit.net/pkg/utilities/string"

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
func (s FileTemplateEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
func (s FileTemplateEngine) prepareToCreate(ctx *application.ApplicationContext, templates []filetemplateStruct.TemplateBaseStruct) ([]filetemplateStruct.TemplateBaseStruct, error) {

	service := ioc.Get[file_template_contract.TemplateInterface]()
	readyToCreateStructs := make([]filetemplateStruct.TemplateBaseStruct, 0)

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
func (s FileTemplateEngine) create(ctx *application.ApplicationContext, templates []filetemplateStruct.TemplateBaseStruct, init bool) error {

	service := ioc.Get[file_template_contract.TemplateInterface]()
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
		fmt.Printf("%v (%d) File Template created\n", template.Header.Name, index)

	}

	return nil
}

func (s FileTemplateEngine) prepareToUpdate(ctx *application.ApplicationContext, templates []filetemplateStruct.TemplateBaseStruct) ([]filetemplateStruct.TemplateBaseStruct, error) {

	service := ioc.Get[file_template_contract.TemplateInterface]()
	readyToUpdateStructs := make([]filetemplateStruct.TemplateBaseStruct, 0)

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
func (s FileTemplateEngine) update(ctx *application.ApplicationContext, templates []filetemplateStruct.TemplateBaseStruct, init bool) error {

	service := ioc.Get[file_template_contract.TemplateInterface]()

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
func (s FileTemplateEngine) prepareToRemove(ctx *application.ApplicationContext, templates []filetemplateStruct.TemplateBaseStruct) ([]filetemplateStruct.TemplateBaseStruct, error) {

	service := ioc.Get[file_template_contract.TemplateInterface]()
	readyToRemoveStructs := make([]filetemplateStruct.TemplateBaseStruct, 0)

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
func (s FileTemplateEngine) remove(ctx *application.ApplicationContext, templates []filetemplateStruct.TemplateBaseStruct, permanent bool) error {

	service := ioc.Get[file_template_contract.TemplateInterface]()

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

func (s FileTemplateEngine) generate(model filetemplateStruct.TemplateBaseStruct) (*filetemplateStruct.TemplateBaseStruct, error) {

	templateService := ioc.Get[file_template_contract.TemplateInterface]()

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

func (s FileTemplateEngine) completeInformation(ctx *application.ApplicationContext, model *filetemplateStruct.TemplateBaseStruct) error {

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
		model.Specifications.Layers = append(model.Specifications.Layers, filetemplate.Layer{})
	}
	return nil
}

func (s FileTemplateEngine) getWorkspace(ctx *application.ApplicationContext, model filetemplateStruct.TemplateBaseStruct) (*workspaceStruct.WorkspaceBaseStruct, error) {

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

func (s FileTemplateEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  "Template.File",
		Order: 4000,
	}
}
func CastArrayToConcrate(data []schemas.SchemaInterface) ([]filetemplateStruct.TemplateBaseStruct, error) {
	r := make([]filetemplateStruct.TemplateBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*filetemplateStruct.TemplateBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected filetemplateStruct.TemplateBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}
