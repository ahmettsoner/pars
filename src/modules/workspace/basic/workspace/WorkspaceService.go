package basic_workspace

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"parsdevkit.net/application/models/layer"
	"parsdevkit.net/application/models/section"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"

	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/modules/project/application_project_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

const (
	DEFAULT_WORKSPACE_PATH string = "workspace"
	CURRENT_WORKSPACE_ID   string = "current_workspace_id"
)

type WorkspaceService struct {
	environment          string
	workspaceRespository *repositories.WorkspaceRepository
	projectRespository   *repositories.ProjectRepository
	settingsRespository  *repositories.SettingsRepository
	projectService       application_project_contract.ProjectInterface
}

func NewWorkspaceService(environment string) basic_workspace_contract.WorkspaceInterface {
	workspaceRespository := ioc.Get[*repositories.WorkspaceRepository]()
	projectRespository := ioc.Get[*repositories.ProjectRepository]()
	settingsRespository := ioc.Get[*repositories.SettingsRepository]()
	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	return &WorkspaceService{
		environment:          environment,
		workspaceRespository: workspaceRespository,
		projectRespository:   projectRespository,
		settingsRespository:  settingsRespository,
		projectService:       projectService,
	}
}
func (s *WorkspaceService) correctWorkspaceName(name string) string {
	correctedName := strings.TrimSpace(name)
	workspaceName := func() string {
		if !_string.IsEmpty(correctedName) {
			return strings.TrimSpace(name)
		}
		return DEFAULT_WORKSPACE_PATH
	}()

	return workspaceName
}
func (s *WorkspaceService) correctOutputPath(output string) (string, error) {
	correctedOutput := strings.TrimSpace(output)
	outputPath, err := func() (*string, error) {
		if !_string.IsEmpty(correctedOutput) {
			correctedOutput = strings.TrimSpace(correctedOutput)
			return &correctedOutput, nil
		}

		currentDir, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		return &currentDir, nil
	}()
	if err != nil {
		return "", err
	}

	return *outputPath, nil
}

func (s WorkspaceService) saveWorkspaceInformation(workspaceModel basic_workspace_payload_structs.WorkspaceBaseStruct) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	jsonData, err := json.Marshal(workspaceModel)
	if err != nil {
		return nil, err
	}

	workspaceEntity := entities.Workspace{
		Name:     workspaceModel.Header.Name,
		Document: string(jsonData),
	}

	err = s.workspaceRespository.Save(&workspaceEntity)
	if err != nil {
		return nil, err
	}

	value, err := s.settingsRespository.GetValue(CURRENT_WORKSPACE_ID)
	if err != nil {
		return nil, err
	}
	if _string.IsEmpty(value) {
		err := s.settingsRespository.SetValue(CURRENT_WORKSPACE_ID, strconv.Itoa(workspaceEntity.ID))
		if err != nil {
			return nil, err
		}
	}

	return &workspaceModel, nil
}

func (s *WorkspaceService) Get(id int) (*basic_workspace_payload_structs.WorkspaceSpecification, error) {
	var workspace *basic_workspace_payload_structs.WorkspaceSpecification

	entity, err := s.workspaceRespository.Get(id)
	if err != nil {
		return nil, err
	}

	if entity != nil {
		err = json.Unmarshal([]byte(entity.Document), &workspace)
	} else {
		workspace = nil
	}

	return workspace, nil
}

func (s WorkspaceService) GetByName(name string) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {
	var workspace *basic_workspace_payload_structs.WorkspaceBaseStruct

	entity, err := s.workspaceRespository.GetByName(name)
	if err != nil {
		return nil, err
	}
	if entity != nil {
		err = json.Unmarshal([]byte(entity.Document), &workspace)
	} else {
		workspace = nil
	}

	return workspace, nil
}

func (s WorkspaceService) ListByNameStartWith(name string) (*([]basic_workspace_payload_structs.WorkspaceBaseStruct), error) {

	entityList, err := s.workspaceRespository.ListByNameStartWith(name)
	if err != nil {
		return nil, err
	}

	workspaceList := make([]basic_workspace_payload_structs.WorkspaceBaseStruct, 0)

	for _, entity := range *entityList {
		var workspace basic_workspace_payload_structs.WorkspaceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &workspace)

		workspaceList = append(workspaceList, workspace)
	}

	return &workspaceList, nil
}

