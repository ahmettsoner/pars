package webapi

import (
	"fmt"
	"log"
	"os"

	"parsdevkit.net/pkg/utilities/json"

	"parsdevkit.net/application/engines"
	group "parsdevkit.net/modules/group/basic_group"

	"parsdevkit.net/components/schema"

	"parsdevkit.net/models"
	_string "parsdevkit.net/pkg/utilities/string"

	dotnetModels "parsdevkit.net/platforms/dotnet/models"

	projectComponent "parsdevkit.net/components/project"

	"parsdevkit.net/components/workspace"

	"github.com/spf13/cobra"
)

var (
	noInit                   bool = true
	name                     string
	workspaceName            string
	projectSet               string
	_package                 string
	platformVersionEnumFlag  dotnetModels.DotnetPlatformVersionEnumFlag
	runtimeVersionEnumFlag   dotnetModels.DotnetRuntimeVersionEnumFlag
	methodologyTypeEnumFlag  models.MethodologyTypeEnumFlag
	designTypeEnumFlag       models.DesignTypeEnumFlag
	architectureTypeEnumFlag models.ArchitectureTypeEnumFlag
	templateTypeEnumFlag     models.TemplateTypeEnumFlag
)

var WebApiCmd = &cobra.Command{
	Use:     "webapi",
	Aliases: []string{"a"},
	Short:   "Initialize new project",
	Long:    `Initialize new project`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	if _string.IsEmpty(name) {
		if len(args) == 0 {
			fmt.Println("Please provide a name for the new project")
			os.Exit(1)
		} else if len(args) > 0 {
			name = args[0]
		}
	}

	if _string.IsEmpty(name) {
		cmd.Help()
		os.Exit(0)
	}

	workspaceName = workspace.GetActiveWorkspaceName(workspaceName)

	projectGroup, _, err := projectComponent.ParseProjectFullName(name)
	if err != nil {
		log.Fatal(err)
	}

	// var structData = struct {
	// 	Group           string
	// 	Name            string
	// 	Set             string
	// 	Package         string
	// 	Path            string
	// 	Workspace       string
	// 	PlatformVersion dotnetModels.DotnetPlatformVersion
	// 	RuntimeVersion  dotnetModels.DotnetRuntimeVersion
	// 	DesignType      models.DesignType
	// 	Architecture    models.ArchitectureType
	// 	Template        models.TemplateType
	// 	Methodology     models.MethodologyType
	// }{
	// 	Group:           projectGroup,
	// 	Name:            projectName,
	// 	Set:             projectSet,
	// 	Package:         _package,
	// 	Path:            projectName,
	// 	PlatformVersion: platformVersionEnumFlag.Value,
	// 	RuntimeVersion:  runtimeVersionEnumFlag.Value,
	// 	Methodology:     methodologyTypeEnumFlag.Value,
	// 	DesignType:      designTypeEnumFlag.Value,
	// 	Architecture:    architectureTypeEnumFlag.Value,
	// 	Template:        templateTypeEnumFlag.Value,
	// 	Workspace:       workspaceName,
	// }

	var templateFilePath = "/dotnet/projects/webapi.yaml.templ"
	if architectureTypeEnumFlag.Value == models.ArchitectureTypes.None {
		if methodologyTypeEnumFlag.Value == models.MethodologyTypes.Basic {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/dotnet/projects/webapi.yaml.templ"
			}
		} else if methodologyTypeEnumFlag.Value == models.MethodologyTypes.Layered {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/dotnet/projects/webapi-layered.yaml.templ"
			}
		} else if methodologyTypeEnumFlag.Value == models.MethodologyTypes.NTier {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/dotnet/projects/webapi-ntier.yaml.templ"
			}
		}
	} else if architectureTypeEnumFlag.Value == models.ArchitectureTypes.Clean {
		if methodologyTypeEnumFlag.Value == models.MethodologyTypes.Basic {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/dotnet/projects/webapi-clean.yaml.templ"
			}
		} else if methodologyTypeEnumFlag.Value == models.MethodologyTypes.Layered {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/dotnet/projects/webapi-layered-clean.yaml.templ"
			}
		} else if methodologyTypeEnumFlag.Value == models.MethodologyTypes.NTier {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/dotnet/projects/webapi-ntier-clean.yaml.templ"
			}
		}
	}

	if !_string.IsEmpty(projectGroup) {
		groupService := group.GroupEngine{}

		var groupStructData = struct {
			Name string
		}{
			Name: projectGroup,
		}
		var groupTemplateFilePath = "/group/group.yaml.templ"
		if err := groupService.CreateGroupsFromTemplate(!noInit, groupStructData, groupTemplateFilePath); err != nil {
			log.Fatal(err)
		}
	}

	result, err := schema.GenerateManifestFilesFromTemplate(templateFilePath)

	if err != nil {
		log.Printf("❌ Error: %v", err)
	} else {
		for _, data := range result {

			if err := data.Validate(); err != nil {
				jsonObject, _ := json.ToJson(data)
				return fmt.Errorf("invalid data: '%s'\n%w", jsonObject, err)
			}

			fmt.Printf("✅ Loaded: %#v\n", data.GetHeader().Name)
		}

		err = engines.DispatchEngineProcess(result)
		if err != nil {
			log.Fatalf("Engine processing failed: %v", err)
		}
	}

}

