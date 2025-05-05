package basic

import (
	"fmt"
	"os"
	"testing"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test"
	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type BasicProjectStructureTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	workspace     string
	faker         *faker.Faker
	set           string
	noCleanOnFail bool
}

func (suite *BasicProjectStructureTestSuite) SetupSuite() {

	suite.faker = faker.NewFaker()

	suite.T().Log("Preparing test suite...")

	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.workspace = suite.faker.Workspace.Name()
	suite.set = suite.faker.Project.Set()

	tempWorkingDir, err := test.CreateTempTestDirectory(testArea)
	require.NoError(suite.T(), err, "Create temporary directory failed")
	suite.testArea = tempWorkingDir
	suite.T().Logf("Creating test location at (%v)", suite.testArea)

	suite.T().Logf("Initializing New Workspace (%v)", suite.workspace)
	common.InitializeNewWorkspace(suite.T(), suite.testArea, suite.workspace, suite.environment)

	suite.T().Logf("Switching to workspace (%v)...", suite.workspace)
	common.SwitchToWorkspace(suite.T(), suite.workspace, suite.environment)

	suite.T().Log("Test suite setup completed")
}
func (suite *BasicProjectStructureTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		common.RemoveWorkspace(suite.T(), suite.workspace, suite.environment)
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *BasicProjectStructureTestSuite) SetupTest()    {}
func (suite *BasicProjectStructureTestSuite) TearDownTest() {}

