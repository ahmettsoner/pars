package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"parsdevkit.net/application/engines"

	"parsdevkit.net/application"
	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"
)

type GroupListOptions struct {
}

var commandOptions GroupListOptions
var maxArgumentCount int = 0

var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List group(s)",
	Long:    `List group(s)`,
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
	var result []string = []string{basic_group_payload_structs.MODULE_KEY}

	appCtx := application.GetContext()

	err := engines.DispatchEngineList(appCtx, result)
	if err != nil {
		return fmt.Errorf("Engine processing failed: %v", err)
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = GroupListOptions{}
}
