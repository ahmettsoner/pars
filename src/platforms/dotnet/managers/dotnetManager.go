package managers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"parsdevkit.net/application/structs"

	_string "parsdevkit.net/pkg/utilities/string"

	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/models"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	"parsdevkit.net/pkg/utilities/file"

	"parsdevkit.net/platforms/core"
	dotnetModels "parsdevkit.net/platforms/dotnet/models"

	"parsdevkit.net/providers"

	mxj "github.com/clbanning/mxj/v2"
)

type DotnetManager struct {
	core.BaseManager
}

func NewDotnetManager() core.ApplicationPlatformManagerInterface {
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

func (s DotnetManager) GetKey() models.PlatformType {
	return models.PlatformTypes.Dotnet
}
func (s DotnetManager) CreateProject(project application_project_payload_structs.ProjectBaseStruct) error {

	dotnetProjectType, err := ProjectTypeToDotnetCLITypeString(models.ProjectType(project.Application.ProjectType))
	if err != nil {
		return fmt.Errorf("xxx: proje oluşturma işlemi sırasında ProjectType bulunamadı? %w", err)
	}
	// if dotnetProjectType == "webapp" {
	// 	dotnetProjectType, err = DotnetWebAppOptionToDotnetCLITypeString(dotnetModels.DotnetWebAppOption(project.Specifications.GetConfiguration().Options))
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	err = providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "new", dotnetProjectType, "--name", project.Specifications.Name, "--output", file.PathWithDot(project.Specifications.GetRelativeProjectPath()))
	if err != nil {
		return err
	}

	return nil
}

func (s DotnetManager) RemoveProject(project application_project_payload_structs.ProjectBaseStruct) error {
	return nil
}

