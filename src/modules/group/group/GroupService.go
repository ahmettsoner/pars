package group

import (
	"encoding/json"
	"fmt"

	"parsdevkit.net/application/ioc"

	"parsdevkit.net/modules/group/group_payload"
	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type GroupService struct {
	groupRespository   *repositories.GroupRepository
	projectRespository *repositories.ProjectRepository
}

func NewGroupService(environment string) contracts.GroupServiceInterface[group_payload.GroupBaseStruct] {

	groupRespository := ioc.Get[*repositories.GroupRepository]()
	projectRespository := ioc.Get[*repositories.ProjectRepository]()

	return &GroupService{
		groupRespository:   groupRespository,
		projectRespository: projectRespository,
	}
}

func (s GroupService) GetByName(name string) (*group_payload.GroupBaseStruct, error) {
	var group *group_payload.GroupBaseStruct

	entity, err := s.groupRespository.GetByName(name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Group getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity != nil {
		err = json.Unmarshal([]byte(entity.Document), &group)
		if err != nil {
			return nil, fmt.Errorf("xxx: Group data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
		}
	} else {
		group = nil
	}

	return group, nil
}

func (s GroupService) Save(model group_payload.GroupBaseStruct) (*group_payload.GroupBaseStruct, error) {

	result, err := s.saveGroupInformation(model)
	if err != nil {
		return nil, fmt.Errorf("xxx: Group kaydedilemedi %+v \n%w", model, err)
	}

	return result, nil
}

func (s GroupService) List() (*([]group_payload.GroupBaseStruct), error) {

	entityList, err := s.groupRespository.List()
	if err != nil {
		return nil, fmt.Errorf("xxx: Group listeleme aşamasında beklenmeyen hata oluştu\n%w", err)
	}

	groupList := make([]group_payload.GroupBaseStruct, 0)

	for _, entity := range *entityList {
		var group group_payload.GroupBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &group)
		if err != nil {
			return nil, fmt.Errorf("xxx: Group data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
		}

		groupList = append(groupList, group)
	}

	return &groupList, nil
}
func (s GroupService) Remove(name string, permanent bool) (*group_payload.GroupBaseStruct, error) {
	//TODO: Geçici olarak tanımlandı, düzenlenecek

	groupGroupEntity, err := s.groupRespository.GetByName(name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Group getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if groupGroupEntity == nil {
		return nil, fmt.Errorf("xxx: Group tanımlı değil '%s'", name)
	}

	projectsBelongsToGroup, err := s.projectRespository.ListByGroup(groupGroupEntity.ID)
	if err != nil {
		return nil, fmt.Errorf("xxx: Group '%s' e ait proje listesi getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}

	if len(*projectsBelongsToGroup) > 0 {
		return nil, fmt.Errorf("xxx: Group '%s' bağlı projeler mevcut %+v", name, projectsBelongsToGroup)
	}

	logrus.Debugf("group %v deleting...", groupGroupEntity.Name)

	err = s.groupRespository.Delete(groupGroupEntity)
	if err != nil {
		return nil, fmt.Errorf("xxx: Group silme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}

	var group group_payload.GroupBaseStruct
	err = json.Unmarshal([]byte(groupGroupEntity.Document), &group)
	if err != nil {
		return nil, fmt.Errorf("xxx: Group data %+v is corrupted or not in the expected format\n%w", groupGroupEntity.Document, err)
	}

	// if !_string.IsEmpty(Specifications.Path) {
	// 	logrus.Debugf("project (%v) files/folders (%v) removing", name, Specifications.GetAbsoluteBaseProjectPath())
	// 	if err := os.RemoveAll(project.Specifications.GetAbsoluteBaseProjectPath()); err != nil {
	// 		return nil, err
	// 	}
	// 	logrus.Debugf("project (%v) files/folders removed", projectName)
	// } else {
	// 	logrus.Debugf("You should delete project files for project (%v)", projectName)
	// }

	return &group, nil
}

func (s *GroupService) IsExists(name string) (bool, error) {

	groupGroupEntity, err := s.groupRespository.GetByName(name)
	if err != nil {
		return false, fmt.Errorf("xxx: Group varlığı sorgulama aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if groupGroupEntity == nil {
		return false, nil
	}

	return true, nil
}
func (s GroupService) GetHash(name string) (string, error) {

	entity, err := s.groupRespository.GetByName(name)
	if err != nil {
		return "", fmt.Errorf("xxx: Group getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return "", fmt.Errorf("xxx: Group tanımlı değil '%s' Hash bilgisi alınamıyor", name)
	}

	return entity.Hash, nil
}

func (s GroupService) saveGroupInformation(groupModel group_payload.GroupBaseStruct) (*group_payload.GroupBaseStruct, error) {

	jsonData, err := json.Marshal(groupModel)
	if err != nil {
		return nil, fmt.Errorf("xxx: Group json'a dönüştürülemedi %+v\n%w", groupModel, err)
	}

	groupEntity := entities.Group{
		Name:     groupModel.Header.Name,
		Document: string(jsonData),
	}

	err = s.groupRespository.Save(&groupEntity)
	if err != nil {
		return nil, fmt.Errorf("xxx: Group kayıt aşamasında beklenmeyen hata oluştu %+v\n%w", groupEntity, err)
	}

	return &groupModel, nil
}
