package group

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/core/utilities/encrypt"
	"parsdevkit.net/modules/group/group/structs"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/core/utils"
)

type GroupEngine struct{}

func (s GroupEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*structs.GroupBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s GroupEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	groups := make([]structs.GroupBaseStruct, 0, len(data))

	for _, item := range data {
		group, ok := item.(*structs.GroupBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Process: expected GroupBaseStruct, got %T", item)
		}
		groups = append(groups, *group)
	}

	return s.createGroups(groups, false)
}
func (s GroupEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	groups := make([]structs.GroupBaseStruct, 0, len(data))

	for _, item := range data {
		group, ok := item.(*structs.GroupBaseStruct)
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

func (s GroupEngine) createGroups(groups []structs.GroupBaseStruct, init bool) error {

	groupsReadyToCreate := make([]structs.GroupBaseStruct, 0)
	groupsForUpdate := make([]structs.GroupBaseStruct, 0)
	groupService := NewGroupService(utils.GetEnvironment())

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

func (s GroupEngine) removeGroups(groups []structs.GroupBaseStruct, permanent bool) error {

	GroupEngine := NewGroupService(utils.GetEnvironment())
	groupsReadyToDelete := make([]structs.GroupBaseStruct, 0)
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
