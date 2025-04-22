package services

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"parsdevkit.net/structs/project"
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"

	platformsCommon "parsdevkit.net/platforms/common"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type ApplicationProjectService struct {
	ApplicationProjectServiceInterface
	workspaceRespository *repositories.WorkspaceRepository
	groupRespository     *repositories.GroupRepository
	projectRespository   *repositories.ProjectRepository
	settingsRespository  *repositories.SettingsRepository
}

func NewApplicationProjectService(environment string) *ApplicationProjectService {
	return &ApplicationProjectService{
		workspaceRespository: repositories.NewWorkspaceRepository(environment),
		groupRespository:     repositories.NewGroupRepository(environment),
		projectRespository:   repositories.NewProjectRepository(environment),
		settingsRespository:  repositories.NewSettingsRepository(environment),
	}
}

// OK!
func (s *ApplicationProjectService) Create(model applicationproject.ProjectBaseStruct, init bool) (*applicationproject.ProjectBaseStruct, error) {

	logrus.Debugf("project %v creating", model.Name)

	if _, err := s.saveProjectInformation(model); err != nil {
		_, err := s.rollbackSaveProjectInformation(model)
		if err != nil {
			logrus.Warnf("rollback %v failed too!", model.Name)
		}
		return nil, err
	}

	if init {
		_, err := s.GenerateProject(model)
		if err != nil {
			logrus.Warnf("rollback %v failed too!", model.Name)
		}
	}

	return &model, nil
}

