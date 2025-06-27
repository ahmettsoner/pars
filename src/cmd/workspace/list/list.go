package list

import (
	"fmt"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

var maxArgumentCount int = 0

var ListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Example: `  pars workspace list [flags]
  pars wl [flags]`,
	Short:   "List workspace project(s)",
	Long:    `List workspace project(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) > maxArgumentCount {
		return fmt.Errorf("There is no argument supported")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {
	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	workspaceList, err := workspaceService.List()
	if err != nil {
		return fmt.Errorf("Failed to retrieve workspace\n%w", err)
	}

	fmt.Printf("(%d) workspace available\n", len(*workspaceList))

	activeWorkspace, err := workspaceService.GetActiveWorkspace()
	if err != nil {
		return fmt.Errorf("Failed to find Active Workspace\n%w", err)
	}

	selectedWorkspace, err := workspaceService.GetSelectedWorkspace()
	if err != nil {
		return fmt.Errorf("Failed to find Selected Workspace\n%w", err)
	}

	fmt.Println()
	if activeWorkspace != nil && selectedWorkspace != nil {
		if activeWorkspace.Header.Name == selectedWorkspace.Header.Name {
			fmt.Printf("* %v (active & selected)\n", activeWorkspace.Header.Name)
		} else {
			fmt.Printf("* %v (active)\n", activeWorkspace.Header.Name)
			fmt.Printf("%v (selected)\n", selectedWorkspace.Header.Name)
		}
	} else if activeWorkspace != nil {
		fmt.Printf("* %v (active)\n", activeWorkspace.Header.Name)
	} else if selectedWorkspace != nil {
		fmt.Printf("* %v (selected)\n", selectedWorkspace.Header.Name)
	}

	for _, workspace := range *workspaceList {
		if (activeWorkspace == nil || activeWorkspace.Header.Name != workspace.Header.Name) &&
			(selectedWorkspace == nil || selectedWorkspace.Header.Name != workspace.Header.Name) {
			fmt.Println(workspace.Header.Name)
		}
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
}
