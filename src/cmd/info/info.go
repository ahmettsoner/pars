package info

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"parsdevkit.net/application"
	_string "parsdevkit.net/pkg/utilities/string"
)

var InfoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{""},
	Short:   "About Pars",
	Long:    `About Pars`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	textFormat := "%-20s: %v\n"
	fmt.Println("New generation SDK")
	fmt.Printf(textFormat, "Stage", application.GetStage())
	fmt.Printf(textFormat, "Version", application.GetVersion())
	fmt.Printf(textFormat, "Platform", application.GetPlatform())
	fmt.Printf(textFormat, "OS", runtime.GOOS)
	fmt.Printf(textFormat, "Architecture", runtime.GOARCH)

	environment := application.GetEnvironment()
	if _string.IsEmpty(environment) {
		environment = "default"
	}
	fmt.Printf(textFormat, "Environment", environment)

	if application.GetStage() == string(application.StageTypes.None) {
		fmt.Printf(textFormat, "Codebase Path", application.GetCodeBaseLocation())
	}
	fmt.Printf(textFormat, "Executable Path", application.GetExecutableLocation())
	fmt.Printf(textFormat, "Config Directory", application.GetConfigLocation())
	fmt.Printf(textFormat, "Data Directory", application.GetDataLocation())
}
