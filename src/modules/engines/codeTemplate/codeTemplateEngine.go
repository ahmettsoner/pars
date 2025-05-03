package codeTemplate

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	textTemplate "text/template"

	codetemplate "parsdevkit.net/structs/template/code-template"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/sirupsen/logrus"
)

type CodeTemplateEngine struct{}

func (s CodeTemplateEngine) CreateTemplatesFromTemplate(init bool, data any, templateFiles ...string) error {

	var allTemplates []codetemplate.TemplateBaseStruct = make([]codetemplate.TemplateBaseStruct, 0)

	for _, templateFilePath := range templateFiles {

		var tmplFile = filepath.Join(utils.GetManagerTemplatesLocation(), templateFilePath)
		tmplContent, err := os.ReadFile(tmplFile)
		if err != nil {
			log.Fatal(err)
		}
		var outputBuffer bytes.Buffer
		err = textTemplate.Must(textTemplate.New("TemplateFromTemplate").Parse(string(tmplContent))).Execute(&outputBuffer, data)
		if err != nil {
			log.Fatal(err)
		}
		mainStr := outputBuffer.String()

		groupSerializer := CodeTemplateSerializer{}
		templates, err := groupSerializer.GetTemplateStructsFromString(mainStr)
		if err != nil {
			return err
		}
		allTemplates = append(allTemplates, templates...)
	}

	if err := s.CreateTemplates(allTemplates, init); err != nil {
		return err
	}

	return nil
}
func (s CodeTemplateEngine) CreateTemplatesFromFile(init bool, files ...string) error {
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
		groupSerializer := CodeTemplateSerializer{}
		templatesFromFile, err := groupSerializer.GetTemplateStructsFromFile(allFiles...)
		if err != nil {
			return err
		}

		logrus.Debugf("found %v template", len(templatesFromFile))
		if err := s.CreateTemplates(templatesFromFile, init); err != nil {
			return err
		}
	}
	return nil
}

func (s CodeTemplateEngine) RemoveTemplatesFromFile(permanent bool, files ...string) error {
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

		groupSerializer := CodeTemplateSerializer{}
		templatesFromFile, err := groupSerializer.GetTemplateStructsFromFile(allFiles...)
		if err != nil {
			return err
		}

		if err := s.RemoveTemplates(templatesFromFile, permanent); err != nil {
			return err
		}
	}
	return nil
}
func (s CodeTemplateEngine) CreateTemplates(templates []codetemplate.TemplateBaseStruct, init bool) error {

	templatesReadyToCreate := make([]codetemplate.TemplateBaseStruct, 0)
	templatesForUpdate := make([]codetemplate.TemplateBaseStruct, 0)
	templateService := services.NewCodeTemplateService(utils.GetEnvironment())

	for _, template := range templates {
		if ok := templateService.IsExists(template.Name, template.Specifications.Workspace); ok {
			newModelHash, err := utils.CalculateHashFromObject(template)
			if err != nil {
				return err
			}
			structHash := templateService.GetHash(template.Name)

			if newModelHash != structHash {
				templatesForUpdate = append(templatesForUpdate, template)
			}
		} else {
			templatesReadyToCreate = append(templatesReadyToCreate, template)
		}
	}
	logrus.Debugf("'%d' template(s) detected that will create", len(templatesReadyToCreate))
	logrus.Debugf("'%d' template(s) detected that will update", len(templatesForUpdate))

	logrus.Debugf("creating %v new templates ", len(templatesReadyToCreate))
	logrus.Debugf("updating %v templates ", len(templatesForUpdate))
	for _, template := range templatesReadyToCreate {

		fmt.Printf("Creating %v Template\n", template.Name)

		if _, err := templateService.Save(template); err != nil {
			return err
		}

		if _, err := s.Generate(template); err != nil {
			return err
		}

		fmt.Printf("%v Template created\n", template.Name)
	}

	logrus.Debugf("updating %v templates ", len(templatesForUpdate))
	for _, template := range templatesForUpdate {

		if _, err := templateService.Save(template); err != nil {
			return err
		}

		if _, err := s.Generate(template); err != nil {
			return err
		}

		fmt.Printf("%v Template updated\n", template.Name)
	}

	return nil
}

func (s CodeTemplateEngine) RemoveTemplates(templates []codetemplate.TemplateBaseStruct, permanent bool) error {

	templateService := services.NewCodeTemplateService(utils.GetEnvironment())
	templatesReadyToDelete := make([]codetemplate.TemplateBaseStruct, 0)
	for _, template := range templates {
		if ok := templateService.IsExists(template.Name, template.Specifications.Workspace); ok {
			templatesReadyToDelete = append(templatesReadyToDelete, template)
		}
	}

	for _, template := range templatesReadyToDelete {
		if _, err := templateService.Remove(template.Name, template.Specifications.Workspace, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Template deleted\n", template.Name)

	}

	return nil
}

func (s CodeTemplateEngine) Generate(model codetemplate.TemplateBaseStruct) (*codetemplate.TemplateBaseStruct, error) {

	templateService := services.NewCodeTemplateService(utils.GetEnvironment())

	result, err := templateService.GetByName(model.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	templateEngine := NewCodeTemplateOperations(utils.GetEnvironment())
	err = templateEngine.GenerateByTemplate(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
