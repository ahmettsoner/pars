package project

import (
	"github.com/spf13/cobra"
)

type ProjectOptions struct {
	Name string
}

var commandOptions ProjectOptions

var ProjectCommand = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Edit project(s)",
	Long:    `Edit project(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {
	return nil
}

func init() {
	ProjectCommand.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")
}