func (s WorkspaceService) Save(model basic_workspace_payload_structs.WorkspaceBaseStruct) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	workspaceName := s.correctWorkspaceName(model.Specifications.Name)
	outputPath, err := s.correctOutputPath(model.Specifications.Path)
	if err != nil {
		return nil, err
	}

	existingWorkspace, err := s.IsDirectoryReserved(outputPath)
	if err != nil {
		return nil, err
	}
	if existingWorkspace != nil {
		return nil, errors.New("Workspace directory reserved to (" + existingWorkspace.Header.Name + ")")
	}

	worskpaceExistingEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}

	if worskpaceExistingEntity != nil {
		return nil, errors.New("Workspace name (" + worskpaceExistingEntity.Name + ") already using")
	}

	isExists := s.directoryExists(model.Specifications.GetAbsolutePath())
	if isExists {
		isDirectory, err := s.isDirectory(model.Specifications.GetAbsolutePath())
		if err != nil {
			return nil, err
		}
		if !isDirectory {
			return nil, errors.New("Not valid directory specified: " + model.Specifications.GetAbsolutePath())
		}
		// return nil, errors.New("There is a workspace folder with same name: " + workspaceName)
	}

	if err := os.MkdirAll(model.Specifications.GetAbsolutePath(), os.ModePerm); err != nil {
		return nil, err
	}

	if err := os.Mkdir(model.Specifications.GetCodeBaseFolder(), os.ModePerm); err != nil {
		return nil, err
	}
	logrus.Debugf("Creating workspace codebase folder: %s", model.Specifications.GetCodeBaseFolder())

	if err := os.Mkdir(model.Specifications.GetTemplatesFolder(), os.ModePerm); err != nil {
		return nil, err
	}

	if err := os.Mkdir(model.Specifications.GetResourcesFolder(), os.ModePerm); err != nil {
		return nil, err
	}

	result, err := s.saveWorkspaceInformation(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// func (s *WorkspaceService) Create(name string, output string) (*workspace.WorkspaceSpecification, error) {
// 	workspaceName := s.correctWorkspaceName(name)
// 	outputPath, err := s.correctOutputPath(output)
// 	if err != nil {
// 		return nil, err
// 	}

// 	existingWorkspace, err := s.IsDirectoryReserved(outputPath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if existingWorkspace != nil {
// 		return nil, errors.New("Current directory reserved to (" + existingWorkspace.Name + ")")
// 	}

// 	worskpaceExistingEntity, err := s.workspaceRespository.GetByName(workspaceName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if worskpaceExistingEntity != nil {
// 		return nil, errors.New("Workspace name (" + worskpaceExistingEntity.Name + ") already using")
// 	}

// 	workspaceDir := filepath.Join(outputPath, workspaceName)
// 	isExists := s.directoryExists(workspaceDir)
// 	if isExists {
// 		return nil, errors.New("There is a workspace folder with same name: " + workspaceName)
// 	}

// 	workspace := workspace.NewWorkspaceSpecification(0, workspaceName, workspaceDir)

// 	if err := os.Mkdir(workspace.Path, os.ModePerm); err != nil {
// 		return nil, err
// 	}

// 	if err := os.Mkdir(workspace.GetCodeBaseFolder(), os.ModePerm); err != nil {
// 		return nil, err
// 	}
// 	logrus.Debugf("Creating workspace codebase folder: %s", workspace.GetCodeBaseFolder())

// 	if err := os.Mkdir(workspace.GetTemplatesFolder(), os.ModePerm); err != nil {
// 		return nil, err
// 	}

// 	if err := os.Mkdir(workspace.GetResourcesFolder(), os.ModePerm); err != nil {
// 		return nil, err
// 	}

// 	newWorkspacen := mapper.WorkspaceModelToEntity(workspace)
// 	newWorkspace := &newWorkspacen
// 	err = s.workspaceRespository.Create(newWorkspace)
// 	if err != nil {
// 		if err := os.RemoveAll(DEFAULT_WORKSPACE_PATH); err != nil {
// 			return nil, err
// 		}
// 		return nil, err
// 	}

// 	value, err := s.settingsRespository.GetValue(CURRENT_WORKSPACE_ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if _string.IsEmpty(value) {
// 		err := s.settingsRespository.SetValue(CURRENT_WORKSPACE_ID, strconv.Itoa(newWorkspace.ID))
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return &workspace, nil
// }

func (s WorkspaceService) List() (*([]basic_workspace_payload_structs.WorkspaceBaseStruct), error) {

	entityList, err := s.workspaceRespository.List()
	if err != nil {
		return nil, err
	}

	workspaceList := make([]basic_workspace_payload_structs.WorkspaceBaseStruct, 0)

	for _, entity := range *entityList {
		var workspace basic_workspace_payload_structs.WorkspaceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &workspace)

		workspaceList = append(workspaceList, workspace)
	}

	return &workspaceList, nil
}

