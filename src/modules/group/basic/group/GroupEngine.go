package basic_group

import (
	"fmt"

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
func (s GroupEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
func (s GroupEngine) create(ctx *application.ApplicationContext, groups []basic_group_payload_structs.GroupBaseStruct, init bool) error {

	readyToCreateStructs, err := s.prepareToCreate(ctx, groups)
	if err != nil {
		return err
	}

	for _, group := range readyToCreateStructs {

		fmt.Printf("\n\n════════════════════════════════════\n")
		fmt.Printf("📦 Processing: %s.%s\n\n", group.Header.Name, group.GetKey())

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
func (s GroupEngine) update(ctx *application.ApplicationContext, groups []basic_group_payload_structs.GroupBaseStruct, init bool) error {

	readyToUpdateStructs, err := s.prepareToUpdate(ctx, groups)
	if err != nil {
		return err
	}
	for _, group := range readyToUpdateStructs {

		fmt.Printf("\n\n════════════════════════════════════\n")
		fmt.Printf("📦 Processing: %s.%s\n\n", group.Header.Name, group.GetKey())

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
func (s GroupEngine) prepareToRemove(ctx *application.ApplicationContext, groups []basic_group_payload_structs.GroupBaseStruct) ([]basic_group_payload_structs.GroupBaseStruct, error) {

	service := ioc.Get[basic_group_contract.GroupInterface]()
	readyToRemoveStructs := make([]basic_group_payload_structs.GroupBaseStruct, 0)

	for _, group := range groups {
		if err := s.completeInformation(ctx, &group); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(group.Header.Name)
		if err != nil {
			return nil, err
		}
		if ok {
			readyToRemoveStructs = append(readyToRemoveStructs, group)
		}
	}
	logrus.Debugf("'%d' group(s) detected that will remove", len(readyToRemoveStructs))

	return readyToRemoveStructs, nil
}
func (s GroupEngine) remove(ctx *application.ApplicationContext, groups []basic_group_payload_structs.GroupBaseStruct, permanent bool) error {

	readyToRemoveStructs, err := s.prepareToRemove(ctx, groups)
	if err != nil {
		return err
	}

	for _, group := range readyToRemoveStructs {

		fmt.Printf("\n\n════════════════════════════════════\n")
		fmt.Printf("📦 Processing: %s.%s\n\n", group.Header.Name, group.GetKey())

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
		Name:  "Group",
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
