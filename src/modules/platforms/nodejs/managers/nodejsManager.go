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
	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/platforms/core"
	nodejsModels "parsdevkit.net/platforms/nodejs/models"

	"parsdevkit.net/providers"

	"parsdevkit.net/core/utils"

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
	if utils.IsEmpty(platform.Version) {
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

func (s NodeJSManager) CreateProject(project applicationproject.ProjectSpecification) error {
	if utils.IsEmpty(project.Group) {
		if len(project.Package) > 0 {
			err := providers.NPMExecute(project.GetAbsoluteProjectPath(), "init", "--scope", fmt.Sprintf("@%v", project.Package[len(project.Package)-1]), "--yes")
			if err != nil {
				return err
			}
		} else {
			err := providers.NPMExecute(project.GetAbsoluteProjectPath(), "init", "--yes")
			if err != nil {
				return err
			}
		}
	} else {
		if len(project.Package) > 0 {
			err := providers.NPMExecute(project.GetAbsoluteGroupPath(), "init", "--scope", fmt.Sprintf("@%v", project.Package[len(project.Package)-1]), "--yes", "--workspace", filepath.Join(project.GetProjectPath()))
			if err != nil {
				return err
			}
		} else {
			err := providers.NPMExecute(project.GetAbsoluteGroupPath(), "init", "--yes", "--workspace", filepath.Join(project.GetProjectPath()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s NodeJSManager) RemoveProject(project applicationproject.ProjectSpecification) error {
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

func (s NodeJSManager) BuildProject(project applicationproject.ProjectSpecification) error {
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

func (s NodeJSManager) CleanProject(project applicationproject.ProjectSpecification) error {
	return nil
}

func (s NodeJSManager) InstallProject(project applicationproject.ProjectSpecification) error {
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

func (s NodeJSManager) TestProject(project applicationproject.ProjectSpecification) error {
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

func (s NodeJSManager) PackageProject(project applicationproject.ProjectSpecification) error {
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

func (s NodeJSManager) RunProject(project applicationproject.ProjectSpecification) error {
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

func (s NodeJSManager) CreateGroup(project applicationproject.ProjectSpecification) error {
	if !utils.IsEmpty(project.Group) {
		if len(project.GroupObject.Package) > 0 {
			err := providers.NPMExecute(project.GetAbsoluteGroupPath(), "init", "--scope", fmt.Sprintf("@%v", project.GroupObject.Package[len(project.GroupObject.Package)-1]), "--yes")
			if err != nil {
				return err
			}
		} else {
			err := providers.NPMExecute(project.GetAbsoluteGroupPath(), "init", "--yes")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s NodeJSManager) DeleteGroup(project applicationproject.ProjectSpecification) {
	fmt.Println("Please remove group manually")
}

func (s NodeJSManager) AddToGroup(project applicationproject.ProjectSpecification) error {

	return nil
}

func (s NodeJSManager) RemoveFromGroup(project applicationproject.ProjectSpecification) error {
	// err := providers.NGExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "new", "--name", project.GroupObject.Name, "--create-application", "false", "--skip-install", "true", "--skip-git", "true", "--skip-tests", "true", "--routing", "--new-project-root", "")
	// if err != nil {
	// 	return err
	// }
	fmt.Println("Please remove application manually")
	return nil
}

func (s NodeJSManager) CreateProjectFolder(project applicationproject.ProjectSpecification, paths ...string) (string, error) {
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

func (s NodeJSManager) AddPackageToProject(project applicationproject.ProjectSpecification, packages []applicationproject.Package) error {
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

func (s NodeJSManager) RemovePackageFromProject(project applicationproject.ProjectSpecification, packages []applicationproject.Package) error {
	for _, _package := range packages {

		err := providers.NPMExecute(project.GetAbsoluteProjectPath(), "uninstall", _package.Name)

		if err != nil {
			return err
		}
	}
	return nil
}

func (s NodeJSManager) AddReferenceToProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error {

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

func (s NodeJSManager) RemoveReferenceFromProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error {

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

func (s NodeJSManager) IsProjectFileExists(project applicationproject.ProjectSpecification) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.GetAbsoluteProjectPath(), s.GetProjectFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s NodeJSManager) GetProjectFileName(project applicationproject.ProjectSpecification) string {
	return "package.json"
}

func (s NodeJSManager) IsGroupFileExists(project applicationproject.ProjectSpecification) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.GetAbsoluteGroupPath(), s.GetGroupFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s NodeJSManager) GetGroupFileName(project applicationproject.ProjectSpecification) string {
	return "package.json"
}

func (s NodeJSManager) ListLayersFromProject(projectSpecification applicationproject.ProjectSpecification) ([]applicationproject.Layer, error) {

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

func (s NodeJSManager) HasLayerOnProject(project applicationproject.ProjectSpecification, layer string) (bool, error) {

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

func (s NodeJSManager) RemoveDefaultFiles(project applicationproject.ProjectSpecification) error {
	var paths []string = []string{}
	// projectPath := project.GetAbsoluteProjectPath()

	var projectType models.ProjectType = models.ProjectType(project.ProjectType)
	if projectType == models.ProjectTypes.Library {
	} else if projectType == models.ProjectTypes.SPA {
	}

	return s.FileRemover(paths...)
}

func (s NodeJSManager) AddFolderToProjectDefinition(project applicationproject.ProjectSpecification, paths ...string) {

}
func (s NodeJSManager) RemoveFolderFromProjectDefinition(project applicationproject.ProjectSpecification, paths ...string) {

}

func (s NodeJSManager) GetProjectFileRelativePath(project applicationproject.ProjectSpecification) string {

	return filepath.Join(project.GetRelativeProjectPath(), s.GetProjectFileName(project))
}

func (s NodeJSManager) HasReferenceOnProject(project applicationproject.ProjectSpecification, reference applicationproject.ProjectSpecification) (bool, error) {

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

func (s NodeJSManager) ListReferencesFromProject(projectSpecification applicationproject.ProjectSpecification) ([]applicationproject.ProjectSpecification, error) {

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

func (s NodeJSManager) IsProjectFolderExists(project applicationproject.ProjectSpecification) (bool, error) {

	stat, err := os.Stat(project.GetAbsoluteProjectPath()) // check if the project directory exists

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

func (s NodeJSManager) PrintPackage(packages []string) string {
	var nonEmptyPackages []string
	for _, pkg := range packages {
		if pkg != "" {
			nonEmptyPackages = append(nonEmptyPackages, pkg)
		}
	}
	return strings.Join(nonEmptyPackages, ".")
}
func (s NodeJSManager) PrintDataType(dataType objectresource.DataType) string {
	result := ""
	if dataType.Category == objectresource.DataTypeCategories.Value {
		switch dataType.Name {
		case string(objectresource.ValueTypes.ShortInt):
			result = "number"
		case string(objectresource.ValueTypes.Int):
			result = "number"
		case string(objectresource.ValueTypes.LongInt):
			result = "bigint"
		case string(objectresource.ValueTypes.Double):
			result = "number"
		case string(objectresource.ValueTypes.Decimal):
			result = "number"
		case string(objectresource.ValueTypes.Float):
			result = "number"
		case string(objectresource.ValueTypes.Char):
			result = "string"
		case string(objectresource.ValueTypes.String):
			result = "string"
		case string(objectresource.ValueTypes.Date):
			result = "string"
		case string(objectresource.ValueTypes.DateTime):
			result = "string"
		case string(objectresource.ValueTypes.Time):
			result = "string"
		case string(objectresource.ValueTypes.Boolean):
			result = "boolean"
		case string(objectresource.ValueTypes.Byte):
			result = "number"
		case string(objectresource.ValueTypes.ShortBlob):
			result = "number"
		case string(objectresource.ValueTypes.Blob):
			result = "number"
		case string(objectresource.ValueTypes.LongBlob):
			result = "number"
		default:
			return "Unknown"
		}
	} else if dataType.Category == objectresource.DataTypeCategories.Reference {
		if !utils.IsEmpty(dataType.Package.Alias) {
			result = fmt.Sprintf("%v.%v", dataType.Package.Alias, dataType.Name)
		} else {
			result = dataType.Name
		}
	} else if dataType.Category == objectresource.DataTypeCategories.Resource {
		if !utils.IsEmpty(dataType.Package.Alias) {
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

	if dataType.Modifier == objectresource.ModifierTypes.Array {
		result = fmt.Sprintf("%v[]", result)
	}

	return result
}

func (s NodeJSManager) PrintVisibility(visibility objectresource.VisibilityType) string {
	switch visibility {
	case objectresource.VisibilityTypeTypes.Public:
		return "public"
	case objectresource.VisibilityTypeTypes.Protected:
		return "protected"
	case objectresource.VisibilityTypeTypes.Private:
		return "private"
	default:
		return "Unknown"
	}
}
