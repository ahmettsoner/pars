package project

import (
	"parsdevkit.net/providers"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/spf13/cobra"
)

var (
	workspaceName string
)

var WorkspaceCommand = &cobra.Command{
	Use:     "workspace",
	Aliases: []string{"w"},
	Short:   "Workspace workspace",
	Long:    `Workspace workspace`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {

	path := parsCMDCommon.GetActiveWorkspacePath(workspaceName)

	providers.VSCodeExecute("", path)
}

func init() {
	WorkspaceCommand.Flags().StringVarP(&workspaceName, "name", "n", "", "Workspace name")
	// WorkspaceCommand.Flags().StringVarP(&force, "force", "", "", "Force to delete")

	// if err := WorkspaceCommand.MarkFlagRequired("force"); err != nil {
	// 	fmt.Println(err)
	// }
}
