package generate

import (
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generate a component",
	Long:    `Geenrate a component`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}
