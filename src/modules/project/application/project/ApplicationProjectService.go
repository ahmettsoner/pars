package application_project

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/models"
	"parsdevkit.net/pkg/utilities/file"
	_string "parsdevkit.net/pkg/utilities/string"
	"parsdevkit.net/structs/project"
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
	applicationProject "parsdevkit.net/application/structs/project"
	projectComponent "parsdevkit.net/components/project"
	"parsdevkit.net/platforms/core"
)

type ApplicationProjectService struct {
	workspaceRespository *repositories.WorkspaceRepository
	groupRespository     *repositories.GroupRepository
	projectRespository   *repositories.ProjectRepository
	settingsRespository  *repositories.SettingsRepository
	platformRegistry     map[models.PlatformType]func() core.ManagerInterface
}

func NewApplicationProjectService(environment string, platformRegistry map[models.PlatformType]func() core.ManagerInterface) contracts.ProjectServiceInterface[applicationproject.ProjectBaseStruct] {
	workspaceRespository := ioc.Get[*repositories.WorkspaceRepository]()
	groupRespository := ioc.Get[*repositories.GroupRepository]()
	projectRespository := ioc.Get[*repositories.ProjectRepository]()
	settingsRespository := ioc.Get[*repositories.SettingsRepository]()

	return &ApplicationProjectService{
		workspaceRespository: workspaceRespository,
		groupRespository:     groupRespository,
		projectRespository:   projectRespository,
		settingsRespository:  settingsRespository,
		platformRegistry:     platformRegistry,
	}
}

// OK!
func (s *ApplicationProjectService) Create(model applicationproject.ProjectBaseStruct, init bool) (*applicationproject.ProjectBaseStruct, error) {

	logrus.Debugf("project %v creating", model.Header.Name)

	if _, err := s.saveProjectInformation(model); err != nil {
		logrus.Warnf("yyy: rolling back the creating %v!", model.Header.Name)
		_, err := s.rollbackSaveProjectInformation(model)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project Information kayıt sırasında meydana gelen hata için uygulanan rollback hatası oluştu: '%s'\n%w", model.Header.Name, err)
		}
		return nil, fmt.Errorf("xxx: Application Project Information kayıt sırasında hata meydana geldi: '%s'\n%w", model.Header.Name, err)
	}

	if init {
		_, err := s.GenerateProject(model)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project oluşturma sırasında hata meydana geldi: '%s'\n%w", model.Header.Name, err)
		}
	}

	return &model, nil
}

func (s ApplicationProjectService) GetByName(name string) (*applicationproject.ProjectBaseStruct, error) {
	var template *applicationproject.ProjectBaseStruct

	entity, err := s.projectRespository.GetByName(name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity != nil {
		err = json.Unmarshal([]byte(entity.Document), &template)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
		}
	} else {
		template = nil
	}

	return template, nil
}
func (s *ApplicationProjectService) GetPlatformManager(platform models.PlatformType) (core.ManagerInterface, error) {
	factory, ok := s.platformRegistry[platform]
	if !ok {
		return nil, fmt.Errorf("xxx: Platform Manager bulunamadı '%s'", platform)
	}
	manager := factory()
	return manager, nil
}
func (s *ApplicationProjectService) GenerateProject(model applicationproject.ProjectBaseStruct) (*applicationproject.ProjectBaseStruct, error) {

	result, err := s.GetByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project oluştururken, proje getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", model.Header.Name, err)
	}

	if result == nil {
		return nil, fmt.Errorf("xxx: Application Project tanımlı değil '%s'", model.Header.Name)
	}

	if _, err := s.CreateProjectFolder(model); err != nil {
		return nil, fmt.Errorf("xxx: Application Project oluştururken, proje klasörü hazırlama aşamasında beklenmeyen hata oluştu '%s'\n%w", model.Header.Name, err)
	}

	projectManager, err := s.GetPlatformManager(model.Specifications.Platform.Type)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project oluştururken, Platform Manager bulunamadı '%s'\n%w", model.Header.Name, err)
	}
	if !_string.IsEmpty(model.Specifications.Group) {
		groupStatus, err := projectManager.IsGroupFileExists(model.Specifications)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project oluştururken, grup path kontrolü aşamasında beklenmeyen hata oluştu '%s'\n%w", model.Header.Name, err)
		}
		if !groupStatus {
			err := projectManager.CreateGroup(model.Specifications)
			if err != nil {
				return nil, fmt.Errorf("Group %v cannot created for %v\n%w", model.Specifications.Group, model.Specifications.Name, err)
			}
		}
	}

	if err := projectManager.CreateProject(model.Specifications); err != nil {
		_, err := s.rollbackSaveProjectInformation(model)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project oluştururken meydana gelen hata için uygulanan rollback hatası oluştu: '%s'\n%w", model.Header.Name, err)
		}
		return nil, fmt.Errorf("xxx: Application Project oluştururken hata meydana geldi: '%s'\n%w", model.Header.Name, err)
	}

	if !_string.IsEmpty(model.Specifications.Group) {
		err := projectManager.AddToGroup(model.Specifications)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project oluştururken, gruba ekleme işlemi sırasında hata meydana geldi: '%s'\n%w", model.Header.Name, err)
		}
	}

	if err := projectManager.RemoveDefaultFiles(model.Specifications); err != nil {
		return nil, fmt.Errorf("xxx: Application Project oluştururma aşamasında default files kaldırma işlemi sırasında hata meydana geldi: '%s'\n%w", model.Header.Name, err)
	}

	if _, err := s.CreateAllProjectFolders(model); err != nil {
		return nil, fmt.Errorf("xxx: Application Project oluştururken, projelerin dizinleri oluşturma sırasında hata meydana geldi: '%s'\n%w", model.Header.Name, err)
	}

	if model.Specifications.Configuration.Dependencies != nil {
		err := s.AddPackageToProject(model, model.Specifications.Configuration.Dependencies...)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project oluştururken, projelerin paketi ekleme sırasında hata meydana geldi: '%s'\n%w", model.Header.Name, err)
		}
	}

	if model.Specifications.Configuration.References != nil {
		err := s.AddReferenceToProject(model, model.Specifications.Configuration.References...)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project oluştururken, projelerin paketi ekleme sırasında hata meydana geldi: '%s'\n%w", model.Header.Name, err)
		}
	}

	return result, nil
}
func (s ApplicationProjectService) AddPackageToProject(model applicationproject.ProjectBaseStruct, packages ...applicationProject.Package) error {
	projectManager, err := s.GetPlatformManager(model.Specifications.Platform.Type)
	if err != nil {
		return fmt.Errorf("xxx: Application Project Paket eklerken, Platform Manager bulunamadı '%s'\n%w", model.Header.Name, err)
	}

	err = projectManager.AddPackageToProject(model.Specifications, packages)
	if err != nil {
		return fmt.Errorf("xxx: Application Project Paket eklerken hata oluştu: '%s' Bağımlılıklar: '%+v'\n%w", model.Header.Name, packages, err)
	}
	return nil

}
func (s ApplicationProjectService) RemovePackageToProject(model applicationproject.ProjectBaseStruct, packages ...applicationProject.Package) error {
	projectManager, err := s.GetPlatformManager(model.Specifications.Platform.Type)
	if err != nil {
		return fmt.Errorf("xxx: Application Project Paket kaldırırken, Platform Manager bulunamadı '%s'\n%w", model.Header.Name, err)
	}

	err = projectManager.RemovePackageFromProject(model.Specifications, packages)
	if err != nil {
		return fmt.Errorf("xxx: Application Project Paket kaldırırken hata oluştu: '%s' Bağımlılıklar: '%+v'\n%w", model.Header.Name, packages, err)
	}
	return nil
}
func (s ApplicationProjectService) AddReferenceToProject(model applicationproject.ProjectBaseStruct, references ...applicationproject.ProjectBaseStruct) error {
	projectManager, err := s.GetPlatformManager(model.Specifications.Platform.Type)
	if err != nil {
		return fmt.Errorf("xxx: Application Project Reference eklerken, Platform Manager bulunamadı '%s'\n%w", model.Header.Name, err)
	}

	for _, ref := range references {
		err := projectManager.AddReferenceToProject(model.Specifications, []applicationproject.ProjectSpecification{ref.Specifications})
		if err != nil {
			return fmt.Errorf("xxx: Application Project Reference eklerken hata oluştu: '%s' Bağımlılıklar: '%+v'\n%w", model.Header.Name, references, err)
		}
	}

	return nil
}
func (s ApplicationProjectService) RemoveReferenceFromProject(model applicationproject.ProjectBaseStruct, references ...applicationproject.ProjectBaseStruct) error {
	projectManager, err := s.GetPlatformManager(model.Specifications.Platform.Type)
	if err != nil {
		return fmt.Errorf("xxx: Application Project Reference kaldırırken, Platform Manager bulunamadı '%s'\n%w", model.Header.Name, err)
	}

	for _, ref := range references {
		err := projectManager.RemoveReferenceFromProject(model.Specifications, []applicationproject.ProjectSpecification{ref.Specifications})
		if err != nil {
			return fmt.Errorf("xxx: Application Project Reference kaldırırken hata oluştu: '%s' Bağımlılıklar: '%+v'\n%w", model.Header.Name, references, err)
		}
	}

	return nil
}

