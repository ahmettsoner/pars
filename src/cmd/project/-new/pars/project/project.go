package project

import (
	"fmt"
	"log"
	"os"

	"parsdevkit.net/engines/applicationProject"
	"parsdevkit.net/engines/group"
	"parsdevkit.net/structs/project"

	parsModels "parsdevkit.net/platforms/pars/models"

	"parsdevkit.net/core/utils"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/spf13/cobra"
)

var (
	noInit                  bool = true
	name                    string
	workspaceName           string
	projectSet              string
	_package                string
	platformVersionEnumFlag parsModels.ParsPlatformVersionEnumFlag
)

var ProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
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
		PlatformVersion parsModels.ParsPlatformVersion
	}{
		Group:           projectGroup,
		Name:            projectName,
		Set:             projectSet,
		Package:         _package,
		Path:            projectName,
		Workspace:       workspaceName,
		PlatformVersion: platformVersionEnumFlag.Value,
	}

	var templateFilePath = "/pars/projects/project.yaml.templ"

	if !utils.IsEmpty(projectGroup) {
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
	projectService := applicationProject.ApplicationProjectEngine{}
	if err := projectService.CreateProjectsFromTemplate(!noInit, structData, templateFilePath); err != nil {
		log.Fatal(err)
	}
}

func init() {
	ProjectCmd.Flags().BoolVarP(&noInit, "no-init", "", false, "Create project but do not initialize")

	ProjectCmd.Flags().StringVarP(&name, "name", "n", "", "Project name")

	platformVersionValues := parsModels.ParsPlatformVersionToArray()
	platformVersionEnumFlag.Value = parsModels.ParsPlatformVersions.BetaV1
	ProjectCmd.PersistentFlags().VarP(&platformVersionEnumFlag, "platform", "", fmt.Sprintf("Select platform version %v", platformVersionValues))

	ProjectCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
	ProjectCmd.Flags().StringVarP(&projectSet, "project-set", "s", "", "Project Set")
	ProjectCmd.Flags().StringVarP(&_package, "package", "p", "", "Package")

}
