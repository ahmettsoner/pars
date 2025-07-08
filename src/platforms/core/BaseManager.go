package core

import (
	"os"
	"strings"

	"parsdevkit.net/models"
	_string "parsdevkit.net/pkg/utilities/string"

	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type BaseManager struct {
	ApplicationPlatformManagerInterface
	PackageDelimiter string
}

func BaseManagerNew(packageDelimiter string) BaseManager {
	return BaseManager{
		PackageDelimiter: packageDelimiter,
	}
}

func (s BaseManager) GetDefaultPlatformProjectType(model application_project_payload_structs.ProjectBaseStruct) models.ProjectType {
	return models.ProjectTypes.Library
}
func (s *BaseManager) FileRemover(paths ...string) error {
	for _, v := range paths {
		if err := os.RemoveAll(v); err != nil {
			return err
		}
	}
	return nil
}

func (s BaseManager) IsProjectFolderExists(project application_project_payload_structs.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(project.Specifications.GetAbsoluteProjectPath()) // check if the project directory exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return stat.IsDir(), nil
	}
}

func (s BaseManager) IsGroupFolderExists(project application_project_payload_structs.ProjectBaseStruct) (bool, error) {

	stat, err := os.Stat(project.Specifications.GetAbsoluteGroupPath()) // check if the project directory exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return stat.IsDir(), nil
	}
}

func (s BaseManager) IsLayerFolderExists(project application_project_payload_structs.ProjectBaseStruct, layer string) (bool, error) {

	layerPath := project.Specifications.GetAbsoluteProjectLayerPath(layer)
	stat, err := os.Stat(layerPath)

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return stat.IsDir(), nil
	}
}

func (s BaseManager) IsLayerFoldersExists(project application_project_payload_structs.ProjectBaseStruct) (bool, error) {
	for _, v := range project.Specifications.Layers {
		state, err := s.IsLayerFolderExists(project, v.Name)
		if err != nil {
			return false, err
		}
		if !state {
			return false, nil
		}
	}

	return true, nil
}

func (s BaseManager) IsGroupExists(project application_project_payload_structs.ProjectBaseStruct, controlFile string) (bool, error) {
	if _string.IsEmpty(project.Specifications.Group) {
		return false, nil
	}

	_, err := os.Stat(controlFile)

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (s *BaseManager) GetGroupPackage(project application_project_payload_structs.ProjectBaseStruct) string {

	result := strings.Join(project.Specifications.GroupObject.Package, s.PackageDelimiter)

	return result
}

func (s *BaseManager) GetProjectPackage(project application_project_payload_structs.ProjectBaseStruct) string {

	result := strings.Join(project.Specifications.GetAllPackage(), s.PackageDelimiter)

	return result
}
func (s BaseManager) NormalizeText(text string) string {

	return text
}
