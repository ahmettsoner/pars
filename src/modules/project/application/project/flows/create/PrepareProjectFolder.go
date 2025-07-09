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

type PrepareProjectFolder struct{ flowx.BaseStep }

func (s *PrepareProjectFolder) Name() string { return "PrepareProjectFolder" }

func (s *PrepareProjectFolder) Run(ctx context.Context, fc *flowx.FlowContext) error {

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
		if _, err := service.CreateProjectFolder(project); err != nil {
			return err
		}

	}
	return nil
}

func (s *PrepareProjectFolder) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
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
		if _, err := service.DeleteProjectFolder(project); err != nil {
			return err
		}

		fc.Log("Rollback: Creating project %s", project.Header.Name)
	}
	return nil
}
