package file_template

import (
	"fmt"

	"parsdevkit.net/application/bus"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"

	layerPkg "parsdevkit.net/application/models/layer"
	"parsdevkit.net/internal/flowx"
	create_steps "parsdevkit.net/modules/template/file_template/flows/create"
	remove_steps "parsdevkit.net/modules/template/file_template/flows/remove"
	update_steps "parsdevkit.net/modules/template/file_template/flows/update"
	describe_printer "parsdevkit.net/modules/template/file_template/printers/describe"
	list_printer "parsdevkit.net/modules/template/file_template/printers/list"
	"parsdevkit.net/modules/template/file_template_contract"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	file_template_payload_events "parsdevkit.net/modules/template/file_template_payload/events"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"

	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/application/schemas"
	"parsdevkit.net/pkg/utilities/encrypt"

	"github.com/sirupsen/logrus"
)

type FileTemplateEngine struct{}

func (s FileTemplateEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*file_template_payload_structs.TemplateBaseStruct)
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

	readyToCreateStructs, err := s.prepareToCreate(ctx, dataStruct)
	if err != nil {
		return err
	}

	err = s.create(ctx, readyToCreateStructs, true)
	if err != nil {
		return err
	}

	readyToUpdateStructs, err := s.prepareToUpdate(ctx, dataStruct)
	if err != nil {
		return err
	}
	err = s.update(ctx, readyToUpdateStructs, true)
	if err != nil {
		return err
	}

	return nil
}
func (s FileTemplateEngine) prepareToCreate(ctx *application.ApplicationContext, templates []file_template_payload_structs.TemplateBaseStruct) ([]file_template_payload_structs.TemplateBaseStruct, error) {

	service := ioc.Get[file_template_contract.TemplateInterface]()
	readyToCreateStructs := make([]file_template_payload_structs.TemplateBaseStruct, 0)

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
func (s FileTemplateEngine) create(ctx *application.ApplicationContext, models []file_template_payload_structs.TemplateBaseStruct, init bool) error {

	for _, template := range models {

		fmt.Printf("\n🛠️  Creating: %s.%s\n\n", template.Header.Name, template.GetKey())

		templateFlow := flowx.NewFlow("CreateNewTemplate").
			Step(&create_steps.SaveTemplate{}).
			Step(&create_steps.GenerateTemplateContent{})

		fc := flowx.NewContextWithData(map[string]any{
			"init":     init,
			"template": template,
		})

		if err := templateFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Template Create işleminde hata oluştu: %w", &err)
		}

		bus.PublishEvent(file_template_payload_events.TemplateCreated{Data: template})
	}

	return nil
}

func (s FileTemplateEngine) prepareToUpdate(ctx *application.ApplicationContext, templates []file_template_payload_structs.TemplateBaseStruct) ([]file_template_payload_structs.TemplateBaseStruct, error) {

	service := ioc.Get[file_template_contract.TemplateInterface]()
	readyToUpdateStructs := make([]file_template_payload_structs.TemplateBaseStruct, 0)

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
func (s FileTemplateEngine) update(ctx *application.ApplicationContext, models []file_template_payload_structs.TemplateBaseStruct, init bool) error {

	for _, template := range models {

		fmt.Printf("\n🛠️  Updating: %s.%s\n\n", template.Header.Name, template.GetKey())

		templateFlow := flowx.NewFlow("UpdateExistingTemplate").
			Step(&update_steps.UpdateTemplate{}).
			Step(&update_steps.GenerateTemplateContent{})

		fc := flowx.NewContextWithData(map[string]any{
			"init":     init,
			"template": template,
		})

		if err := templateFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Template Update işleminde hata oluştu: %w", &err)
		}
	}
	return nil
}

func (s FileTemplateEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	dataStruct, err := CastArrayToConcrate(data)
	if err != nil {
		return err
	}

	readyToDestroyStructs, err := s.prepareToDestroy(ctx, dataStruct)
	if err != nil {
		return err
	}

	err = s.remove(ctx, readyToDestroyStructs, true)
	if err != nil {
		return err
	}

	return nil
}
func (s FileTemplateEngine) prepareToDestroy(ctx *application.ApplicationContext, templates []file_template_payload_structs.TemplateBaseStruct) ([]file_template_payload_structs.TemplateBaseStruct, error) {

	service := ioc.Get[file_template_contract.TemplateInterface]()
	readyToDestroyStructs := make([]file_template_payload_structs.TemplateBaseStruct, 0)

	for _, template := range templates {
		if err := s.completeInformation(ctx, &template); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(template.Header.Name, template.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if ok {
			readyToDestroyStructs = append(readyToDestroyStructs, template)
		}
	}
	logrus.Debugf("'%d' template(s) detected that will destroy", len(readyToDestroyStructs))

	return readyToDestroyStructs, nil
}
func (s FileTemplateEngine) remove(ctx *application.ApplicationContext, models []file_template_payload_structs.TemplateBaseStruct, permanent bool) error {

	for _, template := range models {

		fmt.Printf("\n🛠️  Removing: %s.%s\n\n", template.Header.Name, template.GetKey())

		templateFlow := flowx.NewFlow("DestroyExistingTemplate").
			Step(&remove_steps.DeleteTemplate{}).
			Step(&remove_steps.ClearTemplateHistory{})

		fc := flowx.NewContextWithData(map[string]any{
			"permanent": permanent,
			"template":  template,
		})

		if err := templateFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Template Destroy işleminde hata oluştu: %w", err)
		}

	}
	return nil
}