func (s ApplicationProjectService) GetByName(name string) (*applicationproject.ProjectBaseStruct, error) {
	var template *applicationproject.ProjectBaseStruct

	entity, err := s.projectRespository.GetByName(name)
	if err != nil {
		return nil, err
	}
	if entity != nil {
		err = json.Unmarshal([]byte(entity.Document), &template)
	} else {
		template = nil
	}

	return template, nil
}
func (s *ApplicationProjectService) GenerateProject(model applicationproject.ProjectBaseStruct) (*applicationproject.ProjectBaseStruct, error) {

	result, err := s.GetByName(model.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	if _, err := s.CreateProjectFolder(model); err != nil {
		return nil, err
	}

	projectManager := platformsCommon.ManagerFactory(model.Specifications.Platform.Type)
	if !utils.IsEmpty(model.Specifications.Group) {
		groupStatus, err := projectManager.IsGroupFileExists(model.Specifications)
		if err != nil {
			return nil, err
		}
		if !groupStatus {
			err := projectManager.CreateGroup(model.Specifications)
			if err != nil {
				logrus.Debugf("Group %v cannot created for %v", model.Specifications.Group, model.Specifications.Name)
				return nil, err
			}
		}
	}

	if err := projectManager.CreateProject(model.Specifications); err != nil {
		s.rollbackSaveProjectInformation(model)
		_, err := s.rollbackSaveProjectInformation(model)
		if err != nil {
			logrus.Warnf("rollback %v failed too!", model.Name)
		}
		return nil, err
	}

	if !utils.IsEmpty(model.Specifications.Group) {
		err := projectManager.AddToGroup(model.Specifications)
		if err != nil {
			return nil, err
		}
	}

	if err := projectManager.RemoveDefaultFiles(model.Specifications); err != nil {
		return nil, err
	}

	if _, err := s.CreateAllProjectFolders(model); err != nil {
		return nil, err
	}

	if model.Specifications.Configuration.Dependencies != nil {
		err := s.AddPackageToProject(model, model.Specifications.Configuration.Dependencies...)
		if err != nil {
			return nil, err
		}
	}

	if model.Specifications.Configuration.References != nil {
		err := s.AddReferenceToProject(model, model.Specifications.Configuration.References...)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
func (s ApplicationProjectService) AddPackageToProject(model applicationproject.ProjectBaseStruct, packages ...applicationproject.Package) error {
	projectManager := platformsCommon.ManagerFactory(model.Specifications.Platform.Type)
	return projectManager.AddPackageToProject(model.Specifications, packages)
}
func (s ApplicationProjectService) RemovePackageToProject(model applicationproject.ProjectBaseStruct, packages ...applicationproject.Package) error {
	projectManager := platformsCommon.ManagerFactory(model.Specifications.Platform.Type)
	return projectManager.RemovePackageFromProject(model.Specifications, packages)
}
func (s ApplicationProjectService) AddReferenceToProject(model applicationproject.ProjectBaseStruct, references ...applicationproject.ProjectBaseStruct) error {
	projectManager := platformsCommon.ManagerFactory(model.Specifications.Platform.Type)

	for _, ref := range references {
		err := projectManager.AddReferenceToProject(model.Specifications, []applicationproject.ProjectSpecification{ref.Specifications})
		if err != nil {
			return err
		}
	}

	return nil
}
func (s ApplicationProjectService) RemoveReferenceFromProject(model applicationproject.ProjectBaseStruct, references ...applicationproject.ProjectBaseStruct) error {
	projectManager := platformsCommon.ManagerFactory(model.Specifications.Platform.Type)

	for _, ref := range references {
		err := projectManager.RemoveReferenceFromProject(model.Specifications, []applicationproject.ProjectSpecification{ref.Specifications})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s ApplicationProjectService) CreateProjectFolder(model applicationproject.ProjectBaseStruct, paths ...string) (string, error) {
	folders := utils.CombinePaths([]string{model.Specifications.GetAbsoluteProjectPath()}, paths)
	foldersRelative := utils.CombinePaths(paths)

	folderPath := filepath.Join(folders...)

	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return "", err
	}

	projectManager := platformsCommon.ManagerFactory(model.Specifications.Platform.Type)
	if len(paths) > 0 {
		projectManager.AddFolderToProjectDefinition(model.Specifications, paths...)
	}
	foldersRelativePath := filepath.Join(foldersRelative...)
	return foldersRelativePath, nil
}
func (s ApplicationProjectService) DeleteProjectFolder(model applicationproject.ProjectBaseStruct, paths ...string) (string, error) {
	folders := utils.CombinePaths([]string{model.Specifications.GetAbsoluteProjectPath()}, paths)
	foldersRelative := utils.CombinePaths(paths)

	folderPath := filepath.Join(folders...)

	logrus.Debugf("Deleting project folder %s", folderPath)
	if err := os.RemoveAll(folderPath); err != nil {
		return "", err
	}

	projectManager := platformsCommon.ManagerFactory(model.Specifications.Platform.Type)
	if len(paths) > 0 {
		projectManager.RemoveFolderFromProjectDefinition(model.Specifications, paths...)
	}
	foldersRelativePath := filepath.Join(foldersRelative...)
	return foldersRelativePath, nil
}

func (s ApplicationProjectService) CreateLayerFolder(project applicationproject.ProjectBaseStruct, layers ...applicationproject.Layer) error {
	for _, layer := range layers {
		_, err := s.CreateProjectFolder(project, layer.Path)
		if err != nil {
			return err
		}
	}

	return nil
}
func (s ApplicationProjectService) DeleteLayerFolder(project applicationproject.ProjectBaseStruct, layers ...applicationproject.Layer) error {
	for _, layer := range layers {
		_, err := s.DeleteProjectFolder(project, layer.Path)
		if err != nil {
			return err
		}
	}

	return nil
}
func (s ApplicationProjectService) CreateAllProjectFolders(project applicationproject.ProjectBaseStruct) ([]string, error) {

	folders := make([]string, 0)
	for _, value := range project.Specifications.Configuration.Layers {

		createdFolder, err := s.CreateProjectFolder(project, value.GetPathAsArray()...)
		if err != nil {
			return nil, err
		}

		folders = append(folders, createdFolder)
	}
	return folders, nil
}
func (s *ApplicationProjectService) AddFileToLayer(model applicationproject.ProjectBaseStruct, layer string, paths []string, filename string, content string) (*applicationproject.ProjectBaseStruct, error) {

	logrus.Debugf("file %v creating for project %v on layer %v", filename, model.Name, layer)

	var projectLayer *applicationproject.Layer = nil
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
			return nil, err
		}
		_, fileState := os.Stat(fullFilePath)
		var file *os.File
		if os.IsNotExist(fileState) {
			file, fileState = os.Create(fullFilePath)
			if fileState != nil {
				return nil, fileState
			}
			defer file.Close()

			_, writeError := file.Write(fileContent)
			if writeError != nil {
				return nil, writeError
			}
		} else {
			file, fileState = os.OpenFile(fullFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
			if fileState != nil {
				return nil, fileState
			}
			defer file.Close()

			_, writeError := file.Write(fileContent)
			if writeError != nil {
				return nil, writeError
			}
		}
	}

	return &model, nil
}

// OK!
func (s *ApplicationProjectService) List() (*([]applicationproject.ProjectBaseStruct), error) {

	entityList, err := s.projectRespository.ListByKind(string(project.StructKinds.Application))
	if err != nil {
		return nil, err
	}

	projectList := make([]applicationproject.ProjectBaseStruct, 0)

	for _, entity := range *entityList {
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)

		projectList = append(projectList, project)
	}

	return &projectList, nil
}

// OK!
func (s *ApplicationProjectService) ListBySet(set string) (*([]applicationproject.ProjectBaseStruct), error) {

	entityList, err := s.projectRespository.ListBySet(set)
	if err != nil {
		return nil, err
	}

	projectList := make([]applicationproject.ProjectBaseStruct, 0)

	for _, entity := range *entityList {
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)

		projectList = append(projectList, project)
	}

	return &projectList, nil
}