func (s ApplicationProjectService) CreateProjectFolder(model applicationproject.ProjectBaseStruct, paths ...string) (string, error) {
	folders := file.CombinePaths([]string{model.Specifications.GetAbsoluteProjectPath()}, paths)
	foldersRelative := file.CombinePaths(paths)

	folderPath := filepath.Join(folders...)

	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("xxx: Application Project proje klasörü oluştururken hata oluştu: '%s' Path: '%+v'\n%w", model.Header.Name, folderPath, err)
	}

	projectManager, err := s.GetPlatformManager(model.Specifications.Platform.Type)
	if err != nil {
		return "", fmt.Errorf("xxx: Application Project proje klasörü tanımlanırken, Platform Manager bulunamadı '%s'\n%w", model.Header.Name, err)
	}

	if len(paths) > 0 {
		err := projectManager.AddFolderToProjectDefinition(model.Specifications, paths...)
		if err != nil {
			return "", fmt.Errorf("xxx: Application Project proje klasörü tanımlanırken hata oluştu: '%s' Path: '%+v'\n%w", model.Header.Name, folderPath, err)
		}
	}
	foldersRelativePath := filepath.Join(foldersRelative...)
	return foldersRelativePath, nil
}
func (s ApplicationProjectService) DeleteProjectFolder(model applicationproject.ProjectBaseStruct, paths ...string) (string, error) {
	folders := file.CombinePaths([]string{model.Specifications.GetAbsoluteProjectPath()}, paths)
	foldersRelative := file.CombinePaths(paths)

	folderPath := filepath.Join(folders...)

	logrus.Debugf("Deleting project folder %s", folderPath)
	if err := os.RemoveAll(folderPath); err != nil {
		return "", fmt.Errorf("xxx: Application Project proje klasörü silinirken hata oluştu: '%s' Path: '%+v'\n%w", model.Header.Name, folderPath, err)
	}

	projectManager, err := s.GetPlatformManager(model.Specifications.Platform.Type)
	if err != nil {
		return "", fmt.Errorf("xxx: Application Project proje klasörü kaldırılırken, Platform Manager bulunamadı '%s'\n%w", model.Header.Name, err)
	}

	if len(paths) > 0 {
		err := projectManager.RemoveFolderFromProjectDefinition(model.Specifications, paths...)

		if err != nil {
			return "", fmt.Errorf("xxx: Application Project proje klasörü kaldırılırken hata oluştu: '%s' Path: '%+v'\n%w", model.Header.Name, folderPath, err)
		}
	}
	foldersRelativePath := filepath.Join(foldersRelative...)
	return foldersRelativePath, nil
}

