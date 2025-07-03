package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"parsdevkit.net/components/environment"
)

type EnvironmentListOptions struct {
}

var commandOptions EnvironmentListOptions

var ListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List environment project(s)",
	Long:    `List environment project(s)`,
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
	environmentService := environment.NewEnvironmentService()
	environmentlist, err := environmentService.List()
	if err != nil {
		return fmt.Errorf("Failed list environments\n%w", err)
	}

	fmt.Printf("(%d) environment available\n", (len(environmentlist) + 1))
	fmt.Println("* Default")
	for _, e := range environmentlist {
		fmt.Printf("- %v\n", e)
	}
	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = EnvironmentListOptions{}
}
