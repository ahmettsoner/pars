package applicationProject

import (
	"fmt"
	"os"
	"strings"

	"parsdevkit.net/engines"
	"parsdevkit.net/structs"
	"parsdevkit.net/structs/group"
	"parsdevkit.net/structs/project"
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type ApplicationProjectSerializer struct{}

// TODO: Burda GetWebApiProjects gibi daha specific işlemler gerçekleştirilebilir, text içinde varsa getirir. bu değerlendirilmeli, iç içe olması yerine daha kontrollü süreç yürütülebilir
func (s ApplicationProjectSerializer) GetProjectStuctsFromString(workspaceName string, data string) ([]applicationproject.ProjectBaseStruct, error) {
	projects := make([]applicationproject.ProjectBaseStruct, 0)

	yamlLines := strings.Split(string(data), "---")
	logrus.Debugf("%d section found", len(yamlLines))

	for _, line := range yamlLines {
		// fmt.Println(line)
		var header structs.Header
		if err := yaml.Unmarshal([]byte(line), &header); err != nil {
			return nil, err
		}
		logrus.Debugf("project (%v) will be imported", header.Name)

		if header.Type == structs.StructTypes.Project {

			var projectHeader project.Header
			if err := yaml.Unmarshal([]byte(line), &projectHeader); err != nil {
				return nil, err
			}
			if projectHeader.Kind == project.StructKinds.Application {
				var projectDefinitionStruct = applicationproject.ProjectBaseStruct{}
				if err := yaml.Unmarshal([]byte(line), &projectDefinitionStruct); err != nil {
					return nil, err
				}

				logrus.Debugf("project (%v) imported successfully", projectDefinitionStruct.Name)

				rawProject := projectDefinitionStruct
				if err := s.CompleteProjectInformation(workspaceName, &rawProject); err != nil {
					return nil, err
				}

				projects = append(projects, rawProject)
			}
		}
	}
	return projects, nil
}
func (s ApplicationProjectSerializer) GetProjectStuctsFromFile(workspaceName string, files ...string) ([]applicationproject.ProjectBaseStruct, error) {
	projects := make([]applicationproject.ProjectBaseStruct, 0)
	for _, file := range files {

		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		logrus.Debugf("file '%v' importing", file)
		projectStructs, err := s.GetProjectStuctsFromString(workspaceName, string(data))
		if err != nil {
			return nil, err
		}
		projects = append(projects, projectStructs...)

	}
	logrus.Debugf("%d project imported", len(projects))
	return projects, nil
}

func (s ApplicationProjectSerializer) CompleteProjectInformation(workspaceName string, project *applicationproject.ProjectBaseStruct) error {

	logrus.Debugf("filling project (%v) information", project.Name)

	activeWorkspace, err := s.GetWorkspace(workspaceName, *project)
	if err != nil {
		return err
	}

	//WARN: Doğru mu oldu?
	project.Specifications.Workspace = activeWorkspace.Name
	project.Specifications.WorkspaceObject = activeWorkspace.Specifications
	logrus.Debugf("workspace (%v) detected for (%v)", activeWorkspace.Name, project.Name)

	group, err := s.GetGroup(*project)
	if err != nil {
		return err
	}

	//WARN: Doğru mu oldu?
	project.Specifications.Group = group.Name
	project.Specifications.GroupObject = *&group.Specifications
	logrus.Debugf("group (%v) detected for (%v)", group.Name, project.Name)

	projectReferences, err := s.GetProjectReferences(*project)
	if err != nil {
		return err
	}
	project.Specifications.Configuration.References = projectReferences
	logrus.Debugf("project references (%d) restored for (%v)", len(project.Specifications.Configuration.References), project.Specifications.Path)

	return nil
}

func (s ApplicationProjectSerializer) GetWorkspace(defaultWorkspaceName string, project applicationproject.ProjectBaseStruct) (*workspace.WorkspaceBaseStruct, error) {
	//TODO: Bu şekilde interface'ten tip dönüşümü tamamlanamadı, yapı buna dönüştürülmeli
	// if projectStruct, ok := project.(project.ProjectBaseStruct); !ok {
	// 	return nil, fmt.Errorf("incompatible model type: expected %T, got %T", project, projectStruct)
	// } else {
	// }

	workspaceName := project.Specifications.Workspace
	if utils.IsEmpty(workspaceName) {
		workspaceName = defaultWorkspaceName
	}

	var result *workspace.WorkspaceBaseStruct = nil

	if !utils.IsEmpty(workspaceName) {
		workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
		workspace, err := workspaceService.GetByName(workspaceName)
		if err != nil {
			return nil, err
		}
		if workspace == nil {
			return nil, fmt.Errorf("workspace name (%v) is not correct", workspaceName)
		}
		result = workspace
	} else {
		appContext := engines.GetContext()
		result = appContext.CurrentWorkspace
	}

	return result, nil
}
func (s ApplicationProjectSerializer) GetGroup(project applicationproject.ProjectBaseStruct) (*group.GroupBaseStruct, error) {
	result := group.GroupBaseStruct{}

	groupName := project.Specifications.Group

	if !utils.IsEmpty(groupName) {
		groupService := services.NewGroupService(utils.GetEnvironment())
		group, err := groupService.GetByName(groupName)
		if err != nil {
			return nil, err
		}
		if group == nil {
			return nil, fmt.Errorf("group name (%v) is not correct", groupName)
		}
		result = *group
	}

	return &result, nil
}

func (s ApplicationProjectSerializer) GetProjectReferences(prj applicationproject.ProjectBaseStruct) ([]applicationproject.ProjectBaseStruct, error) {

	projectReferences := make([]applicationproject.ProjectBaseStruct, 0)

	for _, reference := range prj.Specifications.Configuration.References {
		logrus.Debugf("reference (%v) processing for (%v)", reference.Name, prj.Name)

		selectedProject, err := s.GetProjectReference(prj, reference)
		if err != nil {
			return nil, err
		}

		if selectedProject != nil {
			projectReferences = append(projectReferences, *selectedProject)
		} else {
			projectReferences = append(projectReferences, reference)
		}
	}

	return projectReferences, nil
}

func (s ApplicationProjectSerializer) GetProjectReference(prj applicationproject.ProjectBaseStruct, reference applicationproject.ProjectBaseStruct) (*applicationproject.ProjectBaseStruct, error) {

	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	var projectReference *applicationproject.ProjectBaseStruct = nil

	logrus.Debugf("reference (%v) processing for (%v)", reference.Name, prj.Name)

	if reference.Specifications.ID == 0 {

		logrus.Debugf("found new project defination for (%v) referenced by (%v)", reference.Name, prj.Name)

		if reference.Specifications.WorkspaceObject.ID == 0 {
			if utils.IsEmpty(reference.Specifications.Workspace) {
				reference.Specifications.Workspace = prj.Specifications.Workspace
				logrus.Debugf("decided to using same workspace (%v) for reference (%v)", reference.Specifications.Workspace, reference.Name)
			}
			selectedWorkspace, err := workspaceService.GetByName(reference.Specifications.Workspace)
			if err != nil {
				return nil, err
			}
			reference.Specifications.Workspace = selectedWorkspace.Name
			reference.Specifications.WorkspaceObject = selectedWorkspace.Specifications
			logrus.Debugf("different workspace (%v) for reference (%v)", reference.Specifications.Workspace, reference.Name)
		}

		selectedProject, err := projectService.GetByFullNameWorkspace(reference.GetFullName(), reference.Specifications.Workspace)
		if err != nil {
			return nil, err
		}

		if selectedProject != nil {
			projectReference = selectedProject
		} else {
			projectReference = &reference
		}
	}

	return projectReference, nil
}