func (s DotnetManager) BuildProject(project application_project_payload_structs.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err

	}

	if !_string.IsEmpty(project.Specifications.Group) {

		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "build", project.Specifications.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "build", filepath.Join(project.Specifications.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s DotnetManager) CleanProject(project application_project_payload_structs.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "clean", project.Specifications.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "clean", filepath.Join(project.Specifications.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s DotnetManager) InstallProject(project application_project_payload_structs.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "restore", project.Specifications.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "restore", filepath.Join(project.Specifications.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s DotnetManager) TestProject(project application_project_payload_structs.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "test", project.Specifications.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "test", filepath.Join(project.Specifications.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s DotnetManager) PackageProject(project application_project_payload_structs.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err
	}

	if !_string.IsEmpty(project.Specifications.Group) {
		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "publish", project.Specifications.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "publish", filepath.Join(project.Specifications.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s DotnetManager) RunProject(project application_project_payload_structs.ProjectBaseStruct) error {
	groupStatus, err := s.IsGroupFileExists(project)
	if err != nil {
		return err

	}

	if !_string.IsEmpty(project.Specifications.Group) {

		if !groupStatus {
			return errors.New("Project group (" + project.Specifications.Group + ") is not correct")
		} else {
			err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "run", "--project", project.Specifications.GetRelativeProjectPath())
			if err != nil {
				return err
			}
		}
	} else {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "run", "--project", filepath.Join(project.Specifications.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s DotnetManager) RemoveDefaultFiles(project application_project_payload_structs.ProjectBaseStruct) error {
	var paths []string = []string{}
	projectPath := project.Specifications.GetAbsoluteProjectPath()

	var projectType models.ProjectType = models.ProjectType(project.Application.ProjectType)
	paths = append(paths, filepath.Join(projectPath, "obj"))
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
func (s *DotnetManager) addHelloWorld(project application_project_payload_structs.ProjectBaseStruct) error {

	// projectPath := project.Specifications.GetProjectPath()
	// var projectType models.ProjectType = models.ProjectType(project.Specifications.GetSchema().ProjectType)

	// templateManager := DotnetTemplateEngine{}
	// if projectType == models.ProjectTypes.Library {
	// 	outputFile := filepath.Join(projectPath, "SampleClass.cs")

	// 	data := resources.Class{
	// 		Package: project.Specifications.Name,
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
	// 			Package: project.Specifications.Name,
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
	// 		Package: project.Specifications.Name,
	// 		Name:    "Program",
	// 	}
	// 	err := templateManager.AddSimpleConsoleClass(data.Name, outputFile, data)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func (s DotnetManager) GetPlatformVersion(platform application_project_payload_structs.Platform) dotnetModels.DotnetPlatformVersion {
	if _string.IsEmpty(platform.Version) {
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

func (s DotnetManager) CreateGroup(project application_project_payload_structs.ProjectBaseStruct) error {
	return providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "new", "sln", "--name", project.Specifications.GroupObject.Name, "--output", file.PathWithDot(project.Specifications.GetRelativeGroupPath()))
}

func (s DotnetManager) DeleteGroup(project application_project_payload_structs.ProjectBaseStruct) {
}

func (s DotnetManager) AddToGroup(project application_project_payload_structs.ProjectBaseStruct) error {
	return providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "sln", s.GetGroupFileRelativePath(project), "add", s.GetProjectFileRelativePath(project))
}

func (s DotnetManager) RemoveFromGroup(project application_project_payload_structs.ProjectBaseStruct) error {
	return providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "sln", s.GetGroupFileRelativePath(project), "remove", s.GetProjectFileRelativePath(project))
}

func (s DotnetManager) AddFolderToProjectDefinition(project application_project_payload_structs.ProjectBaseStruct, paths ...string) error {
	for _, path := range paths {

		if _string.IsEmpty(path) {
			return nil
		}

		projectFile := filepath.Join(project.Specifications.GetAbsoluteProjectPath(), s.GetProjectFileName(project))
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

	return nil
}
func (s DotnetManager) RemoveFolderFromProjectDefinition(project application_project_payload_structs.ProjectBaseStruct, paths ...string) error {
	for _, path := range paths {

		if _string.IsEmpty(path) {
			return nil
		}

		projectFile := filepath.Join(project.Specifications.GetAbsoluteProjectPath(), s.GetProjectFileName(project))
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

	return nil
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

func (s DotnetManager) AddDependenciesToProject(project application_project_payload_structs.ProjectBaseStruct, dependencies []applicationProject.Dependency) error {

	for _, _package := range dependencies {

		commandArgs := []string{"add", s.GetProjectFileRelativePath(project), "package", _package.Name}

		if !_string.IsEmpty(_package.Version) {
			commandArgs = append(commandArgs, []string{"--version", _package.Version}...)

		}

		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), commandArgs...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s DotnetManager) ListDependenciesFromProject(project application_project_payload_structs.ProjectBaseStruct) ([]applicationProject.Dependency, error) {

	commandArgs := []string{"list", s.GetProjectFileRelativePath(project), "package"}

	output, err := providers.DotnetExecuteWithOutput(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), commandArgs...)
	if err != nil {
		return nil, err
	}

	pattern := regexp.MustCompile(`>\s(.*?)\s+(.*?)\s`)

	matches := pattern.FindAllStringSubmatch(output, -1)

	dependencies := make([]applicationProject.Dependency, 0)
	for _, match := range matches {
		dependencies = append(dependencies, applicationProject.NewDependency(string(match[1]), string(match[2])))
	}

	return dependencies, nil
}

func (s DotnetManager) RemoveDependenciesFromProject(project application_project_payload_structs.ProjectBaseStruct, dependencies []applicationProject.Dependency) error {

	for _, _package := range dependencies {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "remove", s.GetProjectFileRelativePath(project), "package", _package.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s DotnetManager) AddReferenceToProject(project application_project_payload_structs.ProjectBaseStruct, references []application_project_payload_structs.ProjectBaseStruct) error {

	for _, reference := range references {
		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "add", s.GetProjectFileRelativePath(project), "reference", s.GetProjectFileRelativePath(reference))
		if err != nil {
			return err
		}
	}

	return nil
}

func (s DotnetManager) RemoveReferenceFromProject(project application_project_payload_structs.ProjectBaseStruct, references []application_project_payload_structs.ProjectBaseStruct) error {

	for _, reference := range references {

		err := providers.DotnetExecute(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), "remove", s.GetProjectFileRelativePath(project), "reference", s.GetProjectFileRelativePath(reference))
		if err != nil {
			return err
		}
	}

	return nil
}

func (s DotnetManager) GetProjectFileRelativePath(project application_project_payload_structs.ProjectBaseStruct) string {

	return filepath.Join(project.Specifications.GetRelativeProjectPath(), s.GetProjectFileName(project))
}

func (s DotnetManager) GetProjectFileAbsolutePath(project application_project_payload_structs.ProjectBaseStruct) string {

	return filepath.Join(project.Specifications.GetAbsoluteProjectPath(), s.GetProjectFileName(project))
}

func (s DotnetManager) GetGroupFileRelativePath(project application_project_payload_structs.ProjectBaseStruct) string {

	return filepath.Join(project.Specifications.GetRelativeGroupPath(), s.GetGroupFileName(project))
}

func (s DotnetManager) GetGroupFileAbsolutePath(project application_project_payload_structs.ProjectBaseStruct) string {

	return filepath.Join(project.Specifications.GetAbsoluteGroupPath(), s.GetGroupFileName(project))
}

func (s DotnetManager) IsProjectFileExists(project application_project_payload_structs.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(s.GetProjectFileAbsolutePath(project)) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}

func (s DotnetManager) GetProjectFileName(project application_project_payload_structs.ProjectBaseStruct) string {
	return fmt.Sprintf("%v.csproj", project.Specifications.Name)
}

func (s DotnetManager) IsGroupFileExists(project application_project_payload_structs.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(filepath.Join(project.Specifications.GetAbsoluteGroupPath(), s.GetGroupFileName(project))) // check if the project file exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !stat.IsDir(), nil
	}
}
func (s DotnetManager) GetGroupFileName(project application_project_payload_structs.ProjectBaseStruct) string {
	return fmt.Sprintf("%v.sln", project.Specifications.GroupObject.Name)
}
func (s DotnetManager) HasDependencyOnProject(project application_project_payload_structs.ProjectBaseStruct, _package applicationProject.Dependency) (bool, error) {

	dependencies, err := s.ListDependenciesFromProject(project)
	if err != nil {
		return false, err
	}

	packageState := false

	for _, projectDependency := range dependencies {
		if projectDependency.Name == _package.Name && (_string.IsEmpty(_package.Version) || (projectDependency.Version == _package.Version)) {
			packageState = true
			break
		}
	}

	return packageState, nil
}

func (s DotnetManager) HasReferenceOnProject(project application_project_payload_structs.ProjectBaseStruct, reference application_project_payload_structs.ProjectBaseStruct) (bool, error) {

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

func (s DotnetManager) ListProjectsFromGroup(proj application_project_payload_structs.ProjectBaseStruct) ([]application_project_payload_structs.ProjectBaseStruct, error) {

	groupFile := filepath.Join(proj.Specifications.GetAbsoluteGroupPath(), fmt.Sprintf("%v.sln", proj.Specifications.Group))

	data, err := os.ReadFile(groupFile)
	if err != nil {
		log.Fatal(err)
	}

	pattern := regexp.MustCompile(`Project(.*?)\s=\s"(.*?)",\s"(.*?).csproj",\s(.*?)\nEndProject`)

	matches := pattern.FindAllStringSubmatch(string(data), -1)

	projects := make([]application_project_payload_structs.ProjectBaseStruct, 0)
	for _, match := range matches {
		projects = append(projects, application_project_payload_structs.ProjectBaseStruct{
			Specifications: application_project_payload_structs.ProjectSpecification{
				ProjectIdentifier: applicationProject.ProjectIdentifier{
					Name:      string(match[2]),
					Path:      []string{filepath.Dir(string(match[3]))},
					Group:     proj.Specifications.Group,
					Workspace: proj.Specifications.Workspace},
			},
		},
		)
	}

	return projects, nil
}

func (s DotnetManager) HasProjectOnGroup(project application_project_payload_structs.ProjectBaseStruct) (bool, error) {

	projects, err := s.ListProjectsFromGroup(project)
	if err != nil {
		return false, err
	}

	projectState := false

	for _, groupProject := range projects {
		if groupProject.Specifications.Name == project.Specifications.Name && groupProject.Specifications.GetRelativeProjectPath() == project.Specifications.GetRelativeProjectPath() {
			projectState = true
			break
		}
	}

	return projectState, nil
}

func (s DotnetManager) ListReferencesFromProject(project application_project_payload_structs.ProjectBaseStruct) ([]application_project_payload_structs.ProjectBaseStruct, error) {

	commandArgs := []string{"list", s.GetProjectFileRelativePath(project), "reference"}

	output, err := providers.DotnetExecuteWithOutput(string(s.GetPlatformVersion(project.Specifications.Platform)), project.Specifications.GetCodeBasePath(), commandArgs...)
	if err != nil {
		return nil, err
	}

	pattern := regexp.MustCompile(`(.*?)\.csproj`)

	matches := pattern.FindAllStringSubmatch(output, -1)

	references := make([]application_project_payload_structs.ProjectBaseStruct, 0)
	for _, match := range matches {
		for _, projectReference := range project.Specifications.References {
			relativeToReference, err := file.FindRelativePath(project.Specifications.GetAbsoluteProjectPath(), projectReference.Specifications.GetAbsoluteProjectPath())
			if err != nil {
				return nil, err
			}

			pathWithProjectName := filepath.Join(relativeToReference, s.GetProjectFileName(projectReference))

			unifiedPath := string(match[0])
			if runtime.GOOS != "windows" {
				unifiedPath = strings.ReplaceAll(unifiedPath, `\`, `/`)
			}

			// OS'ye uygun hale getir (slashes normalize edilir)
			normalizedPath := filepath.Clean(unifiedPath)
			if normalizedPath == pathWithProjectName {
				references = append(references, projectReference)
				break
			}

		}
	}

	return references, nil
}

func (s DotnetManager) ListFoldersFromProjectDefinition(project application_project_payload_structs.ProjectBaseStruct) ([]string, error) {

	groupFile := filepath.Join(project.Specifications.GetAbsoluteProjectPath(), s.GetProjectFileName(project))

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

func (s DotnetManager) ListLayersFromProject(project application_project_payload_structs.ProjectBaseStruct) ([]applicationProject.Layer, error) {

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

func (s DotnetManager) HasLayerOnProject(project application_project_payload_structs.ProjectBaseStruct, layer string) (bool, error) {

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

func (s DotnetManager) PrintDependencies(dependencies []string) string {
	var nonEmptyDependencies []string
	for _, pkg := range dependencies {
		if pkg != "" {
			nonEmptyDependencies = append(nonEmptyDependencies, pkg)
		}
	}
	return strings.Join(nonEmptyDependencies, ".")
}
func (s DotnetManager) PrintDataType(dataType structs.DataType) string {
	result := ""
	if dataType.Category == structs.DataTypeCategories.Value {
		switch dataType.Name {
		case string(structs.ValueTypes.ShortInt):
			result = "short"
		case string(structs.ValueTypes.Int):
			result = "int"
		case string(structs.ValueTypes.LongInt):
			result = "long"
		case string(structs.ValueTypes.Double):
			result = "double"
		case string(structs.ValueTypes.Decimal):
			result = "decimal"
		case string(structs.ValueTypes.Float):
			result = "float"
		case string(structs.ValueTypes.Char):
			result = "char"
		case string(structs.ValueTypes.String):
			result = "string"
		case string(structs.ValueTypes.Date):
			result = "DateTime"
		case string(structs.ValueTypes.DateTime):
			result = "DateTime"
		case string(structs.ValueTypes.Time):
			result = "DateTime"
		case string(structs.ValueTypes.Boolean):
			result = "bool"
		case string(structs.ValueTypes.Byte):
			result = "byte"
		case string(structs.ValueTypes.ShortBlob):
			result = "byte[]"
		case string(structs.ValueTypes.Blob):
			result = "byte[]"
		case string(structs.ValueTypes.LongBlob):
			result = "byte[]"
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

func (s DotnetManager) PrintVisibility(visibility structs.VisibilityType) string {
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
