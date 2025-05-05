package dataResource

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	textTemplate "text/template"

	"parsdevkit.net/engines/fileTemplate"
	dataresource "parsdevkit.net/structs/resource/data-resource"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/sirupsen/logrus"
)

type DataResourceEngine struct{}

func (s DataResourceEngine) CreateResourcesFromTemplate(init bool, data any, resourceFiles ...string) error {

	var allResources []dataresource.ResourceBaseStruct = make([]dataresource.ResourceBaseStruct, 0)

	for _, resourceFilePath := range resourceFiles {

		var tmplFile = filepath.Join(utils.GetManagerTemplatesLocation(), resourceFilePath)
		tmplContent, err := os.ReadFile(tmplFile)
		if err != nil {
			log.Fatal(err)
		}
		var outputBuffer bytes.Buffer
		err = textTemplate.Must(textTemplate.New("ResourceFromResource").Parse(string(tmplContent))).Execute(&outputBuffer, data)
		if err != nil {
			log.Fatal(err)
		}
		mainStr := outputBuffer.String()

		groupSerializer := DataResourceSerializer{}
		resources, err := groupSerializer.GetResourceStructsFromString(mainStr)
		if err != nil {
			return err
		}
		allResources = append(allResources, resources...)
	}

	if err := s.CreateResources(allResources, init); err != nil {
		return err
	}

	return nil
}
func (s DataResourceEngine) CreateResourcesFromFile(init bool, files ...string) error {
	if len(files) > 0 {

		allFiles := make([]string, 0)
		for _, path := range files {
			if !utils.IsEmpty(path) {
				files, err := utils.GetFilesInPath(path)
				if err != nil {
					return err
				}

				allFiles = append(allFiles, files...)
			}
		}

		logrus.Debugf("found %v files", len(allFiles))
		groupSerializer := DataResourceSerializer{}
		resourcesFromFile, err := groupSerializer.GetResourceStructsFromFile(allFiles...)
		if err != nil {
			return err
		}

		logrus.Debugf("found %v resource", len(resourcesFromFile))
		if err := s.CreateResources(resourcesFromFile, init); err != nil {
			return err
		}
	}
	return nil
}

func (s DataResourceEngine) RemoveResourcesFromFile(permanent bool, files ...string) error {
	if len(files) > 0 {

		allFiles := make([]string, 0)
		for _, path := range files {
			if !utils.IsEmpty(path) {

				files, err := utils.GetFilesInPath(path)
				if err != nil {
					return err
				}

				allFiles = append(allFiles, files...)
			}
		}

		groupSerializer := DataResourceSerializer{}
		resourcesFromFile, err := groupSerializer.GetResourceStructsFromFile(allFiles...)
		if err != nil {
			return err
		}

		if err := s.RemoveResources(resourcesFromFile, permanent); err != nil {
			return err
		}
	}
	return nil
}
func (s DataResourceEngine) CreateResources(resources []dataresource.ResourceBaseStruct, init bool) error {

	resourcesReadyToCreate := make([]dataresource.ResourceBaseStruct, 0)
	resourcesForUpdate := make([]dataresource.ResourceBaseStruct, 0)
	resourceService := services.NewDataResourceService(utils.GetEnvironment())

	for _, resource := range resources {
		if ok := resourceService.IsExists(resource.Name, resource.Specifications.Workspace); ok {
			newModelHash, err := utils.CalculateHashFromObject(resource)
			if err != nil {
				return err
			}
			structHash := resourceService.GetHash(resource.Name)

			if newModelHash != structHash {
				resourcesForUpdate = append(resourcesForUpdate, resource)
			}
		} else {
			resourcesReadyToCreate = append(resourcesReadyToCreate, resource)
		}
	}
	logrus.Debugf("'%d' resource(s) detected that will create", len(resourcesReadyToCreate))
	logrus.Debugf("'%d' resource(s) detected that will update", len(resourcesForUpdate))

	logrus.Debugf("creating %v new resources ", len(resourcesReadyToCreate))
	logrus.Debugf("updating %v resources ", len(resourcesForUpdate))
	for _, resource := range resourcesReadyToCreate {

		if _, err := resourceService.Save(resource); err != nil {
			return err
		}

		if _, err := s.Generate(resource); err != nil {
			return err
		}

		fmt.Printf("%v Resource created\n", resource.Name)
	}

	logrus.Debugf("updating %v resources ", len(resourcesForUpdate))
	for _, resource := range resourcesForUpdate {

		if _, err := resourceService.Save(resource); err != nil {
			return err
		}

		fmt.Printf("%v Resource updated\n", resource.Name)
	}

	return nil
}

func (s DataResourceEngine) RemoveResources(resources []dataresource.ResourceBaseStruct, permanent bool) error {

	resourceService := services.NewDataResourceService(utils.GetEnvironment())
	resourcesReadyToDelete := make([]dataresource.ResourceBaseStruct, 0)
	for _, resource := range resources {
		if ok := resourceService.IsExists(resource.Name, resource.Specifications.Workspace); ok {
			resourcesReadyToDelete = append(resourcesReadyToDelete, resource)
		}
	}

	for _, resource := range resourcesReadyToDelete {

		if _, err := resourceService.Remove(resource.Name, resource.Specifications.Workspace, true, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Resource deleted\n", resource.Name)

	}

	return nil
}

func (s DataResourceEngine) Generate(model dataresource.ResourceBaseStruct) (*dataresource.ResourceBaseStruct, error) {

	resourceService := services.NewDataResourceService(utils.GetEnvironment())

	result, err := resourceService.GetByName(model.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	templateOperations := fileTemplate.NewFileTemplateOperations(utils.GetEnvironment())
	err = templateOperations.GenerateByResource(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
