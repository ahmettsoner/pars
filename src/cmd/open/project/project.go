package project

import (
	"fmt"
	"log"
	"os"

	"parsdevkit.net/structs/project"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/providers"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/repositories"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	name          string
	workspaceName string
	force         string
)

var ProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Project Information",
	Long:    `Project Information`,
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

	groupRespository := repositories.NewGroupRepository(utils.GetEnvironment())
	groupId := 0
	projectGroupEntity, err := groupRespository.GetByName(projectGroup)
	if err != nil {
		log.Fatal(err)
	}
	if projectGroupEntity != nil {
		groupId = projectGroupEntity.ID
	}

	if groupId > 0 {
		logrus.Debugf("project (%v) in the group (%v)", projectName, projectGroup)
	}

	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	if utils.IsEmpty(projectName) && groupId > 0 {
		projectEntities, err := projectService.ListByFullNameWorkspace(fmt.Sprintf("%v/", projectGroup), workspaceName)
		if err != nil {
			log.Fatal(err)
		}

		if len(*projectEntities) > 0 {
			providers.VSCodeExecute("", (*projectEntities)[0].Specifications.GetAbsoluteGroupPath())
		}
	} else {
		project, err := projectService.GetByFullNameWorkspace(name, workspaceName)
		if err != nil {
			log.Fatal(err)
		}

		if groupId == 0 {
			providers.VSCodeExecute("", project.Specifications.GetAbsoluteProjectPath())
		} else {
			providers.VSCodeExecute("", project.Specifications.GetAbsoluteGroupPath())
		}
	}

}

func init() {
	addSubCommands()
}

func addSubCommands() {
	ProjectCmd.Flags().StringVarP(&name, "name", "n", "", "Project name")

	ProjectCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
}
