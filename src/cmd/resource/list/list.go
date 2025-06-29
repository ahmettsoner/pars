package list

import (
	"fmt"

	"parsdevkit.net/operation/services"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
	"parsdevkit.net/application"
)

type ListOptions struct {
	Workspace string
}

var commandOptions ListOptions
var maxArgumentCount int = 0

var ListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List resource(s)",
	Long:    `List resource(s)`,
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

	checkGlobals := utils.IsEmpty(commandOptions.Workspace)
	objectResourceService := services.NewObjectResourceService(utils.GetEnvironment())
	dataResourceService := services.NewDataResourceService(utils.GetEnvironment())

	if checkGlobals {
		fmt.Println("*** Global Resources ***")
		fmt.Println()

		commandOptions.Workspace = "None"

		objectResourceList, err := objectResourceService.ListByWorkspace(commandOptions.Workspace)
		if err != nil {
			return fmt.Errorf("Failed to retrieve global Object resources\n%w", err)
		}

		fmt.Printf("(%d) object resource available\n\n", len(*objectResourceList))
		for _, resource := range *objectResourceList {
			fmt.Printf("- %v\n", resource.GetFullInformation())
		}

		fmt.Println()
		fmt.Println("--------------------------")
		fmt.Println()

		dataResourceList, err := dataResourceService.ListByWorkspace(commandOptions.Workspace)
		if err != nil {
			return fmt.Errorf("Failed to list global Data resources\n%w", err)
		}

		fmt.Printf("(%d) data resource available\n\n", len(*dataResourceList))
		for _, resource := range *dataResourceList {
			fmt.Printf("- %v\n", resource.GetFullInformation())
		}

		commandOptions.Workspace = ""
		fmt.Println()
	}

	fmt.Println("*** Workspace Specific Resources ***")
	fmt.Println()

	appCtx := application.GetContext()
	if appCtx == nil {
		return fmt.Errorf("xxx: Current workspace bulunamadı")
	}
	commandOptions.Workspace = parsCMDCommon.GetActiveWorkspaceName(appCtx, commandOptions.Workspace)

	objectResourceList, err := objectResourceService.ListByWorkspace(commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to list Active Workspace Object resources\n%w", err)
	}

	fmt.Printf("(%d) object resource available\n\n", len(*objectResourceList))
	for _, resource := range *objectResourceList {
		fmt.Printf("- %v\n", resource.GetFullInformation())
	}

	fmt.Println()
	fmt.Println("--------------------------")
	fmt.Println()

	dataResourceList, err := dataResourceService.ListByWorkspace(commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to list Active Workspace Data resources\n%w", err)
	}

	fmt.Printf("(%d) data resource available\n\n", len(*dataResourceList))
	for _, resource := range *dataResourceList {
		fmt.Printf("- %v\n", resource.GetFullInformation())
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = ListOptions{}
}
