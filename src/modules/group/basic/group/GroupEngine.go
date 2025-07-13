package basic_group

import (
	"fmt"

	"parsdevkit.net/modules/project/application_project_contract"

	"github.com/sirupsen/logrus"

	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"
	"parsdevkit.net/pkg/utilities/encrypt"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/internal/flowx"
	create_steps "parsdevkit.net/modules/group/basic_group/flows/create"
	remove_steps "parsdevkit.net/modules/group/basic_group/flows/remove"
	update_steps "parsdevkit.net/modules/group/basic_group/flows/update"
	describe_printer "parsdevkit.net/modules/group/basic_group/printers/describe"
	list_printer "parsdevkit.net/modules/group/basic_group/printers/list"
	"parsdevkit.net/modules/group/basic_group_contract"
	_string "parsdevkit.net/pkg/utilities/string"
)

type GroupEngine struct{}

func (s GroupEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*basic_group_payload_structs.GroupBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s GroupEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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

func (s GroupEngine) prepareToCreate(ctx *application.ApplicationContext, groups []basic_group_payload_structs.GroupBaseStruct) ([]basic_group_payload_structs.GroupBaseStruct, error) {

	service := ioc.Get[basic_group_contract.GroupInterface]()
	readyToCreateStructs := make([]basic_group_payload_structs.GroupBaseStruct, 0)

	for _, group := range groups {
		if err := s.completeInformation(ctx, &group); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(group.Header.Name)
		if err != nil {
			return nil, err
		}
		if !ok {
			readyToCreateStructs = append(readyToCreateStructs, group)
		}
	}
	logrus.Debugf("'%d' group(s) detected that will create", len(readyToCreateStructs))

	return readyToCreateStructs, nil
}
func (s GroupEngine) create(ctx *application.ApplicationContext, models []basic_group_payload_structs.GroupBaseStruct, init bool) error {

	for _, group := range models {

		fmt.Printf("\n🛠️  Creating: %s.%s\n\n", group.Header.Name, group.GetKey())

		groupFlow := flowx.NewFlow("CreateNewGroup").
			Step(&create_steps.SaveGroup{})

		fc := flowx.NewContextWithData(map[string]any{
			"init":  init,
			"group": group,
		})

		if err := groupFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Group Create işleminde hata oluştu: %w", &err)
		}
	}

	return nil
}

func (s GroupEngine) prepareToUpdate(ctx *application.ApplicationContext, groups []basic_group_payload_structs.GroupBaseStruct) ([]basic_group_payload_structs.GroupBaseStruct, error) {

	service := ioc.Get[basic_group_contract.GroupInterface]()
	readyToUpdateStructs := make([]basic_group_payload_structs.GroupBaseStruct, 0)

	for _, group := range groups {
		if err := s.completeInformation(ctx, &group); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(group.Header.Name)
		if err != nil {
			return nil, err
		}
		if ok {
			newModelHash, err := encrypt.CalculateHashFromObject(group)
			if err != nil {
				return nil, err
			}
			structHash, err := service.GetHash(group.Header.Name)
			if err != nil {
				return nil, err
			}

			if newModelHash != structHash {
				readyToUpdateStructs = append(readyToUpdateStructs, group)
			}
		}
	}
	logrus.Debugf("'%d' group(s) detected that will update", len(readyToUpdateStructs))

	return readyToUpdateStructs, nil
}
func (s GroupEngine) update(ctx *application.ApplicationContext, models []basic_group_payload_structs.GroupBaseStruct, init bool) error {

	for _, group := range models {

		fmt.Printf("�️ Updating: %s.%s\n\n", group.Header.Name, group.GetKey())

		groupFlow := flowx.NewFlow("UpdateExistingGroup").
			Step(&update_steps.UpdateGroup{})

		fc := flowx.NewContextWithData(map[string]any{
			"init":  init,
			"group": group,
		})

		if err := groupFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Group Update işleminde hata oluştu: %w", &err)
		}
	}
	return nil
}

func (s GroupEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	dataStruct, err := CastArrayToConcrate(data)
	if err != nil {
		return err
	}

	readyToRemoveStructs, err := s.prepareToDestroy(ctx, dataStruct)
	if err != nil {
		return err
	}

	err = s.remove(ctx, readyToRemoveStructs, true)
	if err != nil {
		return err
	}

	return nil
}
func (s GroupEngine) prepareToDestroy(ctx *application.ApplicationContext, groups []basic_group_payload_structs.GroupBaseStruct) ([]basic_group_payload_structs.GroupBaseStruct, error) {

	service := ioc.Get[basic_group_contract.GroupInterface]()
	readyToDestroyStructs := make([]basic_group_payload_structs.GroupBaseStruct, 0)

	for _, group := range groups {
		if err := s.completeInformation(ctx, &group); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(group.Header.Name)
		if err != nil {
			return nil, err
		}
		if ok {
			readyToDestroyStructs = append(readyToDestroyStructs, group)
		}
	}
	logrus.Debugf("'%d' group(s) detected that will destroy", len(readyToDestroyStructs))

	return readyToDestroyStructs, nil
}

