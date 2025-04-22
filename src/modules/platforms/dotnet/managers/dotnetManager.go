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
	dotnetModels "parsdevkit.net/platforms/dotnet/models"

	"parsdevkit.net/providers"

	"parsdevkit.net/core/utils"

	mxj "github.com/clbanning/mxj/v2"
)

type DotnetManager struct {
	core.BaseManager
}

func NewDotnetManager() DotnetManager {
	return DotnetManager{
		core.BaseManagerNew(".")}
}

func ProjectTypeToDotnetCLITypeString(c models.ProjectType) (string, error) {
	switch c {
	case models.ProjectTypes.Library:
		return "classlib", nil
	case models.ProjectTypes.WebApi:
		return "webapi", nil
	case models.ProjectTypes.Console:
		return "console", nil
	case models.ProjectTypes.WebApp:
		return "webapp", nil
	default:
		return "", fmt.Errorf("error: %v is not defined for %v", c, models.ProjectTypes)
	}
}

func DotnetWebAppOptionToDotnetCLITypeString(c dotnetModels.DotnetWebAppOption) (string, error) {
	switch c {
	case dotnetModels.DotnetWebAppOptions.MVC:
		return "mvc", nil
	case dotnetModels.DotnetWebAppOptions.Razor:
		return "razor", nil
	default:
		return "", fmt.Errorf("error: %v is not defined for %v", c, dotnetModels.DotnetWebAppOptions)
	}
}

func (s DotnetManager) CreateProject(project applicationproject.ProjectSpecification) error {

	dotnetProjectType, err := ProjectTypeToDotnetCLITypeString(models.ProjectType(project.ProjectType))
	if err != nil {
		return err
	}
	// if dotnetProjectType == "webapp" {
	// 	dotnetProjectType, err = DotnetWebAppOptionToDotnetCLITypeString(dotnetModels.DotnetWebAppOption(project.GetConfiguration().Options))
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	err = providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "new", dotnetProjectType, "--name", project.Name, "--output", utils.PathWithDot(project.GetRelativeProjectPath()))
	if err != nil {
		return err
	}

	return nil
}

func (s DotnetManager) RemoveProject(project applicationproject.ProjectSpecification) error {
	return nil
}

