package remove

import (
	"fmt"
	"os"

	"parsdevkit.net/components/workspace"
	_string "parsdevkit.net/pkg/utilities/string"
	basic_task_payload_structs "parsdevkit.net/modules/task/basic_task_payload/structs"

	"parsdevkit.net/application"

	"github.com/spf13/cobra"
)

type RemoveOptions struct {
	Names     []string
	Workspace string
	Force     string
}

var commandOptions RemoveOptions

var RemoveCmd = &cobra.Command{
	Use:               "remove name [name]...",
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


	var result []string = []string{basic_task_payload_structs.MODULE_KEY}

	for _, name := range commandOptions.Names {
		appCtx := application.GetContext()

		err := engines.DispatchEngineRemove(appCtx, result, name)
		if err != nil {
			return fmt.Errorf("Engine processing failed: %v", err)
		}
	}

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
