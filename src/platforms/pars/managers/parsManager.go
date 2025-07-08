package managers

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	applicationProject "parsdevkit.net/application/structs/project"

	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"

	"parsdevkit.net/models"
	_string "parsdevkit.net/pkg/utilities/string"
	"parsdevkit.net/platforms/core"
	parsModels "parsdevkit.net/platforms/pars/models"
)

type ParsManager struct {
	core.BaseManager
}

func NewParsManager() core.ApplicationPlatformManagerInterface {
	return ParsManager{
		core.BaseManagerNew(":"),
	}
}
func (s ParsManager) GetKey() models.PlatformType {
	return models.PlatformTypes.Pars
}

func (s ParsManager) GetPlatformVersion(platform application_project_payload_structs.Platform) parsModels.ParsPlatformVersion {
	if _string.IsEmpty(platform.Version) {
		platformVersion := parsModels.ParsPlatformVersions.BetaV1

		return platformVersion
	} else {
		platformVersion, err := parsModels.ParsPlatformVersionEnumFromString(platform.Version)
		if err != nil {
			panic(err)
		}
		return platformVersion
	}
}

func (s ParsManager) CreateProject(project application_project_payload_structs.ProjectBaseStruct) error {

	if _, err := s.CreateProjectFolder(project); err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		groupStatus, err := s.IsGroupFileExists(project)
		if err != nil {
			return err
		}
		if !groupStatus {
			err := s.CreateGroup(project)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// if _, err := s.CreateProjectFolderStructure2(project); err != nil {
	// 	return err
	// }

	return nil
}

func (s ParsManager) RemoveProject(project application_project_payload_structs.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := s.RemoveFromGroup(project)
			if err != nil {
				log.Fatal(err)
			}
			return nil
		}
	}
	return nil
}

func (s ParsManager) BuildProject(project application_project_payload_structs.ProjectBaseStruct) error {

	return nil
}

func (s ParsManager) CleanProject(project application_project_payload_structs.ProjectBaseStruct) error {

	return nil
}

func (s ParsManager) InstallProject(project application_project_payload_structs.ProjectBaseStruct) error {

	return nil
}

func (s ParsManager) TestProject(project application_project_payload_structs.ProjectBaseStruct) error {

	return nil
}

func (s ParsManager) PackageProject(project application_project_payload_structs.ProjectBaseStruct) error {

	return nil
}

func (s ParsManager) RunProject(project application_project_payload_structs.ProjectBaseStruct) error {

	return nil
}

func (s *ParsManager) removeClassLibraryDefaultFiles(project application_project_payload_structs.ProjectBaseStruct) error {
	var paths []string = []string{}

	return s.FileRemover(paths...)
}

func (s ParsManager) CreateGroup(project application_project_payload_structs.ProjectBaseStruct) error {

	return nil
}

func (s ParsManager) DeleteGroup(project application_project_payload_structs.ProjectBaseStruct) {
}

func (s ParsManager) AddToGroup(project application_project_payload_structs.ProjectBaseStruct) error {

	return nil
}

func (s ParsManager) RemoveFromGroup(project application_project_payload_structs.ProjectBaseStruct) error {

	return nil
}

func (s ParsManager) CreateProjectFolder(project application_project_payload_structs.ProjectBaseStruct, paths ...string) (string, error) {
	var folders []string
	var foldersRelative []string = []string{}
	folders = append(folders, project.Specifications.GetAbsoluteProjectPath())
	for _, path := range paths {
		folders = append(folders, path)
		foldersRelative = append(foldersRelative, path)
	}
	foldePath := filepath.Join(folders...)

	if err := os.MkdirAll(foldePath, os.ModePerm); err != nil {
		return "", err
	}

	foldersRelativePath := filepath.Join(foldersRelative...)
	return foldersRelativePath, nil
}

func (s ParsManager) AddDependenciesToProject(project application_project_payload_structs.ProjectBaseStruct, dependencies []applicationProject.Dependency) error {

	return nil
}

func (s ParsManager) RemoveDependenciesFromProject(project application_project_payload_structs.ProjectBaseStruct, dependencies []applicationProject.Dependency) error {

	return nil
}

func (s ParsManager) AddReferenceToProject(project application_project_payload_structs.ProjectBaseStruct, references []application_project_payload_structs.ProjectBaseStruct) error {

	return nil
}

func (s ParsManager) RemoveReferenceFromProject(project application_project_payload_structs.ProjectBaseStruct, references []application_project_payload_structs.ProjectBaseStruct) error {

	return nil
}

func (s ParsManager) IsProjectFileExists(project application_project_payload_structs.ProjectBaseStruct) (bool, error) {
	return s.IsProjectFolderExists(project)
}

func (s ParsManager) IsGroupFileExists(project application_project_payload_structs.ProjectBaseStruct) (bool, error) {

	return s.IsGroupFolderExists(project)
}
