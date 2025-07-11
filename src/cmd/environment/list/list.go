package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"parsdevkit.net/application/engines"

	"parsdevkit.net/application"
	basic_environment_payload_structs "parsdevkit.net/modules/environment/basic_environment_payload/structs"
)

type EnvironmentListOptions struct {
}

var commandOptions EnvironmentListOptions

var ListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List environment(s)",
	Long:    `List environment(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	var result []string = []string{basic_environment_payload_structs.MODULE_KEY}

	appCtx := application.GetContext()

	err := engines.DispatchEngineList(appCtx, result)
	if err != nil {
		return fmt.Errorf("Engine processing failed: %v", err)
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = EnvironmentListOptions{}
}
