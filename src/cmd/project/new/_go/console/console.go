package console

import (
	"fmt"
	"log"

	"gorm.io/gorm/schema"
	"parsdevkit.net/models"

	"parsdevkit.net/application/engines"

	projectComponent "parsdevkit.net/components/project"

	"parsdevkit.net/components/schema"
	"parsdevkit.net/components/workspace"
	_string "parsdevkit.net/pkg/utilities/string"

	"github.com/spf13/cobra"
	"parsdevkit.net/pkg/utilities/json"
)

type NewOptions struct {
	NoInit          bool
	Name            string
	Workspace       string
	Set             string
	Package         string
	PlatformVersion models.GoPlatformVersionEnumFlag
	RuntimeVersion  models.GoRuntimeVersionEnumFlag
	Methodolog      models.MethodologyTypeEnumFlag
	Design          models.DesignTypeEnumFlag
	Architecture    models.ArchitectureTypeEnumFlag
	Template        models.TemplateTypeEnumFlag
}

var commandOptions = NewOptions{
	NoInit: true,
}

var maxArgumentCount int = 0

var ConsoleCmd = &cobra.Command{
	Use:     "console",
	Aliases: []string{"c"},
	Short:   "Initialize new project",
	Long:    `Initialize new project`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: project name is required. Provide it with '--name' or as an argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("error: too many arguments. Only project name is expected.")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}

	if _string.IsEmpty(commandOptions.Workspace) {
		commandOptions.Workspace = workspace.GetActiveWorkspaceName("")
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

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
	// 	PlatformVersion models.GoPlatformVersion
	// 	RuntimeVersion  models.GoRuntimeVersion
	// 	DesignType      models.DesignType
	// 	Architecture    models.ArchitectureType
	// 	Template        models.TemplateType
	// 	Methodology     models.MethodologyType
	// }{
	// 	Group:           projectGroup,
	// 	Name:            commandOptions.Name,
	// 	Set:             commandOptions.Set,
	// 	Package:         commandOptions.Package,
	// 	Path:            commandOptions.Name,
	// 	PlatformVersion: commandOptions.PlatformVersion.Value,
	// 	RuntimeVersion:  commandOptions.RuntimeVersion.Value,
	// 	Methodology:     commandOptions.Methodology.Value,
	// 	DesignType:      commandOptions.Design.Value,
	// 	Architecture:    commandOptions.Architecture.Value,
	// 	Template:        commandOptions.Template.Value,
	// 	Workspace:       commandOptions.Workspace,
	// }

	var templateFilePath = "/go/projects/console.yaml.templ"
	if commandOptions.Architecture.Value == models.ArchitectureTypes.None {
		if commandOptions.Methodology.Value.Methodology.Value == models.MethodologyTypes.Basic {
			if commandOptions.Design.Value == models.DesignTypes.Classic {
				templateFilePath = "/go/projects/console.yaml.templ"
			}
		} else if commandOptions.Methodology.Value == models.MethodologyTypes.Layered {
			if commandOptions.Design.Value == models.DesignTypes.Classic {
				templateFilePath = "/go/projects/console-layered.yaml.templ"
			}
		} else if commandOptions.Methodology.Value == models.MethodologyTypes.NTier {
			if commandOptions.Design.Value == models.DesignTypes.Classic {
				templateFilePath = "/go/projects/console-ntier.yaml.templ"
			}
		}
	} else if commandOptions.Architecture.Value == models.ArchitectureTypes.Clean {
		if commandOptions.Methodology.Value == models.MethodologyTypes.Basic {
			if commandOptions.Design.Value == models.DesignTypes.Classic {
				templateFilePath = "/go/projects/console-clean.yaml.templ"
			}
		} else if commandOptions.Methodology.Value == models.MethodologyTypes.Layered {
			if commandOptions.Design.Value == models.DesignTypes.Classic {
				templateFilePath = "/go/projects/console-layered-clean.yaml.templ"
			}
		} else if commandOptions.Methodology.Value == models.MethodologyTypes.NTier {
			if commandOptions.Design.Value == models.DesignTypes.Classic {
				templateFilePath = "/go/projects/console-ntier-clean.yaml.templ"
			}
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
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = NewOptions{
		NoInit: true,
	}
}

func init() {
	ConsoleCmd.Flags().BoolVarP(&commandOptions.NoInit, "no-init", "", false, "Create project but do not initialize")

	ConsoleCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")

	ConsoleCmd.Flags().StringVarP(&commandOptions.WorkspaceName, "workspace", "w", "", "Workspace name")
	ConsoleCmd.Flags().StringVarP(&commandOptions.Set, "project-set", "s", "", "Project Set")
	ConsoleCmd.Flags().StringVarP(&commandOptions.Package, "package", "p", "", "Package")

	platformVersionValues := models.GoPlatformVersionToArray()
	platformVersionEnumFlag.Value = models.GoPlatformVersions.Go121
	ConsoleCmd.PersistentFlags().VarP(&commandOptions.PlatformVersion, "platform", "", fmt.Sprintf("Select platform version %v", platformVersionValues))

	runtimeVersionValues := models.GoRuntimeVersionToArray()
	runtimeVersionEnumFlag.Value = models.GoRuntimeVersions.Go121
	ConsoleCmd.PersistentFlags().VarP(&runtimeVersionEnumFlag, "runtime", "", fmt.Sprintf("Select runtime version %v", runtimeVersionValues))

	methodologyTypeValues := models.MethodologyTypeToArray()
	commandOptions.Methodology.Value = models.MethodologyTypes.Basic
	ConsoleCmd.PersistentFlags().VarP(&commandOptions.MethodologyType, "methodology", "m", fmt.Sprintf("Select a methodology %v", methodologyTypeValues))

	designTypeValues := models.DesignTypeToArray()
	commandOptions.Design.Value = models.DesignType(models.DesignTypes.Classic)
	ConsoleCmd.PersistentFlags().VarP(&commandOptions.DesignType, "design", "d", fmt.Sprintf("Select a design %v", designTypeValues))

	architectureTypeValues := models.ArchitectureTypeToArray()
	commandOptions.Architecture.Value = models.ArchitectureTypes.None
	ConsoleCmd.PersistentFlags().VarP(&commandOptions.ArchitectureType, "architecture", "a", fmt.Sprintf("Select a architecture %v", architectureTypeValues))

	validEnumValues := models.TemplateTypeToArray()
	commandOptions.Template.Value = models.TemplateTypes.Simple
	ConsoleCmd.PersistentFlags().VarP(&templateTypeEnumFlag, "template", "t", fmt.Sprintf("Select a template type %v", validEnumValues))
}
