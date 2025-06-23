package list

import (
	"fmt"
	"log"

	"parsdevkit.net/operation/services"

	"github.com/spf13/cobra"
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
}

func validateArgs(cmd *cobra.Command, args []string) error {
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {
	environmentService := services.NewEnvironmentService()
	environmentlist, err := environmentService.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("(%d) environment available\n", (len(environmentlist) + 1))

	fmt.Println("* Default")
	for _, e := range environmentlist {
		fmt.Printf("- %v\n", e)
	}
	return nil
}
