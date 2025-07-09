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

type SaveProject struct{ flowx.BaseStep }

func (s *SaveProject) Name() string { return "SaveProject" }

func (s *SaveProject) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

func (s *SaveProject) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[application_project_contract.ProjectInterface]()

	project, ok := flowx.Get[application_project_payload_structs.ProjectBaseStruct](fc, "project")
	if !ok {
		panic(fmt.Errorf("xxx: project parametresi hatalı tipte"))
	}

	logrus.Warnf("yyy: rolling back the creating %v!", project.Header.Name)
	if _, err := service.UndoSaveProject(project); err != nil {
		return fmt.Errorf("xxx: Application Project Information kayıt sırasında meydana gelen hata için uygulanan rollback hatası oluştu: '%s'\n%w", &project.Header.Name, &err)
	}

	fc.Log("Rollback: Persist project %s", project.Header.Name)
	return nil
}
