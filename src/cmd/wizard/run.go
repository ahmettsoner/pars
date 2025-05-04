package wizard

import (
	"github.com/spf13/cobra"
)

var WizardCmd = &cobra.Command{
	Use:   "wizard",
	Short: "Project Wizard",
	Long:  `Project Wizard`,
	Run:   executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}


