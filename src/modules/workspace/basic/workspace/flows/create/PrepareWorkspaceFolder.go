package create

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
)

type PrepareWorkspaceFolder struct{ flowx.BaseStep }

func (s *PrepareWorkspaceFolder) Name() string { return "PrepareWorkspaceFolder" }

func (s *PrepareWorkspaceFolder) Run(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()

	workspace, ok := flowx.Get[basic_workspace_payload_structs.WorkspaceBaseStruct](fc, "workspace")
	if !ok {
		panic(fmt.Errorf("xxx: workspace parametresi hatalı tipte"))
	}

	logrus.Debugf("trying to create %v", workspace.Header.Name)
	if _, err := service.CreateWorkspaceFolder(workspace); err != nil {
		return err
	}

	return nil
}

func (s *PrepareWorkspaceFolder) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
