package containerize

import (
	"github.com/spf13/cobra"
)

type ContainerizeOptions struct {
}

var commandOptions ContainerizeOptions

var ContainerizeCmd = &cobra.Command{
	Use:     "containerize",
	Aliases: []string{"c"},
	Short:   "Containerize project(s)",
	Long:    `Containerize project(s)`,
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
	commandOptions = ContainerizeOptions{}
}