func (s DotnetManager) BuildProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err

	}

	if !utils.IsEmpty(project.Group) {

		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "build", project.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "build", filepath.Join(project.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s DotnetManager) CleanProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "clean", project.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "clean", filepath.Join(project.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s DotnetManager) InstallProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "restore", project.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "restore", filepath.Join(project.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s DotnetManager) TestProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "test", project.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "test", filepath.Join(project.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s DotnetManager) PackageProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !utils.IsEmpty(project.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "publish", project.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "publish", filepath.Join(project.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s DotnetManager) RunProject(project applicationproject.ProjectSpecification) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err

	}

	if !utils.IsEmpty(project.Group) {

		if !groupStatus {
			return errors.New("Project group (" + project.Group + ") is not correct")
		} else {
			err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "run", "--project", project.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "run", "--project", filepath.Join(project.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s DotnetManager) RemoveDefaultFiles(project applicationproject.ProjectSpecification) error {
	var paths []string = []string{}
	projectPath := project.GetAbsoluteProjectPath()

	var projectType models.ProjectType = models.ProjectType(project.ProjectType)
	if projectType == models.ProjectTypes.Library {
		paths = append(paths, filepath.Join(projectPath, "Class1.cs"))
	} else if projectType == models.ProjectTypes.WebApi {
		paths = append(paths, filepath.Join(projectPath, "WeatherForecast.cs"))
		paths = append(paths, filepath.Join(projectPath, "Controllers", "WeatherForecastController.cs"))
	} else if projectType == models.ProjectTypes.Console {
		paths = append(paths, filepath.Join(projectPath, "Program.cs"))
	}

	return s.FileRemover(paths...)
}
func (s *DotnetManager) addHelloWorld(project applicationproject.ProjectSpecification) error {

	// projectPath := project.GetProjectPath()
	// var projectType models.ProjectType = models.ProjectType(project.GetSchema().ProjectType)

	// templateManager := DotnetTemplateEngine{}
	// if projectType == models.ProjectTypes.Library {
	// 	outputFile := filepath.Join(projectPath, "SampleClass.cs")

	// 	data := resources.Class{
	// 		Package: project.Name,
	// 		Name:    "SampleClass",
	// 	}
	// 	err := templateManager.AddSimpleClass(data.Name, outputFile, data)
	// 	if err != nil {
	// 		return err
	// 	}
	// } else if projectType == models.ProjectTypes.WebApi {
	// 	outputFile := filepath.Join(projectPath, "Controllers", "SampleController.cs")

	// 	data := resources.Controller{
	// 		Class: resources.Class{
	// 			Package: project.Name,
	// 			Name:    "SampleController",
	// 		},
	// 	}

	// 	err := templateManager.AddSimpleControllerClass(data.Name, outputFile, data)
	// 	if err != nil {
	// 		return err
	// 	}
	// } else if projectType == models.ProjectTypes.Console {
	// 	outputFile := filepath.Join(projectPath, "Program.cs")

	// 	data := resources.Class{
	// 		Package: project.Name,
	// 		Name:    "Program",
	// 	}
	// 	err := templateManager.AddSimpleConsoleClass(data.Name, outputFile, data)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func (s DotnetManager) GetPlatformVersion(platform applicationproject.Platform) dotnetModels.DotnetPlatformVersion {
	if utils.IsEmpty(platform.Version) {
		platformVersion := dotnetModels.DotnetPlatformVersions.Net8

		return platformVersion
	} else {
		platformVersion, err := dotnetModels.DotnetPlatformVersionEnumFromString(platform.Version)
		if err != nil {
			panic(err)
		}
		return platformVersion
	}
}

func (s DotnetManager) CreateGroup(project applicationproject.ProjectSpecification) error {
	return providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "new", "sln", "--name", project.GroupObject.Name, "--output", utils.PathWithDot(project.GetRelativeGroupPath()))
}

func (s DotnetManager) DeleteGroup(project applicationproject.ProjectSpecification) {
}

func (s DotnetManager) AddToGroup(project applicationproject.ProjectSpecification) error {
	return providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "sln", s.GetGroupFileRelativePath(project), "add", s.GetProjectFileRelativePath(project))
}

func (s DotnetManager) RemoveFromGroup(project applicationproject.ProjectSpecification) error {
	return providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "sln", s.GetGroupFileRelativePath(project), "remove", s.GetProjectFileRelativePath(project))
}

