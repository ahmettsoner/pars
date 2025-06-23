package list

import (
	"fmt"
	"log"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List group(s)",
	Long:    `List group(s)`,
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
	groupService := services.NewGroupService(utils.GetEnvironment())
	groupList, err := groupService.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("(%d) group available\n\n", len(*groupList))
	for _, group := range *groupList {
		fmt.Printf("- %v\n", group.Name)
	}

	return nil
}
