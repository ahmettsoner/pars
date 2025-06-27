package distribute

import (
	"github.com/spf13/cobra"
)

type DistributeOptions struct {
}

var commandOptions DistributeOptions

var DistributeCmd = &cobra.Command{
	Use:     "distribute",
	Aliases: []string{""},
	Short:   "Distribute project(s)",
	Long:    `Distribute project(s)`,
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
	cmd.Help()
	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = DistributeOptions{}
}
