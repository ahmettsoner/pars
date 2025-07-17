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

type UpdateProjectReference struct {
	flowx.BaseStep
	oldReference applicationProject.Reference
	newReference applicationProject.Reference
}

func NewUpdateProjectReference(old applicationProject.Reference, new applicationProject.Reference) *UpdateProjectReference {
	return &UpdateProjectReference{
		oldReference: old,
		newReference: new,
	}
}

func (s *UpdateProjectReference) Name() string {
	return fmt.Sprintf("UpdateProjectReference: %s", s.oldReference.Header.Name)
}

func (s *UpdateProjectReference) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

		if err := service.RemoveReferenceFromProject(project, s.oldReference); err != nil {
			return err
		}

		if err := service.AddReferenceToProject(project, s.newReference); err != nil {
			return err
		}
	}
	return nil
}

func (s *UpdateProjectReference) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
