package remove

import (
	"fmt"
	"os"

	"parsdevkit.net/components/workspace"
	"parsdevkit.net/engines/commonTask"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

type RemoveOptions struct {
	Names     []string
	Workspace string
	Force     string
}

var commandOptions RemoveOptions

var RemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"r"},
	Short:   "Task Information",
	Long:    `Task Information`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {

	if len(commandOptions.Names) == 0 && len(args) == 0 &&  {
		return fmt.Errorf("error: task name is required.")
	}

	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if len(commandOptions.Names) == 0 && len(args) > 0 {
		commandOptions.Names = args
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {
		checkGlobals := utilities.IsEmpty(commandOptions.Workspace)
		taskService := services.NewCommonTaskService(utils.GetEnvironment())

		for _, name := range commandOptions.Names {
			if checkGlobals {
				commandOptions.Workspace = "None"

				if taskService.IsExists(name, commandOptions.Workspace) {
					task, err := taskService.Remove(name, commandOptions.Workspace, true)
					if err != nil {
						return fmt.Errorf("Failed to remove task(s) '%s'\n%w", commandOptions.Names[], err)
					}
					fmt.Println("Task (" + task.Name + ") deleted permanently")
				}

				commandOptions.Workspace = ""
			}

			commandOptions.Workspace = workspace.GetActiveWorkspaceName(commandOptions.Workspace)

			if taskService.IsExists(name, commandOptions.Workspace) {
				task, err := taskService.Remove(name, commandOptions.Workspace, true)
				if err != nil {
					return fmt.Errorf("Failed to remove task(s) '%s'\n%w", name, err)
				}
				fmt.Println("Task (" + task.Name + ") deleted permanently")
			}
		}
		fmt.Fprintf(os.Stdout, "✔ task(s) '%v' removed successfully\n", commandOptions.Names)

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = RemoveOptions{}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	RemoveCmd.Flags().StringSliceVarP(&commandOptions.Names, "name", "n", nil, "Template names")
}
