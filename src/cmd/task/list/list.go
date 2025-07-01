package list

import (
	"fmt"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/components/workspace"
	_string "parsdevkit.net/core/utilities/string"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	Workspace string
}

var commandOptions ListOptions
var maxArgumentCount int = 0

var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List task(s)",
	Long:    `List task(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) > maxArgumentCount {
		return fmt.Errorf("There is no argument supported")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	checkGlobals := _string.IsEmpty(commandOptions.Workspace)
	taskService := services.NewCommonTaskService(utils.GetEnvironment())

	if checkGlobals {
		fmt.Println()
		fmt.Println("*** Global Tasks ***")
		fmt.Println()

		commandOptions.Workspace = "None"

		taskList, err := taskService.ListByWorkspace(commandOptions.Workspace)
		if err != nil {
			return fmt.Errorf("Failed to retrieve global tasks\n%w", err)
		}

		fmt.Printf("(%d) task available\n\n", len(*taskList))
		for _, task := range *taskList {
			fmt.Printf("- %v\n", task.Name)
		}

		commandOptions.Workspace = ""
	}

	fmt.Println()
	fmt.Println("*** Workspace Specific Tasks ***")
	fmt.Println()

	commandOptions.Workspace = workspace.GetActiveWorkspaceName(commandOptions.Workspace)

	taskList, err := taskService.ListByWorkspace(commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to list Active Workspace tasks\n%w", err)
	}

	fmt.Printf("(%d) task available\n\n", len(*taskList))
	for _, task := range *taskList {
		fmt.Printf("- %v\n", task.Name)
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = ListOptions{}
}
