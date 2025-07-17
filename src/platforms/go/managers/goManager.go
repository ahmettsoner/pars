package managers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_string "parsdevkit.net/pkg/utilities/string"

	applicationProject "parsdevkit.net/application/structs/project"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	"parsdevkit.net/pkg/utilities/file"

	"parsdevkit.net/models"
	"parsdevkit.net/platforms/core"
	goModels "parsdevkit.net/platforms/go/models"

	"parsdevkit.net/providers"
)

type GoManager struct {
	core.BaseManager
}

func NewGoManager() core.ApplicationPlatformManagerInterface {
	return GoManager{
		core.BaseManagerNew("/")}
}
func (s GoManager) GetKey() models.PlatformType {
	return models.PlatformTypes.GO
}

func (s GoManager) GetPlatformVersion(platform applicationProject.Platform) goModels.GoPlatformVersion {
	if _string.IsEmpty(platform.Version) {
		platformVersion := goModels.GoPlatformVersions.Go121

		return platformVersion
	} else {
		platformVersion, err := goModels.GoPlatformVersionEnumFromString(platform.Version)
		if err != nil {
			panic(err)
		}
		return platformVersion
	}
}
func (s GoManager) CreateProject(project application_project_payload_structs.ProjectBaseStruct) error {

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

	err := providers.GoExecute(project.Specifications.GetAbsoluteProjectPath(), "mod", "init", s.GetProjectPackage(project.Specifications))
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

	// if project.Specifications.Dependencies != nil {
	// 	s.AddDependenciesToProject(project, project.Specifications.Dependencies)
	// }

	// if project.Specifications.References != nil {
	// 	s.AddReferenceToProject(project, project.Specifications.References)
	// }

	return nil
}

func (s GoManager) RemoveProject(project application_project_payload_structs.ProjectBaseStruct) error {
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

func (s GoManager) BuildProject(project application_project_payload_structs.ProjectBaseStruct) error {
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

func (s GoManager) CleanProject(project application_project_payload_structs.ProjectBaseStruct) error {
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

func (s GoManager) InstallProject(project application_project_payload_structs.ProjectBaseStruct) error {
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

func (s GoManager) TestProject(project application_project_payload_structs.ProjectBaseStruct) error {
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

func (s GoManager) PackageProject(project application_project_payload_structs.ProjectBaseStruct) error {
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

func (s GoManager) RunProject(project application_project_payload_structs.ProjectBaseStruct) error {
	if !_string.IsEmpty(project.Specifications.Group) {
		groupStatus, err := s.IsGroupFileExists(project)
		if err != nil {
			return err

		}

		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		}
	}

	err := providers.GoExecute(project.Specifications.GetAbsoluteGroupPath(), "run", s.GetProjectPackage(project.Specifications))
	if err != nil {
		return err
	}
	return nil
}

func (s GoManager) CreateGroup(project application_project_payload_structs.ProjectBaseStruct) error {

	err := providers.GoExecute(project.Specifications.GetAbsoluteGroupPath(), "mod", "init", s.GetGroupPackage(project))
	if err != nil {
		return err
	}

	return nil
}

func (s GoManager) DeleteGroup(project application_project_payload_structs.ProjectBaseStruct) {
}

func (s GoManager) AddToGroup(project application_project_payload_structs.ProjectBaseStruct) error {

	relativeProjectPath, err := file.FindRelativePath(project.Specifications.GetAbsoluteGroupPath(), project.Specifications.GetAbsoluteProjectPath())
	if err != nil {
		return err
	}
	err = providers.GoExecute(project.Specifications.GetAbsoluteGroupPath(), "mod", "edit", "-replace", fmt.Sprintf("%v=%v", s.GetProjectPackage(project.Specifications), relativeProjectPath))
	if err != nil {
		return err
	}

	err = providers.GoExecute(project.Specifications.GetAbsoluteGroupPath(), "get", s.GetProjectPackage(project.Specifications))
	if err != nil {
		return err
	}

	return nil
}

func (s GoManager) RemoveFromGroup(project application_project_payload_structs.ProjectBaseStruct) error {
	err := providers.GoExecute(project.Specifications.GetAbsoluteGroupPath(), "mod", "tidy")
	if err != nil {
		return err
	}

	return nil
}

func (s GoManager) CreateProjectFolder(project application_project_payload_structs.ProjectBaseStruct, paths ...string) (string, error) {
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

func (s GoManager) AddDependenciesToProject(project application_project_payload_structs.ProjectBaseStruct, dependencies []applicationProject.Dependency) error {

	for _, _package := range dependencies {
		err := providers.GoExecute(project.Specifications.GetAbsoluteProjectPath(), "get", _package.GetFullName())
		if err != nil {
			return err
		}
	}

	return nil
}

func (s GoManager) RemoveDependenciesFromProject(project application_project_payload_structs.ProjectBaseStruct, dependencies []applicationProject.Dependency) error {
	fmt.Printf("remove package not implemented yet")
	// for _, _package := range dependencies {
	// }

	return nil
}

func (s GoManager) AddReferenceToProject(project application_project_payload_structs.ProjectBaseStruct, references []applicationProject.Reference) error {

	for _, reference := range references {

		relativeProjectPath, err := file.FindRelativePath(project.Specifications.GetAbsoluteProjectPath(), reference.Specifications.GetAbsoluteProjectPath())
		if err != nil {
			return err
		}
		err = providers.GoExecute(project.Specifications.GetAbsoluteProjectPath(), "mod", "edit", "-replace", fmt.Sprintf("%v=%v", s.GetProjectPackage(reference.Specifications), relativeProjectPath))
		if err != nil {
			return err
		}

		err = providers.GoExecute(project.Specifications.GetAbsoluteProjectPath(), "get", s.GetProjectPackage(reference.Specifications))
		if err != nil {
			return err
		}
	}

	return nil
}

func (s GoManager) RemoveReferenceFromProject(project application_project_payload_structs.ProjectBaseStruct, references []applicationProject.Reference) error {
	fmt.Printf("Remove reference not implemented yet")
	// for _, reference := range references {
	// }

	return nil
}
func (s GoManager) IsProjectFileExists(project application_project_payload_structs.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.Specifications.GetAbsoluteProjectPath(), s.GetProjectFileName(project.Specifications))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s GoManager) GetProjectFileName(specifications applicationProject.ProjectSpecification) string {
	return fmt.Sprintf("go.mod")
}

func (s GoManager) IsGroupFileExists(project application_project_payload_structs.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.Specifications.GetAbsoluteGroupPath(), s.GetGroupFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s GoManager) GetGroupFileName(project application_project_payload_structs.ProjectBaseStruct) string {
	return fmt.Sprintf("go.mod")
}
