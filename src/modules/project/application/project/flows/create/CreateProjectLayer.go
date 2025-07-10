package create

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type CreateProjectLayer struct {
	flowx.BaseStep
	layer applicationProject.Layer
}

func NewCreateProjectLayer(layer applicationProject.Layer) *CreateProjectLayer {
	return &CreateProjectLayer{
		layer: layer,
	}
}

func (s *CreateProjectLayer) Name() string {
	return fmt.Sprintf("CreateProjectLayer: %s", s.layer.Name)
}

func (s *CreateProjectLayer) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

		logrus.Debugf("trying to create %v", project.Header.Name)
		if _, err := service.CreateProjectFolder(project, s.layer.GetPathAsArray()...); err != nil {
			return err
		}
	}
	return nil
}

func (s *CreateProjectLayer) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
