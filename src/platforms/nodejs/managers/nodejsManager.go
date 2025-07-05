package managers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"parsdevkit.net/application/structs"

	_string "parsdevkit.net/pkg/utilities/string"

	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/models"
	applicationproject "parsdevkit.net/modules/project/application_project_payload"

	"parsdevkit.net/platforms/core"
	nodejsModels "parsdevkit.net/platforms/nodejs/models"

	"parsdevkit.net/providers"

	"parsdevkit.net/pkg/utilities/file"

	"github.com/sirupsen/logrus"
)

type NodeJSManager struct {
	core.BaseManager
}

func NewNodeJSManager() NodeJSManager {
	return NodeJSManager{
		core.BaseManagerNew("/"),
	}
}

func ProjectTypeToNodeJSCLITypeString(c nodejsModels.NodeJSProjectType) (string, error) {
	switch c {
	case nodejsModels.NodeJSProjectTypes.Library:
		return "library", nil
	default:
		return "", fmt.Errorf("error: %v is not defined for %v", c, nodejsModels.NodeJSProjectTypes)
	}
}

func (s NodeJSManager) GetPlatformVersion(platform applicationproject.Platform) nodejsModels.NodeJSPlatformVersion {
	if _string.IsEmpty(platform.Version) {
		platformVersion := nodejsModels.NodeJSPlatformVersions.V17

		return platformVersion
	} else {
		platformVersion, err := nodejsModels.NodeJSPlatformVersionEnumFromString(platform.Version)
		if err != nil {
			panic(err)
		}
		return platformVersion
	}
}

