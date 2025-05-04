package core

import (
	"os"
	"strings"

	applicationproject "parsdevkit.net/structs/project/application-project"

	"parsdevkit.net/core/utils"
)

type BaseManager struct {
	ManagerInterface
	PackageDelimiter string
}

func BaseManagerNew(packageDelimiter string) BaseManager {
	return BaseManager{
		PackageDelimiter: packageDelimiter,
	}
}

func (s *BaseManager) FileRemover(paths ...string) error {
	for _, v := range paths {
		if err := os.RemoveAll(v); err != nil {
			return err
		}
	}
	return nil
}

func (s BaseManager) IsProjectFolderExists(project applicationproject.ProjectSpecification) (bool, error) {

	stat, err := os.Stat(project.GetAbsoluteProjectPath()) // check if the project directory exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return stat.IsDir(), nil
	}
}

func (s BaseManager) IsGroupFolderExists(project applicationproject.ProjectSpecification) (bool, error) {

	stat, err := os.Stat(project.GetAbsoluteGroupPath()) // check if the project directory exists

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return stat.IsDir(), nil
	}
}

func (s BaseManager) IsLayerFolderExists(project applicationproject.ProjectSpecification, layer string) (bool, error) {

	layerPath := project.GetAbsoluteProjectLayerPath(layer)
	stat, err := os.Stat(layerPath)

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return stat.IsDir(), nil
	}
}

func (s BaseManager) IsLayerFoldersExists(project applicationproject.ProjectSpecification) (bool, error) {
	for _, v := range project.Configuration.Layers {
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

func (s BaseManager) IsGroupExists(project applicationproject.ProjectSpecification, controlFile string) (bool, error) {
	if utils.IsEmpty(project.Group) {
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

func (s *BaseManager) GetGroupPackage(project applicationproject.ProjectSpecification) string {

	result := strings.Join(project.GroupObject.Package, s.PackageDelimiter)

	return result
}

func (s *BaseManager) GetProjectPackage(project applicationproject.ProjectSpecification) string {

	result := strings.Join(project.GetAllPackage(), s.PackageDelimiter)

	return result
}
func (s BaseManager) NormalizeText(text string) string {

	return text
}
