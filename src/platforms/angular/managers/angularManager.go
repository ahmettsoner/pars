package managers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	_string "parsdevkit.net/pkg/utilities/string"

	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/models"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"

	angularModels "parsdevkit.net/platforms/angular/models"
	"parsdevkit.net/platforms/core"

	"parsdevkit.net/providers"

	"parsdevkit.net/pkg/utilities/file"

	"github.com/sirupsen/logrus"
)

type AngularManager struct {
	core.BaseManager
}

func NewAngularManager() core.ApplicationPlatformManagerInterface {
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
func (s AngularManager) GetKey() models.PlatformType {
	return models.PlatformTypes.Angular
}
func (s AngularManager) GetPlatformVersion(platform application_project_payload_structs.Platform) angularModels.AngularPlatformVersion {
	if _string.IsEmpty(platform.Version) {
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

func (s AngularManager) CreateProject(project application_project_payload_structs.ProjectBaseStruct) error {
	_, err := ProjectTypeToAngularCLITypeString(angularModels.AngularProjectType(project.Application.ProjectType))
	if err != nil {
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
		//TODO: Burda path deesteği getirmek için, project.Specifications.Name bilgisi, project.Specifications.Path ile birleştirilmeli
		err = providers.NGExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetAbsoluteGroupPath(), "generate", "application", "--name", project.Specifications.Name)
		if err != nil {
			return err
		}
	} else {
		err := providers.NGExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetAbsoluteGroupPath(), "new", "--name", project.Specifications.Name, "--skip-install", "true", "--skip-git", "true", "--skip-tests", "true", "--routing", "--new-project-root", "")
		if err != nil {
			return err
		}
	}

	return nil
}

func (s AngularManager) RemoveProject(project application_project_payload_structs.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			logrus.Debugf("Project removing from group (%v).", project.Specifications.Group)
			err := s.RemoveFromGroup(project)
			if err != nil {
				log.Fatal(err)
			}
			return nil
		}
	}
	return nil
}

func (s AngularManager) BuildProject(project application_project_payload_structs.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.NGExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetAbsoluteGroupPath(), "build", project.Specifications.Name)
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.NGExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetAbsoluteGroupPath(), "build")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s AngularManager) CleanProject(project application_project_payload_structs.ProjectBaseStruct) error {
	return nil
}

func (s AngularManager) InstallProject(project application_project_payload_structs.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.NPMExecute(project.Specifications.GetAbsoluteGroupPath(), "install", "--force")
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.NPMExecute(project.Specifications.GetAbsoluteGroupPath(), "install", "--force")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s AngularManager) TestProject(project application_project_payload_structs.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.NGExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetAbsoluteGroupPath(), "test", project.Specifications.Name)
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.NGExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetAbsoluteProjectPath(), "test")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s AngularManager) PackageProject(project application_project_payload_structs.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.NGExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetAbsoluteGroupPath(), "build", project.Specifications.Name)
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.NGExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetAbsoluteProjectPath(), "build")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s AngularManager) RunProject(project application_project_payload_structs.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.NGExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetAbsoluteGroupPath(), "serve", project.Specifications.Name, "--open")
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.NGExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetAbsoluteProjectPath(), "serve", "--open")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s AngularManager) CreateGroup(project application_project_payload_structs.ProjectBaseStruct) error {
	err := providers.NGExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "new", "--name", project.Specifications.GroupObject.Name, "--create-application", "false", "--skip-install", "true", "--skip-git", "true", "--skip-tests", "true", "--routing", "--new-project-root", "")
	if err != nil {
		return err
	}
	return nil
}

func (s AngularManager) DeleteGroup(project application_project_payload_structs.ProjectBaseStruct) {
	fmt.Println("Please remove group manually")
}

func (s AngularManager) AddToGroup(project application_project_payload_structs.ProjectBaseStruct) error {

	return nil
}

func (s AngularManager) RemoveFromGroup(project application_project_payload_structs.ProjectBaseStruct) error {
	// err := providers.NGExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "new", "--name", project.Specifications.GroupObject.Name, "--create-application", "false", "--skip-install", "true", "--skip-git", "true", "--skip-tests", "true", "--routing", "--new-project-root", "")
	// if err != nil {
	// 	return err
	// }
	fmt.Println("Please remove application manually")
	return nil
}

func (s AngularManager) CreateProjectFolder(project application_project_payload_structs.ProjectBaseStruct, paths ...string) (string, error) {
	var folders []string
	folders = append(folders, project.Specifications.GetAbsoluteProjectPath())
	for _, path := range paths {
		folders = append(folders, path)
	}
	foldePath := filepath.Join(folders...)

	if err := os.RemoveAll(foldePath); err != nil {
		return "", err
	}

	return foldePath, nil
}

