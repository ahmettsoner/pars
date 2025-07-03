package group

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"parsdevkit.net/modules/group/group_payload"
	"parsdevkit.net/pkg/utilities/encrypt"

	"parsdevkit.net/application"
	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"
)

type GroupEngine struct{}

func (s GroupEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*group_payload.GroupBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s GroupEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	groups := make([]group_payload.GroupBaseStruct, 0, len(data))

	for _, item := range data {
		group, ok := item.(*group_payload.GroupBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Process: expected GroupBaseStruct, got %T", item)
		}
		groups = append(groups, *group)
	}

	return s.createGroups(groups, false)
}
func (s GroupEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	groups := make([]group_payload.GroupBaseStruct, 0, len(data))

	for _, item := range data {
		group, ok := item.(*group_payload.GroupBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Destroy: expected GroupBaseStruct, got %T", item)
		}
		groups = append(groups, *group)
	}

	return s.removeGroups(groups, false)
}
func (s GroupEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  "Group",
		Order: 1000,
	}
}

func (s GroupEngine) createGroups(groups []group_payload.GroupBaseStruct, init bool) error {

	groupsReadyToCreate := make([]group_payload.GroupBaseStruct, 0)
	groupsForUpdate := make([]group_payload.GroupBaseStruct, 0)
	groupService := ioc.Get[contracts.GroupServiceInterface[group_payload.GroupBaseStruct]]()

	for _, group := range groups {
		ok, err := groupService.IsExists(group.Header.Name)
		if err != nil {
			return err
		}
		if ok {
			newModelHash, err := encrypt.CalculateHashFromObject(group)
			if err != nil {
				return err
			}
			structHash, err := groupService.GetHash(group.Header.Name)
			if err != nil {
				return err
			}

			if newModelHash != structHash {
				groupsForUpdate = append(groupsForUpdate, group)
			}
		} else {
			groupsReadyToCreate = append(groupsReadyToCreate, group)
		}
	}
	logrus.Debugf("'%d' group(s) detected that will create", len(groupsReadyToCreate))
	logrus.Debugf("'%d' group(s) detected that will update", len(groupsForUpdate))

	logrus.Debugf("creating %v new groups ", len(groupsReadyToCreate))
	logrus.Debugf("updating %v groups ", len(groupsForUpdate))
	for _, group := range groupsReadyToCreate {

		if _, err := groupService.Save(group); err != nil {
			return err
		}

		fmt.Printf("%v Group created\n", group.Header.Name)
	}

	logrus.Debugf("updating %v groups ", len(groupsForUpdate))
	for _, group := range groupsForUpdate {

		if _, err := groupService.Save(group); err != nil {
			return err
		}

		fmt.Printf("%v Group updated\n", group.Header.Name)
	}

	return nil
}

func (s GroupEngine) removeGroups(groups []group_payload.GroupBaseStruct, permanent bool) error {

	GroupEngine := ioc.Get[contracts.GroupServiceInterface[group_payload.GroupBaseStruct]]()
	groupsReadyToDelete := make([]group_payload.GroupBaseStruct, 0)
	for _, group := range groups {
		ok, err := GroupEngine.IsExists(group.Header.Name)
		if err != nil {
			return err
		}
		if ok {
			groupsReadyToDelete = append(groupsReadyToDelete, group)
		}
	}

	for _, group := range groupsReadyToDelete {

		if _, err := GroupEngine.Remove(group.Header.Name, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Group deleted\n", group.Header.Name)

	}

	return nil
}
