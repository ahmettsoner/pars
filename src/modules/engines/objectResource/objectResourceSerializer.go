package objectResource

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/core/utils"
	"parsdevkit.net/engines"
	"parsdevkit.net/operation/services"
	"parsdevkit.net/structs"
	"parsdevkit.net/structs/resource"
	objectresource "parsdevkit.net/structs/resource/object-resource"
	"parsdevkit.net/structs/workspace"

	"gopkg.in/yaml.v3"
)

type ObjectResourceSerializer struct{}

func (s ObjectResourceSerializer) GetResourceStructsFromString(data string) ([]objectresource.ResourceBaseStruct, error) {
	resources := make([]objectresource.ResourceBaseStruct, 0)

	yamlLines := strings.Split(string(data), "---")

	for _, line := range yamlLines {
		var header structs.Header
		if err := yaml.Unmarshal([]byte(line), &header); err != nil {
			return nil, err
		}

		if header.Type == structs.StructTypes.Resource {
			var resourceHeader resource.Header
			if err := yaml.Unmarshal([]byte(line), &resourceHeader); err != nil {
				return nil, err
			}
			if resourceHeader.Kind == resource.StructKinds.Object {
				var resourceDefinitionStruct = objectresource.ResourceBaseStruct{}
				if err := yaml.Unmarshal([]byte(line), &resourceDefinitionStruct); err != nil {
					return nil, err
				}

				rawResource := resourceDefinitionStruct
				if err := s.CompleteResourceInformation(&rawResource); err != nil {
					return nil, err
				}

				resources = append(resources, rawResource)
			}
		}
	}

	return resources, nil
}

func (s ObjectResourceSerializer) GetResourceStructsFromFile(files ...string) ([]objectresource.ResourceBaseStruct, error) {
	resources := make([]objectresource.ResourceBaseStruct, 0)
	for _, file := range files {

		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		resourceStructs, err := s.GetResourceStructsFromString(string(data))
		if err != nil {
			return nil, err
		}
		resources = append(resources, resourceStructs...)

	}
	return resources, nil
}

func (s ObjectResourceSerializer) CompleteResourceInformation(resource *objectresource.ResourceBaseStruct) error {

	logrus.Debugf("filling resource (%v) information", resource.Name)

	if strings.ToLower(resource.Specifications.Workspace) != strings.ToLower("None") {
		activeWorkspace, err := s.GetWorkspace(*resource)
		if err != nil {
			return err
		}

		//WARN: Doğru mu oldu?
		resource.Specifications.Workspace = activeWorkspace.Name
		resource.Specifications.WorkspaceObject = activeWorkspace.Specifications
		logrus.Debugf("workspace (%v) detected for (%v)", activeWorkspace.Name, resource.Name)
	}

	return nil
}

func (s ObjectResourceSerializer) GetWorkspace(resource objectresource.ResourceBaseStruct) (*workspace.WorkspaceBaseStruct, error) {

	appContext := engines.GetContext()
	var result workspace.WorkspaceBaseStruct = *appContext.CurrentWorkspace

	workspaceName := resource.Specifications.Workspace

	if !utils.IsEmpty(workspaceName) {
		workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
		workspace, err := workspaceService.GetByName(workspaceName)
		if err != nil {
			return nil, err
		}
		if workspace == nil {
			return nil, fmt.Errorf("workspace name (%v) is not correct", workspaceName)
		}
		result = *workspace
	}

	return &result, nil
}