func (s WorkspaceService) Remove(name string, force bool, permanent bool) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {
	//TODO: Geçici olarak tanımlandı, düzenlenecek

	workspaceName := s.correctWorkspaceName(name)
	workspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}
	if workspaceEntity == nil {
		return nil, errors.New("invalid workspace workspace")
	}

	// projectsBelongsToWorkspace, err := s.projectService.ListByWorkspace(workspaceName)
	// if err != nil {
	// 	return nil, err
	// }

	// if len(*projectsBelongsToWorkspace) > 0 {
	// 	if force {
	// 		for _, project := range *projectsBelongsToWorkspace {
	// 			s.projectService.Remove(project.GetFullName(), project.Specifications.Workspace, force, permanent)
	// 		}
	// 	} else {
	// 		return nil, errors.New(fmt.Sprintf("workspace (%v) has related projects", workspaceName))
	// 	}
	// }

	logrus.Debugf("workspace %v deleting...", workspaceEntity.Name)

	err = s.workspaceRespository.Delete(workspaceEntity)
	if err != nil {
		if err := os.RemoveAll(DEFAULT_WORKSPACE_PATH); err != nil {
			return nil, err
		}
		return nil, err
	}

	var workspace basic_workspace_payload_structs.WorkspaceBaseStruct
	err = json.Unmarshal([]byte(workspaceEntity.Document), &workspace)
	if err != nil {
		return nil, err
	}

	value, err := s.settingsRespository.Get(CURRENT_WORKSPACE_ID)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(value.Value) == strconv.Itoa(workspaceEntity.ID) {
		err := s.settingsRespository.Delete(value)
		if err != nil {
			return nil, err
		}
	}

	err = os.RemoveAll(workspace.Specifications.Path)
	if err != nil {
		return nil, err
	}

	err = s.projectRespository.DeleteByWorkspace(workspaceName)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

// func (s *WorkspaceService) Remove(name string, permanent bool) (*workspace.WorkspaceSpecification, error) {
// 	workspaceName := s.correctWorkspaceName(name)
// 	workspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if workspaceEntity == nil {
// 		return nil, errors.New("Workspace name (" + name + ") is not correct")
// 	}

// 	err = s.workspaceRespository.Delete(workspaceEntity)
// 	if err != nil {
// 		if err := os.RemoveAll(DEFAULT_WORKSPACE_PATH); err != nil {
// 			return nil, err
// 		}
// 		return nil, err
// 	}

// 	value, err := s.settingsRespository.Get(CURRENT_WORKSPACE_ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if strings.TrimSpace(value.Value) == strconv.Itoa(workspaceEntity.ID) {
// 		err := s.settingsRespository.Delete(value)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	err = os.RemoveAll(workspaceEntity.Path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = s.projectRespository.DeleteByWorkspaceID(workspaceEntity.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	workspace := mapper.WorkspaceEntityToModel(*workspaceEntity)

// 	return &workspace, nil
// }