func (s AngularManager) AddDependenciesToProject(project application_project_payload_structs.ProjectBaseStruct, dependencies []applicationProject.Dependency) error {
	for _, _package := range dependencies {

		packageName := _package.Name

		if !_string.IsEmpty(_package.Version) {
			packageName = fmt.Sprintf("%s@%s", packageName, _package.Version)
		}

		err := providers.NPMExecute(project.Specifications.GetAbsoluteProjectPath(), "install", packageName)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s AngularManager) RemoveDependenciesFromProject(project application_project_payload_structs.ProjectBaseStruct, dependencies []applicationProject.Dependency) error {

	for _, _package := range dependencies {

		err := providers.NPMExecute(project.Specifications.GetAbsoluteProjectPath(), "uninstall", _package.Name)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s AngularManager) AddReferenceToProject(project application_project_payload_structs.ProjectBaseStruct, references []application_project_payload_structs.ProjectBaseStruct) error {

	for _, reference := range references {
		relativePath, err := file.FindRelativePath(project.Specifications.GetAbsoluteProjectPath(), reference.Specifications.GetAbsoluteProjectPath())
		if err != nil {
			return err
		}

		// err = providers.NPMExecute(project.Specifications.GetAbsoluteProjectPath(), "link", relativePath)
		err = providers.NPMExecute(project.Specifications.GetAbsoluteProjectPath(), "install", "file:"+relativePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s AngularManager) RemoveReferenceFromProject(project application_project_payload_structs.ProjectBaseStruct, references []application_project_payload_structs.ProjectBaseStruct) error {

	for _, reference := range references {

		relativePath, err := file.FindRelativePath(project.Specifications.GetAbsoluteProjectPath(), reference.Specifications.GetAbsoluteProjectPath())
		if err != nil {
			return err
		}

		// err = providers.NPMExecute(project.Specifications.GetAbsoluteProjectPath(), "unlink", relativePath)
		err = providers.NPMExecute(project.Specifications.GetAbsoluteProjectPath(), "uninstall", "file:"+relativePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s AngularManager) IsProjectFileExists(project application_project_payload_structs.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.Specifications.GetAbsoluteProjectPath(), s.GetProjectFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s AngularManager) GetProjectFileName(project application_project_payload_structs.ProjectBaseStruct) string {
	return fmt.Sprintf("tsconfig.app.json")
}

func (s AngularManager) IsGroupFileExists(project application_project_payload_structs.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.Specifications.GetAbsoluteGroupPath(), s.GetGroupFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s AngularManager) GetGroupFileName(project application_project_payload_structs.ProjectBaseStruct) string {
	return fmt.Sprintf("package.json")
}

func (s AngularManager) ListLayersFromProject(project application_project_payload_structs.ProjectBaseStruct) ([]applicationProject.Layer, error) {

	folders, err := s.ListFoldersFromProjectDefinition(project)
	if err != nil {
		return nil, err
	}

	layers := make([]applicationProject.Layer, 0)
	for _, folder := range folders {
		for _, projectLayer := range project.Specifications.Layers {

			if filepath.Join(folder) == filepath.Join(projectLayer.Path) {
				layers = append(layers, projectLayer)
				break
			}

		}
	}

	return layers, nil
}

func (s AngularManager) HasLayerOnProject(project application_project_payload_structs.ProjectBaseStruct, layer string) (bool, error) {

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

func (s AngularManager) RemoveDefaultFiles(project application_project_payload_structs.ProjectBaseStruct) error {
	var paths []string = []string{}
	// projectPath := project.Specifications.GetAbsoluteProjectPath()

	var projectType models.ProjectType = models.ProjectType(project.Application.ProjectType)
	if projectType == models.ProjectTypes.Library {
	} else if projectType == models.ProjectTypes.SPA {
	}

	return s.FileRemover(paths...)
}

func (s AngularManager) AddFolderToProjectDefinition(project application_project_payload_structs.ProjectBaseStruct, paths ...string) error {
	return nil
}
func (s AngularManager) RemoveFolderFromProjectDefinition(project application_project_payload_structs.ProjectBaseStruct, paths ...string) error {
	return nil
}

func (s AngularManager) GetProjectFileRelativePath(project application_project_payload_structs.ProjectBaseStruct) string {

	return filepath.Join(project.Specifications.GetRelativeProjectPath(), s.GetProjectFileName(project))
}

func (s AngularManager) HasReferenceOnProject(project application_project_payload_structs.ProjectBaseStruct, reference application_project_payload_structs.ProjectBaseStruct) (bool, error) {

	references, err := s.ListReferencesFromProject(project)
	if err != nil {
		return false, err
	}

	referenceState := false

	for _, projectReference := range references {

		if projectReference.Specifications.Name == reference.Specifications.Name && projectReference.Specifications.GetAbsoluteProjectPath() == reference.Specifications.GetAbsoluteProjectPath() {
			referenceState = true
			break
		}
	}

	return referenceState, nil
}

func (s AngularManager) ListReferencesFromProject(project application_project_payload_structs.ProjectBaseStruct) ([]application_project_payload_structs.ProjectBaseStruct, error) {

	output, err := providers.NPMExecuteWithOutput(project.Specifications.GetAbsoluteProjectPath(), "list", "--link")
	if err != nil {
		return nil, err
	}

	pattern := regexp.MustCompile("([`└+]--|──)\\s+(.*?)@(.*?)\\s*->\\s*(.*?)\\n")

	matches := pattern.FindAllStringSubmatch(output, -1)

	references := make([]application_project_payload_structs.ProjectBaseStruct, 0)
	for _, match := range matches {
		for _, projectReference := range project.Specifications.References {
			relativeToReference, err := file.FindRelativePath(project.Specifications.GetAbsoluteProjectPath(), projectReference.Specifications.GetAbsoluteProjectPath())
			if err != nil {
				return nil, err
			}

			pathFromProjectSource := filepath.Clean(string(match[4]))
			pathFromStruct := filepath.Clean(relativeToReference)
			if pathFromProjectSource == pathFromStruct {
				references = append(references, projectReference)
				break
			}

		}
	}

	return references, nil
}

func (s AngularManager) IsProjectFolderExists(project application_project_payload_structs.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(project.Specifications.GetAbsoluteProjectPath()) // check if the project directory exists

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