func init() {
	WebApiCmd.Flags().BoolVarP(&noInit, "no-init", "", false, "Create project but do not initialize")

	WebApiCmd.Flags().StringVarP(&name, "name", "n", "", "Project name")

	WebApiCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
	WebApiCmd.Flags().StringVarP(&projectSet, "project-set", "s", "", "Project Set")
	WebApiCmd.Flags().StringVarP(&_package, "package", "p", "", "Package")

	platformVersionValues := dotnetModels.DotnetPlatformVersionToArray()
	platformVersionEnumFlag.Value = dotnetModels.DotnetPlatformVersions.Net8
	WebApiCmd.PersistentFlags().VarP(&platformVersionEnumFlag, "platform", "", fmt.Sprintf("Select platform version %v", platformVersionValues))

	runtimeVersionValues := dotnetModels.DotnetRuntimeVersionToArray()
	runtimeVersionEnumFlag.Value = dotnetModels.DotnetRuntimeVersions.Net8
	WebApiCmd.PersistentFlags().VarP(&runtimeVersionEnumFlag, "runtime", "", fmt.Sprintf("Select runtime version %v", runtimeVersionValues))

	methodologyTypeValues := models.MethodologyTypeToArray()
	methodologyTypeEnumFlag.Value = models.MethodologyTypes.Basic
	WebApiCmd.PersistentFlags().VarP(&methodologyTypeEnumFlag, "methodology", "m", fmt.Sprintf("Select a methodology %v", methodologyTypeValues))

	designTypeValues := models.DesignTypeToArray()
	designTypeEnumFlag.Value = models.DesignType(models.DesignTypes.Classic)
	WebApiCmd.PersistentFlags().VarP(&designTypeEnumFlag, "design", "d", fmt.Sprintf("Select a design %v", designTypeValues))

	architectureTypeValues := models.ArchitectureTypeToArray()
	architectureTypeEnumFlag.Value = models.ArchitectureTypes.None
	WebApiCmd.PersistentFlags().VarP(&architectureTypeEnumFlag, "architecture", "a", fmt.Sprintf("Select a architecture %v", architectureTypeValues))

	validEnumValues := models.TemplateTypeToArray()
	templateTypeEnumFlag.Value = models.TemplateTypes.Simple
	WebApiCmd.PersistentFlags().VarP(&templateTypeEnumFlag, "template", "t", fmt.Sprintf("Select a template type %v", validEnumValues))
}
