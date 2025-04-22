package managers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"parsdevkit.net/models"
	applicationproject "parsdevkit.net/structs/project/application-project"

	angularModels "parsdevkit.net/platforms/angular/models"
	"parsdevkit.net/platforms/core"

	"parsdevkit.net/providers"

	"parsdevkit.net/core/utils"

	"github.com/sirupsen/logrus"
)

type AngularManager struct {
	core.BaseManager
}

func NewAngularManager() AngularManager {
	return AngularManager{
		core.BaseManagerNew("/"),
	}
}

func ProjectTypeToAngularCLITypeString(c angularModels.AngularProjectType) (string, error) {
	switch c {
	case angularModels.AngularProjectTypes.Library:
		return "library", nil
	case angularModels.AngularProjectTypes.SPA:
		return "spa", nil
	default:
		return "", fmt.Errorf("error: %v is not defined for %v", c, angularModels.AngularProjectTypes)
	}
}

func (s AngularManager) GetPlatformVersion(platform applicationproject.Platform) angularModels.AngularPlatformVersion {
	if utils.IsEmpty(platform.Version) {
		platformVersion := angularModels.AngularPlatformVersions.V17

		return platformVersion
	} else {
		platformVersion, err := angularModels.AngularPlatformVersionEnumFromString(platform.Version)
		if err != nil {
			panic(err)
		}
		return platformVersion
	}
}

func (s AngularManager) CreateProject(project applicationproject.ProjectSpecification) error {
	_, err := ProjectTypeToAngularCLITypeString(angularModels.AngularProjectType(project.ProjectType))
	if err != nil {
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
		//TODO: Burda path deesteği getirmek için, project.Name bilgisi, project.Path ile birleştirilmeli
		err = providers.NGExecute(string(s.GetPlatformVersion(project.Platform)), project.GetAbsoluteGroupPath(), "generate", "application", "--name", project.Name)
		if err != nil {
			return err
		}
	} else {
		err := providers.NGExecute(string(s.GetPlatformVersion(project.Platform)), project.GetAbsoluteGroupPath(), "new", "--name", project.Name, "--skip-install", "true", "--skip-git", "true", "--skip-tests", "true", "--routing", "--new-project-root", "")
		if err != nil {
			return err
		}
	}

	return nil
}

func (s AngularManager) RemoveProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			logrus.Debugf("Project removing from group (%v).", project.Group)
			err := s.RemoveFromGroup(project)
			if err != nil {
				log.Fatal(err)
			}
			return nil
		}
	}
	return nil
}

func (s AngularManager) BuildProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.NGExecute(string(s.GetPlatformVersion(project.Platform)), project.GetAbsoluteGroupPath(), "build", project.Name)
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.NGExecute(string(s.GetPlatformVersion(project.Platform)), project.GetAbsoluteGroupPath(), "build")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s AngularManager) CleanProject(project applicationproject.ProjectSpecification) error {
	return nil
}

func (s AngularManager) InstallProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.NPMExecute(project.GetAbsoluteGroupPath(), "install", "--force")
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.NPMExecute(project.GetAbsoluteGroupPath(), "install", "--force")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s AngularManager) TestProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.NGExecute(string(s.GetPlatformVersion(project.Platform)), project.GetAbsoluteGroupPath(), "test", project.Name)
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.NGExecute(string(s.GetPlatformVersion(project.Platform)), project.GetAbsoluteProjectPath(), "test")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s AngularManager) PackageProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.NGExecute(string(s.GetPlatformVersion(project.Platform)), project.GetAbsoluteGroupPath(), "build", project.Name)
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.NGExecute(string(s.GetPlatformVersion(project.Platform)), project.GetAbsoluteProjectPath(), "build")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s AngularManager) RunProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.NGExecute(string(s.GetPlatformVersion(project.Platform)), project.GetAbsoluteGroupPath(), "serve", project.Name, "--open")
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.NGExecute(string(s.GetPlatformVersion(project.Platform)), project.GetAbsoluteProjectPath(), "serve", "--open")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s AngularManager) CreateGroup(project applicationproject.ProjectSpecification) error {
	err := providers.NGExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "new", "--name", project.GroupObject.Name, "--create-application", "false", "--skip-install", "true", "--skip-git", "true", "--skip-tests", "true", "--routing", "--new-project-root", "")
	if err != nil {
		return err
	}
	return nil
}

func (s AngularManager) DeleteGroup(project applicationproject.ProjectSpecification) {
	fmt.Println("Please remove group manually")
}

func (s AngularManager) AddToGroup(project applicationproject.ProjectSpecification) error {

	return nil
}

func (s AngularManager) RemoveFromGroup(project applicationproject.ProjectSpecification) error {
	// err := providers.NGExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "new", "--name", project.GroupObject.Name, "--create-application", "false", "--skip-install", "true", "--skip-git", "true", "--skip-tests", "true", "--routing", "--new-project-root", "")
	// if err != nil {
	// 	return err
	// }
	fmt.Println("Please remove application manually")
	return nil
}

func (s AngularManager) CreateProjectFolder(project applicationproject.ProjectSpecification, paths ...string) (string, error) {
	var folders []string
	folders = append(folders, project.GetAbsoluteProjectPath())
	for _, path := range paths {
		folders = append(folders, path)
	}
	foldePath := filepath.Join(folders...)

	if err := os.RemoveAll(foldePath); err != nil {
		return "", err
	}

	return foldePath, nil
}