func (s ApplicationProjectService) CreateLayerFolder(project applicationproject.ProjectBaseStruct, layers ...applicationProject.Layer) error {
	for _, layer := range layers {
		_, err := s.CreateProjectFolder(project, layer.Path)
		if err != nil {
			return fmt.Errorf("xxx: Application Project layer klasörü oluştururken hata oluştu: '%s'\n%w", project.Header.Name, err)
		}
	}

	return nil
}
func (s ApplicationProjectService) DeleteLayerFolder(project applicationproject.ProjectBaseStruct, layers ...applicationProject.Layer) error {
	for _, layer := range layers {
		_, err := s.DeleteProjectFolder(project, layer.Path)
		if err != nil {
			return fmt.Errorf("xxx: Application Project layer klasörü kaldırılırken hata oluştu: '%s'\n%w", project.Header.Name, err)
		}
	}

	return nil
}
func (s ApplicationProjectService) CreateAllProjectFolders(project applicationproject.ProjectBaseStruct) ([]string, error) {
	folders := make([]string, 0)
	for _, value := range project.Specifications.Configuration.Layers {

		createdFolder, err := s.CreateProjectFolder(project, value.GetPathAsArray()...)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project layer klasörü oluştururken hata oluştu: '%s'\n%w", project.Header.Name, err)
		}

		folders = append(folders, createdFolder)
	}
	return folders, nil
}
func (s *ApplicationProjectService) AddFileToLayer(model applicationproject.ProjectBaseStruct, layer string, paths []string, filename string, content string) (*applicationproject.ProjectBaseStruct, error) {

	logrus.Debugf("file %v creating for project %v on layer %v", filename, model.Header.Name, layer)

	var projectLayer *applicationProject.Layer = nil
	for _, layerItem := range model.Specifications.Configuration.Layers {
		if layerItem.Name == layer {
			projectLayer = &layerItem
			break
		}
	}

	if projectLayer != nil {
		fileContent := []byte(content)
		layerPathFolder := model.Specifications.GetAbsoluteProjectLayerPath(layer)
		fileDirPathFolder := filepath.Join(paths...)
		fileDirPath := filepath.Join(layerPathFolder, fileDirPathFolder)

		fullFilePath := filepath.Join(fileDirPath, filename)

		err := os.MkdirAll(fileDirPath, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project layer'a dosya eklerken hata oluştu: '%s' Path: '%+v'\n%w", model.Header.Name, fileDirPath, err)
		}
		_, fileState := os.Stat(fullFilePath)
		var file *os.File
		if os.IsNotExist(fileState) {
			file, fileState = os.Create(fullFilePath)
			if fileState != nil {
				return nil, fmt.Errorf("xxx: Application Project layer'a yeni dosya eklerken durum kontrolü hatası oluştu: '%s' Path: '%+v'\n%w", model.Header.Name, fileDirPath, fileState)
			}
			defer file.Close()

			_, writeError := file.Write(fileContent)
			if writeError != nil {
				return nil, fmt.Errorf("xxx: Application Project layer'a yeni dosya eklerken yazma işlemi hatası oluştu: '%s' Path: '%+v'\n%w", model.Header.Name, fileDirPath, writeError)
			}
		} else {
			file, fileState = os.OpenFile(fullFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
			if fileState != nil {
				return nil, fmt.Errorf("xxx: Application Project layer'da dosya güncellerken durum kontrolü hatası oluştu: '%s' Path: '%+v'\n%w", model.Header.Name, fileDirPath, fileState)
			}
			defer file.Close()

			_, writeError := file.Write(fileContent)
			if writeError != nil {
				return nil, fmt.Errorf("xxx: Application Project layer'da dosya güncellerken yazma işlemi hatası oluştu: '%s' Path: '%+v'\n%w", model.Header.Name, fileDirPath, writeError)
			}
		}
	}

	return &model, nil
}

// OK!
func (s *ApplicationProjectService) List() (*([]applicationproject.ProjectBaseStruct), error) {

	entityList, err := s.projectRespository.ListByKind(string(project.ProjectKinds.Application))
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project listeleme aşamasında beklenmeyen hata oluştu\n%w", err)
	}

	projectList := make([]applicationproject.ProjectBaseStruct, 0)

	for _, entity := range *entityList {
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
		}

		projectList = append(projectList, project)
	}

	return &projectList, nil
}

// OK!
func (s *ApplicationProjectService) ListBySet(set string) (*([]applicationproject.ProjectBaseStruct), error) {

	entityList, err := s.projectRespository.ListBySet(set)
	if err != nil {
		return nil, fmt.Errorf("xxx: Set'a ait Application Project listeleme aşamasında beklenmeyen hata oluştu, '%s'\n%w", set, err)
	}

	projectList := make([]applicationproject.ProjectBaseStruct, 0)

	for _, entity := range *entityList {
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, fmt.Errorf("xxx: Set'a ait Application Project data %+v is corrupted or not in the expected format, '%s'\n%w", entity.Document, set, err)
		}

		projectList = append(projectList, project)
	}

	return &projectList, nil
}

// OK!
func (s *ApplicationProjectService) ListByGroupName(group string) (*([]applicationproject.ProjectBaseStruct), error) {

	entityList, err := s.projectRespository.ListByGroupName(group)
	if err != nil {
		return nil, fmt.Errorf("xxx: Group'a ait Application Project listeleme aşamasında beklenmeyen hata oluştu, '%s'\n%w", group, err)
	}

	projectList := make([]applicationproject.ProjectBaseStruct, 0)

	for _, entity := range *entityList {
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, fmt.Errorf("xxx: Group'a ait Application Project data %+v is corrupted or not in the expected format, '%s'\n%w", entity.Document, group, err)
		}

		projectList = append(projectList, project)
	}

	return &projectList, nil
}

// OK!
func (s *ApplicationProjectService) ListBySetAndLayers(set string, layers ...string) (*([]applicationproject.ProjectBaseStruct), error) {

	entityList, err := s.projectRespository.ListBySetAndLayers(set, layers...)
	if err != nil {
		return nil, fmt.Errorf("xxx: Set ve Layer'a ait Application Project listeleme aşamasında beklenmeyen hata oluştu, '%s', '%s'\n%w", set, layers, err)
	}

	projectList := make([]applicationproject.ProjectBaseStruct, 0)

	for _, entity := range *entityList {
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, fmt.Errorf("xxx: Set ve Layer'a ait Application Project data %+v is corrupted or not in the expected format, '%s', '%s'\n%w", entity.Document, set, layers, err)
		}

		projectList = append(projectList, project)
	}

	return &projectList, nil
}

// OK!
func (s *ApplicationProjectService) ListByWorkspace(workspaceName string) (*([]applicationproject.ProjectBaseStruct), error) {

	entityList, err := s.projectRespository.ListByWorkspace(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Workspace'a ait Application Project listeleme aşamasında beklenmeyen hata oluştu, '%s'\n%w", workspaceName, err)
	}

	projectList := make([]applicationproject.ProjectBaseStruct, 0)

	for _, entity := range *entityList {
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, fmt.Errorf("xxx: Workspace'a ait Application Project data %+v is corrupted or not in the expected format, '%s'\n%w", entity.Document, workspaceName, err)
		}

		projectList = append(projectList, project)
	}

	return &projectList, nil

}