func (s GroupEngine) List(ctx *application.ApplicationContext) error {

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
func (s GroupEngine) prepareToList(ctx *application.ApplicationContext) ([]basic_group_payload_structs.GroupBaseStruct, error) {

	service := ioc.Get[basic_group_contract.GroupInterface]()

	groupList, err := service.List()
	if err != nil {
		return nil, err
	}

	return *groupList, nil
}
func (s GroupEngine) list(ctx *application.ApplicationContext, models []basic_group_payload_structs.GroupBaseStruct) error {

	var viewModels []list_printer.ViewModel = make([]list_printer.ViewModel, 0)
	for _, e := range models {
		resource := list_printer.ViewModel{
			Name:    e.Header.Name,
			Tags:    e.Header.Metadata.Tags,
			Path:    e.Specifications.Path,
			Package: e.Specifications.Package,
		}

		viewModels = append(viewModels, resource)
	}

	fmt.Printf("\n🛠️  Group List (%d):\n\n", len(viewModels))

	printer := list_printer.ListGroup{Groups: viewModels}
	printer.Print()

	return nil
}

func (s GroupEngine) Describe(ctx *application.ApplicationContext, args ...any) error {
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
func (s GroupEngine) prepareToDescribe(ctx *application.ApplicationContext, args ...any) ([]basic_group_payload_structs.GroupBaseStruct, error) {

	service := ioc.Get[basic_group_contract.GroupInterface]()

	var groupList []basic_group_payload_structs.GroupBaseStruct = make([]basic_group_payload_structs.GroupBaseStruct, 0)
	for _, v := range args {
		if a, ok := v.(string); ok {
			group, err := service.GetByName(a)
			if err != nil {
				return nil, err
			}
			if group != nil {
				groupList = append(groupList, *group)

			} else {
				return nil, fmt.Errorf("xxx: Group '%s' bulunamadı", a)
			}
		} else {
			return nil, fmt.Errorf("xxx: Group argümanı doğru değil")
		}
	}

	return groupList, nil
}
func (s GroupEngine) describe(ctx *application.ApplicationContext, models []basic_group_payload_structs.GroupBaseStruct) error {

	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	for _, group := range models {

		projectList, err := projectService.ListByGroupName(group.Header.Name)
		if err != nil {
			return fmt.Errorf("Failed to retrieve group projects '%s'\n%w", group.Header.Name, err)
		}

		resourceProjectList := []describe_printer.ProjectViewModel{}
		for _, e := range *projectList {
			projectName := e.Header.Name
			if !_string.IsEmpty(e.Specifications.Set) {
				projectName = fmt.Sprintf("%s (%s)", e.Header.Name, e.Specifications.Set)
			}
			resourceProject := describe_printer.ProjectViewModel{
				Name: projectName,
			}
			resourceProjectList = append(resourceProjectList, resourceProject)
		}

		viewModel := describe_printer.ViewModel{
			Name:     group.Header.Name,
			Tags:     group.Header.Metadata.Tags,
			Path:     group.Specifications.Path,
			Package:  group.Specifications.Package,
			Projects: resourceProjectList,
		}

		fmt.Printf("\n🛠️  Details for: %s\n\n", viewModel.Name)

		printer := describe_printer.DescribeGroup{Group: viewModel}

		if err := printer.Print(); err != nil {
			return fmt.Errorf("xxx: Group Print sırasında hata oluştu: %w", &err)
		}
	}

	return nil
}

func (s GroupEngine) Remove(ctx *application.ApplicationContext, args ...any) error {

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
func (s GroupEngine) prepareToRemove(ctx *application.ApplicationContext, args ...any) ([]basic_group_payload_structs.GroupBaseStruct, error) {

	service := ioc.Get[basic_group_contract.GroupInterface]()
	readyToRemoveStructs := make([]basic_group_payload_structs.GroupBaseStruct, 0)

	for _, a := range args {
		group, err := service.GetByName(a.(string))
		if err != nil {
			return nil, err
		}
		readyToRemoveStructs = append(readyToRemoveStructs, *group)
	}
	logrus.Debugf("'%d' group(s) detected that will remove", len(readyToRemoveStructs))

	return readyToRemoveStructs, nil
}
func (s GroupEngine) remove(ctx *application.ApplicationContext, models []basic_group_payload_structs.GroupBaseStruct, permanent bool) error {

	for _, group := range models {

		fmt.Printf("�️ Removing: %s.%s\n\n", group.Header.Name, group.GetKey())

		groupFlow := flowx.NewFlow("RemoveExistingGroup").
			Step(&remove_steps.DeleteGroup{})

		fc := flowx.NewContextWithData(map[string]any{
			"permanent": permanent,
			"group":     group,
		})

		if err := groupFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Group Remove işleminde hata oluştu: %w", &err)
		}
	}

	return nil
}

func (s GroupEngine) completeInformation(ctx *application.ApplicationContext, model *basic_group_payload_structs.GroupBaseStruct) error {

	logrus.Debugf("filling group (%v) information", model.Header.Name)

	if _string.IsEmpty(model.Specifications.Name) {
		model.Specifications.Name = model.Header.Name
	}

	return nil
}

func (s GroupEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  basic_group_payload_structs.MODULE_KEY,
		Order: 1000,
	}
}
func CastArrayToConcrate(data []schemas.SchemaInterface) ([]basic_group_payload_structs.GroupBaseStruct, error) {
	r := make([]basic_group_payload_structs.GroupBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*basic_group_payload_structs.GroupBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected basic_group_payload_structs.GroupBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}
