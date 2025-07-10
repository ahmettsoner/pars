package remove

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
)

type DeleteWorkspace struct{ flowx.BaseStep }

func (s *DeleteWorkspace) Name() string { return "DeleteWorkspace" }

func (s *DeleteWorkspace) Run(ctx context.Context, fc *flowx.FlowContext) error {

	permanent, ok := flowx.Get[bool](fc, "permanent")
	if !ok {
		panic(fmt.Errorf("xxx: permanent parametresi hatalı tipte"))
	}

	if permanent {
		service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()

		workspace, ok := flowx.Get[basic_workspace_payload_structs.WorkspaceBaseStruct](fc, "workspace")
		if !ok {
			panic(fmt.Errorf("xxx: workspace parametresi hatalı tipte"))
		}

		logrus.Debugf("trying to create %v", workspace.Header.Name)
		if _, err := service.DeleteWorkspace(workspace); err != nil {
			return err
		}

	}
	return nil
}

func (s *DeleteWorkspace) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
