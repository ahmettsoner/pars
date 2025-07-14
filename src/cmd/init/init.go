package init

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"parsdevkit.net/application/engines"

	"parsdevkit.net/application/schemas"

	"parsdevkit.net/pkg/utilities/json"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/application/ioc"

	"github.com/spf13/cobra"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"

	"parsdevkit.net/application"
)

type InitOptions struct {
	Name string
	Path string
}

var commandOptions InitOptions
var maxArgumentCount int = 2
var workspaceService basic_workspace_contract.WorkspaceInterface

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
	workspaceService = ioc.Get[basic_workspace_contract.WorkspaceInterface]()

	if len(args) > 0 {
		commandOptions.Name = args[0]
	}
	if len(args) > 1 {
		commandOptions.Path = args[1]
	}

	if _string.IsEmpty(commandOptions.Name) {
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
	if _string.IsEmpty(commandOptions.Path) {
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

	var result []schemas.SchemaInterface = make([]schemas.SchemaInterface, 0)

	result = append(result, basic_workspace_payload_structs.NewWorkspaceBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Workspace,
			"",
			commandOptions.Name,
			schemas.Metadata{},
		),
		basic_workspace_payload_structs.NewWorkspaceSpecification(
			0,
			commandOptions.Name,
			commandOptions.Path,
		)))

	var loadedSchemas []string = make([]string, 0)
	for _, data := range result {

		if err := data.Validate(); err != nil {
			jsonObject, _ := json.ToJson(data)
			return fmt.Errorf("invalid data: '%s'\n%w", jsonObject, err)
		}

		loadedSchemas = append(loadedSchemas, fmt.Sprintf("%s.%s", data.GetHeader().Name, data.GetKey()))
	}
	fmt.Printf("Loaded Schemas: %s\n", _string.Concat(", ", loadedSchemas...))
	fmt.Printf("────────────────────────────────────\n")

	appCtx := application.GetContext()
	if appCtx == nil {
		return fmt.Errorf("xxx: Current workspace bulunamadı")
	}
	err := engines.DispatchEngineInit(appCtx, result)
	if err != nil {
		return fmt.Errorf("Engine processing failed: %v", err)
	}

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
