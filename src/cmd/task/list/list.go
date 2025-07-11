package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	basic_task_payload_structs "parsdevkit.net/modules/task/basic_task_payload/structs"
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

	var result []string = []string{
		basic_task_payload_structs.MODULE_KEY,
	}

	appCtx := application.GetContext()

	err := engines.DispatchEngineList(appCtx, result)
	if err != nil {
		return fmt.Errorf("Engine processing failed: %v", err)
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = ListOptions{}
}
