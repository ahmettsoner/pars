package basic_group

import (
	"encoding/json"
	"fmt"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/models/layer"
	"parsdevkit.net/application/models/section"
	"parsdevkit.net/application/platforms"
	"parsdevkit.net/application/schemas"
	applicationGroup "parsdevkit.net/application/structs/group"
	"parsdevkit.net/models"
	"parsdevkit.net/modules/group/basic_group_contract"
	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type GroupService struct {
	groupRespository   *repositories.GroupRepository
	projectRespository *repositories.ProjectRepository
}

func NewGroupService(environment string) basic_group_contract.GroupInterface {

	groupRespository := ioc.Get[*repositories.GroupRepository]()
	projectRespository := ioc.Get[*repositories.ProjectRepository]()

	return &GroupService{
		groupRespository:   groupRespository,
		projectRespository: projectRespository,
	}
}

func (s GroupService) GetByName(name string) (*basic_group_payload_structs.GroupBaseStruct, error) {
	var group *basic_group_payload_structs.GroupBaseStruct

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

func (s GroupService) Save(model basic_group_payload_structs.GroupBaseStruct) (*basic_group_payload_structs.GroupBaseStruct, error) {

	result, err := s.saveGroupInformation(model)
	if err != nil {
		return nil, fmt.Errorf("xxx: Group kaydedilemedi %+v \n%w", model, err)
	}

	return result, nil
}

func (s GroupService) List() (*([]basic_group_payload_structs.GroupBaseStruct), error) {

	entityList, err := s.groupRespository.List()
	if err != nil {
		return nil, fmt.Errorf("xxx: Group listeleme aşamasında beklenmeyen hata oluştu\n%w", err)
	}

	groupList := make([]basic_group_payload_structs.GroupBaseStruct, 0)

	for _, entity := range *entityList {
		var group basic_group_payload_structs.GroupBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &group)
		if err != nil {
			return nil, fmt.Errorf("xxx: Group data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
		}

		groupList = append(groupList, group)
	}

	return &groupList, nil
}
func (s GroupService) Remove(name string, permanent bool) (*basic_group_payload_structs.GroupBaseStruct, error) {
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

	var group basic_group_payload_structs.GroupBaseStruct
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

func (s GroupService) saveGroupInformation(groupModel basic_group_payload_structs.GroupBaseStruct) (*basic_group_payload_structs.GroupBaseStruct, error) {

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

func (s GroupService) SaveGroup(model basic_group_payload_structs.GroupBaseStruct) (*basic_group_payload_structs.GroupBaseStruct, error) {

	jsonData, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	groupEntity := entities.Group{
		Name:     model.Header.Name,
		Document: string(jsonData),
	}

	err = s.groupRespository.Save(&groupEntity)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (s *GroupService) UndoSaveGroup(model basic_group_payload_structs.GroupBaseStruct) (*basic_group_payload_structs.GroupBaseStruct, error) {

	err := s.groupRespository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Code Group silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}

	return &model, nil
}

func (s *GroupService) DeleteGroup(model basic_group_payload_structs.GroupBaseStruct) (*basic_group_payload_structs.GroupBaseStruct, error) {

	logrus.Debugf("template %v removing", model.Header.Name)

	err := s.groupRespository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Code Group silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}
	logrus.Debugf("template (%v) information removed", model)

	return &model, nil
}

func (s *GroupService) Context(workspace, project, resource, template schemas.SchemaInterface, layer layer.LayerIdentifier, section section.SectionIdentifier) interface{} {

	if model, ok := project.(application_project_payload_structs.ProjectBaseStruct); ok {
		return GroupComposite{
			Group:    s.structToModel(model.Specifications.Platform.Type, model.Specifications.GroupObject),
			Original: model.Specifications.GroupObject,
		}
	}

	return nil
}
func (s *GroupService) structToModel(platform models.PlatformType, model applicationGroup.GroupIdentifier /*basic_group_payload_structs.GroupBaseStruct*/) Group {
	dependencies := model.Package
	manager := platforms.Get[application_project_payload_structs.ProjectBaseStruct](platform)

	var result Group = Group{
		Package: manager.PrintDependencies(dependencies),
		Name:    model.Name,
	}
	return result
}

type GroupComposite struct {
	Group
	Original applicationGroup.GroupIdentifier /*basic_group_payload_structs.GroupBaseStruct*/
}

type Group struct {
	Name    string
	Package string
}