func (s AngularManager) AddPackageToProject(project applicationproject.ProjectSpecification, packages []applicationproject.Package) error {
	for _, _package := range packages {

		packageName := _package.Name

		if !utils.IsEmpty(_package.Version) {
			packageName = fmt.Sprintf("%s@%s", packageName, _package.Version)
		}

		err := providers.NPMExecute(project.GetAbsoluteProjectPath(), "install", packageName)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s AngularManager) RemovePackageFromProject(project applicationproject.ProjectSpecification, packages []applicationproject.Package) error {

	for _, _package := range packages {

		err := providers.NPMExecute(project.GetAbsoluteProjectPath(), "uninstall", _package.Name)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s AngularManager) AddReferenceToProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error {

	for _, reference := range references {
		relativePath, err := utils.FindRelativePath(project.GetAbsoluteProjectPath(), reference.GetAbsoluteProjectPath())
		if err != nil {
			return err
		}

		err = providers.NPMExecute(project.GetAbsoluteProjectPath(), "link", relativePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s AngularManager) RemoveReferenceFromProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error {

	for _, reference := range references {

		relativePath, err := utils.FindRelativePath(project.GetAbsoluteProjectPath(), reference.GetAbsoluteProjectPath())
		if err != nil {
			return err
		}

		err = providers.NPMExecute(project.GetAbsoluteProjectPath(), "unlink", relativePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s AngularManager) IsProjectFileExists(project applicationproject.ProjectSpecification) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.GetAbsoluteProjectPath(), s.GetProjectFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s AngularManager) GetProjectFileName(project applicationproject.ProjectSpecification) string {
	return fmt.Sprintf("tsconfig.app.json")
}

func (s AngularManager) IsGroupFileExists(project applicationproject.ProjectSpecification) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.GetAbsoluteGroupPath(), s.GetGroupFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s AngularManager) GetGroupFileName(project applicationproject.ProjectSpecification) string {
	return fmt.Sprintf("package.json")
}

func (s AngularManager) ListLayersFromProject(projectSpecification applicationproject.ProjectSpecification) ([]applicationproject.Layer, error) {

	folders, err := s.ListFoldersFromProjectDefinition(projectSpecification)
	if err != nil {
		return nil, err
	}

	layers := make([]applicationproject.Layer, 0)
	for _, folder := range folders {
		for _, projectLayer := range projectSpecification.Configuration.Layers {

			if filepath.Join(folder) == filepath.Join(projectLayer.Path) {
				layers = append(layers, projectLayer)
				break
			}

		}
	}

	return layers, nil
}

func (s AngularManager) HasLayerOnProject(project applicationproject.ProjectSpecification, layer string) (bool, error) {

	layers, err := s.ListLayersFromProject(project)
	if err != nil {
		return false, err
	}

	layerState := false

	for _, projectLayer := range layers {
		if projectLayer.Name == layer {
			layerState = true
			break
		}
	}

	return layerState, nil
}

func (s AngularManager) RemoveDefaultFiles(project applicationproject.ProjectSpecification) error {
	var paths []string = []string{}
	// projectPath := project.GetAbsoluteProjectPath()

	var projectType models.ProjectType = models.ProjectType(project.ProjectType)
	if projectType == models.ProjectTypes.Library {
	} else if projectType == models.ProjectTypes.SPA {
	}

	return s.FileRemover(paths...)
}

func (s AngularManager) AddFolderToProjectDefinition(project applicationproject.ProjectSpecification, paths ...string) {

}
func (s AngularManager) RemoveFolderFromProjectDefinition(project applicationproject.ProjectSpecification, paths ...string) {

}

func (s AngularManager) GetProjectFileRelativePath(project applicationproject.ProjectSpecification) string {

	return filepath.Join(project.GetRelativeProjectPath(), s.GetProjectFileName(project))
}

func (s AngularManager) HasReferenceOnProject(project applicationproject.ProjectSpecification, reference applicationproject.ProjectSpecification) (bool, error) {

	references, err := s.ListReferencesFromProject(project)
	if err != nil {
		return false, err
	}

	referenceState := false

	for _, projectReference := range references {

		if projectReference.Name == reference.Name && projectReference.GetAbsoluteProjectPath() == reference.GetAbsoluteProjectPath() {
			referenceState = true
			break
		}
	}

	return referenceState, nil
}

func (s AngularManager) ListReferencesFromProject(projectSpecification applicationproject.ProjectSpecification) ([]applicationproject.ProjectSpecification, error) {

	output, err := providers.NPMExecuteWithOutput(projectSpecification.GetAbsoluteProjectPath(), "list", "--link")
	if err != nil {
		return nil, err
	}

	pattern := regexp.MustCompile("([`└+]--|──)\\s+(.*?)@(.*?)extraneous\\s*->\\s*(.*?)\\n")

	matches := pattern.FindAllStringSubmatch(output, -1)

	references := make([]applicationproject.ProjectSpecification, 0)
	for _, match := range matches {
		for _, projectReference := range projectSpecification.Configuration.References {
			relativeToReference, err := utils.FindRelativePath(projectSpecification.GetAbsoluteProjectPath(), projectReference.Specifications.GetAbsoluteProjectPath())
			if err != nil {
				return nil, err
			}

			pathFromProjectSource := filepath.Clean(string(match[4]))
			pathFromStruct := filepath.Clean(relativeToReference)
			if pathFromProjectSource == pathFromStruct {
				references = append(references, projectReference.Specifications)
				break
			}

		}
	}

	return references, nil
}

func (s AngularManager) IsProjectFolderExists(project applicationproject.ProjectSpecification) (bool, error) {

	stat, err := os.Stat(project.GetAbsoluteProjectPath()) // check if the project directory exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return stat.IsDir(), nil
	}
}

func (s AngularManager) NormalizeText(input string) string {
	input = strings.ToLower(input)

	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	input = reg.ReplaceAllString(input, "-")

	input = strings.Trim(input, "-")

	return input
}
