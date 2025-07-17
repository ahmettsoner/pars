package basic_workspace

import (
	"parsdevkit.net/components/template"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"

	"parsdevkit.net/application/contracts"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
)

type WorkspaceContextProvider struct {
	environment string
}

func NewWorkspaceContextProvider(environment string) basic_workspace_contract.ContextProviderInterface {
	return &WorkspaceContextProvider{
		environment: environment,
	}
}

func (s WorkspaceContextProvider) GetConfig() contracts.ContextProviderConfig {
	return contracts.ContextProviderConfig{
		Name: basic_workspace_payload_structs.MODULE_KEY,
	}
}
func (s *WorkspaceContextProvider) Context(source template.ContextProviderSource) interface{} {
	if model, ok := source.Workspace.(basic_workspace_payload_structs.WorkspaceBaseStruct); ok {
		return WorkspaceComposite{
			Workspace: s.structToModel(model),
			Original:  model,
		}
	}

	return nil
}
func (s *WorkspaceContextProvider) structToModel(model basic_workspace_payload_structs.WorkspaceBaseStruct) Workspace {

	var result Workspace = Workspace{
		Name: model.Header.Name,
	}

	return result
}

type WorkspaceComposite struct {
	Workspace
	Original basic_workspace_payload_structs.WorkspaceBaseStruct
}

type Workspace struct {
	Name string
}