// OK!
func (s *ApplicationProjectService) ListByGroupName(group string) (*([]applicationproject.ProjectBaseStruct), error) {

	entityList, err := s.projectRespository.ListByGroupName(group)
	if err != nil {
		return nil, err
	}

	projectList := make([]applicationproject.ProjectBaseStruct, 0)

	for _, entity := range *entityList {
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)

		projectList = append(projectList, project)
	}

	return &projectList, nil
}

// OK!
func (s *ApplicationProjectService) ListBySetAndLayers(set string, layers ...string) (*([]applicationproject.ProjectBaseStruct), error) {

	entityList, err := s.projectRespository.ListBySetAndLayers(set, layers...)
	if err != nil {
		return nil, err
	}

	projectList := make([]applicationproject.ProjectBaseStruct, 0)

	for _, entity := range *entityList {
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)

		projectList = append(projectList, project)
	}

	return &projectList, nil
}

// OK!
func (s *ApplicationProjectService) ListByWorkspace(workspaceName string) (*([]applicationproject.ProjectBaseStruct), error) {

	entityList, err := s.projectRespository.ListByWorkspace(workspaceName)
	if err != nil {
		return nil, err
	}

	projectList := make([]applicationproject.ProjectBaseStruct, 0)

	for _, entity := range *entityList {
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)

		projectList = append(projectList, project)
	}

	return &projectList, nil

}

func (s *ApplicationProjectService) ListIndividualByWorkspace(workspaceName string) (*([]applicationproject.ProjectBaseStruct), error) {

	var projectList []applicationproject.ProjectBaseStruct = []applicationproject.ProjectBaseStruct{}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}
	if projectWorkspaceEntity == nil {
		return nil, errors.New("invalid project workspace")
	}

	entity, err := s.projectRespository.GetIndividualByWorkspaceName(workspaceName)
	if err != nil {
		return nil, err
	}

	if entity != nil {

		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, err
		}

		projectList = append(projectList, project)
	}

	return &projectList, nil
}

