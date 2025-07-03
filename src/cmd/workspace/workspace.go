package workspace

import (
	"fmt"
	"log"
	"os"

	"parsdevkit.net/cmd/workspace/describe"
	"parsdevkit.net/cmd/workspace/list"
	"parsdevkit.net/cmd/workspace/remove"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/ioc"

	"github.com/spf13/cobra"
)

type WorkspaceOptions struct {
	SwitchTo string
}

var commandOptions WorkspaceOptions

var WorkspaceCmd = &cobra.Command{
	Use:     "workspace",
	Aliases: []string{"w"},
	Short:   "Workspace information",
	Long:    `Workspace information`,
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
	if !_string.IsEmpty(commandOptions.SwitchTo) {

		workspaceService := ioc.Get[contracts.WorkspaceServiceInterface[workspace.WorkspaceBaseStruct]]()
		_, err := workspaceService.ChangeCurrentWorkspace(commandOptions.SwitchTo)
		if err != nil {
			return fmt.Errorf("Failed to switch '%s'\n%w", commandOptions.SwitchTo, err)
		}

		fmt.Fprintf(os.Stdout, "✔ Swithched to: '%v' removed successfully\n", commandOptions.SwitchTo)
	} else {
		cmd.Help()
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = WorkspaceOptions{}
}

func init() {
	addSubCommands()

	WorkspaceCmd.Flags().StringVarP(&commandOptions.SwitchTo, "switch", "s", "", "Switch to workspace")
	WorkspaceCmd.RegisterFlagCompletionFunc("switch", switchFlagCompletion)
}

func addSubCommands() {
	WorkspaceCmd.AddCommand(list.ListCommand)
	WorkspaceCmd.AddCommand(describe.DescribeCmd)
	WorkspaceCmd.AddCommand(remove.RemoveCmd)
}

func switchFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	workspaceService := ioc.Get[contracts.WorkspaceServiceInterface[workspace.WorkspaceBaseStruct]]()
	workspaceList, err := workspaceService.List()
	if err != nil {
		log.Fatal(err)
	}

	var workspaces = make([]string, 0)
	for _, workspace := range *workspaceList {
		workspaces = append(workspaces, workspace.Header.Name)
	}

	return workspaces, cobra.ShellCompDirectiveNoFileComp
}