func (s FileTemplateEngine) List(ctx *application.ApplicationContext) error {
	readyToListStructs, err := s.prepareToList(ctx)
	if err != nil {
		return err
	}

	err = s.list(ctx, readyToListStructs)
	if err != nil {
		return err
	}

	return nil
}
func (s FileTemplateEngine) prepareToList(ctx *application.ApplicationContext) ([]file_template_payload_structs.TemplateBaseStruct, error) {

	service := ioc.Get[file_template_contract.TemplateInterface]()

	groupList, err := service.List()
	if err != nil {
		return nil, err
	}

	return *groupList, nil
}
func (s FileTemplateEngine) list(ctx *application.ApplicationContext, models []file_template_payload_structs.TemplateBaseStruct) error {

	var viewModels []list_printer.ViewModel = make([]list_printer.ViewModel, 0)
	for _, e := range models {
		resource := list_printer.ViewModel{
			Name: e.Header.Name,
			Tags: e.Header.Metadata.Tags,
		}

		viewModels = append(viewModels, resource)
	}

	fmt.Printf("\n🛠️  File Template List (%d):\n\n", len(viewModels))

	printer := list_printer.ListTemplate{Templates: viewModels}
	printer.Print()

	return nil
}

func (s FileTemplateEngine) Describe(ctx *application.ApplicationContext, args ...any) error {
	readyToDescribeStructs, err := s.prepareToDescribe(ctx, args...)
	if err != nil {
		return err
	}

	err = s.describe(ctx, readyToDescribeStructs)
	if err != nil {
		return err
	}

	return nil
}
func (s FileTemplateEngine) prepareToDescribe(ctx *application.ApplicationContext, args ...any) ([]file_template_payload_structs.TemplateBaseStruct, error) {

	service := ioc.Get[file_template_contract.TemplateInterface]()

	var readyToDescribeStructs []file_template_payload_structs.TemplateBaseStruct = make([]file_template_payload_structs.TemplateBaseStruct, 0)
	for _, v := range args {
		if a, ok := v.(string); ok {
			template, err := service.GetByName(a)
			if err != nil {
				return nil, err
			}
			if template != nil && template.Header.Kind == file_template_payload_structs.TEMPLATE_KIND {

				readyToDescribeStructs = append(readyToDescribeStructs, *template)
			} else {
				return nil, fmt.Errorf("xxx: File Template '%s' bulunamadı", a)
			}
		} else {
			return nil, fmt.Errorf("xxx: File Template argümanı doğru değil")
		}
	}

	return readyToDescribeStructs, nil
}
func (s FileTemplateEngine) describe(ctx *application.ApplicationContext, models []file_template_payload_structs.TemplateBaseStruct) error {

	for _, template := range models {

		resource := describe_printer.ViewModel{
			Name: template.Header.Name,
			Tags: template.Header.Metadata.Tags,
		}

		fmt.Printf("\n🛠️  Details for: %s\n\n", resource.Name)

		printer := describe_printer.DescribeTemplate{Template: resource}

		if err := printer.Print(); err != nil {
			return fmt.Errorf("xxx: File Template Print sırasında hata oluştu: %w", &err)
		}
	}

	return nil
}
func (s FileTemplateEngine) Remove(ctx *application.ApplicationContext, args ...any) error {

	readyToRemoveStructs, err := s.prepareToRemove(ctx, args...)
	if err != nil {
		return err
	}

	err = s.remove(ctx, readyToRemoveStructs, true)
	if err != nil {
		return err
	}

	return nil
}
func (s FileTemplateEngine) prepareToRemove(ctx *application.ApplicationContext, args ...any) ([]file_template_payload_structs.TemplateBaseStruct, error) {

	service := ioc.Get[file_template_contract.TemplateInterface]()
	readyToRemoveStructs := make([]file_template_payload_structs.TemplateBaseStruct, 0)

	for _, a := range args {
		template, err := service.GetByName(a.(string))
		if err != nil {
			return nil, err
		}
		if template != nil && template.Header.Kind == file_template_payload_structs.TEMPLATE_KIND {
			readyToRemoveStructs = append(readyToRemoveStructs, *template)
		}
	}
	logrus.Debugf("'%d' template(s) detected that will remove", len(readyToRemoveStructs))

	return readyToRemoveStructs, nil
}

func (s FileTemplateEngine) completeInformation(ctx *application.ApplicationContext, model *file_template_payload_structs.TemplateBaseStruct) error {

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
		model.Specifications.Layers = append(model.Specifications.Layers, layerPkg.Layer{})
	}
	return nil
}

func (s FileTemplateEngine) getWorkspace(ctx *application.ApplicationContext, model file_template_payload_structs.TemplateBaseStruct) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	workspaceName := model.Specifications.Workspace
	if _string.IsEmpty(workspaceName) {
		workspaceName = ctx.CurrentWorkspace.Name
	}

	var result *basic_workspace_payload_structs.WorkspaceBaseStruct = nil

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
		Name:  file_template_payload_structs.MODULE_KEY,
		Order: 4000,
	}
}
func CastArrayToConcrate(data []schemas.SchemaInterface) ([]file_template_payload_structs.TemplateBaseStruct, error) {
	r := make([]file_template_payload_structs.TemplateBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*file_template_payload_structs.TemplateBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected file_template_payload_structs.TemplateBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}
