package project

import (
	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/ioc"
	applicationproject "parsdevkit.net/structs/project/application-project"

	"fmt"
	"log"
	"os"

	"parsdevkit.net/structs/project"

	_string "parsdevkit.net/core/utilities/string"

	"parsdevkit.net/providers"

	"parsdevkit.net/application"
	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/contexts"
	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/components/workspace"

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
		appCtx := application.GetContext()
		if appCtx == nil {
			return fmt.Errorf("xxx: Current workspace bulunamadı")
		}
		commandOptions.Workspace = workspace.GetActiveWorkspaceName(appCtx, "")
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	projectGroup, projectName, err := project.ParseProjectFullName(commandOptions.Name)

	dbContext := contexts.NewDbContext(utils.GetEnvironment())
	groupRespository := repositories.NewGroupRepository(dbContext)
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

	projectService := ioc.Get[contracts.ProjectServiceInterface[applicationproject.ProjectBaseStruct]]()
	if _string.IsEmpty(projectName) && groupId > 0 {
		projectEntities, err := projectService.ListByFullNameWorkspace(fmt.Sprintf("%v/", projectGroup), commandOptions.Workspace)
		if err != nil {
			return fmt.Errorf("Failed to open project '%s'\n%w", commandOptions.Name, err)
		}

		if len(*projectEntities) > 0 {
			providers.VSCodeExecute("", (*projectEntities)[0].Specifications.GetAbsoluteGroupPath())
		}
	} else {
		project, err := projectService.GetByFullNameWorkspace(commandOptions.Name, commandOptions.Workspace)
		if err != nil {
			return fmt.Errorf("Failed to open project '%s'\n%w", commandOptions.Name, err)
		}

		if groupId == 0 {
			providers.VSCodeExecute("", project.Specifications.GetAbsoluteProjectPath())
		} else {
			providers.VSCodeExecute("", project.Specifications.GetAbsoluteGroupPath())
		}
	}

	fmt.Fprintf(os.Stdout, "✔ Project '%v' opened successfully\n", commandOptions.Name)
	return nil

}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = ProjectOptions{}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	ProjectCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")

	ProjectCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
}