func (s NodeJSManager) CreateProject(project applicationproject.ProjectBaseStruct) error {
	if _string.IsEmpty(project.Specifications.Group) {
		if len(project.Specifications.Package) > 0 {
			err := providers.NPMExecute(project.Specifications.GetAbsoluteProjectPath(), "init", "--scope", fmt.Sprintf("@%v", project.Specifications.Package[len(project.Specifications.Package)-1]), "--yes")
			if err != nil {
				return err
			}
		} else {
			err := providers.NPMExecute(project.Specifications.GetAbsoluteProjectPath(), "init", "--yes")
			if err != nil {
				return err
			}
		}
	} else {
		if len(project.Specifications.Package) > 0 {
			err := providers.NPMExecute(project.Specifications.GetAbsoluteGroupPath(), "init", "--scope", fmt.Sprintf("@%v", project.Specifications.Package[len(project.Specifications.Package)-1]), "--yes", "--workspace", filepath.Join(project.Specifications.GetProjectPath()))
			if err != nil {
				return err
			}
		} else {
			err := providers.NPMExecute(project.Specifications.GetAbsoluteGroupPath(), "init", "--yes", "--workspace", filepath.Join(project.Specifications.GetProjectPath()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s NodeJSManager) RemoveProject(project applicationproject.ProjectBaseStruct) error {
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

func (s NodeJSManager) BuildProject(project applicationproject.ProjectBaseStruct) error {
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

func (s NodeJSManager) CleanProject(project applicationproject.ProjectBaseStruct) error {
	return nil
}

func (s NodeJSManager) InstallProject(project applicationproject.ProjectBaseStruct) error {
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

func (s NodeJSManager) TestProject(project applicationproject.ProjectBaseStruct) error {
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

func (s NodeJSManager) PackageProject(project applicationproject.ProjectBaseStruct) error {
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

func (s NodeJSManager) RunProject(project applicationproject.ProjectBaseStruct) error {
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

func (s NodeJSManager) CreateGroup(project applicationproject.ProjectBaseStruct) error {
	if !_string.IsEmpty(project.Specifications.Group) {
		if len(project.Specifications.GroupObject.Package) > 0 {
			err := providers.NPMExecute(project.Specifications.GetAbsoluteGroupPath(), "init", "--scope", fmt.Sprintf("@%v", project.Specifications.GroupObject.Package[len(project.Specifications.GroupObject.Package)-1]), "--yes")
			if err != nil {
				return err
			}
		} else {
			err := providers.NPMExecute(project.Specifications.GetAbsoluteGroupPath(), "init", "--yes")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s NodeJSManager) DeleteGroup(project applicationproject.ProjectBaseStruct) {
	fmt.Println("Please remove group manually")
}

func (s NodeJSManager) AddToGroup(project applicationproject.ProjectBaseStruct) error {

	return nil
}

func (s NodeJSManager) RemoveFromGroup(project applicationproject.ProjectBaseStruct) error {
	// err := providers.NGExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "new", "--name", project.Specifications.GroupObject.Name, "--create-application", "false", "--skip-install", "true", "--skip-git", "true", "--skip-tests", "true", "--routing", "--new-project-root", "")
	// if err != nil {
	// 	return err
	// }
	fmt.Println("Please remove application manually")
	return nil
}

func (s NodeJSManager) CreateProjectFolder(project applicationproject.ProjectBaseStruct, paths ...string) (string, error) {
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

func (s NodeJSManager) AddDependenciesToProject(project applicationproject.ProjectBaseStruct, dependencies []applicationProject.Package) error {
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

func (s NodeJSManager) RemoveDependenciesFromProject(project applicationproject.ProjectBaseStruct, dependencies []applicationProject.Package) error {
	for _, _package := range dependencies {

		err := providers.NPMExecute(project.Specifications.GetAbsoluteProjectPath(), "uninstall", _package.Name)

		if err != nil {
			return err
		}
	}
	return nil
}

func (s NodeJSManager) AddReferenceToProject(project applicationproject.ProjectBaseStruct, references []applicationproject.ProjectBaseStruct) error {

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

func (s NodeJSManager) RemoveReferenceFromProject(project applicationproject.ProjectBaseStruct, references []applicationproject.ProjectBaseStruct) error {

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

func (s NodeJSManager) IsProjectFileExists(project applicationproject.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.Specifications.GetAbsoluteProjectPath(), s.GetProjectFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s NodeJSManager) GetProjectFileName(project applicationproject.ProjectBaseStruct) string {
	return "package.json"
}

func (s NodeJSManager) IsGroupFileExists(project applicationproject.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.Specifications.GetAbsoluteGroupPath(), s.GetGroupFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s NodeJSManager) GetGroupFileName(project applicationproject.ProjectBaseStruct) string {
	return "package.json"
}

func (s NodeJSManager) ListLayersFromProject(project applicationproject.ProjectBaseStruct) ([]applicationProject.Layer, error) {

	folders, err := s.ListFoldersFromProjectDefinition(project)
	if err != nil {
		return nil, err
	}

	layers := make([]applicationProject.Layer, 0)
	for _, folder := range folders {
		for _, projectLayer := range project.Specifications.Configuration.Layers {

			if filepath.Join(folder) == filepath.Join(projectLayer.Path) {
				layers = append(layers, projectLayer)
				break
			}

		}
	}

	return layers, nil
}

func (s NodeJSManager) HasLayerOnProject(project applicationproject.ProjectBaseStruct, layer string) (bool, error) {

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

func (s NodeJSManager) RemoveDefaultFiles(project applicationproject.ProjectBaseStruct) error {
	var paths []string = []string{}
	// projectPath := project.Specifications.GetAbsoluteProjectPath()

	var projectType models.ProjectType = models.ProjectType(project.Specifications.ProjectType)
	if projectType == models.ProjectTypes.Library {
	} else if projectType == models.ProjectTypes.SPA {
	}

	return s.FileRemover(paths...)
}

func (s NodeJSManager) AddFolderToProjectDefinition(project applicationproject.ProjectBaseStruct, paths ...string) error {
	return nil
}
func (s NodeJSManager) RemoveFolderFromProjectDefinition(project applicationproject.ProjectBaseStruct, paths ...string) error {
	return nil
}

func (s NodeJSManager) GetProjectFileRelativePath(project applicationproject.ProjectBaseStruct) string {

	return filepath.Join(project.Specifications.GetRelativeProjectPath(), s.GetProjectFileName(project))
}

func (s NodeJSManager) HasReferenceOnProject(project applicationproject.ProjectBaseStruct, reference applicationproject.ProjectBaseStruct) (bool, error) {

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

func (s NodeJSManager) ListReferencesFromProject(project applicationproject.ProjectBaseStruct) ([]applicationproject.ProjectBaseStruct, error) {

	output, err := providers.NPMExecuteWithOutput(project.Specifications.GetAbsoluteProjectPath(), "list", "--link")
	if err != nil {
		return nil, err
	}

	pattern := regexp.MustCompile("([`└+]--|──)\\s+(.*?)@(.*?)\\s*->\\s*(.*?)\\n")

	matches := pattern.FindAllStringSubmatch(output, -1)

	references := make([]applicationproject.ProjectBaseStruct, 0)
	for _, match := range matches {
		for _, projectReference := range project.Specifications.Configuration.References {
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

func (s NodeJSManager) IsProjectFolderExists(project applicationproject.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(project.Specifications.GetAbsoluteProjectPath()) // check if the project directory exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return stat.IsDir(), nil
	}
}

func (s NodeJSManager) NormalizeText(input string) string {
	input = strings.ToLower(input)

	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	input = reg.ReplaceAllString(input, "-")

	input = strings.Trim(input, "-")

	return input
}

func (s NodeJSManager) PrintDependencies(dependencies []string) string {
	var nonEmptyPackages []string
	for _, pkg := range dependencies {
		if pkg != "" {
			nonEmptyPackages = append(nonEmptyPackages, pkg)
		}
	}
	return strings.Join(nonEmptyPackages, ".")
}
func (s NodeJSManager) PrintDataType(dataType structs.DataType) string {
	result := ""
	if dataType.Category == structs.DataTypeCategories.Value {
		switch dataType.Name {
		case string(structs.ValueTypes.ShortInt):
			result = "number"
		case string(structs.ValueTypes.Int):
			result = "number"
		case string(structs.ValueTypes.LongInt):
			result = "bigint"
		case string(structs.ValueTypes.Double):
			result = "number"
		case string(structs.ValueTypes.Decimal):
			result = "number"
		case string(structs.ValueTypes.Float):
			result = "number"
		case string(structs.ValueTypes.Char):
			result = "string"
		case string(structs.ValueTypes.String):
			result = "string"
		case string(structs.ValueTypes.Date):
			result = "string"
		case string(structs.ValueTypes.DateTime):
			result = "string"
		case string(structs.ValueTypes.Time):
			result = "string"
		case string(structs.ValueTypes.Boolean):
			result = "boolean"
		case string(structs.ValueTypes.Byte):
			result = "number"
		case string(structs.ValueTypes.ShortBlob):
			result = "number"
		case string(structs.ValueTypes.Blob):
			result = "number"
		case string(structs.ValueTypes.LongBlob):
			result = "number"
		default:
			return "Unknown"
		}
	} else if dataType.Category == structs.DataTypeCategories.Reference {
		if !_string.IsEmpty(dataType.Package.Alias) {
			result = fmt.Sprintf("%v.%v", dataType.Package.Alias, dataType.Name)
		} else {
			result = dataType.Name
		}
	} else if dataType.Category == structs.DataTypeCategories.Resource {
		if !_string.IsEmpty(dataType.Package.Alias) {
			result = fmt.Sprintf("%v.%v", dataType.Package.Alias, dataType.Name)
		} else {
			result = dataType.Name
		}
	}

	if dataType.Generics != nil && len(dataType.Generics) > 0 {
		generics := []string{}

		for _, genericType := range dataType.Generics {
			generic := s.PrintDataType(genericType)
			generics = append(generics, generic)
		}

		result = fmt.Sprintf("%v<%v>", result, strings.Join(generics, ", "))
	}

	if dataType.Modifier == structs.ModifierTypes.Array {
		result = fmt.Sprintf("%v[]", result)
	}

	return result
}

func (s NodeJSManager) PrintVisibility(visibility structs.VisibilityType) string {
	switch visibility {
	case structs.VisibilityTypeTypes.Public:
		return "public"
	case structs.VisibilityTypeTypes.Protected:
		return "protected"
	case structs.VisibilityTypeTypes.Private:
		return "private"
	default:
		return "Unknown"
	}
}
