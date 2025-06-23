package browse

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

type BrowseOptions struct {
	URL string
}

var commandOptions BrowseOptions
var maxArgumentCount int = 1

var BrowseCmd = &cobra.Command{
	Use:     "browse",
	Aliases: []string{""},
	Short:   "Browse project(s)",
	Long:    `Browse project(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.URL) && len(args) == 0 {
		return fmt.Errorf("Please provide an access URL using the --url flag or as argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.URL) && len(args) > 0 {
		commandOptions.URL = args[0]
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {
	err := openBrowser(commandOptions.URL)
	if err != nil {
		return fmt.Errorf("❌ Failed to open URL: %v\n", err)
	}

	fmt.Fprintf(os.Stdout, "✔ URL opened: %s\n", commandOptions.URL)
	return nil
}

func init() {
	BrowseCmd.Flags().StringVarP(&commandOptions.URL, "url", "u", "", "Url")
}

func openBrowser(url string) error {
	if url == "" {
		return fmt.Errorf("URL is empty")
	}

	var openCmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		openCmd = exec.Command("open", url)
	case "windows":
		openCmd = exec.Command("cmd", "/c", "start", "", url)
	default:
		openCmd = exec.Command("xdg-open", url)
	}

	return openCmd.Start()
}
