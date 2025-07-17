package update

import (
	"context"
	"fmt"

	"parsdevkit.net/application/ioc"
	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type DeleteProjectReference struct {
	flowx.BaseStep
	reference applicationProject.Reference
}

func NewDeleteProjectReference(reference applicationProject.Reference) *DeleteProjectReference {
	return &DeleteProjectReference{
		reference: reference,
	}
}

func (s *DeleteProjectReference) Name() string {
	return fmt.Sprintf("DeleteProjectReference: %s", s.reference.Header.Name)
}

func (s *DeleteProjectReference) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

		if err := service.RemoveReferenceFromProject(project, s.reference); err != nil {
			return err
		}
	}
	return nil
}

func (s *DeleteProjectReference) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
