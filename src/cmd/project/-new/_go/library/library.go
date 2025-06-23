package library

import (
	"fmt"
	"log"
	"os"

	"parsdevkit.net/models"
	"parsdevkit.net/structs/project"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/engines"

	"github.com/spf13/cobra"
)

var (
	noInit                   bool = true
	name                     string
	workspaceName            string
	projectSet               string
	_package                 string
	platformVersionEnumFlag  models.GoPlatformVersionEnumFlag
	runtimeVersionEnumFlag   models.GoRuntimeVersionEnumFlag
	methodologyTypeEnumFlag  models.MethodologyTypeEnumFlag
	designTypeEnumFlag       models.DesignTypeEnumFlag
	architectureTypeEnumFlag models.ArchitectureTypeEnumFlag
	templateTypeEnumFlag     models.TemplateTypeEnumFlag
)

var LibraryCmd = &cobra.Command{
	Use:     "library",
	Aliases: []string{"l"},
	Short:   "Initialize new project",
	Long:    `Initialize new project`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	if utils.IsEmpty(name) {
		if len(args) == 0 {
			fmt.Println("Please provide a name for the new project")
			os.Exit(1)
		} else if len(args) > 0 {
			name = args[0]
		}
	}

	if utils.IsEmpty(name) {
		cmd.Help()
		os.Exit(0)
	}

	workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	if workspaceID == 0 {
		log.Fatal("There are no active workspace, please initialize or switch to available workspace")
	}

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
		PlatformVersion models.GoPlatformVersion
		RuntimeVersion  models.GoRuntimeVersion
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

	var templateFilePath = "/go/projects/library.yaml.templ"
	if architectureTypeEnumFlag.Value == models.ArchitectureTypes.None {
		if methodologyTypeEnumFlag.Value == models.MethodologyTypes.Basic {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/go/projects/library.yaml.templ"
			}
		} else if methodologyTypeEnumFlag.Value == models.MethodologyTypes.Layered {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/go/projects/library-layered.yaml.templ"
			}
		} else if methodologyTypeEnumFlag.Value == models.MethodologyTypes.NTier {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/go/projects/library-ntier.yaml.templ"
			}
		}
	} else if architectureTypeEnumFlag.Value == models.ArchitectureTypes.Clean {
		if methodologyTypeEnumFlag.Value == models.MethodologyTypes.Basic {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/go/projects/library-clean.yaml.templ"
			}
		} else if methodologyTypeEnumFlag.Value == models.MethodologyTypes.Layered {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/go/projects/library-layered-clean.yaml.templ"
			}
		} else if methodologyTypeEnumFlag.Value == models.MethodologyTypes.NTier {
			if designTypeEnumFlag.Value == models.DesignTypes.Classic {
				templateFilePath = "/go/projects/library-ntier-clean.yaml.templ"
			}
		}
	}

	engine := engines.Engine{}
	if err := engine.CreateProjectsFromTemplate(!noInit, structData, templateFilePath); err != nil {
		log.Fatal(err)
	}
}

func init() {
	LibraryCmd.Flags().BoolVarP(&noInit, "no-init", "", false, "Create project but do not initialize")

	LibraryCmd.Flags().StringVarP(&name, "name", "n", "", "Project name")

	LibraryCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
	LibraryCmd.Flags().StringVarP(&projectSet, "project-set", "s", "", "Project Set")
	LibraryCmd.Flags().StringVarP(&_package, "package", "p", "", "Package")

	platformVersionValues := models.GoPlatformVersionToArray()
	platformVersionEnumFlag.Value = models.GoPlatformVersions.Go121
	LibraryCmd.PersistentFlags().VarP(&platformVersionEnumFlag, "platform", "", fmt.Sprintf("Select platform version %v", platformVersionValues))

	runtimeVersionValues := models.GoRuntimeVersionToArray()
	runtimeVersionEnumFlag.Value = models.GoRuntimeVersions.Go121
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

	validEnumValues := models.TemplateTypeToArray()
	templateTypeEnumFlag.Value = models.TemplateTypes.Simple
	LibraryCmd.PersistentFlags().VarP(&templateTypeEnumFlag, "template", "t", fmt.Sprintf("Select a template type %v", validEnumValues))
}