// OK!
func (s *ApplicationProjectService) ListByFullNameWorkspace(name string, workspaceName string) (*([]applicationproject.ProjectBaseStruct), error) {

	var projectList []applicationproject.ProjectBaseStruct = []applicationproject.ProjectBaseStruct{}

	projectGroup, projectName, err := project.ParseProjectFullName(name)
	if err != nil {
		log.Fatal(err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}
	if projectWorkspaceEntity == nil {
		return nil, errors.New("invalid project workspace")
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, err
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if utils.IsEmpty(projectName) && !utils.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, err
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
				return nil, err
			}

			projectList = append(projectList, project)
		}

		return &projectList, nil
	} else {
		entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
		if err != nil {
			return nil, err
		}
		if entity == nil {
			return nil, errors.New("Project name (" + name + ") is not correct")
		}
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, err
		}

		projectList = append(projectList, project)

		return &projectList, nil
	}
}

// OK!
func (s *ApplicationProjectService) GetByFullNameWorkspace(name string, workspaceName string) (*applicationproject.ProjectBaseStruct, error) {

	var result applicationproject.ProjectBaseStruct = applicationproject.ProjectBaseStruct{}
	logrus.Debugf("trying to find project with fullname (%v)", name)

	projectGroup, projectName, err := project.ParseProjectFullName(name)
	if err != nil {
		log.Fatal(err)
	}
	logrus.Debugf("fullname parsed to name (%v) and group (%v)", projectName, projectGroup)

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}
	if projectWorkspaceEntity == nil {
		return nil, errors.New("invalid project workspace")
	}

	if !utils.IsEmpty(projectName) {

		groupId := 0
		projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
		if err != nil {
			return nil, err
		}
		if projectGroupEntity != nil {
			groupId = projectGroupEntity.ID
		}

		if groupId > 0 {
			logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
		}

		entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
		if err != nil {
			return nil, err
		}
		if entity == nil {
			return nil, nil
		}

		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, err
		}

		result = project

	}
	return &result, nil
}
func (s *ApplicationProjectService) CheckIfWorkingOnProject() (*applicationproject.ProjectBaseStruct, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return s.IsDirectoryReserved(workingDir)
}

// OK!
func (s *ApplicationProjectService) ValidateProjectStructure(model applicationproject.ProjectSpecification) (bool, error) {

	projectManager := platformsCommon.ManagerFactory(model.Platform.Type)

	if !utils.IsEmpty(model.Group) {
		state, err := projectManager.IsGroupFolderExists(model)
		if err != nil {
			return false, err
		}
		if !state {
			return false, nil
		}

		state, err = projectManager.IsGroupFileExists(model)
		if err != nil {
			return false, err
		}
		if !state {
			return false, nil
		}
	}

	state, err := projectManager.IsProjectFolderExists(model)
	if err != nil {
		return false, err
	}
	if !state {
		return false, nil
	}

	state, err = projectManager.IsProjectFileExists(model)
	if err != nil {
		return false, err
	}
	if !state {
		return false, nil
	}

	state, err = projectManager.IsLayerFoldersExists(model)
	if err != nil {
		return false, err
	}
	if !state {
		return false, nil
	}

	return true, nil
}
func (s *ApplicationProjectService) ValidateProjectDependency(model applicationproject.ProjectSpecification, _package applicationproject.Package) (bool, error) {

	projectManager := platformsCommon.ManagerFactory(model.Platform.Type)

	packages, err := projectManager.ListPackagesFromProject(model)
	if err != nil {
		return false, err
	}

	for _, projectPackage := range packages {
		if projectPackage.Name == _package.Name && projectPackage.Version == _package.Version {
			return true, nil
		}
	}

	return false, nil
}

func (s *ApplicationProjectService) ValidateProjectDependencies(model applicationproject.ProjectSpecification) (bool, error) {

	projectManager := platformsCommon.ManagerFactory(model.Platform.Type)

	for _, _package := range model.Configuration.Dependencies {
		isValid, err := projectManager.HasPackageOnProject(model, _package)
		if err != nil {
			return false, err
		}
		if !isValid {
			return false, nil
		}
	}

	return true, nil
}

