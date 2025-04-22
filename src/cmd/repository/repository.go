package repository

import (
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "repository",
	Short: "Version Control options for a project",
	Long:  `Version Control options for a project`,
	Run:   executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}


