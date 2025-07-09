package create

import (
	"context"
	"fmt"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type RemoveProjectFiles struct{ flowx.BaseStep }

func (s *RemoveProjectFiles) Name() string { return "RemoveProjectFiles" }

func (s *RemoveProjectFiles) Run(ctx context.Context, fc *flowx.FlowContext) error {

	permanent, ok := flowx.Get[bool](fc, "permanent")
	if !ok {
		panic(fmt.Errorf("xxx: permanent parametresi hatalı tipte"))
	}

	if permanent {
		service := ioc.Get[application_project_contract.ProjectInterface]()
		project, ok := flowx.Get[application_project_payload_structs.ProjectBaseStruct](fc, "project")
		if !ok {
			panic(fmt.Errorf("xxx: project parametresi hatalı tipte"))
		}

		if _, err := service.RemoveProjectFiles(project); err != nil {
			return err
		}

	}
	return nil
}

func (s *RemoveProjectFiles) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
