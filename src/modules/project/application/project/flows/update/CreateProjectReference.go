package update

import (
	"context"
	"fmt"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type CreateProjectReference struct {
	flowx.BaseStep
	reference application_project_payload_structs.ProjectBaseStruct
}

func NewCreateProjectReference(reference application_project_payload_structs.ProjectBaseStruct) *CreateProjectReference {
	return &CreateProjectReference{
		reference: reference,
	}
}

func (s *CreateProjectReference) Name() string {
	return fmt.Sprintf("CreateProjectReference: %s", s.reference.Header.Name)
}

func (s *CreateProjectReference) Run(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[application_project_contract.ProjectInterface]()

	init, ok := flowx.Get[bool](fc, "init")
	if !ok {
		panic(fmt.Errorf("xxx: init parametresi hatalı tipte"))
	}

	if init {
		project, ok := flowx.Get[application_project_payload_structs.ProjectBaseStruct](fc, "project")
		if !ok {
			panic(fmt.Errorf("xxx: project parametresi hatalı tipte"))
		}

		if err := service.AddReferenceToProject(project, s.reference); err != nil {
			return err
		}
	}
	return nil
}

func (s *CreateProjectReference) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
