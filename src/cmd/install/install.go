package install

import (
	"fmt"
	"os"

	"parsdevkit.net/application"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/project/application_project_contract"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/components/workspace"

	"github.com/spf13/cobra"
)

type InstallOptions struct {
	Name      string
	Workspace string
}

var commandOptions InstallOptions
var maxArgumentCount int = 1

var InstallCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{""},
	Short:   "Install project(s)",
	Long:    `Install project(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: name is required. Provide it with '--name' or as an argument.")
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
		var workspaceName, err = workspace.GetActiveWorkspaceNameV2(appCtx, commandOptions.Workspace)
		if err != nil {
			return fmt.Errorf("failed to find active workspace '%s'\n%w", commandOptions.Name, err)
		}
		commandOptions.Workspace = workspaceName
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	_, err := projectService.Install(commandOptions.Name, commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to install project '%s' packages\n%w", commandOptions.Name, err)
	}

	fmt.Fprintf(os.Stdout, "✔ Project '%v' packages installed successfully\n", commandOptions.Name)

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = InstallOptions{}
}

func init() {
	InstallCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")

	InstallCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
	// RemoveCommand.Flags().StringVarP(&force, "force", "", "", "Force to delete")

	// if err := RemoveCommand.MarkFlagRequired("force"); err != nil {
	// 	fmt.Println(err)
	// }
}