func (s *WorkspaceService) IsExists(name string) (bool, error) {

	workspaceWorkspaceEntity, err := s.workspaceRespository.GetByName(name)
	if err != nil {
		return false, fmt.Errorf("xxx: Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if workspaceWorkspaceEntity == nil {
		return false, nil
	}

	return true, nil
}

func (s *WorkspaceService) GetActiveWorkspace() (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return s.IsDirectoryReserved(workingDir)
}

func (s *WorkspaceService) IsDirectoryReserved(path string) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {
	directories := s.getDirectories(path)

	entityList, err := s.workspaceRespository.ListByStartWithPath(directories)
	if err != nil {
		return nil, err
	}

	workspace := basic_workspace_payload_structs.WorkspaceBaseStruct{}

	if len(*entityList) > 0 {
		err = json.Unmarshal([]byte((*entityList)[0].Document), &workspace)
		return &workspace, nil
	} else {
		return nil, nil
	}
}

func (s *WorkspaceService) GetSelectedWorkspace() (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {
	value, err := s.settingsRespository.GetValue(CURRENT_WORKSPACE_ID)
	if err != nil {
		return nil, err
	}
	if _string.IsEmpty(value) {
		return nil, nil
	}

	currentWorkspaceID, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}

	entity, err := s.workspaceRespository.Get(currentWorkspaceID)
	if err != nil {
		return nil, err
	}
	if entity != nil {
		var workspace basic_workspace_payload_structs.WorkspaceBaseStruct
		err = json.Unmarshal([]byte((*entity).Document), &workspace)
		return &workspace, nil
	} else {
		return nil, nil
	}

	return nil, nil
}

func (s *WorkspaceService) ChangeCurrentWorkspace(name string) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {
	workspaceName := s.correctWorkspaceName(name)
	workspaceEntity, err := s.workspaceRespository.GetByName(workspaceName)
	if err != nil {
		return nil, err
	}
	if workspaceEntity == nil {
		return nil, errors.New("Workspace name (" + name + ") is not correct")
	}

	err = s.settingsRespository.SetValue(CURRENT_WORKSPACE_ID, strconv.Itoa(workspaceEntity.ID))
	if err != nil {
		return nil, err
	}

	var workspace basic_workspace_payload_structs.WorkspaceBaseStruct
	err = json.Unmarshal([]byte((*workspaceEntity).Document), &workspace)

	return &workspace, nil
}

func (s *WorkspaceService) findProjectsInWorkspace(workspaceDirectory string) ([]string, error) {
	var folders []string

	rootDir := filepath.Clean(workspaceDirectory)
	err := filepath.WalkDir(workspaceDirectory, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != rootDir {
			relativePath, err := filepath.Rel(rootDir, path)
			if err == nil && !strings.HasPrefix(relativePath, "..") {
				folders = append(folders, relativePath)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return folders, err
}

func (s *WorkspaceService) directoryExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (s *WorkspaceService) isDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	return (info != nil && info.IsDir()), err
}
func (s *WorkspaceService) getDirectories(currentDir string) []string {
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

func (s WorkspaceService) GetHash(name string) (string, error) {

	entity, err := s.workspaceRespository.GetByName(name)
	if err != nil {
		return "", fmt.Errorf("xxx: Workspace getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return "", fmt.Errorf("xxx: Workspace tanımlı değil '%s' Hash bilgisi alınamıyor\n%w", name, err)
	}

	return entity.Hash, nil
}

func (s WorkspaceService) SaveWorkspace(model basic_workspace_payload_structs.WorkspaceBaseStruct) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	jsonData, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	workspaceEntity := entities.Workspace{
		Name:     model.Header.Name,
		Document: string(jsonData),
	}

	err = s.workspaceRespository.Save(&workspaceEntity)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (s *WorkspaceService) UndoSaveWorkspace(model basic_workspace_payload_structs.WorkspaceBaseStruct) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	err := s.workspaceRespository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Code Workspace silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}

	return &model, nil
}

func (s *WorkspaceService) DeleteWorkspace(model basic_workspace_payload_structs.WorkspaceBaseStruct) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	logrus.Debugf("template %v removing", model.Header.Name)

	err := s.workspaceRespository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Code Workspace silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}
	logrus.Debugf("template (%v) information removed", model)

	return &model, nil
}

func (s WorkspaceService) CreateWorkspaceFolder(model basic_workspace_payload_structs.WorkspaceBaseStruct) (string, error) {

	outputPath, err := s.correctOutputPath(model.Specifications.Path)
	if err != nil {
		return "", err
	}

	existingWorkspace, err := s.IsDirectoryReserved(outputPath)
	if err != nil {
		return "", err
	}
	if existingWorkspace != nil && existingWorkspace.Header.Name != model.Header.Name {
		return "", errors.New("Workspace directory reserved to (" + existingWorkspace.Header.Name + ")")
	}

	isExists := s.directoryExists(model.Specifications.GetAbsolutePath())
	if isExists {
		isDirectory, err := s.isDirectory(model.Specifications.GetAbsolutePath())
		if err != nil {
			return "", err
		}
		if !isDirectory {
			return "", errors.New("Not valid directory specified: " + model.Specifications.GetAbsolutePath())
		}
		// return "", errors.New("There is a workspace folder with same name: " + workspaceName)
	}

	if err := os.MkdirAll(model.Specifications.GetAbsolutePath(), os.ModePerm); err != nil {
		return "", err
	}

	if err := os.Mkdir(model.Specifications.GetCodeBaseFolder(), os.ModePerm); err != nil {
		return "", err
	}
	logrus.Debugf("Creating workspace codebase folder: %s", model.Specifications.GetCodeBaseFolder())

	if err := os.Mkdir(model.Specifications.GetTemplatesFolder(), os.ModePerm); err != nil {
		return "", err
	}

	if err := os.Mkdir(model.Specifications.GetResourcesFolder(), os.ModePerm); err != nil {
		return "", err
	}

	return outputPath, nil
}

func (s *WorkspaceService) Context(workspace, project, resource, template schemas.SchemaInterface, layer layer.LayerIdentifier, section section.SectionIdentifier) interface{} {
	if model, ok := workspace.(basic_workspace_payload_structs.WorkspaceBaseStruct); ok {
		return WorkspaceComposite{
			Workspace: s.structToModel(model),
			Original:  model,
		}
	}

	return nil
}
func (s *WorkspaceService) structToModel(model basic_workspace_payload_structs.WorkspaceBaseStruct) Workspace {

	var result Workspace = Workspace{
		Name: model.Header.Name,
	}

	return result
}

type WorkspaceComposite struct {
	Workspace
	Original basic_workspace_payload_structs.WorkspaceBaseStruct
}

type Workspace struct {
	Name string
}
