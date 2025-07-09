package create

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type CreateProjectLayers struct{ flowx.BaseStep }

func (s *CreateProjectLayers) Name() string { return "CreateProjectLayers" }

func (s *CreateProjectLayers) Run(ctx context.Context, fc *flowx.FlowContext) error {

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
		if _, err := service.CreateAllProjectFolders(project); err != nil {
			return err
		}
	}
	return nil
}

func (s *CreateProjectLayers) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