func (s *ApplicationProjectService) ValidateProjectReferences(model applicationproject.ProjectSpecification) (bool, error) {

	projectManager := platformsCommon.ManagerFactory(model.Platform.Type)

	for _, reference := range model.Configuration.References {
		isValid, err := projectManager.HasReferenceOnProject(model, reference.Specifications)
		if err != nil {
			return false, err
		}
		if !isValid {
			return false, nil
		}
	}

	return true, nil
}
func (s *ApplicationProjectService) IsProjectFileExists(model applicationproject.ProjectSpecification) (bool, error) {

	projectManager := platformsCommon.ManagerFactory(model.Platform.Type)
	state, err := projectManager.IsProjectFileExists(model)
	return state, err
}

func (s *ApplicationProjectService) IsDirectoryReserved(path string) (*applicationproject.ProjectBaseStruct, error) {
	directories := s.getDirectories(path)

	entityList, err := s.projectRespository.ListByStartWithPath(directories)
	if err != nil {
		return nil, err
	}

	projectList := make([]applicationproject.ProjectBaseStruct, 0)

	// for _, entity := range *entityList {
	// 	var project = s.projectMapper.ProjectEntityToStruct(entity)

	// 	projectList = append(projectList, project)
	// }

	for _, entity := range *entityList {
		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)

		projectList = append(projectList, project)
	}

	if len(projectList) > 0 {
		return &projectList[0], nil
	} else {
		return nil, nil
	}
}

func (s *ApplicationProjectService) Remove(name string, workspaceName string, force bool, permanent bool) (*applicationproject.ProjectSpecification, error) {
	logrus.Debugf("project(s) will be %v removing", name)
	projectGroup, projectName, err := project.ParseProjectFullName(name)
	if err != nil {
		log.Fatal(err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}
	if projectWorkspaceEntity == nil {
		return nil, errors.New("invalid project workspace")
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, err
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}
	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if utils.IsEmpty(projectName) && !utils.IsEmpty(projectGroup) {
		logrus.Debugf("all group projects %v removing", projectName)
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, err
		}

		var latestValue *applicationproject.ProjectSpecification
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, err
			}
			latestValue, _ = s.Remove(project.GetFullName(), workspaceName, force, permanent)
		}

		return latestValue, nil
	} else {

		logrus.Debugf("project %v removing", projectName)
		entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
		if err != nil {
			return nil, err
		}
		if entity == nil {
			logrus.Errorf("project '%v' not found in group '%v' (%d)", projectName, projectGroup, groupId)
			return nil, errors.New("Project name (" + name + ") is not correct")
		}
		// entity.Workspace = projectWorkspaceEntity

		var project applicationproject.ProjectBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &project)
		if err != nil {
			return nil, err
		}

		projectManager := platformsCommon.ManagerFactory(project.Specifications.Platform.Type)

		logrus.Debugf("project (%v) content removing", projectName)
		err = projectManager.RemoveProject(project.Specifications)
		if err != nil {
			return nil, err
		}

		groupStatus, err := projectManager.IsGroupFileExists(project.Specifications)
		if err != nil {
			return nil, err
		}

		if !utils.IsEmpty(project.Specifications.Group) {
			if !groupStatus {
				return nil, errors.New("Project group (" + project.Specifications.Group + ") is not correct")
			} else {
				err := projectManager.RemoveFromGroup(project.Specifications)
				if err != nil {
					log.Fatal(err)
					return nil, err
				}
			}
		}

		logrus.Debugf("project (%v) content removed", projectName)

		if !utils.IsEmpty(project.Specifications.Group) && len(project.Specifications.Path) > 0 {
			logrus.Debugf("project (%v) files/folders (%v) removing", projectName, project.Specifications.GetAbsoluteBaseProjectPath())
			if err := os.RemoveAll(project.Specifications.GetAbsoluteBaseProjectPath()); err != nil {
				return nil, err
			}
			logrus.Debugf("project (%v) files/folders removed", projectName)
		} else {
			logrus.Debugf("You should delete project files for project (%v)", projectName)
		}

		logrus.Debugf("project (%v) information removing", projectName)
		err = s.projectRespository.Delete(entity)
		if err != nil {
			return nil, err
		}
		logrus.Debugf("project (%v) information removed", projectName)

		if !utils.IsEmpty(project.Specifications.Group) {
			count, err := s.projectRespository.CountByWorkspaceIDAndGroup(workspaceName, projectGroup)
			if err != nil {
				return nil, err
			}

			if count == 0 {
				if !utils.IsEmpty(project.Specifications.Group) && len(project.Specifications.Path) > 0 {
					logrus.Debugf("project group (%v) has no other project inside, all things removing belong to group", projectGroup)
					logrus.Debugf("removing path %v \n project: %v, group: %v", project.Specifications.GetAbsoluteGroupPath(), project.Name, project.Specifications.Group)
					if err := os.RemoveAll(project.Specifications.GetAbsoluteBaseGroupPath()); err != nil {
						return nil, err
					}
				} else {
					logrus.Debugf("You should delete group files for grpup (%v)", projectGroup)
				}
				projectManager.DeleteGroup(project.Specifications)
			}
		}

		return &project.Specifications, nil
	}
}

