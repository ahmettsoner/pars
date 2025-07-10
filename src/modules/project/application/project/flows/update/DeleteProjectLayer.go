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

type DeleteProjectLayer struct {
	flowx.BaseStep
	layer applicationProject.Layer
}

func NewDeleteProjectLayer(layer applicationProject.Layer) *DeleteProjectLayer {
	return &DeleteProjectLayer{
		layer: layer,
	}
}

func (s *DeleteProjectLayer) Name() string {
	return fmt.Sprintf("DeleteProjectLayer: %s", s.layer.Name)
}

func (s *DeleteProjectLayer) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

		if err := service.DeleteLayerFolder(project, s.layer); err != nil {
			return err
		}
	}
	return nil
}

func (s *DeleteProjectLayer) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
