package library

import (
	"fmt"
	"log"
	"os"

	"parsdevkit.net/engines/group"
	v2 "parsdevkit.net/engines/v2"
	"parsdevkit.net/models"
	_string "parsdevkit.net/pkg/utilities/string"
	"parsdevkit.net/pkg/utils/json"
	"parsdevkit.net/structs/project"

	nodejsModels "parsdevkit.net/platforms/nodejs/models"

	"parsdevkit.net/components/workspace"

	"github.com/spf13/cobra"
)

var (
	noInit                   bool = true
	name                     string
	workspaceName            string
	projectSet               string
	_package                 string
	platformVersionEnumFlag  nodejsModels.nodejsPlatformVersionEnumFlag
	runtimeVersionEnumFlag   nodejsModels.NodeJSRuntimeVersionEnumFlag
	templateTypeEnumFlag     models.TemplateTypeEnumFlag
	methodologyTypeEnumFlag  models.MethodologyTypeEnumFlag
	designTypeEnumFlag       models.DesignTypeEnumFlag
	architectureTypeEnumFlag models.ArchitectureTypeEnumFlag
)

var LibraryCmd = &cobra.Command{
	Use:     "library",
	Aliases: []string{"l"},
	Short:   "Initialize new project",
	Long:    `Initialize new project`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	if _string.IsEmpty(name) {
		if len(args) == 0 {
			fmt.Println("Please provide a project name")
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

	projectGroup, projectName, err := project.ParseProjectFullName(name)
	if err != nil {
		log.Fatal(err)
	}

	var structData = struct {
		Group           string
		Name            string
		Set             string
		Package         string
		Path            string
		Workspace       string
		PlatformVersion nodejsModels.nodejsPlatformVersion
		RuntimeVersion  nodejsModels.NodeJSRuntimeVersion
		DesignType      models.DesignType
		Architecture    models.ArchitectureType
		Template        models.TemplateType
		Methodology     models.MethodologyType
	}{
		Group:           projectGroup,
		Name:            projectName,
		Set:             projectSet,
		Package:         _package,
		Path:            projectName,
		PlatformVersion: platformVersionEnumFlag.Value,
		RuntimeVersion:  runtimeVersionEnumFlag.Value,
		Methodology:     methodologyTypeEnumFlag.Value,
		DesignType:      designTypeEnumFlag.Value,
		Architecture:    architectureTypeEnumFlag.Value,
		Template:        templateTypeEnumFlag.Value,
		Workspace:       workspaceName,
	}

	var templateFilePath = "/nodejs/projects/library.yaml.templ"
	if architectureTypeEnumFlag.Value == models.ArchitectureTypes.None {
		if methodologyTypeEnumFlag.Value == models.MethodologyTypes.Basic {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/nodejs/projects/library.yaml.templ"
			}
		} else if methodologyTypeEnumFlag.Value == models.MethodologyTypes.Layered {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/nodejs/projects/library-layered.yaml.templ"
			}
		} else if methodologyTypeEnumFlag.Value == models.MethodologyTypes.NTier {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/nodejs/projects/library-ntier.yaml.templ"
			}
		}
	} else if architectureTypeEnumFlag.Value == models.ArchitectureTypes.Clean {
		if methodologyTypeEnumFlag.Value == models.MethodologyTypes.Basic {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/nodejs/projects/library-clean.yaml.templ"
			}
		} else if methodologyTypeEnumFlag.Value == models.MethodologyTypes.Layered {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/nodejs/projects/library-layered-clean.yaml.templ"
			}
		} else if methodologyTypeEnumFlag.Value == models.MethodologyTypes.NTier {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/nodejs/projects/library-ntier-clean.yaml.templ"
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

	result, err := v2.GenerateManifestFilesFromTemplate(templateFilePath)

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

		err = v2.DispatchEngineProcess(result)
		if err != nil {
			log.Fatalf("Engine processing failed: %v", err)
		}
	}

	fmt.Fprintf(os.Stdout, "✔ Schema(s) '%v' applied successfully\n", commandOptions.FilePaths)
}

func init() {
	LibraryCmd.Flags().BoolVarP(&noInit, "no-init", "", false, "Create project but do not initialize")

	LibraryCmd.Flags().StringVarP(&name, "name", "n", "", "Project name")

	LibraryCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
	LibraryCmd.Flags().StringVarP(&projectSet, "project-set", "s", "", "Project Set")
	LibraryCmd.Flags().StringVarP(&_package, "package", "p", "", "Package")

	platformVersionValues := nodejsModels.nodejsPlatformVersionToArray()
	platformVersionEnumFlag.Value = nodejsModels.nodejsPlatformVersions.V16
	LibraryCmd.PersistentFlags().VarP(&platformVersionEnumFlag, "platform", "", fmt.Sprintf("Select platform version %v", platformVersionValues))

	runtimeVersionValues := nodejsModels.NodeJSRuntimeVersionToArray()
	runtimeVersionEnumFlag.Value = nodejsModels.NodeJSRuntimeVersions.V21
	LibraryCmd.PersistentFlags().VarP(&runtimeVersionEnumFlag, "runtime", "", fmt.Sprintf("Select runtime version %v", runtimeVersionValues))

	methodologyTypeValues := models.MethodologyTypeToArray()
	methodologyTypeEnumFlag.Value = models.MethodologyTypes.Basic
	LibraryCmd.PersistentFlags().VarP(&methodologyTypeEnumFlag, "methodology", "m", fmt.Sprintf("Select a methodology %v", methodologyTypeValues))

	designTypeValues := models.DesignTypeToArray()
	designTypeEnumFlag.Value = models.DesignType(models.DesignTypes.Classic)
	LibraryCmd.PersistentFlags().VarP(&designTypeEnumFlag, "design", "d", fmt.Sprintf("Select a design %v", designTypeValues))

	architectureTypeValues := models.ArchitectureTypeToArray()
	architectureTypeEnumFlag.Value = models.ArchitectureTypes.None
	LibraryCmd.PersistentFlags().VarP(&architectureTypeEnumFlag, "architecture", "a", fmt.Sprintf("Select a architecture %v", architectureTypeValues))

	templateTypeValues := models.TemplateTypeToArray()
	templateTypeEnumFlag.Value = models.TemplateTypes.Simple
	LibraryCmd.PersistentFlags().VarP(&templateTypeEnumFlag, "template", "t", fmt.Sprintf("Select a template type %v", templateTypeValues))
}
