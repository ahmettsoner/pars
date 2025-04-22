package list

import (
	"fmt"
	"log"

	"parsdevkit.net/operation/services"

	"github.com/spf13/cobra"
)

var ListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List environment project(s)",
	Long:    `List environment project(s)`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
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
}