func (s *ApplicationProjectService) ListIndividualByWorkspace(workspaceName string) (*([]applicationproject.ProjectBaseStruct), error) {

	var projectList []applicationproject.ProjectBaseStruct = []applicationproject.ProjectBaseStruct{}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Workspace'a ait bağımsız Application Project listeleme aşamasında beklenmeyen hata oluştu, '%s'\n%w", workspaceName, err)
	}
	if projectWorkspaceEntity == nil {
		return nil, fmt.Errorf("xxx: Application Project Workspace tanımlı değil '%s'\n%w", workspaceName, err)
	}

	entity, err := s.projectRespository.GetIndividualByWorkspaceName(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Workspace'a ait Bağımsız Application Project listeleme aşamasında beklenmeyen hata oluştu, '%s'\n%w", workspaceName, err)
	}

	if entity != nil {

		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, fmt.Errorf("xxx: Workspace'a ait Bağımsız Application Project data %+v is corrupted or not in the expected format, '%s'\n%w", entity.Document, workspaceName, err)
		}

		projectList = append(projectList, project)
	}

	return &projectList, nil
}

// OK!
func (s *ApplicationProjectService) ListByFullNameWorkspace(name string, workspaceName string) (*([]applicationproject.ProjectBaseStruct), error) {

	var projectList []applicationproject.ProjectBaseStruct = []applicationproject.ProjectBaseStruct{}

	projectGroup, projectName, err := projectComponent.ParseProjectFullName(name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Full proje id çözümlenemedi: '%s'\n%w", name, err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectWorkspaceEntity == nil {
		return nil, fmt.Errorf("xxx: Application Project Workspace tanımlı değil '%s'", workspaceName)
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Grup getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if _string.IsEmpty(projectName) && !_string.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, fmt.Errorf("xxx: Gruba ait Application Project Listeleme aşamasında beklenmeyen hata oluştu '%s'\n%w", projectGroup, err)
		}

		// for _, entity := range *projectEntities {
		// 	entity.Workspace = projectWorkspaceEntity
		// 	var project = s.projectMapper.ProjectEntityToStruct(entity)

		// 	projectList = append(projectList, project)
		// }

		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, fmt.Errorf("xxx: Gruba ait Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
			}

			projectList = append(projectList, project)
		}

		return &projectList, nil
	} else {
		entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
		if err != nil {
			return nil, fmt.Errorf("xxx: Gruba ait Application Project Listeleme aşamasında beklenmeyen hata oluştu '%s'\n%w", projectGroup, err)
		}
		if entity == nil {
			return nil, nil
		}
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, fmt.Errorf("xxx: Gruba ait Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
		}

		projectList = append(projectList, project)

		return &projectList, nil
	}
}

// OK!
func (s *ApplicationProjectService) GetByFullNameWorkspace(name string, workspaceName string) (*applicationproject.ProjectBaseStruct, error) {

	var result applicationproject.ProjectBaseStruct = applicationproject.ProjectBaseStruct{}
	logrus.Debugf("trying to find project with fullname (%v)", name)

	projectGroup, projectName, err := projectComponent.ParseProjectFullName(name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Full proje id çözümlenemedi: '%s'\n%w", name, err)
	}
	logrus.Debugf("fullname parsed to name (%v) and group (%v)", projectName, projectGroup)

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectWorkspaceEntity == nil {
		return nil, fmt.Errorf("xxx: Application Project Workspace tanımlı değil '%s'", workspaceName)
	}

	if !_string.IsEmpty(projectName) {

		groupId := 0
		projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project Grup getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
		}
		if projectGroupEntity != nil {
			groupId = projectGroupEntity.ID
		}

		if groupId > 0 {
			logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
		}

		entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
		if err != nil {
			return nil, fmt.Errorf("xxx: Gruba ait Application Project Listeleme aşamasında beklenmeyen hata oluştu '%s'\n%w", projectGroup, err)
		}
		if entity == nil {
			return nil, nil
		}

		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, fmt.Errorf("xxx: Gruba ait Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
		}

		result = project

	}
	return &result, nil
}
func (s *ApplicationProjectService) CheckIfWorkingOnProject() (*applicationproject.ProjectBaseStruct, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("xxx: Geçerli dizinin tespiti sırasında hata oluştu! \n%w", err)
	}

	project, err := s.IsDirectoryReserved(workingDir)
	if err != nil {
		return nil, fmt.Errorf("xxx: Geçerli dizinin reserve kontrolü sırasında hata oluştu! \n%w", err)
	}

	return project, nil
}

// OK!
func (s *ApplicationProjectService) ValidateProjectStructure(model applicationproject.ProjectBaseStruct) (bool, error) {

	projectManager, err := s.GetPlatformManager(model.Specifications.Platform.Type)
	if err != nil {
		return false, fmt.Errorf("xxx: Application Project File Structure Validasyon, Platform Manager bulunamadı '%s'\n%w", model.Specifications.Name, err)
	}

	if !_string.IsEmpty(model.Specifications.Group) {
		state, err := projectManager.IsGroupFolderExists(model.Specifications)
		if err != nil {
			return false, fmt.Errorf("xxx: Application Project File Structure Validasyon sırasında Group dizin(ler)i varlığı kontrol edilirken hata oluştu: '%s'\n%w", model.Specifications.Name, err)
		}
		if !state {
			return false, nil
		}

		state, err = projectManager.IsGroupFileExists(model.Specifications)
		if err != nil {
			return false, fmt.Errorf("xxx: Application ProjeProject File Structure ct Validasyon sırasında Group dosya(lar/s)ı varlığı kontrol edilirken hata oluştu: '%s'\n%w", model.Specifications.Name, err)
		}
		if !state {
			return false, nil
		}
	}

	state, err := projectManager.IsProjectFolderExists(model.Specifications)
	if err != nil {
		return false, fmt.Errorf("xxx: Application Project File Structure  Validasyon sırasında dizin(ler)i varlığı kontrol edilirken hata oluştu: '%s'\n%w", model.Specifications.Name, err)
	}
	if !state {
		return false, nil
	}

	state, err = projectManager.IsProjectFileExists(model.Specifications)
	if err != nil {
		return false, fmt.Errorf("xxx: Application Project File Structure  Validasyon sırasında dosya(lar/s)ı varlığı kontrol edilirken hata oluştu: '%s'\n%w", model.Specifications.Name, err)
	}
	if !state {
		return false, nil
	}

	state, err = projectManager.IsLayerFoldersExists(model.Specifications)
	if err != nil {
		return false, fmt.Errorf("xxx: Application Project File Structure  Validasyon sırasında Layer dizin(ler)i varlığı kontrol edilirken hata oluştu: '%s'\n%w", model.Specifications.Name, err)
	}
	if !state {
		return false, nil
	}

	return true, nil
}
func (s *ApplicationProjectService) ValidateProjectDependency(model applicationproject.ProjectBaseStruct, _package applicationProject.Package) (bool, error) {

	projectManager, err := s.GetPlatformManager(model.Specifications.Platform.Type)
	if err != nil {
		return false, fmt.Errorf("xxx: Application Project Dependency Validasyon sırasında paket listesi alınırken, Platform Manager bulunamadı '%s'\n%w", model.Specifications.Name, err)
	}

	packages, err := projectManager.ListPackagesFromProject(model.Specifications)
	if err != nil {
		return false, fmt.Errorf("xxx: Application Project Dependency Validasyon sırasında paket listesi alınırken hata oluştu: '%s'\n%w", model.Specifications.Name, err)
	}

	for _, projectPackage := range packages {
		if projectPackage.Name == _package.Name && projectPackage.Version == _package.Version {
			return true, nil
		}
	}

	return false, nil
}

