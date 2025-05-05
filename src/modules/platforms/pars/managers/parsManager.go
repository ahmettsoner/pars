package managers

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	applicationproject "parsdevkit.net/structs/project/application-project"

	"parsdevkit.net/platforms/core"
	"parsdevkit.net/platforms/pars/models"

	"parsdevkit.net/core/utils"
)

type ParsManager struct {
	core.BaseManager
}

func NewParsManager() ParsManager {
	return ParsManager{
		core.BaseManagerNew(":"),
	}
}

func (s ParsManager) GetPlatformVersion(platform applicationproject.Platform) models.ParsPlatformVersion {
	if utils.IsEmpty(platform.Version) {
		platformVersion := models.ParsPlatformVersions.BetaV1

		return platformVersion
	} else {
		platformVersion, err := models.ParsPlatformVersionEnumFromString(platform.Version)
		if err != nil {
			panic(err)
		}
		return platformVersion
	}
}

func (s ParsManager) CreateProject(project applicationproject.ProjectSpecification) error {

	if _, err := s.CreateProjectFolder(project); err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
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

func (s ParsManager) RemoveProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
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

func (s ParsManager) BuildProject(project applicationproject.ProjectSpecification) error {

	return nil
}

func (s ParsManager) CleanProject(project applicationproject.ProjectSpecification) error {

	return nil
}

func (s ParsManager) InstallProject(project applicationproject.ProjectSpecification) error {

	return nil
}

func (s ParsManager) TestProject(project applicationproject.ProjectSpecification) error {

	return nil
}

func (s ParsManager) PackageProject(project applicationproject.ProjectSpecification) error {

	return nil
}

func (s ParsManager) RunProject(project applicationproject.ProjectSpecification) error {

	return nil
}

func (s *ParsManager) removeClassLibraryDefaultFiles(project applicationproject.ProjectSpecification) error {
	var paths []string = []string{}

	return s.FileRemover(paths...)
}

func (s ParsManager) CreateGroup(project applicationproject.ProjectSpecification) error {

	return nil
}

func (s ParsManager) DeleteGroup(project applicationproject.ProjectSpecification) {
}

func (s ParsManager) AddToGroup(project applicationproject.ProjectSpecification) error {

	return nil
}

func (s ParsManager) RemoveFromGroup(project applicationproject.ProjectSpecification) error {

	return nil
}

func (s ParsManager) CreateProjectFolder(project applicationproject.ProjectSpecification, paths ...string) (string, error) {
	var folders []string
	var foldersRelative []string = []string{}
	folders = append(folders, project.GetAbsoluteProjectPath())
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

func (s ParsManager) AddPackageToProject(project applicationproject.ProjectSpecification, packages []applicationproject.Package) error {

	return nil
}

func (s ParsManager) RemovePackageFromProject(project applicationproject.ProjectSpecification, packages []applicationproject.Package) error {

	return nil
}

func (s ParsManager) AddReferenceToProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error {

	return nil
}

func (s ParsManager) RemoveReferenceFromProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error {

	return nil
}

func (s ParsManager) IsProjectFileExists(project applicationproject.ProjectSpecification) (bool, error) {
	return s.IsProjectFolderExists(project)
}

func (s ParsManager) IsGroupFileExists(project applicationproject.ProjectSpecification) (bool, error) {

	return s.IsGroupFolderExists(project)
}