func (s *ApplicationProjectService) IsExists(name string, workspaceName string) bool {
	projectGroup, projectName, err := project.ParseProjectFullName(name)
	if err != nil {
		log.Fatal(err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return false
	}
	if projectWorkspaceEntity == nil {
		return false
	}

	_, err = s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return false
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return false
	}
	if entity == nil {
		return false
	}

	return true
}

func (s ApplicationProjectService) GetHash(name string, workspaceName string) string {
	projectGroup, projectName, err := project.ParseProjectFullName(name)
	if err != nil {
		log.Fatal(err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return ""
	}
	if projectWorkspaceEntity == nil {
		return ""
	}

	_, err = s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return ""
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return ""
	}
	if entity == nil {
		return ""
	}

	return entity.Hash
}

func (s *ApplicationProjectService) Build(name string, workspaceName string) (*applicationproject.ProjectSpecification, error) {
	projectGroup, projectName, err := project.ParseProjectFullName(name)
	if err != nil {
		log.Fatal(err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}
	if projectWorkspaceEntity == nil {
		return nil, errors.New("invalid project workspace")
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, err
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if utils.IsEmpty(projectName) && !utils.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, err
		}

		var latestValue *applicationproject.ProjectSpecification
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, err
			}
			latestValue, _ = s.Build(project.GetFullName(), workspaceName)
		}

		return latestValue, nil
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, errors.New("Project name (" + name + ") is not correct")
	}
	// entity.Workspace = projectWorkspaceEntity
	var project applicationproject.ProjectBaseStruct
	err = json.Unmarshal([]byte(entity.Document), &project)
	if err != nil {
		return nil, err
	}
	projectManager := platformsCommon.ManagerFactory(project.Specifications.Platform.Type)

	err = projectManager.BuildProject(project.Specifications)
	if err != nil {
		log.Fatal(err)
	}

	return &project.Specifications, nil
}

