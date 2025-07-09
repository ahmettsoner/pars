package update

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type UpdateProject struct{ flowx.BaseStep }

func (s *UpdateProject) Name() string { return "UpdateProject" }

func (s *UpdateProject) Run(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[application_project_contract.ProjectInterface]()

	project, ok := flowx.Get[application_project_payload_structs.ProjectBaseStruct](fc, "project")
	if !ok {
		panic(fmt.Errorf("xxx: project parametresi hatalı tipte"))
	}

	logrus.Debugf("trying to create %v", project.Header.Name)
	if _, err := service.SaveProject(project); err != nil {
		return err
	}

	return nil
}

func (s *UpdateProject) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
