package managers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_string "parsdevkit.net/pkg/utilities/string"

	applicationProject "parsdevkit.net/application/structs/project"
	applicationproject "parsdevkit.net/modules/project/application_project_payload"
	"parsdevkit.net/pkg/utilities/file"

	"parsdevkit.net/platforms/core"
	"parsdevkit.net/platforms/go/models"

	"parsdevkit.net/providers"
)

type GoManager struct {
	core.BaseManager
}

func NewGoManager() GoManager {
	return GoManager{
		core.BaseManagerNew("/")}
}

func (s GoManager) GetPlatformVersion(platform applicationproject.Platform) models.GoPlatformVersion {
	if _string.IsEmpty(platform.Version) {
		platformVersion := models.GoPlatformVersions.Go121

		return platformVersion
	} else {
		platformVersion, err := models.GoPlatformVersionEnumFromString(platform.Version)
		if err != nil {
			panic(err)
		}
		return platformVersion
	}
}
func (s GoManager) CreateProject(project applicationproject.ProjectBaseStruct) error {

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

	err := providers.GoExecute(project.Specifications.GetAbsoluteProjectPath(), "mod", "init", s.GetProjectPackage(project))
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		err = s.AddToGroup(project)
		if err != nil {
			log.Fatal(err)
		}
	}

	// if _, err := s.CreateProjectFolderStructure(project); err != nil {
	// 	return err
	// }

	// if project.Specifications.Configuration.Dependencies != nil {
	// 	s.AddDependenciesToProject(project, project.Specifications.Configuration.Dependencies)
	// }

	// if project.Specifications.Configuration.References != nil {
	// 	s.AddReferenceToProject(project, project.Specifications.Configuration.References)
	// }

	return nil
}

func (s GoManager) RemoveProject(project applicationproject.ProjectBaseStruct) error {
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

func (s GoManager) BuildProject(project applicationproject.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.GoExecute(project.Specifications.GetCodeBasePath(), "build")
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.GoExecute(project.Specifications.GetCodeBasePath(), "build")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s GoManager) CleanProject(project applicationproject.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.GoExecute(project.Specifications.GetCodeBasePath(), "clean", filepath.Join(project.Specifications.GetAbsoluteGroupPath(), project.Specifications.Name))
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.GoExecute(project.Specifications.GetCodeBasePath(), "clean", filepath.Join(project.Specifications.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s GoManager) InstallProject(project applicationproject.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.GoExecute(project.Specifications.GetCodeBasePath(), "restore", project.Specifications.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.GoExecute(project.Specifications.GetCodeBasePath(), "restore", filepath.Join(project.Specifications.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s GoManager) TestProject(project applicationproject.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.GoExecute(project.Specifications.GetCodeBasePath(), "test", project.Specifications.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.GoExecute(project.Specifications.GetCodeBasePath(), "test", filepath.Join(project.Specifications.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s GoManager) PackageProject(project applicationproject.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.GoExecute(project.Specifications.GetCodeBasePath(), "publish", project.Specifications.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.GoExecute(project.Specifications.GetCodeBasePath(), "publish", filepath.Join(project.Specifications.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s GoManager) RunProject(project applicationproject.ProjectBaseStruct) error {
	if !_string.IsEmpty(project.Specifications.Group) {
		groupStatus, err := s.IsGroupFileExists(project)
		if err != nil {
			return err

		}

		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		}
	}

	err := providers.GoExecute(project.Specifications.GetAbsoluteGroupPath(), "run", s.GetProjectPackage(project))
	if err != nil {
		return err
	}
	return nil
}

func (s GoManager) CreateGroup(project applicationproject.ProjectBaseStruct) error {

	err := providers.GoExecute(project.Specifications.GetAbsoluteGroupPath(), "mod", "init", s.GetGroupPackage(project))
	if err != nil {
		return err
	}

	return nil
}

func (s GoManager) DeleteGroup(project applicationproject.ProjectBaseStruct) {
}

func (s GoManager) AddToGroup(project applicationproject.ProjectBaseStruct) error {

	relativeProjectPath, err := file.FindRelativePath(project.Specifications.GetAbsoluteGroupPath(), project.Specifications.GetAbsoluteProjectPath())
	if err != nil {
		return err
	}
	err = providers.GoExecute(project.Specifications.GetAbsoluteGroupPath(), "mod", "edit", "-replace", fmt.Sprintf("%v=%v", s.GetProjectPackage(project), relativeProjectPath))
	if err != nil {
		return err
	}

	err = providers.GoExecute(project.Specifications.GetAbsoluteGroupPath(), "get", s.GetProjectPackage(project))
	if err != nil {
		return err
	}

	return nil
}

func (s GoManager) RemoveFromGroup(project applicationproject.ProjectBaseStruct) error {
	err := providers.GoExecute(project.Specifications.GetAbsoluteGroupPath(), "mod", "tidy")
	if err != nil {
		return err
	}

	return nil
}

func (s GoManager) CreateProjectFolder(project applicationproject.ProjectBaseStruct, paths ...string) (string, error) {
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

func (s GoManager) AddDependenciesToProject(project applicationproject.ProjectBaseStruct, dependencies []applicationProject.Package) error {

	for _, _package := range dependencies {
		err := providers.GoExecute(project.Specifications.GetAbsoluteProjectPath(), "get", _package.GetFullName())
		if err != nil {
			return err
		}
	}

	return nil
}

func (s GoManager) RemoveDependenciesFromProject(project applicationproject.ProjectBaseStruct, dependencies []applicationProject.Package) error {
	fmt.Printf("remove package not implemented yet")
	// for _, _package := range dependencies {
	// }

	return nil
}

func (s GoManager) AddReferenceToProject(project applicationproject.ProjectBaseStruct, references []applicationproject.ProjectBaseStruct) error {

	for _, reference := range references {

		relativeProjectPath, err := file.FindRelativePath(project.Specifications.GetAbsoluteProjectPath(), reference.Specifications.GetAbsoluteProjectPath())
		if err != nil {
			return err
		}
		err = providers.GoExecute(project.Specifications.GetAbsoluteProjectPath(), "mod", "edit", "-replace", fmt.Sprintf("%v=%v", s.GetProjectPackage(reference), relativeProjectPath))
		if err != nil {
			return err
		}

		err = providers.GoExecute(project.Specifications.GetAbsoluteProjectPath(), "get", s.GetProjectPackage(reference))
		if err != nil {
			return err
		}
	}

	return nil
}

func (s GoManager) RemoveReferenceFromProject(project applicationproject.ProjectBaseStruct, references []applicationproject.ProjectBaseStruct) error {
	fmt.Printf("Remove reference not implemented yet")
	// for _, reference := range references {
	// }

	return nil
}
func (s GoManager) IsProjectFileExists(project applicationproject.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.Specifications.GetAbsoluteProjectPath(), s.GetProjectFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s GoManager) GetProjectFileName(project applicationproject.ProjectBaseStruct) string {
	return fmt.Sprintf("go.mod")
}

func (s GoManager) IsGroupFileExists(project applicationproject.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.Specifications.GetAbsoluteGroupPath(), s.GetGroupFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s GoManager) GetGroupFileName(project applicationproject.ProjectBaseStruct) string {
	return fmt.Sprintf("go.mod")
}