func (s *ApplicationProjectService) Clean(name string, workspaceName string) (*applicationproject.ProjectSpecification, error) {
	projectGroup, projectName, err := project.ParseProjectFullName(name)
	if err != nil {
		log.Fatal(err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}
	if projectWorkspaceEntity == nil {
		return nil, errors.New("invalid project workspace")
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, err
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if utils.IsEmpty(projectName) && !utils.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, err
		}

		var latestValue *applicationproject.ProjectSpecification
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, err
			}
			latestValue, _ = s.Clean(project.GetFullName(), workspaceName)
		}

		return latestValue, nil
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, errors.New("Project name (" + name + ") is not correct")
	}
	// entity.Workspace = projectWorkspaceEntity

	var project applicationproject.ProjectBaseStruct
	err = json.Unmarshal([]byte(entity.Document), &project)
	if err != nil {
		return nil, err
	}

	projectManager := platformsCommon.ManagerFactory(project.Specifications.Platform.Type)

	err = projectManager.CleanProject(project.Specifications)
	if err != nil {
		log.Fatal(err)
	}

	return &project.Specifications, nil
}
func (s *ApplicationProjectService) Install(name string, workspaceName string) (*applicationproject.ProjectSpecification, error) {
	projectGroup, projectName, err := project.ParseProjectFullName(name)
	if err != nil {
		log.Fatal(err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}
	if projectWorkspaceEntity == nil {
		return nil, errors.New("invalid project workspace")
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, err
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if utils.IsEmpty(projectName) && !utils.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, err
		}

		var latestValue *applicationproject.ProjectSpecification
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, err
			}
			latestValue, _ = s.Install(project.GetFullName(), workspaceName)
		}

		return latestValue, nil
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, errors.New("Project name (" + name + ") is not correct")
	}
	// entity.Workspace = projectWorkspaceEntity

	var project applicationproject.ProjectBaseStruct
	err = json.Unmarshal([]byte(entity.Document), &project)
	if err != nil {
		return nil, err
	}

	projectManager := platformsCommon.ManagerFactory(project.Specifications.Platform.Type)

	err = projectManager.InstallProject(project.Specifications)
	if err != nil {
		log.Fatal(err)
	}

	return &project.Specifications, nil
}

func (s *ApplicationProjectService) Test(name string, workspaceName string) (*applicationproject.ProjectSpecification, error) {
	projectGroup, projectName, err := project.ParseProjectFullName(name)
	if err != nil {
		log.Fatal(err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}
	if projectWorkspaceEntity == nil {
		return nil, errors.New("invalid project workspace")
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, err
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}
	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if utils.IsEmpty(projectName) && !utils.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectName)
		if err != nil {
			return nil, err
		}

		var latestValue *applicationproject.ProjectSpecification
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, err
			}
			latestValue, _ = s.Test(project.GetFullName(), workspaceName)
		}

		return latestValue, nil
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, errors.New("Project name (" + name + ") is not correct")
	}
	// entity.Workspace = projectWorkspaceEntity

	var project applicationproject.ProjectBaseStruct
	err = json.Unmarshal([]byte(entity.Document), &project)
	if err != nil {
		return nil, err
	}

	projectManager := platformsCommon.ManagerFactory(project.Specifications.Platform.Type)

	err = projectManager.TestProject(project.Specifications)
	if err != nil {
		log.Fatal(err)
	}

	return &project.Specifications, nil
}

func (s *ApplicationProjectService) Release(name string, workspaceName string) (*applicationproject.ProjectSpecification, error) {
	projectGroup, projectName, err := project.ParseProjectFullName(name)
	if err != nil {
		log.Fatal(err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}
	if projectWorkspaceEntity == nil {
		return nil, errors.New("invalid project workspace")
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, err
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if utils.IsEmpty(projectName) && !utils.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, err
		}

		var latestValue *applicationproject.ProjectSpecification
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, err
			}
			latestValue, _ = s.Release(project.GetFullName(), workspaceName)
		}

		return latestValue, nil
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, errors.New("Project name (" + name + ") is not correct")
	}
	// entity.Workspace = projectWorkspaceEntity

	var project applicationproject.ProjectBaseStruct
	err = json.Unmarshal([]byte(entity.Document), &project)
	if err != nil {
		return nil, err
	}

	projectManager := platformsCommon.ManagerFactory(project.Specifications.Platform.Type)

	err = projectManager.PackageProject(project.Specifications)
	if err != nil {
		log.Fatal(err)
	}

	return &project.Specifications, nil
}

