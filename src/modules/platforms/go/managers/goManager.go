package managers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	applicationproject "parsdevkit.net/structs/project/application-project"

	"parsdevkit.net/platforms/core"
	"parsdevkit.net/platforms/go/models"

	"parsdevkit.net/providers"

	"parsdevkit.net/core/utils"
)

type GoManager struct {
	core.BaseManager
}

func NewGoManager() GoManager {
	return GoManager{
		core.BaseManagerNew("/")}
}

func (s GoManager) GetPlatformVersion(platform applicationproject.Platform) models.GoPlatformVersion {
	if utils.IsEmpty(platform.Version) {
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
func (s GoManager) CreateProject(project applicationproject.ProjectSpecification) error {

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

	err := providers.GoExecute(project.GetAbsoluteProjectPath(), "mod", "init", s.GetProjectPackage(project))
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		err = s.AddToGroup(project)
		if err != nil {
			log.Fatal(err)
		}
	}

	// if _, err := s.CreateProjectFolderStructure(project); err != nil {
	// 	return err
	// }

	// if project.Configuration.Dependencies != nil {
	// 	s.AddPackageToProject(project, project.Configuration.Dependencies)
	// }

	// if project.Configuration.References != nil {
	// 	s.AddReferenceToProject(project, project.Configuration.References)
	// }

	return nil
}

func (s GoManager) RemoveProject(project applicationproject.ProjectSpecification) error {
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

func (s GoManager) BuildProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.GoExecute(project.GetCodeBasePath(), "build")
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.GoExecute(project.GetCodeBasePath(), "build")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s GoManager) CleanProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.GoExecute(project.GetCodeBasePath(), "clean", filepath.Join(project.GetAbsoluteGroupPath(), project.Name))
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.GoExecute(project.GetCodeBasePath(), "clean", filepath.Join(project.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s GoManager) InstallProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.GoExecute(project.GetCodeBasePath(), "restore", project.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.GoExecute(project.GetCodeBasePath(), "restore", filepath.Join(project.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s GoManager) TestProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.GoExecute(project.GetCodeBasePath(), "test", project.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.GoExecute(project.GetCodeBasePath(), "test", filepath.Join(project.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s GoManager) PackageProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.GoExecute(project.GetCodeBasePath(), "publish", project.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.GoExecute(project.GetCodeBasePath(), "publish", filepath.Join(project.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s GoManager) RunProject(project applicationproject.ProjectSpecification) error {
	if !utils.IsEmpty(project.Group) {
		groupStatus, err := s.IsGroupFileExists(project)
		if err != nil {
			return err

		}

		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		}
	}

	err := providers.GoExecute(project.GetAbsoluteGroupPath(), "run", s.GetProjectPackage(project))
	if err != nil {
		return err
	}
	return nil
}

func (s GoManager) CreateGroup(project applicationproject.ProjectSpecification) error {

	err := providers.GoExecute(project.GetAbsoluteGroupPath(), "mod", "init", s.GetGroupPackage(project))
	if err != nil {
		return err
	}

	return nil
}

func (s GoManager) DeleteGroup(project applicationproject.ProjectSpecification) {
}

func (s GoManager) AddToGroup(project applicationproject.ProjectSpecification) error {

	relativeProjectPath, err := utils.FindRelativePath(project.GetAbsoluteGroupPath(), project.GetAbsoluteProjectPath())
	if err != nil {
		return err
	}
	err = providers.GoExecute(project.GetAbsoluteGroupPath(), "mod", "edit", "-replace", fmt.Sprintf("%v=%v", s.GetProjectPackage(project), relativeProjectPath))
	if err != nil {
		return err
	}

	err = providers.GoExecute(project.GetAbsoluteGroupPath(), "get", s.GetProjectPackage(project))
	if err != nil {
		return err
	}

	return nil
}

func (s GoManager) RemoveFromGroup(project applicationproject.ProjectSpecification) error {
	err := providers.GoExecute(project.GetAbsoluteGroupPath(), "mod", "tidy")
	if err != nil {
		return err
	}

	return nil
}

func (s GoManager) CreateProjectFolder(project applicationproject.ProjectSpecification, paths ...string) (string, error) {
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

func (s GoManager) AddPackageToProject(project applicationproject.ProjectSpecification, packages []applicationproject.Package) error {

	for _, _package := range packages {
		err := providers.GoExecute(project.GetAbsoluteProjectPath(), "get", _package.GetFullName())
		if err != nil {
			return err
		}
	}

	return nil
}

func (s GoManager) RemovePackageFromProject(project applicationproject.ProjectSpecification, packages []applicationproject.Package) error {
	fmt.Printf("remove package not implemented yet")
	// for _, _package := range packages {
	// }

	return nil
}

func (s GoManager) AddReferenceToProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error {

	for _, reference := range references {

		relativeProjectPath, err := utils.FindRelativePath(project.GetAbsoluteProjectPath(), reference.GetAbsoluteProjectPath())
		if err != nil {
			return err
		}
		err = providers.GoExecute(project.GetAbsoluteProjectPath(), "mod", "edit", "-replace", fmt.Sprintf("%v=%v", s.GetProjectPackage(reference), relativeProjectPath))
		if err != nil {
			return err
		}

		err = providers.GoExecute(project.GetAbsoluteProjectPath(), "get", s.GetProjectPackage(reference))
		if err != nil {
			return err
		}
	}

	return nil
}

func (s GoManager) RemoveReferenceFromProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error {
	fmt.Printf("Remove reference not implemented yet")
	// for _, reference := range references {
	// }

	return nil
}
func (s GoManager) IsProjectFileExists(project applicationproject.ProjectSpecification) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.GetAbsoluteProjectPath(), s.GetProjectFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s GoManager) GetProjectFileName(project applicationproject.ProjectSpecification) string {
	return fmt.Sprintf("go.mod")
}

func (s GoManager) IsGroupFileExists(project applicationproject.ProjectSpecification) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.GetAbsoluteGroupPath(), s.GetGroupFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s GoManager) GetGroupFileName(project applicationproject.ProjectSpecification) string {
	return fmt.Sprintf("go.mod")
}
