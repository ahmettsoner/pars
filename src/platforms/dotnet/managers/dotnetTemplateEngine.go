package managers

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"parsdevkit.net/platforms/core"

	"parsdevkit.net/core/utils"
)

type DotnetTemplateEngine struct{}

func (s *DotnetTemplateEngine) AddSimpleClass(templateName string, outputFile string, data any) error {
	var templateFilePath = "/dotnet/net70/SimpleClass.cs.templ"
	var tmplFile = filepath.Join(utils.GetManagerTemplatesLocation(), templateFilePath)

	return s.RenderTemplate(tmplFile, templateName, outputFile, data)
}
func (s *DotnetTemplateEngine) AddSimpleConsoleClass(templateName string, outputFile string, data any) error {
	var templateFilePath = "/dotnet/net70/SimpleConsoleClass.cs.templ"
	var tmplFile = filepath.Join(utils.GetManagerTemplatesLocation(), templateFilePath)

	return s.RenderTemplate(tmplFile, templateName, outputFile, data)
}
func (s *DotnetTemplateEngine) AddSimpleControllerClass(templateName string, outputFile string, data any) error {
	var templateFilePath = "/dotnet/net70/SimpleControllerClass.cs.templ"
	var tmplFile = filepath.Join(utils.GetManagerTemplatesLocation(), templateFilePath)

	return s.RenderTemplate(tmplFile, templateName, outputFile, data)
}
func (s *DotnetTemplateEngine) RenderTemplate(tmplFile string, templateName string, outputFile string, data any) error {
	mainStr, err := s.GetContent(tmplFile, templateName, data)
	if err != nil {
		return err
	}

	infoData := core.GenerateMessageData{
		CreatedAt: time.Now().Format("2006-01-02"),
		CreatedBy: "Pars",
		Version:   "1.0.0",
	}
	var infoTemplateFilePath = "/dotnet/generateMessage.templ"
	var infoTmplFile = filepath.Join(utils.GetManagerTemplatesLocation(), infoTemplateFilePath)
	infoStr, err := s.GetContent(infoTmplFile, "info", infoData)
	if err != nil {
		return err
	}

	content := fmt.Sprintf("%s\r\n\r\n%s", infoStr, mainStr)

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	fmt.Printf("Template output written to %s\n", outputFile)

	return nil
}

func (*DotnetTemplateEngine) GetContent(tmplFile string, templateName string, data any) (string, error) {
	tmplContent, err := os.ReadFile(tmplFile)
	if err != nil {
		return "", err
	}

	var outputBuffer bytes.Buffer
	err = template.Must(template.New(templateName).Parse(string(tmplContent))).Execute(&outputBuffer, data)
	if err != nil {
		return "", err
	}
	mainStr := outputBuffer.String()

	return mainStr, nil
}
