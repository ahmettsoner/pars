package init

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

type InitOptions struct {
	Name string
	Path string
}

var commandOptions InitOptions
var maxArgumentCount int = 2
var workspaceService *services.WorkspaceService

var InitCmd = &cobra.Command{
	Use:               "init [name] [path]",
	Aliases:           []string{"i"},
	Short:             "Initialize new Pars workspace",
	Long:              `Create new workspace for Pars, that contains one or more project(s)`,
	Args:              validateArgs,
	PreRunE:           prepareFunc,
	RunE:              executeFunc,
	PostRun:           afterFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {

	if len(args) > maxArgumentCount {
		return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	workspaceService = services.NewWorkspaceService(utils.GetEnvironment())

	if len(args) > 0 {
		commandOptions.Name = args[0]
	}
	if len(args) > 1 {
		commandOptions.Path = args[1]
	}

	if utils.IsEmpty(commandOptions.Name) {
		commandOptions.Name = "workspace"
		existingDefaultNamedWorkspaces, err := workspaceService.ListByNameStartWith(commandOptions.Name)
		if err != nil {
			log.Fatal(err)
		}

		existingDefaultNamedWorkspaceCount := len(*existingDefaultNamedWorkspaces)
		if existingDefaultNamedWorkspaceCount > 0 {
			commandOptions.Name = fmt.Sprintf("%v_%d", commandOptions.Name, existingDefaultNamedWorkspaceCount)
		}
		// fmt.Printf("Default name (%v), \n", name)
	}
	if utils.IsEmpty(commandOptions.Path) {
		commandOptions.Path = commandOptions.Name
	}

	if !filepath.IsAbs(commandOptions.Path) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Current working directory not recognized:", err)
			os.Exit(1)
		}
		commandOptions.Path = filepath.Join(cwd, commandOptions.Path)
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	workspace, err := workspaceService.Save(workspace.NewWorkspaceBaseStruct(structs.NewHeader(structs.StructTypes.Workspace, commandOptions.Name, structs.Metadata{}), workspace.NewWorkspaceSpecification(0, commandOptions.Name, commandOptions.Path)))
	if err != nil {
		return fmt.Errorf("Failed to initialize workspace '%s'\n%w", commandOptions.Name, err)
	}

	fmt.Fprintf(os.Stdout, "✔ New workspace (%v) created at: %v\n", workspace.Specifications.Name, workspace.Specifications.Path)

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = InitOptions{}
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) == 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	if len(args) == 1 {
		dirs, _ := filepath.Glob(filepath.Join(toComplete, "*"))
		completions := []string{}

		for _, dir := range dirs {
			if info, err := os.Stat(dir); err == nil && info.IsDir() {
				dir, _ := GetLastComponent(dir)
				completions = append(completions, dir)
			}
		}
		return completions, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveDefault
	}

	return nil, cobra.ShellCompDirectiveNoFileComp
}
func GetLastComponent(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if info.IsDir() {
		// Get the last directory name
		return filepath.Base(filepath.Clean(path)), nil
	}

	return "", nil
}
