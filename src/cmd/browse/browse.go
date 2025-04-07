package browse

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

var (
	url string
)

var BrowseCmd = &cobra.Command{
	Use:     "browse",
	Aliases: []string{""},
	Short:   "Browse project(s)",
	Long:    `Browse project(s)`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	if utils.IsEmpty(url) {
		if len(args) == 0 {
			fmt.Println("Please provide a url")
			os.Exit(1)
		} else if len(args) > 0 {
			url = args[0]
		}
	}

	if utils.IsEmpty(url) {
		cmd.Help()
		os.Exit(0)
	}

	switch runtime.GOOS {
	case "darwin":
		exec.Command("open", url).Start()
	case "windows":
		exec.Command("cmd", "/c", "start", url).Start()
	default:
		exec.Command("xdg-open", url).Start()
	}

	fmt.Println("Url (" + url + ") opened")
}

func init() {
	BrowseCmd.Flags().StringVarP(&url, "url", "u", "", "Url")

}
