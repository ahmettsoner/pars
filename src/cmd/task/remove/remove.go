package remove

import (
	"fmt"
	"log"
	"os"

	parsCMDCommon "parsdevkit.net/core/cmd"
	"parsdevkit.net/engines/commonTask"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

type RemoveOptions struct {
	Names     []string
	Workspace string
	Force     string
	FilePaths []string
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
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(commandOptions.Names) == 0 && len(args) == 0 {
		return fmt.Errorf("error: project name is required. Provide it with '--name' or as an argument.")
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
	if len(commandOptions.FilePaths) > 0 {
		taskService := commonTask.CommonTaskEngine{}
		if err := taskService.RemoveTasksFromFile(true, commandOptions.FilePaths...); err != nil {
			log.Fatal(err)
		}
	} else {

		if len(commandOptions.Names) == 0 {
			if len(args) == 0 {
				fmt.Println("Please provide a name for the remove the remove")
				os.Exit(1)
			} else if len(args) > 0 {
				commandOptions.Names = args
			}
		}

		checkGlobals := utils.IsEmpty(commandOptions.Workspace)
		taskService := services.NewCommonTaskService(utils.GetEnvironment())

		for _, name := range commandOptions.Names {
			if checkGlobals {
				commandOptions.Workspace = "None"

				if taskService.IsExists(name, commandOptions.Workspace) {
					task, err := taskService.Remove(name, commandOptions.Workspace, true)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Task (" + task.Name + ") deleted permanently")
				}

				commandOptions.Workspace = ""
			}

			commandOptions.Workspace = parsCMDCommon.GetActiveWorkspaceName(commandOptions.Workspace)

			if taskService.IsExists(name, commandOptions.Workspace) {
				task, err := taskService.Remove(name, commandOptions.Workspace, true)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Task (" + task.Name + ") deleted permanently")
			}
		}
	}

	return nil
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	RemoveCmd.Flags().StringSliceVarP(&commandOptions.Names, "name", "n", nil, "Template names")

	RemoveCmd.Flags().StringSliceVarP(&commandOptions.FilePaths, "files", "f", nil, "Comma-separated list of declaration files")
}
