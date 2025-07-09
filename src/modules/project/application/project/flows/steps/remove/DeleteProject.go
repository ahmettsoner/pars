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

type DeleteProject struct{ flowx.BaseStep }

func (s *DeleteProject) Name() string { return "DeleteProject" }

func (s *DeleteProject) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

		logrus.Debugf("trying to create %v", project.Header.Name)
		if _, err := service.DeleteProject(project); err != nil {
			return err
		}

		fmt.Printf("%v Project created\n", project.Header.Name)
	}
	return nil
}

func (s *DeleteProject) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
