package shared_task

import (
	"fmt"

	shared_template_payload_structs "parsdevkit.net/modules/template/shared_template_payload/structs"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	create_steps "parsdevkit.net/modules/template/shared_template/flows/create"
	remove_steps "parsdevkit.net/modules/template/shared_template/flows/remove"
	update_steps "parsdevkit.net/modules/template/shared_template/flows/update"
	describe_printer "parsdevkit.net/modules/template/shared_template/printers/describe"
	list_printer "parsdevkit.net/modules/template/shared_template/printers/list"
	"parsdevkit.net/pkg/utilities/encrypt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/modules/template/shared_template_contract"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"

	_string "parsdevkit.net/pkg/utilities/string"
)

type SharedTemplateEngine struct{}

func (s SharedTemplateEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*shared_template_payload_structs.TemplateBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}

func (s SharedTemplateEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
func (s SharedTemplateEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
func (s SharedTemplateEngine) prepareToCreate(ctx *application.ApplicationContext, templates []shared_template_payload_structs.TemplateBaseStruct) ([]shared_template_payload_structs.TemplateBaseStruct, error) {

	service := ioc.Get[shared_template_contract.TemplateInterface]()
	readyToCreateStructs := make([]shared_template_payload_structs.TemplateBaseStruct, 0)

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
func (s SharedTemplateEngine) create(ctx *application.ApplicationContext, templates []shared_template_payload_structs.TemplateBaseStruct, init bool) error {

	readyToCreateStructs, err := s.prepareToCreate(ctx, templates)
	if err != nil {
		return err
	}

	for _, template := range readyToCreateStructs {

		fmt.Printf("\n🛠️  Creating: %s.%s\n\n", template.Header.Name, template.GetKey())

		templateFlow := flowx.NewFlow("CreateNewTemplate").
			Step(&create_steps.SaveTemplate{})

		fc := flowx.NewContextWithData(map[string]any{
			"init":     init,
			"template": template,
		})

		if err := templateFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Template Create işleminde hata oluştu: %w", &err)
		}

	}

	return nil
}

func (s SharedTemplateEngine) prepareToUpdate(ctx *application.ApplicationContext, templates []shared_template_payload_structs.TemplateBaseStruct) ([]shared_template_payload_structs.TemplateBaseStruct, error) {

	service := ioc.Get[shared_template_contract.TemplateInterface]()
	readyToUpdateStructs := make([]shared_template_payload_structs.TemplateBaseStruct, 0)

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
func (s SharedTemplateEngine) update(ctx *application.ApplicationContext, templates []shared_template_payload_structs.TemplateBaseStruct, init bool) error {

	readyToUpdateStructs, err := s.prepareToUpdate(ctx, templates)
	if err != nil {
		return err
	}
	for _, template := range readyToUpdateStructs {

		fmt.Printf("\n🛠️  Updating: %s.%s\n\n", template.Header.Name, template.GetKey())

		templateFlow := flowx.NewFlow("UpdateExistingTemplate").
			Step(&update_steps.UpdateTemplate{})

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
func (s SharedTemplateEngine) prepareToRemove(ctx *application.ApplicationContext, templates []shared_template_payload_structs.TemplateBaseStruct) ([]shared_template_payload_structs.TemplateBaseStruct, error) {

	service := ioc.Get[shared_template_contract.TemplateInterface]()
	readyToRemoveStructs := make([]shared_template_payload_structs.TemplateBaseStruct, 0)

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

	return readyToRemoveStructs, nil
}
func (s SharedTemplateEngine) remove(ctx *application.ApplicationContext, templates []shared_template_payload_structs.TemplateBaseStruct, permanent bool) error {

	readyToRemoveStructs, err := s.prepareToRemove(ctx, templates)
	if err != nil {
		return err
	}

	for _, template := range readyToRemoveStructs {

		fmt.Printf("\n🛠️  Removing: %s.%s\n\n", template.Header.Name, template.GetKey())

		templateFlow := flowx.NewFlow("RemoveExistingTemplate").
			Step(&remove_steps.DeleteTemplate{})

		fc := flowx.NewContextWithData(map[string]any{
			"permanent": permanent,
			"template":  template,
		})

		if err := templateFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Template Remove işleminde hata oluştu: %w", err)
		}

	}

	return nil
}

func (s SharedTemplateEngine) completeInformation(ctx *application.ApplicationContext, model *shared_template_payload_structs.TemplateBaseStruct) error {

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

	return nil
}

func (s SharedTemplateEngine) getWorkspace(ctx *application.ApplicationContext, model shared_template_payload_structs.TemplateBaseStruct) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

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

func (s SharedTemplateEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  shared_template_payload_structs.MODULE_KEY,
		Order: 4000,
	}
}
func CastArrayToConcrate(data []schemas.SchemaInterface) ([]shared_template_payload_structs.TemplateBaseStruct, error) {
	r := make([]shared_template_payload_structs.TemplateBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*shared_template_payload_structs.TemplateBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected shared_template_payload_structs.TemplateBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}

func (s SharedTemplateEngine) List(ctx *application.ApplicationContext) error {
	err := s.list(ctx)
	if err != nil {
		return err
	}

	return nil
}
func (s SharedTemplateEngine) prepareToList(ctx *application.ApplicationContext) ([]list_printer.ViewModel, error) {

	service := ioc.Get[shared_template_contract.TemplateInterface]()

	var readyToListStructs []list_printer.ViewModel = make([]list_printer.ViewModel, 0)
	groupList, err := service.List()
	if err != nil {
		return nil, err
	}

	for _, e := range *groupList {
		resource := list_printer.ViewModel{
			Name: e.Header.Name,
			Tags: e.Header.Metadata.Tags,
		}

		readyToListStructs = append(readyToListStructs, resource)
	}

	return readyToListStructs, nil
}
func (s SharedTemplateEngine) list(ctx *application.ApplicationContext) error {

	readyToListStructs, err := s.prepareToList(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("\n🛠️  Shared Template List (%d):\n\n", len(readyToListStructs))

	printer := list_printer.ListTemplate{Templates: readyToListStructs}
	printer.Print()

	return nil
}

func (s SharedTemplateEngine) Describe(ctx *application.ApplicationContext, args ...any) error {
	err := s.describe(ctx, args...)
	if err != nil {
		return err
	}

	return nil
}
func (s SharedTemplateEngine) prepareToDescribe(ctx *application.ApplicationContext, args ...any) ([]describe_printer.ViewModel, error) {

	service := ioc.Get[shared_template_contract.TemplateInterface]()

	var readyToDescribeStructs []describe_printer.ViewModel = make([]describe_printer.ViewModel, 0)
	for _, v := range args {
		if a, ok := v.(string); ok {
			group, err := service.GetByName(a)
			if err != nil {
				return nil, err
			}
			if group != nil {

				resource := describe_printer.ViewModel{
					Name: group.Header.Name,
					Tags: group.Header.Metadata.Tags,
				}

				readyToDescribeStructs = append(readyToDescribeStructs, resource)
			} else {
				return nil, fmt.Errorf("xxx: Shared Template '%s' bulunamadı", a)
			}
		} else {
			return nil, fmt.Errorf("xxx: Shared Template argümanı doğru değil")
		}
	}

	return readyToDescribeStructs, nil
}
func (s SharedTemplateEngine) describe(ctx *application.ApplicationContext, args ...any) error {

	readyToDescribeStructs, err := s.prepareToDescribe(ctx, args...)
	if err != nil {
		return err
	}

	for _, group := range readyToDescribeStructs {

		fmt.Printf("\n🛠️  Details for: %s\n\n", group.Name)

		printer := describe_printer.DescribeTemplate{Template: group}

		if err := printer.Print(); err != nil {
			return fmt.Errorf("xxx: Shared Template Print sırasında hata oluştu: %w", &err)
		}
	}

	return nil
}
