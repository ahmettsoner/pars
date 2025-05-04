package engines

import (
	"bytes"
	"fmt"
	"text/template"

	"parsdevkit.net/templates/engines/functions"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/operation/services"
)

func TemplateEngine(templateFile string, data any) (string, error) {
	var outputBuffer bytes.Buffer

	funcMap := template.FuncMap{
		"context": func() ContextFuncs {
			return ContextFuncs{}
		},
		"engine": func() EngineFuncs {
			return EngineFuncs{}
		},
		"console": func() functions.ConsoleFuncs {
			return functions.ConsoleFuncs{}
		},
		"math": func() functions.MathFuncs {
			return functions.MathFuncs{}
		},
		"string": func() functions.StringFuncs {
			return functions.StringFuncs{}
		},
		"time": func() functions.TimeFuncs {
			return functions.TimeFuncs{}
		},
		"data": func() functions.DataFuncs {
			return functions.DataFuncs{}
		},
		"regexp": func() functions.RegexpFuncs {
			return functions.RegexpFuncs{}
		},
		"base64": func() functions.Base64Funcs {
			return functions.Base64Funcs{}
		},
		"array": func() functions.ArrayFuncs {
			return functions.ArrayFuncs{}
		},
		"map": func() functions.MapFuncs {
			return functions.MapFuncs{}
		},
		"json": func() functions.JsonFuncs {
			return functions.JsonFuncs{}
		},
		"convert": func() functions.ConvertFuncs {
			return functions.ConvertFuncs{}
		},
	}

	tmpl := template.New("ResourceFromContent").Funcs(funcMap)

	sharedTemplateService := services.NewSharedTemplateService(utils.GetEnvironment())
	sharedTemplateList, err := sharedTemplateService.List()
	if err != nil {
		return "", err
	}

	for _, sharedTempl := range *sharedTemplateList {
		tempDef := fmt.Sprintf("{{define \"%v\"}}%v{{end}}", sharedTempl.Name, sharedTempl.Specifications.Template.Content)
		_, err := tmpl.Parse(tempDef)
		if err != nil {
			return "", err
		}
	}

	builtinTemplates := []string{}
	for _, tmplContent := range builtinTemplates {
		_, err := tmpl.Parse(tmplContent)
		if err != nil {
			return "", err
		}
	}

	_, err = tmpl.Parse(templateFile)
	if err != nil {
		return "", err
	}

	if err := tmpl.Execute(&outputBuffer, data); err != nil {
		return "", err
	}

	return outputBuffer.String(), nil
}