func (s *ApplicationProjectService) ValidateProjectDependencies(model applicationproject.ProjectBaseStruct) (bool, error) {

	projectManager, err := s.GetPlatformManager(model.Specifications.Platform.Type)
	if err != nil {
		return false, fmt.Errorf("xxx: Application Project Dependency Validasyon sırasında projede paket kontrol edilirken, Platform Manager bulunamadı '%s'\n%w", model.Specifications.Name, err)
	}

	for _, _package := range model.Specifications.Configuration.Dependencies {
		isValid, err := projectManager.HasPackageOnProject(model.Specifications, _package)
		if err != nil {
			return false, fmt.Errorf("xxx: Application Project Dependency Validasyon sırasında projede paket kontrol edilirken hata oluştu: '%s'\n%w", model.Specifications.Name, err)
		}
		if !isValid {
			return false, nil
		}
	}

	return true, nil
}

func (s *ApplicationProjectService) ValidateProjectReferences(model applicationproject.ProjectBaseStruct) (bool, error) {

	projectManager, err := s.GetPlatformManager(model.Specifications.Platform.Type)
	if err != nil {
		return false, fmt.Errorf("xxx: Application Project Dependency Validasyon sırasında projede paket kontrol edilirken, Platform Manager bulunamadı '%s'\n%w", model.Specifications.Name, err)
	}

	for _, reference := range model.Specifications.Configuration.References {
		isValid, err := projectManager.HasReferenceOnProject(model.Specifications, reference.Specifications)
		if err != nil {
			return false, fmt.Errorf("xxx: Application Project Reference Validasyon sırasında projede referans kontrol edilirken hata oluştu: '%s'\n%w", model.Specifications.Name, err)
		}
		if !isValid {
			return false, nil
		}
	}

	return true, nil
}
func (s *ApplicationProjectService) IsProjectFileExists(model applicationproject.ProjectSpecification) (bool, error) {

	projectManager, err := s.GetPlatformManager(model.Platform.Type)
	if err != nil {
		return false, fmt.Errorf("xxx: Application Project proje dosyaları kontrolünde, Platform Manager bulunamadı '%s'\n%w", model.Name, err)
	}
	state, err := projectManager.IsProjectFileExists(model)
	if err != nil {
		return false, fmt.Errorf("xxx: Application Project proje dosyaları kontrolünde hata oluştu: '%s'\n%w", model.Name, err)
	}
	return state, err
}

func (s *ApplicationProjectService) IsDirectoryReserved(path string) (*applicationproject.ProjectBaseStruct, error) {
	directories := s.getDirectories(path)

	entityList, err := s.projectRespository.ListByStartWithPath(directories)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project reserve dizin dosyaları kontrolünde, belirtilen dizinde bulunan projeler listelenirken hata oluştu: '%s'\n%w", path, err)
	}

	projectList := make([]applicationproject.ProjectBaseStruct, 0)

	// for _, entity := range *entityList {
	// 	var project = s.projectMapper.ProjectEntityToStruct(entity)

	// 	projectList = append(projectList, project)
	// }

	for _, entity := range *entityList {
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project reserve dizin dosyaları kontrolünde Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
		}

		projectList = append(projectList, project)
	}

	if len(projectList) > 0 {
		return &projectList[0], nil
	} else {
		return nil, nil
	}
}

