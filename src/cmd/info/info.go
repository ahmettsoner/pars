package info

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"parsdevkit.net/core/utilities"
	"parsdevkit.net/core/utils"
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
	fmt.Printf(textFormat, "Stage", utils.GetStage())
	fmt.Printf(textFormat, "Version", utils.GetVersion())
	fmt.Printf(textFormat, "Platform", utils.GetPlatform())
	fmt.Printf(textFormat, "OS", runtime.GOOS)
	fmt.Printf(textFormat, "Architecture", runtime.GOARCH)

	environment := utils.GetEnvironment()
	if utilities.IsEmpty(environment) {
		environment = "default"
	}
	fmt.Printf(textFormat, "Environment", environment)

	if utils.GetStage() == string(utils.StageTypes.None) {
		fmt.Printf(textFormat, "Codebase Path", utils.GetCodeBaseLocation())
	}
	fmt.Printf(textFormat, "Executable Path", utils.GetExecutableLocation())
	fmt.Printf(textFormat, "Config Directory", utils.GetConfigLocation())
	fmt.Printf(textFormat, "Data Directory", utils.GetDataLocation())
}
