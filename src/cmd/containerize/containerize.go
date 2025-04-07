package containerize

import (
	"github.com/spf13/cobra"
)

var ContainerizeCmd = &cobra.Command{
	Use:     "containerize",
	Aliases: []string{"c"},
	Short:   "Containerize project(s)",
	Long:    `Containerize project(s)`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}


