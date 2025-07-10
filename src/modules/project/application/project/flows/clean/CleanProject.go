package clean

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type CleanProject struct{ flowx.BaseStep }

func (s *CleanProject) Name() string { return "CleanProject" }

func (s *CleanProject) Run(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[application_project_contract.ProjectInterface]()

	project, ok := flowx.Get[application_project_payload_structs.ProjectBaseStruct](fc, "project")
	if !ok {
		panic(fmt.Errorf("xxx: project parametresi hatalı tipte"))
	}

	logrus.Debugf("trying to clean %v", project.Header.Name)
	if _, err := service.CleanV2(project.Header.Name, project.Specifications.Workspace); err != nil {
		return err
	}

	return nil
}

func (s *CleanProject) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