func (s *ApplicationProjectService) Remove(name string, workspaceName string, force bool, permanent bool) (*applicationproject.ProjectBaseStruct, error) {
	logrus.Debugf("project(s) will be %v removing", name)
	projectGroup, projectName, err := projectComponent.ParseProjectFullName(name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Full proje id çözümlenemedi: '%s'\n%w", name, err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectWorkspaceEntity == nil {
		return nil, fmt.Errorf("xxx: Application Project Workspace tanımlı değil '%s'", workspaceName)
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Grup getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}
	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if _string.IsEmpty(projectName) && !_string.IsEmpty(projectGroup) {
		logrus.Debugf("all group projects %v removing", projectName)
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, fmt.Errorf("xxx: Gruba ait Application Project Listeleme aşamasında beklenmeyen hata oluştu '%s'\n%w", projectGroup, err)
		}

		var latestValue *applicationproject.ProjectBaseStruct
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, fmt.Errorf("xxx: Gruba ait Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
			}
			latestValue, err = s.Remove(project.GetFullName(), workspaceName, force, permanent)
			if err != nil {
				return nil, fmt.Errorf("xxx: Gruba ait Application Project Silme aşamasında beklenmeyen hata oluştu '%s'\n%w", projectGroup, err)
			}
		}

		return latestValue, nil
	} else {

		logrus.Debugf("project %v removing", projectName)
		entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
		if err != nil {
			return nil, fmt.Errorf("xxx: Gruba ait Application Project Listeleme aşamasında beklenmeyen hata oluştu '%s'\n%w", projectName, err)
		}
		if entity == nil {
			logrus.Errorf("project '%v' not found in group '%v' (%d)", projectName, projectGroup, groupId)
			return nil, errors.New("Project name (" + name + ") is not correct")
		}
		// entity.Workspace = projectWorkspaceEntity

		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, fmt.Errorf("xxx: Gruba ait Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
		}

		projectManager, err := s.GetPlatformManager(project.Specifications.Platform.Type)
		if err != nil {
			return nil, fmt.Errorf("xxx: Gruba ait Application Project Silme aşamasında, Platform Manager bulunamadı '%s'\n%w", projectName, err)
		}

		logrus.Debugf("project (%v) content removing", projectName)
		err = projectManager.RemoveProject(project.Specifications)
		if err != nil {
			return nil, fmt.Errorf("xxx: Gruba ait Application Project Silme aşamasında beklenmeyen hata oluştu '%s'\n%w", projectName, err)
		}

		groupStatus, err := projectManager.IsGroupFileExists(project.Specifications)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project grup dosyaları kontrolünde hata oluştu: '%s'\n%w", projectName, err)
		}

		if !_string.IsEmpty(project.Specifications.Group) {
			if !groupStatus {
				return nil, errors.New("Project group (" + project.Specifications.Group + ") is not correct")
			} else {
				err := projectManager.RemoveFromGroup(project.Specifications)
				if err != nil {
					return nil, fmt.Errorf("xxx: Gruba ait Application Project Gruptan kaldırma işleminde beklenmeyen hata oluştu '%s'\n%w", projectName, err)
				}
			}
		}

		logrus.Debugf("project (%v) content removed", projectName)

		if !_string.IsEmpty(project.Specifications.Group) && len(project.Specifications.Path) > 0 {
			logrus.Debugf("project (%v) files/folders (%v) removing", projectName, project.Specifications.GetAbsoluteBaseProjectPath())
			if err := os.RemoveAll(project.Specifications.GetAbsoluteBaseProjectPath()); err != nil {
				return nil, fmt.Errorf("xxx: Application Project proje klasörü silinirken hata oluştu: '%s' Path: '%+v'\n%w", projectName, project.Specifications.GetAbsoluteBaseProjectPath(), err)
			}
			logrus.Debugf("project (%v) files/folders removed", projectName)
		} else {
			logrus.Debugf("You should delete project files for project (%v)", projectName)
		}

		logrus.Debugf("project (%v) information removing", projectName)
		err = s.projectRespository.Delete(entity)
		if err != nil {
			return nil, fmt.Errorf("xxx: Application Project silme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
		}
		logrus.Debugf("project (%v) information removed", projectName)

		if !_string.IsEmpty(project.Specifications.Group) {
			count, err := s.projectRespository.CountByWorkspaceIDAndGroup(workspaceName, projectGroup)
			if err != nil {
				return nil, fmt.Errorf("xxx: Grupta bulunan proje sayısı tespiti aşamasında beklenmeyen hata oluştu '%s'\n%w", projectGroup, err)
			}

			if count == 0 {
				if !_string.IsEmpty(project.Specifications.Group) && len(project.Specifications.Path) > 0 {
					logrus.Debugf("project group (%v) has no other project inside, all things removing belong to group", projectGroup)
					logrus.Debugf("removing path %v \n project: %v, group: %v", project.Specifications.GetAbsoluteGroupPath(), project.Header.Name, project.Specifications.Group)
					if err := os.RemoveAll(project.Specifications.GetAbsoluteBaseGroupPath()); err != nil {
						return nil, fmt.Errorf("xxx: GApplication Project silme aşamasında child proje bulunmayan grubun path'ini silme aşamasında beklenmeyen hata oluştu '%s'\n%w", projectGroup, err)
					}
				} else {
					logrus.Debugf("You should delete group files for grpup (%v)", projectGroup)
				}
				projectManager.DeleteGroup(project.Specifications)
			}
		}

		return &project, nil
	}
}

