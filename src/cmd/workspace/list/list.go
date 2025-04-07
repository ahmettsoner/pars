package list

import (
	"fmt"
	"log"

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
	Short: "List workspace project(s)",
	Long:  `List workspace project(s)`,
	Run:   executeFunc,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > maxArgumentCount {
			return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
		}
		return nil
	},
}

func executeFunc(cmd *cobra.Command, args []string) {
	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	workspaceList, err := workspaceService.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("(%d) workspace available\n", len(*workspaceList))

	activeWorkspace, err := workspaceService.GetActiveWorkspace()
	if err != nil {
		log.Fatal(err)
	}

	selectedWorkspace, err := workspaceService.GetSelectedWorkspace()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	if activeWorkspace != nil && selectedWorkspace != nil {
		if activeWorkspace.Name == selectedWorkspace.Name {
			fmt.Printf("* %v (active & selected)\n", activeWorkspace.Name)
		} else {
			fmt.Printf("* %v (active)\n", activeWorkspace.Name)
			fmt.Printf("%v (selected)\n", selectedWorkspace.Name)
		}
	} else if activeWorkspace != nil {
		fmt.Printf("* %v (active)\n", activeWorkspace.Name)
	} else if selectedWorkspace != nil {
		fmt.Printf("* %v (selected)\n", selectedWorkspace.Name)
	}

	for _, workspace := range *workspaceList {
		if (activeWorkspace == nil || activeWorkspace.Name != workspace.Name) &&
			(selectedWorkspace == nil || selectedWorkspace.Name != workspace.Name) {
			fmt.Println(workspace.Name)
		}
	}
}


