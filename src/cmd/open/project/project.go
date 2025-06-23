package project

import (
	"fmt"
	"log"

	"parsdevkit.net/structs/project"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/providers"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/repositories"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ProjectOptions struct {
	Name      string
	Workspace string
}

var commandOptions ProjectOptions
var maxArgumentCount int = 1

var ProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Project Information",
	Long:    `Project Information`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: project name is required. Provide it with '--name' or as an argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("error: too many arguments. Only project name is expected.")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}

	if utils.IsEmpty(commandOptions.Workspace) {
		commandOptions.Workspace = parsCMDCommon.GetActiveWorkspaceName("")
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	projectGroup, projectName, err := project.ParseProjectFullName(commandOptions.Name)

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
		projectEntities, err := projectService.ListByFullNameWorkspace(fmt.Sprintf("%v/", projectGroup), commandOptions.Workspace)
		if err != nil {
			log.Fatal(err)
		}

		if len(*projectEntities) > 0 {
			providers.VSCodeExecute("", (*projectEntities)[0].Specifications.GetAbsoluteGroupPath())
		}
	} else {
		project, err := projectService.GetByFullNameWorkspace(commandOptions.Name, commandOptions.Workspace)
		if err != nil {
			log.Fatal(err)
		}

		if groupId == 0 {
			providers.VSCodeExecute("", project.Specifications.GetAbsoluteProjectPath())
		} else {
			providers.VSCodeExecute("", project.Specifications.GetAbsoluteGroupPath())
		}
	}

	return nil

}

func init() {
	addSubCommands()
}

func addSubCommands() {
	ProjectCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")

	ProjectCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
}