func (s *ApplicationProjectService) IsExists(name string, workspaceName string) (bool, error) {
	projectGroup, projectName, err := projectComponent.ParseProjectFullName(name)
	if err != nil {
		return false, fmt.Errorf("xxx: Full proje id çözümlenemedi: '%s'\n%w", name, err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return false, fmt.Errorf("xxx: Application Project Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectWorkspaceEntity == nil {
		return false, fmt.Errorf("xxx: Application Project Workspace tanımlı değil '%s'", workspaceName)
	}

	_, err = s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return false, fmt.Errorf("xxx: Application Project Grup getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", projectGroup, err)
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return false, fmt.Errorf("xxx: Application Project getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return false, nil
	}

	return true, nil
}

func (s ApplicationProjectService) GetHash(name string, workspaceName string) (string, error) {
	projectGroup, projectName, err := projectComponent.ParseProjectFullName(name)
	if err != nil {
		return "", fmt.Errorf("xxx: Full proje id çözümlenemedi: '%s'\n%w", name, err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return "", fmt.Errorf("xxx: Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectWorkspaceEntity == nil {
		return "", fmt.Errorf("xxx: Workspace tanımlı değil '%s' Hash bilgisi alınamıyor\n%w", workspaceName, err)
	}

	_, err = s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return "", fmt.Errorf("xxx: Group getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", projectGroup, err)
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return "", fmt.Errorf("xxx: Proje getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return "", fmt.Errorf("xxx: Proje tanımlı değil '%s' Hash bilgisi alınamıyor\n%w", name, err)
	}

	return entity.Hash, nil
}

func (s *ApplicationProjectService) Build(name string, workspaceName string) (*applicationproject.ProjectBaseStruct, error) {
	projectGroup, projectName, err := projectComponent.ParseProjectFullName(name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Full proje id çözümlenemedi: '%s'\n%w", name, err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectWorkspaceEntity == nil {
		return nil, fmt.Errorf("xxx: Workspace tanımlı değil '%s' Hash bilgisi alınamıyor\n%w", workspaceName, err)
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Grup getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if _string.IsEmpty(projectName) && !_string.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, fmt.Errorf("xxx: Gruba ait Application Project Listeleme aşamasında beklenmeyen hata oluştu '%s'\n%w", projectGroup, err)
		}

		var latestValue *applicationproject.ProjectBaseStruct
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, fmt.Errorf("xxx: Gruba ait Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
			}
			latestValue, err = s.Build(project.GetFullName(), workspaceName)
			if err != nil {
				return nil, fmt.Errorf("xxx: Gruba ait Application Project Build aşamasında beklenmeyen hata oluştu '%s'\n%w", projectGroup, err)
			}
		}

		return latestValue, nil
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Proje getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return nil, fmt.Errorf("xxx: Proje tanımlı değil '%s' Build edilemiyor", name)
	}
	// entity.Workspace = projectWorkspaceEntity
	var project applicationproject.ProjectBaseStruct
	err = json.Unmarshal([]byte(entity.Document), &project)
	if err != nil {
		return nil, fmt.Errorf("xxx: Gruba ait Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
	}
	projectManager, err := s.GetPlatformManager(project.Specifications.Platform.Type)
	if err != nil {
		return nil, fmt.Errorf("xxx: Gruba ait Application Project Build aşamasında, Platform Manager bulunamadı '%s'\n%w", name, err)
	}

	err = projectManager.BuildProject(project.Specifications)
	if err != nil {
		return nil, fmt.Errorf("xxx: Gruba ait Application Project Build aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}

	return &project, nil
}

func (s *ApplicationProjectService) CleanV2(name string, workspaceName string) (*applicationproject.ProjectBaseStruct, error) {
	projectGroup, projectName, err := projectComponent.ParseProjectFullName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to parse full project name '%s'\n%w", name, err)
	}

	workspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("failed to get workspace '%s'\n%w", name, err)
	}
	if workspaceEntity == nil {
		return nil, fmt.Errorf("there are no workspace '%s'\n%w", name, err)
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to get group '%s'\n%w", name, err)
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if _string.IsEmpty(projectName) && !_string.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, fmt.Errorf("failed to get project for group '%s' in workspace '%s'\n%w", projectGroup, workspaceName, err)
		}

		var latestValue *applicationproject.ProjectBaseStruct
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, fmt.Errorf("xxx: Gruba ait Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
			}
			latestValue, err = s.Clean(project.GetFullName(), workspaceName)
			if err != nil {
				return nil, fmt.Errorf("xxx: ait Application Project Clean aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
			}
		}

		return latestValue, nil
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Proje getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return nil, fmt.Errorf("xxx: Proje tanımlı değil '%s' Clean edilemiyor", name)
	}
	// entity.Workspace = workspaceEntity

	var project applicationproject.ProjectBaseStruct
	err = json.Unmarshal([]byte(entity.Document), &project)
	if err != nil {
		return nil, fmt.Errorf("xxx: Gruba ait Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
	}

	projectManager, err := s.GetPlatformManager(project.Specifications.Platform.Type)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Clean aşamasında, Platform Manager bulunamadı '%s'\n%w", name, err)
	}

	err = projectManager.CleanProject(project.Specifications)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Clean aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}

	return &project, nil
}

func (s *ApplicationProjectService) Clean(name string, workspaceName string) (*applicationproject.ProjectBaseStruct, error) {
	projectGroup, projectName, err := projectComponent.ParseProjectFullName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to parse full project name '%s'\n%w", name, err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectWorkspaceEntity == nil {
		return nil, fmt.Errorf("there are no workspace '%s'\n%w", name, err)
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to get group '%s'\n%w", name, err)
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if _string.IsEmpty(projectName) && !_string.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, fmt.Errorf("failed to get project for group '%s' in workspace '%s'\n%w", projectGroup, workspaceName, err)
		}

		var latestValue *applicationproject.ProjectBaseStruct
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, fmt.Errorf("xxx: Gruba ait Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
			}
			latestValue, err = s.Clean(project.GetFullName(), workspaceName)
			if err != nil {
				return nil, fmt.Errorf("xxx: ait Application Project Clean aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
			}
		}

		return latestValue, nil
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Proje getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return nil, fmt.Errorf("xxx: Proje tanımlı değil '%s' Clean edilemiyor", name)
	}
	// entity.Workspace = projectWorkspaceEntity

	var project applicationproject.ProjectBaseStruct
	err = json.Unmarshal([]byte(entity.Document), &project)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
	}

	projectManager, err := s.GetPlatformManager(project.Specifications.Platform.Type)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Clean aşamasında, Platform Manager bulunamadı '%s'\n%w", name, err)
	}

	err = projectManager.CleanProject(project.Specifications)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Clean aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}

	return &project, nil
}

func (s *ApplicationProjectService) Install(name string, workspaceName string) (*applicationproject.ProjectBaseStruct, error) {
	projectGroup, projectName, err := projectComponent.ParseProjectFullName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to parse full project name '%s'\n%w", name, err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectWorkspaceEntity == nil {
		return nil, fmt.Errorf("there are no workspace '%s'\n%w", name, err)
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to get group '%s'\n%w", name, err)
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if _string.IsEmpty(projectName) && !_string.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, fmt.Errorf("failed to get project for group '%s' in workspace '%s'\n%w", projectGroup, workspaceName, err)
		}

		var latestValue *applicationproject.ProjectBaseStruct
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, fmt.Errorf("xxx: Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
			}
			latestValue, err = s.Install(project.GetFullName(), workspaceName)
			if err != nil {
				return nil, fmt.Errorf("xxx: ait Application Project Install aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
			}
		}

		return latestValue, nil
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Proje getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return nil, fmt.Errorf("xxx: Proje tanımlı değil '%s' Install edilemiyor", name)
	}
	// entity.Workspace = projectWorkspaceEntity

	var project applicationproject.ProjectBaseStruct
	err = json.Unmarshal([]byte(entity.Document), &project)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
	}

	projectManager, err := s.GetPlatformManager(project.Specifications.Platform.Type)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Install aşamasında, Platform Manager bulunamadı '%s'\n%w", name, err)
	}

	err = projectManager.InstallProject(project.Specifications)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Install aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}

	return &project, nil
}

func (s *ApplicationProjectService) Test(name string, workspaceName string) (*applicationproject.ProjectBaseStruct, error) {
	projectGroup, projectName, err := projectComponent.ParseProjectFullName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to parse full project name '%s'\n%w", name, err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectWorkspaceEntity == nil {
		return nil, fmt.Errorf("there are no workspace '%s'\n%w", name, err)
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to get group '%s'\n%w", name, err)
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}
	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if _string.IsEmpty(projectName) && !_string.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectName)
		if err != nil {
			return nil, fmt.Errorf("failed to get project for group '%s' in workspace '%s'\n%w", projectGroup, workspaceName, err)
		}

		var latestValue *applicationproject.ProjectBaseStruct
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, fmt.Errorf("xxx: Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
			}
			latestValue, err = s.Test(project.GetFullName(), workspaceName)
			if err != nil {
				return nil, fmt.Errorf("xxx: ait Application Project Install aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
			}
		}

		return latestValue, nil
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Proje getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return nil, fmt.Errorf("xxx: Proje tanımlı değil '%s' Test edilemiyor", name)
	}
	// entity.Workspace = projectWorkspaceEntity

	var project applicationproject.ProjectBaseStruct
	err = json.Unmarshal([]byte(entity.Document), &project)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
	}

	projectManager, err := s.GetPlatformManager(project.Specifications.Platform.Type)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Test aşamasında, Platform Manager bulunamadı '%s'\n%w", name, err)
	}

	err = projectManager.TestProject(project.Specifications)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Test aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}

	return &project, nil
}

