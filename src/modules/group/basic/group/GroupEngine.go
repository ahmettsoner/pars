package basic_group

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"parsdevkit.net/modules/group/basic_group_payload"
	"parsdevkit.net/pkg/utilities/encrypt"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/modules/group/basic_group_contract"
	_string "parsdevkit.net/pkg/utilities/string"
)

type GroupEngine struct{}

func (s GroupEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*basic_group_payload.GroupBaseStruct)
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
func (s GroupEngine) prepareToCreate(ctx *application.ApplicationContext, groups []basic_group_payload.GroupBaseStruct) ([]basic_group_payload.GroupBaseStruct, error) {

	service := ioc.Get[basic_group_contract.GroupInterface]()
	readyToCreateStructs := make([]basic_group_payload.GroupBaseStruct, 0)

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
func (s GroupEngine) create(ctx *application.ApplicationContext, groups []basic_group_payload.GroupBaseStruct, init bool) error {

	service := ioc.Get[basic_group_contract.GroupInterface]()
	readyToCreateStructs, err := s.prepareToCreate(ctx, groups)
	if err != nil {
		return err
	}

	for index, group := range readyToCreateStructs {

		logrus.Debugf("trying to create %v", group.Header.Name)
		if _, err := service.Save(group); err != nil {
			return err
		}

		fmt.Printf("%v (%d) Group created\n", group.Header.Name, index)

	}

	return nil
}

func (s GroupEngine) prepareToUpdate(ctx *application.ApplicationContext, groups []basic_group_payload.GroupBaseStruct) ([]basic_group_payload.GroupBaseStruct, error) {

	service := ioc.Get[basic_group_contract.GroupInterface]()
	readyToUpdateStructs := make([]basic_group_payload.GroupBaseStruct, 0)

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
func (s GroupEngine) update(ctx *application.ApplicationContext, groups []basic_group_payload.GroupBaseStruct, init bool) error {

	service := ioc.Get[basic_group_contract.GroupInterface]()

	readyToUpdateStructs, err := s.prepareToUpdate(ctx, groups)
	if err != nil {
		return err
	}
	for _, group := range readyToUpdateStructs {
		if _, err := service.Save(group); err != nil {
			return err
		}
	}
	return nil
}
func (s GroupEngine) prepareToRemove(ctx *application.ApplicationContext, groups []basic_group_payload.GroupBaseStruct) ([]basic_group_payload.GroupBaseStruct, error) {

	service := ioc.Get[basic_group_contract.GroupInterface]()
	readyToRemoveStructs := make([]basic_group_payload.GroupBaseStruct, 0)

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
func (s GroupEngine) remove(ctx *application.ApplicationContext, groups []basic_group_payload.GroupBaseStruct, permanent bool) error {

	service := ioc.Get[basic_group_contract.GroupInterface]()

	readyToRemoveStructs, err := s.prepareToRemove(ctx, groups)
	if err != nil {
		return err
	}

	for _, group := range readyToRemoveStructs {

		if _, err := service.Remove(group.Header.Name, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Group deleted\n", group.Header.Name)

	}

	logrus.Debugf("'%d' group(s) deleting", len(readyToRemoveStructs))

	return nil
}

func (s GroupEngine) completeInformation(ctx *application.ApplicationContext, model *basic_group_payload.GroupBaseStruct) error {

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
func CastArrayToConcrate(data []schemas.SchemaInterface) ([]basic_group_payload.GroupBaseStruct, error) {
	r := make([]basic_group_payload.GroupBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*basic_group_payload.GroupBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected basic_group_payload.GroupBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}
