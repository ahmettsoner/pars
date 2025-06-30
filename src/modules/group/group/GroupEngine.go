package group

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"parsdevkit.net/core"
	"parsdevkit.net/core/schemas"
	"parsdevkit.net/core/utils"
)

type GroupEngine struct{}

func (s GroupEngine) Validate(data []schemas.Schema) bool {
	for _, item := range data {
		_, ok := item.(*GroupBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s GroupEngine) Process(ctx *core.ApplicationContext, data []schemas.Schema) error {
	groups := make([]GroupBaseStruct, 0, len(data))

	for _, item := range data {
		group, ok := item.(*GroupBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Process: expected GroupBaseStruct, got %T", item)
		}
		groups = append(groups, *group)
	}

	return s.createGroups(groups, false)
}
func (s GroupEngine) Destroy(ctx *core.ApplicationContext, data []schemas.Schema) error {
	groups := make([]GroupBaseStruct, 0, len(data))

	for _, item := range data {
		group, ok := item.(*GroupBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Destroy: expected GroupBaseStruct, got %T", item)
		}
		groups = append(groups, *group)
	}

	return s.removeGroups(groups, false)
}

func (s GroupEngine) createGroups(groups []GroupBaseStruct, init bool) error {

	groupsReadyToCreate := make([]GroupBaseStruct, 0)
	groupsForUpdate := make([]GroupBaseStruct, 0)
	groupService := NewGroupService(utils.GetEnvironment())

	for _, group := range groups {
		ok, err := groupService.IsExists(group.Header.Name)
		if err != nil {
			return err
		}
		if ok {
			newModelHash, err := utils.CalculateHashFromObject(group)
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

func (s GroupEngine) removeGroups(groups []GroupBaseStruct, permanent bool) error {

	GroupEngine := NewGroupService(utils.GetEnvironment())
	groupsReadyToDelete := make([]GroupBaseStruct, 0)
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
