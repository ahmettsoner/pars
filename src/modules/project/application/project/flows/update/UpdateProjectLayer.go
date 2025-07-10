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

type UpdateProjectLayer struct {
	flowx.BaseStep
	oldLayer applicationProject.Layer
	newLayer applicationProject.Layer
}

func NewUpdateProjectLayer(old applicationProject.Layer, new applicationProject.Layer) *UpdateProjectLayer {
	return &UpdateProjectLayer{
		oldLayer: old,
		newLayer: new,
	}
}

func (s *UpdateProjectLayer) Name() string {
	return fmt.Sprintf("UpdateProjectLayer: %s", s.oldLayer.Name)
}

func (s *UpdateProjectLayer) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

		if err := service.DeleteLayerFolder(project, s.oldLayer); err != nil {
			return err
		}

		if err := service.CreateLayerFolder(project, s.newLayer); err != nil {
			return err
		}
	}
	return nil
}

func (s *UpdateProjectLayer) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
