package application

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/core/utils"

	applicationproject "parsdevkit.net/structs/project/application-project"
)

type Service struct{}

func (s Service) GetProjectsFromManifest(path ...string) ([]applicationproject.ProjectBaseStruct, error) {
	files := path

	if len(files) > 0 {
		logrus.Debugf("found %v file(s) to create project", len(files))

		allFiles := make([]string, 0)
		for _, path := range files {
			if !utils.IsEmpty(path) {

				files, err := utils.GetFilesInPath(path)
				if err != nil {
					return nil, err
				}

				allFiles = append(allFiles, files...)
			}
		}

		serializer := Serializer{}
		projectsFromFile, err := serializer.GetProjectStuctsFromFile(allFiles...)
		if err != nil {
			return nil, err
		} else {
			return projectsFromFile, nil
		}

	} else {
		logrus.Debugf("no project definition found")
		return make([]applicationproject.ProjectBaseStruct, 0), nil
	}
}

func (s Service) GetProjectsFromManifestWithTemplate(data any, path ...string) ([]applicationproject.ProjectBaseStruct, error) {
	files := path

	if len(files) > 0 {
		logrus.Debugf("found %v file(s) to create project", len(files))

		allFiles := make([]string, 0)
		for _, path := range files {
			if !utils.IsEmpty(path) {

				files, err := utils.GetFilesInPath(path)
				if err != nil {
					return nil, err
				}

				allFiles = append(allFiles, files...)
			}
		}

		serializer := Serializer{}
		projectsFromFile, err := serializer.GetProjectStuctsFromFile(allFiles...)
		if err != nil {
			return nil, err
		} else {
			return projectsFromFile, nil
		}

	} else {
		logrus.Debugf("no project definition found")
		return make([]applicationproject.ProjectBaseStruct, 0), nil
	}
}

func (s Service) CreateProjectsFromTemplate(init bool, data any, templateFiles ...string) error {

	var allProjects []applicationproject.ProjectBaseStruct = make([]applicationproject.ProjectBaseStruct, 0)

	logrus.Debugf("found %v template(s) to create project", len(templateFiles))
	for _, templateFilePath := range templateFiles {

		var tmplFile = filepath.Join(utils.GetManagerTemplatesLocation(), templateFilePath)
		tmplContent, err := os.ReadFile(tmplFile)
		if err != nil {
			log.Fatal(err)
		}
		var outputBuffer bytes.Buffer
		err = template.Must(template.New("ProjectFromTemplate").Parse(string(tmplContent))).Execute(&outputBuffer, data)
		if err != nil {
			log.Fatal(err)
		}
		mainStr := outputBuffer.String()

		projectSerializer := Serializer{}
		projects, err := projectSerializer.GetProjectStuctsFromString(mainStr)
		if err != nil {
			return err
		}
		allProjects = append(allProjects, projects...)
	}
	logrus.Debugf("found %v project(s) to create", len(allProjects))
	if err := s.CreateProjects(allProjects, init); err != nil {
		return err
	}

	return nil
}
