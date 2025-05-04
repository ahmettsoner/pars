package services

import (
	"encoding/json"
	"errors"

	"parsdevkit.net/structs/group"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type GroupService struct {
	GroupServiceInterface
	groupRespository   *repositories.GroupRepository
	projectRespository *repositories.ProjectRepository
}

func NewGroupService(environment string) *GroupService {
	return &GroupService{
		groupRespository:   repositories.NewGroupRepository(environment),
		projectRespository: repositories.NewProjectRepository(environment),
	}
}

func (s GroupService) GetByName(name string) (*group.GroupBaseStruct, error) {
	var group *group.GroupBaseStruct

	entity, err := s.groupRespository.GetByName(name)
	if err != nil {
		return nil, err
	}
	if entity != nil {
		err = json.Unmarshal([]byte(entity.Document), &group)
	} else {
		group = nil
	}

	return group, nil
}

func (s GroupService) Save(model group.GroupBaseStruct) (*group.GroupBaseStruct, error) {

	result, err := s.saveGroupInformation(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s GroupService) List() (*([]group.GroupBaseStruct), error) {

	entityList, err := s.groupRespository.List()
	if err != nil {
		return nil, err
	}

	groupList := make([]group.GroupBaseStruct, 0)

	for _, entity := range *entityList {
		var group group.GroupBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &group)

		groupList = append(groupList, group)
	}

	return &groupList, nil
}
func (s GroupService) Remove(name string, permanent bool) (*group.GroupBaseStruct, error) {
	//TODO: Geçici olarak tanımlandı, düzenlenecek

	groupGroupEntity, err := s.groupRespository.GetByName(name)
	if err != nil {
		return nil, err
	}
	if groupGroupEntity == nil {
		return nil, errors.New("invalid group group")
	}

	projectsBelongsToGroup, err := s.projectRespository.ListByGroup(groupGroupEntity.ID)
	if err != nil {
		return nil, err
	}

	if len(*projectsBelongsToGroup) > 0 {
		return nil, errors.New("group has related projects")
	}

	logrus.Debugf("group %v deleting...", groupGroupEntity.Name)

	err = s.groupRespository.Delete(groupGroupEntity)

	var group group.GroupBaseStruct
	err = json.Unmarshal([]byte(groupGroupEntity.Document), &group)
	if err != nil {
		return nil, err
	}

	// if !utils.IsEmpty(group.Specifications.Path) {
	// 	logrus.Debugf("project (%v) files/folders (%v) removing", name, group.Specifications.GetAbsoluteBaseProjectPath())
	// 	if err := os.RemoveAll(project.Specifications.GetAbsoluteBaseProjectPath()); err != nil {
	// 		return nil, err
	// 	}
	// 	logrus.Debugf("project (%v) files/folders removed", projectName)
	// } else {
	// 	logrus.Debugf("You should delete project files for project (%v)", projectName)
	// }

	return &group, nil
}

func (s *GroupService) IsExists(name string) bool {

	groupGroupEntity, err := s.groupRespository.GetByName(name)
	if err != nil {
		return false
	}
	if groupGroupEntity == nil {
		return false
	}

	return true
}
func (s GroupService) GetHash(name string) string {

	entity, err := s.groupRespository.GetByName(name)
	if err != nil {
		return ""
	}
	if entity == nil {
		return ""
	}

	return entity.Hash
}

func (s GroupService) saveGroupInformation(groupModel group.GroupBaseStruct) (*group.GroupBaseStruct, error) {

	jsonData, err := json.Marshal(groupModel)
	if err != nil {
		return nil, err
	}

	groupEntity := entities.Group{
		Name:     groupModel.Name,
		Document: string(jsonData),
	}

	err = s.groupRespository.Save(&groupEntity)
	if err != nil {
		return nil, err
	}

	return &groupModel, nil
}