func (s *ApplicationProjectService) Run(name string, workspaceName string) (*applicationproject.ProjectSpecification, error) {
	projectGroup, projectName, err := project.ParseProjectFullName(name)
	if err != nil {
		log.Fatal(err)
	}

	projectWorkspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}
	if projectWorkspaceEntity == nil {
		return nil, errors.New("invalid project workspace")
	}

	groupId := 0
	projectGroupEntity, err := s.groupRespository.GetByName(projectGroup)
	if err != nil {
		return nil, err
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	//TODO: iyileştirilecek, kolay çözüm uygulandı
	if utils.IsEmpty(projectName) && !utils.IsEmpty(projectGroup) {
		projectEntities, err := s.projectRespository.ListByWorkspaceNameAndGroup(workspaceName, projectGroup)
		if err != nil {
			return nil, err
		}

		var latestValue *applicationproject.ProjectSpecification
		for _, entity := range *projectEntities {
			var project applicationproject.ProjectBaseStruct
			err = json.Unmarshal([]byte(entity.Document), &project)
			if err != nil {
				return nil, err
			}
			latestValue, _ = s.Run(project.GetFullName(), workspaceName)
		}

		return latestValue, nil
	}

	entity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, projectGroup, workspaceName)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, errors.New("Project name (" + name + ") is not correct")
	}
	// entity.Workspace = projectWorkspaceEntity

	var project applicationproject.ProjectBaseStruct
	err = json.Unmarshal([]byte(entity.Document), &project)
	if err != nil {
		return nil, err
	}

	projectManager := platformsCommon.ManagerFactory(project.Specifications.Platform.Type)

	err = projectManager.RunProject(project.Specifications)
	if err != nil {
		log.Fatal(err)
	}

	return &project.Specifications, nil
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
		return nil, err
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, groupName)
	}

	projectEntity, err := s.projectRespository.GetByNameGroupAndWorkspaceName(projectName, group, workspaceName)
	if err != nil {
		return nil, err
	}

	if projectEntity != nil {
		return nil, errors.New("Project name (" + projectEntity.Name + ") already using")
	}

	return projectEntity, nil
}
func (s *ApplicationProjectService) GetProjectWorkspace(workspaceName string) (*workspace.WorkspaceSpecification, error) {
	workspaceService := NewWorkspaceService(utils.GetEnvironment())
	workspace, err := workspaceService.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}

	if !utils.IsEmpty(workspaceName) {
		workspace, err := workspaceService.GetByName(workspaceName)
		if err != nil {
			return nil, err
		}
		if workspace == nil {
			return nil, errors.New("Workspace name (" + workspaceName + ") is not correct")
		}
	}

	return &workspace.Specifications, nil
}

func (s *ApplicationProjectService) saveProjectInformation(projectModel applicationproject.ProjectBaseStruct) (*applicationproject.ProjectBaseStruct, error) {

	logrus.Debugf("project %v information saving", projectModel.Name)

	jsonData, err := json.Marshal(projectModel)
	if err != nil {
		return nil, err
	}

	projectEntity := entities.Project{
		Name:     projectModel.Name,
		Document: string(jsonData),
	}

	err = s.projectRespository.Save(&projectEntity)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("project %v information saved", projectModel.Name)
	return &projectModel, nil
}

func (s *ApplicationProjectService) rollbackSaveProjectInformation(projectModel applicationproject.ProjectBaseStruct) (*applicationproject.ProjectBaseStruct, error) {

	logrus.Debugf("project %v information rolling back", projectModel.Name)

	jsonData, err := json.Marshal(projectModel)
	if err != nil {
		return nil, err
	}

	projectEntity := entities.Project{
		Name:     projectModel.Name,
		Document: string(jsonData),
	}

	err = s.projectRespository.Delete(&projectEntity)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("project %v information saved", projectModel.Name)
	return &projectModel, nil
}
