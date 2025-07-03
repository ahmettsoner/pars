package list

import (
	"fmt"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/ioc"
	groupStructs "parsdevkit.net/modules/group/group/structs"

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
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {
	groupService := ioc.Get[contracts.GroupServiceInterface[groupStructs.GroupBaseStruct]]()
	groupList, err := groupService.List()
	if err != nil {
		return fmt.Errorf("Failed to retrieve groups\n%w", err)
	}

	fmt.Printf("(%d) group available\n\n", len(*groupList))
	for _, group := range *groupList {
		fmt.Printf("- %v\n", group.Header.Name)
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
}
