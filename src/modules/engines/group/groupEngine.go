package group

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	textTemplate "text/template"

	groupStruct "parsdevkit.net/structs/group"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/sirupsen/logrus"
)

type GroupEngine struct{}

func (s GroupEngine) CreateGroupsFromTemplate(init bool, data any, groupFiles ...string) error {

	var allGroups []groupStruct.GroupBaseStruct = make([]groupStruct.GroupBaseStruct, 0)

	for _, groupFilePath := range groupFiles {

		var tmplFile = filepath.Join(utils.GetManagerTemplatesLocation(), groupFilePath)
		tmplContent, err := os.ReadFile(tmplFile)
		if err != nil {
			log.Fatal(err)
		}
		var outputBuffer bytes.Buffer
		err = textTemplate.Must(textTemplate.New("GroupFromGroup").Parse(string(tmplContent))).Execute(&outputBuffer, data)
		if err != nil {
			log.Fatal(err)
		}
		mainStr := outputBuffer.String()

		groupSerializer := GroupSerializer{}
		groups, err := groupSerializer.GetGroupStructsFromString(mainStr)
		if err != nil {
			return err
		}
		allGroups = append(allGroups, groups...)
	}

	if err := s.CreateGroups(allGroups, init); err != nil {
		return err
	}

	return nil
}
func (s GroupEngine) CreateGroupsFromFile(init bool, files ...string) error {
	if len(files) > 0 {
		allFiles := make([]string, 0)
		for _, path := range files {
			if !utils.IsEmpty(path) {
				files, err := utils.GetFilesInPath(path)
				if err != nil {
					return err
				}

				allFiles = append(allFiles, files...)
			}
		}

		logrus.Debugf("found %v files", len(allFiles))
		groupSerializer := GroupSerializer{}
		groupsFromFile, err := groupSerializer.GetGroupStructsFromFile(allFiles...)
		if err != nil {
			return err
		}

		logrus.Debugf("found %v group", len(groupsFromFile))
		if err := s.CreateGroups(groupsFromFile, init); err != nil {
			return err
		}
	}
	return nil
}

func (s GroupEngine) RemoveGroupsFromFile(permanent bool, files ...string) error {
	if len(files) > 0 {

		allFiles := make([]string, 0)
		for _, path := range files {
			if !utils.IsEmpty(path) {

				files, err := utils.GetFilesInPath(path)
				if err != nil {
					return err
				}

				allFiles = append(allFiles, files...)
			}
		}

		groupSerializer := GroupSerializer{}
		groupsFromFile, err := groupSerializer.GetGroupStructsFromFile(allFiles...)
		if err != nil {
			return err
		}

		if err := s.RemoveGroups(groupsFromFile, permanent); err != nil {
			return err
		}
	}
	return nil
}
func (s GroupEngine) CreateGroups(groups []groupStruct.GroupBaseStruct, init bool) error {

	groupsReadyToCreate := make([]groupStruct.GroupBaseStruct, 0)
	groupsForUpdate := make([]groupStruct.GroupBaseStruct, 0)
	GroupEngine := services.NewGroupService(utils.GetEnvironment())

	for _, group := range groups {
		if ok := GroupEngine.IsExists(group.Name); ok {
			newModelHash, err := utils.CalculateHashFromObject(group)
			if err != nil {
				return err
			}
			structHash := GroupEngine.GetHash(group.Name)

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

		if _, err := GroupEngine.Save(group); err != nil {
			return err
		}

		fmt.Printf("%v Group created\n", group.Name)
	}

	logrus.Debugf("updating %v groups ", len(groupsForUpdate))
	for _, group := range groupsForUpdate {

		if _, err := GroupEngine.Save(group); err != nil {
			return err
		}

		fmt.Printf("%v Group updated\n", group.Name)
	}

	return nil
}

func (s GroupEngine) RemoveGroups(groups []groupStruct.GroupBaseStruct, permanent bool) error {

	GroupEngine := services.NewGroupService(utils.GetEnvironment())
	groupsReadyToDelete := make([]groupStruct.GroupBaseStruct, 0)
	for _, group := range groups {
		if ok := GroupEngine.IsExists(group.Name); ok {
			groupsReadyToDelete = append(groupsReadyToDelete, group)
		}
	}

	for _, group := range groupsReadyToDelete {

		if _, err := GroupEngine.Remove(group.Name, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Group deleted\n", group.Name)

	}

	return nil
}
