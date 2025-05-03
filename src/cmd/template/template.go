package template

import (
	"parsdevkit.net/cmd/template/list"
	"parsdevkit.net/cmd/template/remove"
	"parsdevkit.net/cmd/template/submit"

	"github.com/spf13/cobra"
)

var TemplateCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"t"},
	Short:   "Template Information",
	Long:    `Template Information`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	TemplateCmd.AddCommand(submit.SubmitCmd)
	TemplateCmd.AddCommand(remove.RemoveCmd)
	TemplateCmd.AddCommand(list.ListCmd)
}
