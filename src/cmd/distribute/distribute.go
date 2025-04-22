package distribute

import (
	"github.com/spf13/cobra"
)

var DistributeCmd = &cobra.Command{
	Use:     "distribute",
	Aliases: []string{""},
	Short:   "Distribute project(s)",
	Long:    `Distribute project(s)`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}