func (s *ApplicationProjectService) Release(name string, workspaceName string) (*applicationproject.ProjectBaseStruct, error) {
	projectGroup, projectName, err := projectComponent.ParseProjectFullName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to parse full project name '%s'\n%w", name, err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectWorkspaceEntity == nil {
		return nil, fmt.Errorf("there are no workspace '%s'\n%w", name, err)
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to get group '%s'\n%w", name, err)
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if _string.IsEmpty(projectName) && !_string.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, fmt.Errorf("failed to get project for group '%s' in workspace '%s'\n%w", projectGroup, workspaceName, err)
		}

		var latestValue *applicationproject.ProjectBaseStruct
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, fmt.Errorf("xxx: Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
			}
			latestValue, err = s.Release(project.GetFullName(), workspaceName)
			if err != nil {
				return nil, fmt.Errorf("xxx: ait Application Project Install aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
			}
		}

		return latestValue, nil
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Proje getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return nil, fmt.Errorf("xxx: Proje tanımlı değil '%s' Release edilemiyor", name)
	}
	// entity.Workspace = projectWorkspaceEntity

	var project applicationproject.ProjectBaseStruct
	err = json.Unmarshal([]byte(entity.Document), &project)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
	}

	projectManager, err := s.GetPlatformManager(project.Specifications.Platform.Type)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Release aşamasında, Platform Manager bulunamadı '%s'\n%w", name, err)
	}

	err = projectManager.PackageProject(project.Specifications)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Release aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}

	return &project, nil
}

func (s *ApplicationProjectService) Run(name string, workspaceName string) (*applicationproject.ProjectBaseStruct, error) {
	projectGroup, projectName, err := projectComponent.ParseProjectFullName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to parse full project name '%s'\n%w", name, err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}
	if projectWorkspaceEntity == nil {
		return nil, fmt.Errorf("there are no workspace '%s'\n%w", name, err)
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to get group '%s'\n%w", name, err)
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if _string.IsEmpty(projectName) && !_string.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, fmt.Errorf("failed to get project for group '%s' in workspace '%s'\n%w", projectGroup, workspaceName, err)
		}

		var latestValue *applicationproject.ProjectBaseStruct
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, fmt.Errorf("xxx: Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
			}
			latestValue, err = s.Run(project.GetFullName(), workspaceName)
			if err != nil {
				return nil, fmt.Errorf("xxx: ait Application Project Install aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
			}
		}

		return latestValue, nil
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Proje getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return nil, fmt.Errorf("xxx: Proje tanımlı değil '%s' Run edilemiyor", name)
	}
	// entity.Workspace = projectWorkspaceEntity

	var project applicationproject.ProjectBaseStruct
	err = json.Unmarshal([]byte(entity.Document), &project)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project data %+v is corrupted or not in the expected format\n%w", entity.Document, err)
	}

	projectManager, err := s.GetPlatformManager(project.Specifications.Platform.Type)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Run aşamasında, Platform Manager bulunamadı '%s'\n%w", name, err)
	}

	err = projectManager.RunProject(project.Specifications)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project Run aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}

	return &project, nil
}

func (s *ApplicationProjectService) getDirectories(currentDir string) []string {
	var directories []string

	for {
		directories = append(directories, currentDir)

		parentDir := filepath.Dir(currentDir)

		if parentDir == currentDir {
			break
		}

		currentDir = parentDir
	}

	return directories
}

func (s *ApplicationProjectService) correctProjectName(name string) string {
	correctedName := strings.TrimSpace(name)

	return correctedName
}
func (s *ApplicationProjectService) correctGroupName(name string) string {
	correctedName := strings.TrimSpace(name)

	return correctedName
}
func (s *ApplicationProjectService) getProject(name string, group string, workspaceName string) (*entities.Project, error) {

	projectName := s.correctProjectName(name)
	groupName := s.correctProjectName(group)

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(groupName)
	if err != nil {
		return nil, fmt.Errorf("failed to get group '%s'\n%w", name, err)
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, groupName)
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, group, workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Proje getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}

	if entity != nil {
		return nil, fmt.Errorf("Project name ("+entity.Name+") already using", name)
	}

	return entity, nil
}
func (s *ApplicationProjectService) GetProjectWorkspace(workspaceName string) (*workspace.WorkspaceSpecification, error) {
	workspaceService := ioc.Get[contracts.WorkspaceServiceInterface[workspace.WorkspaceBaseStruct]]()

	workspace, err := workspaceService.GetByName(workspaceName)
	if err != nil {
		return nil, fmt.Errorf("xxx: Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
	}

	if !_string.IsEmpty(workspaceName) {
		workspace, err := workspaceService.GetByName(workspaceName)
		if err != nil {
			return nil, fmt.Errorf("xxx: Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", workspaceName, err)
		}
		if workspace == nil {
			return nil, fmt.Errorf("there are no workspace '%s'\n%w", workspaceName, err)
		}
	}

	return &workspace.Specifications, nil
}

func (s *ApplicationProjectService) saveProjectInformation(projectModel applicationproject.ProjectBaseStruct) (*applicationproject.ProjectBaseStruct, error) {

	logrus.Debugf("project %v information saving", projectModel.Header.Name)

	jsonData, err := json.Marshal(projectModel)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project json'a dönüştürülemedi %+v\n%w", projectModel, err)
	}

	projectEntity := entities.Project{
		Name:     projectModel.Header.Name,
		Document: string(jsonData),
	}

	err = s.projectRespository.Save(&projectEntity)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project kayıt aşamasında beklenmeyen hata oluştu %+v\n%w", projectEntity, err)
	}

	logrus.Debugf("project %v information saved", projectModel.Header.Name)
	return &projectModel, nil
}

func (s *ApplicationProjectService) rollbackSaveProjectInformation(projectModel applicationproject.ProjectBaseStruct) (*applicationproject.ProjectBaseStruct, error) {

	logrus.Debugf("project %v information rolling back", projectModel.Header.Name)

	jsonData, err := json.Marshal(projectModel)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project json'a dönüştürülemedi %+v\n%w", projectModel, err)
	}

	projectEntity := entities.Project{
		Name:     projectModel.Header.Name,
		Document: string(jsonData),
	}

	err = s.projectRespository.Delete(&projectEntity)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project silme aşamasında beklenmeyen hata oluştu %+v\n%w", projectEntity, err)
	}

	logrus.Debugf("project %v information saved", projectModel.Header.Name)
	return &projectModel, nil
}