func (s DotnetManager) AddFolderToProjectDefinition(project applicationproject.ProjectSpecification, paths ...string) {
	for _, path := range paths {

		if utils.IsEmpty(path) {
			return
		}

		projectFile := filepath.Join(project.GetAbsoluteProjectPath(), s.GetProjectFileName(project))
		folderPath := fmt.Sprintf("%v\\", filepath.Join(path))

		data, err := os.ReadFile(projectFile)
		if err != nil {
			log.Fatal(err)
		}

		updatedXML, err := addFolderToItemProperty(data, folderPath)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(projectFile, updatedXML, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func (s DotnetManager) RemoveFolderFromProjectDefinition(project applicationproject.ProjectSpecification, paths ...string) {
	for _, path := range paths {

		if utils.IsEmpty(path) {
			return
		}

		projectFile := filepath.Join(project.GetAbsoluteProjectPath(), s.GetProjectFileName(project))
		folderPath := fmt.Sprintf("%v\\", filepath.Join(path))

		data, err := os.ReadFile(projectFile)
		if err != nil {
			log.Fatal(err)
		}

		updatedXML, err := removeFolderFromItemProperty(data, folderPath)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(projectFile, updatedXML, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func addFolderToItemProperty(xmlContent []byte, folderPath string) ([]byte, error) {
	m, err := mxj.NewMapXml(xmlContent)
	if err != nil {
		return nil, err
	}

	folders, err := m.ValuesForPath("Project.ItemGroup.Folder.-Include")
	if err != nil {
		return nil, err
	}

	for _, f := range folders {
		if f == folderPath {
			return m.XmlIndent("", "    ")
		}
	}

	itemGroups, err := m.ValuesForPath("Project.ItemGroup")
	if err != nil {
		return nil, err
	}

	for _, itemGroup := range itemGroups {
		itemGroupMap := itemGroup.(map[string]interface{})
		folderInterface, ok := itemGroupMap["Folder"]
		if ok {
			switch folder := folderInterface.(type) {
			case []interface{}:
				existingFolders := make([]interface{}, len(folder))
				copy(existingFolders, folder)

				newFolder := map[string]interface{}{
					"-Include": folderPath,
				}
				existingFolders = append(existingFolders, newFolder)

				itemGroupMap["Folder"] = existingFolders
			case map[string]interface{}:
				existingFolders := make([]interface{}, 1)
				existingFolders[0] = folder

				newFolder := map[string]interface{}{
					"-Include": folderPath,
				}
				existingFolders = append(existingFolders, newFolder)

				itemGroupMap["Folder"] = existingFolders
			}
			return m.XmlIndent("", "    ")
		}
	}

	newItemGroup := map[string]interface{}{
		"Folder": map[string]interface{}{
			"-Include": folderPath,
		},
	}
	itemGroups = append(itemGroups, newItemGroup)
	m["Project"].(map[string]interface{})["ItemGroup"] = itemGroups

	return m.XmlIndent("", "    ")
}

func removeFolderFromItemProperty(xmlContent []byte, folderPath string) ([]byte, error) {
	m, err := mxj.NewMapXml(xmlContent)
	if err != nil {
		return nil, err
	}

	folders, err := m.ValuesForPath("Project.ItemGroup.Folder.-Include")
	if err != nil {
		return nil, err
	}

	isExists := false
	for _, f := range folders {
		if f == folderPath {
			isExists = true
			break
		}
	}

	if !isExists {
		return m.XmlIndent("", "    ")
	}

	itemGroups, err := m.ValuesForPath("Project.ItemGroup")
	if err != nil {
		return nil, err
	}

	for _, itemGroup := range itemGroups {
		itemGroupMap, _ := itemGroup.(map[string]interface{})
		folders, ok := itemGroupMap["Folder"]
		if !ok {
			continue
		}

		switch folders := folders.(type) {
		case []interface{}:
			var updatedFolders []interface{}
			for _, folder := range folders {
				folderMap, _ := folder.(map[string]interface{})
				includeAttr, ok := folderMap["-Include"].(string)
				if !ok || includeAttr != folderPath {
					updatedFolders = append(updatedFolders, folder)
				}
			}
			itemGroupMap["Folder"] = updatedFolders
		case map[string]interface{}:
			includeAttr, ok := folders["-Include"].(string)
			if ok && includeAttr == folderPath {
				// Remove the whole ItemGroup if it has only the folder
				if len(itemGroupMap) == 1 {
					delete(m, "Project.ItemGroup")
				} else {
					delete(itemGroupMap, "Folder")
				}
			}
		}
	}

	return m.XmlIndent("", "    ")
}

func (s DotnetManager) AddPackageToProject(project applicationproject.ProjectSpecification, packages []applicationproject.Package) error {

	for _, _package := range packages {

		commandArgs := []string{"add", s.GetProjectFileRelativePath(project), "package", _package.Name}

		if !utils.IsEmpty(_package.Version) {
			commandArgs = append(commandArgs, []string{"--version", _package.Version}...)

		}

		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), commandArgs...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s DotnetManager) ListPackagesFromProject(projectSpecification applicationproject.ProjectSpecification) ([]applicationproject.Package, error) {

	commandArgs := []string{"list", s.GetProjectFileRelativePath(projectSpecification), "package"}

	output, err := providers.DotnetExecuteWithOutput(string(s.GetPlatformVersion(projectSpecification.Platform)), projectSpecification.GetCodeBasePath(), commandArgs...)
	if err != nil {
		return nil, err
	}

	pattern := regexp.MustCompile(`>\s(.*?)\s+(.*?)\s`)

	matches := pattern.FindAllStringSubmatch(output, -1)

	packages := make([]applicationproject.Package, 0)
	for _, match := range matches {
		packages = append(packages, applicationproject.NewPackage(string(match[1]), string(match[2])))
	}

	return packages, nil
}

func (s DotnetManager) RemovePackageFromProject(project applicationproject.ProjectSpecification, packages []applicationproject.Package) error {

	for _, _package := range packages {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "remove", s.GetProjectFileRelativePath(project), "package", _package.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s DotnetManager) AddReferenceToProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error {

	for _, reference := range references {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "add", s.GetProjectFileRelativePath(project), "reference", s.GetProjectFileRelativePath(reference))
		if err != nil {
			return err
		}
	}

	return nil
}

func (s DotnetManager) RemoveReferenceFromProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error {

	for _, reference := range references {

		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Platform)), project.GetCodeBasePath(), "remove", s.GetProjectFileRelativePath(project), "reference", s.GetProjectFileRelativePath(reference))
		if err != nil {
			return err
		}
	}

	return nil
}

func (s DotnetManager) GetProjectFileRelativePath(project applicationproject.ProjectSpecification) string {

	return filepath.Join(project.GetRelativeProjectPath(), s.GetProjectFileName(project))
}

func (s DotnetManager) GetProjectFileAbsolutePath(project applicationproject.ProjectSpecification) string {

	return filepath.Join(project.GetAbsoluteProjectPath(), s.GetProjectFileName(project))
}

func (s DotnetManager) GetGroupFileRelativePath(project applicationproject.ProjectSpecification) string {

	return filepath.Join(project.GetRelativeGroupPath(), s.GetGroupFileName(project))
}

func (s DotnetManager) GetGroupFileAbsolutePath(project applicationproject.ProjectSpecification) string {

	return filepath.Join(project.GetAbsoluteGroupPath(), s.GetGroupFileName(project))
}

func (s DotnetManager) IsProjectFileExists(project applicationproject.ProjectSpecification) (bool, error) {

	stat, err := os.Stat(s.GetProjectFileAbsolutePath(project)) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}

func (s DotnetManager) GetProjectFileName(project applicationproject.ProjectSpecification) string {
	return fmt.Sprintf("%v.csproj", project.Name)
}

func (s DotnetManager) IsGroupFileExists(project applicationproject.ProjectSpecification) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.GetAbsoluteGroupPath(), s.GetGroupFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s DotnetManager) GetGroupFileName(project applicationproject.ProjectSpecification) string {
	return fmt.Sprintf("%v.sln", project.GroupObject.Name)
}
func (s DotnetManager) HasPackageOnProject(project applicationproject.ProjectSpecification, _package applicationproject.Package) (bool, error) {

	packages, err := s.ListPackagesFromProject(project)
	if err != nil {
		return false, err
	}

	packageState := false

	for _, projectPackage := range packages {
		if projectPackage.Name == _package.Name && (utils.IsEmpty(_package.Version) || (projectPackage.Version == _package.Version)) {
			packageState = true
			break
		}
	}

	return packageState, nil
}

func (s DotnetManager) HasReferenceOnProject(project applicationproject.ProjectSpecification, reference applicationproject.ProjectSpecification) (bool, error) {

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

func (s DotnetManager) ListProjectsFromGroup(proj applicationproject.ProjectSpecification) ([]applicationproject.ProjectSpecification, error) {

	groupFile := filepath.Join(proj.GetAbsoluteGroupPath(), fmt.Sprintf("%v.sln", proj.Group))

	data, err := os.ReadFile(groupFile)
	if err != nil {
		log.Fatal(err)
	}

	pattern := regexp.MustCompile(`Project(.*?)\s=\s"(.*?)",\s"(.*?).csproj",\s(.*?)\nEndProject`)

	matches := pattern.FindAllStringSubmatch(string(data), -1)

	projects := make([]applicationproject.ProjectSpecification, 0)
	for _, match := range matches {
		projects = append(projects, applicationproject.ProjectSpecification{
			ProjectIdentifier: applicationproject.ProjectIdentifier{Name: string(match[2]), Group: proj.Group, Workspace: proj.Workspace},
			Path:              []string{filepath.Dir(string(match[3]))},
		},
		)
	}

	return projects, nil
}

func (s DotnetManager) HasProjectOnGroup(project applicationproject.ProjectSpecification) (bool, error) {

	projects, err := s.ListProjectsFromGroup(project)
	if err != nil {
		return false, err
	}

	projectState := false

	for _, groupProject := range projects {
		if groupProject.Name == project.Name && groupProject.GetRelativeProjectPath() == project.GetRelativeProjectPath() {
			projectState = true
			break
		}
	}

	return projectState, nil
}

func (s DotnetManager) ListReferencesFromProject(projectSpecification applicationproject.ProjectSpecification) ([]applicationproject.ProjectSpecification, error) {

	commandArgs := []string{"list", s.GetProjectFileRelativePath(projectSpecification), "reference"}

	output, err := providers.DotnetExecuteWithOutput(string(s.GetPlatformVersion(projectSpecification.Platform)), projectSpecification.GetCodeBasePath(), commandArgs...)
	if err != nil {
		return nil, err
	}

	pattern := regexp.MustCompile(`(.*?)\.csproj`)

	matches := pattern.FindAllStringSubmatch(output, -1)

	references := make([]applicationproject.ProjectSpecification, 0)
	for _, match := range matches {
		for _, projectReference := range projectSpecification.Configuration.References {
			relativeToReference, err := utils.FindRelativePath(projectSpecification.GetAbsoluteProjectPath(), projectReference.Specifications.GetAbsoluteProjectPath())
			if err != nil {
				return nil, err
			}

			pathWithProjectName := filepath.Join(relativeToReference, s.GetProjectFileName(projectReference.Specifications))
			if string(match[0]) == pathWithProjectName {
				references = append(references, projectReference.Specifications)
				break
			}

		}
	}

	return references, nil
}

func (s DotnetManager) ListFoldersFromProjectDefinition(proj applicationproject.ProjectSpecification) ([]string, error) {

	groupFile := filepath.Join(proj.GetAbsoluteProjectPath(), s.GetProjectFileName(proj))

	data, err := os.ReadFile(groupFile)
	if err != nil {
		log.Fatal(err)
	}

	pattern := regexp.MustCompile(`<Folder\sInclude="(.*?)"\/>`)

	matches := pattern.FindAllStringSubmatch(string(data), -1)

	folders := make([]string, 0)
	for _, match := range matches {
		folders = append(folders, string(match[1]))
	}

	return folders, nil
}

func (s DotnetManager) ListLayersFromProject(projectSpecification applicationproject.ProjectSpecification) ([]applicationproject.Layer, error) {

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

func (s DotnetManager) HasLayerOnProject(project applicationproject.ProjectSpecification, layer string) (bool, error) {

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

func (s DotnetManager) PrintPackage(packages []string) string {
	var nonEmptyPackages []string
	for _, pkg := range packages {
		if pkg != "" {
			nonEmptyPackages = append(nonEmptyPackages, pkg)
		}
	}
	return strings.Join(nonEmptyPackages, ".")
}
func (s DotnetManager) PrintDataType(dataType objectresource.DataType) string {
	result := ""
	if dataType.Category == objectresource.DataTypeCategories.Value {
		switch dataType.Name {
		case string(objectresource.ValueTypes.ShortInt):
			result = "short"
		case string(objectresource.ValueTypes.Int):
			result = "int"
		case string(objectresource.ValueTypes.LongInt):
			result = "long"
		case string(objectresource.ValueTypes.Double):
			result = "double"
		case string(objectresource.ValueTypes.Decimal):
			result = "decimal"
		case string(objectresource.ValueTypes.Float):
			result = "float"
		case string(objectresource.ValueTypes.Char):
			result = "char"
		case string(objectresource.ValueTypes.String):
			result = "string"
		case string(objectresource.ValueTypes.Date):
			result = "DateTime"
		case string(objectresource.ValueTypes.DateTime):
			result = "DateTime"
		case string(objectresource.ValueTypes.Time):
			result = "DateTime"
		case string(objectresource.ValueTypes.Boolean):
			result = "bool"
		case string(objectresource.ValueTypes.Byte):
			result = "byte"
		case string(objectresource.ValueTypes.ShortBlob):
			result = "byte[]"
		case string(objectresource.ValueTypes.Blob):
			result = "byte[]"
		case string(objectresource.ValueTypes.LongBlob):
			result = "byte[]"
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

func (s DotnetManager) PrintVisibility(visibility objectresource.VisibilityType) string {
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