func (suite *BasicProjectStructureTestSuite) fTestCreateBasicResource() {

	projectName1 := suite.faker.Project.Name()
	projectLayers1 := make([]string, 0)
	projectLayers1 = append(projectLayers1, suite.faker.Project.Layer())
	projectLayers1 = append(projectLayers1, suite.faker.Project.Layer())
	projectTemplateFile1 := CreateNewProjectFromTemplateFile(common.CommanderTypes.GO, suite.T(), suite.environment, suite.testArea, projectName1, suite.set, projectLayers1, []string{}, []string{})

	projectName2 := suite.faker.Project.Name()
	projectLayers2 := make([]string, 0)
	projectLayers2 = append(projectLayers2, suite.faker.Project.Layer())
	projectLayers2 = append(projectLayers2, suite.faker.Project.Layer())
	projectLayers2 = append(projectLayers2, suite.faker.Project.Layer())
	projectTemplateFile2 := CreateNewProjectFromTemplateFile(common.CommanderTypes.GO, suite.T(), suite.environment, suite.testArea, projectName2, suite.set, projectLayers2, []string{}, []string{projectName1})

	templateName := suite.faker.Project.Name()
	templateServiceInterfaceTemplateFile := CreateNewServiceInterfaceTemplateFromTemplateFile(suite.T(), suite.environment, suite.testArea, fmt.Sprintf("%vServiceInterface", templateName), suite.set, []string{projectLayers1[0]})
	templateServiceTemplateFile := CreateNewServiceTemplateFromTemplateFile(suite.T(), suite.environment, suite.testArea, fmt.Sprintf("%vService", templateName), suite.set, []string{projectLayers2[1]})

	resourceName := suite.faker.Project.Name()
	resourcePath := suite.faker.Project.Path(1)
	resourcePackages := resourcePath

	structLayers := make([]string, 0)
	structLayers = append(structLayers, projectLayers1...)
	structLayers = append(structLayers, projectLayers2...)
	var structData = struct {
		Name    string
		Set     string
		Path    string
		Package string
		Layers  []string
	}{
		Name:    resourceName,
		Set:     suite.set,
		Path:    resourcePath,
		Package: resourcePackages,
		Layers:  structLayers,
	}

	declarationFile := utils.GetTestFileFromCurrentLocation("resources.yaml")
	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)

	common.SubmitResourceFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	suite.T().Cleanup(func() {
		common.RemoveProjectFromFile(common.CommanderTypes.GO, suite.T(), projectTemplateFile1, suite.environment)
		os.Remove(projectTemplateFile1)

		common.RemoveProjectFromFile(common.CommanderTypes.GO, suite.T(), projectTemplateFile2, suite.environment)
		os.Remove(projectTemplateFile2)

		common.RemoveTemplateFromFile(suite.T(), templateServiceInterfaceTemplateFile, suite.environment)
		os.Remove(templateServiceInterfaceTemplateFile)

		common.RemoveTemplateFromFile(suite.T(), templateServiceTemplateFile, suite.environment)
		os.Remove(templateServiceTemplateFile)

		common.RemoveResourceFromFile(suite.T(), templateFile, suite.environment)
		os.Remove(templateFile)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func TestBasicProjectStructureTestSuite(t *testing.T) {
	suite.Run(t, new(BasicProjectStructureTestSuite))
}

func CreateNewProjectFromTemplateFile(commander common.CommanderType, t *testing.T, environment, testArea, name, set string, layers []string, dependencies []string, references []string) string {

	declarationFile := utils.GetTestFileFromCurrentLocation("projects.yaml")

	var structData = struct {
		Name         string
		Platform     string
		Type         string
		Set          string
		Layers       []string
		Dependencies []string
		References   []string
	}{
		Name:         name,
		Platform:     "dotnet",
		Type:         "library",
		Set:          set,
		Layers:       layers,
		References:   references,
		Dependencies: dependencies,
	}

	templateFile := common.CreateTempFileFromTemplate(t, declarationFile, testArea, structData)

	common.SubmitProjectFromFile(commander, t, templateFile, environment)

	return templateFile
}
func CreateNewServiceTemplateFromTemplateFile(t *testing.T, environment, testArea, name, set string, layers []string) string {

	declarationFile := utils.GetTestFileFromCurrentLocation("templates.yaml")

	var structData = struct {
		Name    string
		Set     string
		Path    string
		Output  string
		Layers  []string
		Package string
		Code    string
	}{
		Name:    name,
		Set:     set,
		Path:    "",
		Output:  `Service.cs`,
		Layers:  layers,
		Package: name,
		Code: `namespace {{.Resource.Package}};

      public class {{.Resource.Name}}Service : I{{.Resource.Name}}Service
      {
        {{range .Resource.Attributes}}
        {{.Visibility}} {{.Type}} {{.Name}} { set; get; }
        {{end}}

        public string GreetingsFromPars()
        {
          return "Pars wishes you success :)";
        }

        {{range .Resource.Methods}}
        {{ if gt (len .ReturnTypes) 0 }}
        {{ $parametersLength := len .Parameters }}
        {{.Visibility}} {{ index .ReturnTypes 0 }} {{.Name}}({{range $i, $p := .Parameters}}{{ $p.Name }} {{ $p.Type }}{{if ne $i (math.Sub $parametersLength 1) }}, {{end}}{{end}}) {
          {{ index .ReturnTypes 0 }} result = default({{ index .ReturnTypes 0 }})

          return result
        }
        {{end}}
        {{else}}
        {{.Visibility}} void {{.Name}} {
        }
        {{end}}
      }`,
	}

	templateFile := common.CreateTempFileFromTemplate(t, declarationFile, testArea, structData)

	common.SubmitTemplateFromFile(common.CommanderTypes.GO, t, templateFile, environment)

	return templateFile
}
func CreateNewServiceInterfaceTemplateFromTemplateFile(t *testing.T, environment, testArea, name, set string, layers []string) string {

	declarationFile := utils.GetTestFileFromCurrentLocation("templates.yaml")

	var structData = struct {
		Name    string
		Set     string
		Path    string
		Output  string
		Layers  []string
		Package string
		Code    string
	}{
		Name:    name,
		Set:     set,
		Path:    "",
		Output:  `IService.cs`,
		Layers:  layers,
		Package: name,
		Code: `namespace {{.Resource.Package}};

      public interface I{{.Resource.Name}}Service
      {
        {{range .Resource.Attributes}}
        {{.Type}} {{.Name}} { set; get; }
        {{end}}


        {{range .Resource.Methods}}
        {{ if gt (len .ReturnTypes) 0 }}
        {{ $parametersLength := len .Parameters }}
        {{ index .ReturnTypes 0 }} {{.Name}} ({{range $i, $p := .Parameters}}{{ $p.Name }} {{ $p.Type }}{{if ne $i (math.Sub $parametersLength 1) }}, {{end}}{{end}});
        {{end}}
        {{else}}
        {{.Visibility}} void {{.Name}};
        {{end}}
      }`,
	}

	templateFile := common.CreateTempFileFromTemplate(t, declarationFile, testArea, structData)

	common.SubmitTemplateFromFile(common.CommanderTypes.GO, t, templateFile, environment)

	return templateFile
}
