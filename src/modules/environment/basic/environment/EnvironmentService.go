package basic_environment

import (
	"os"
	"path/filepath"
	"regexp"

	"parsdevkit.net/application"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs/environment"
	applicationEnvironment "parsdevkit.net/application/structs/environment"
	"parsdevkit.net/modules/environment/basic_environment_contract"
	"parsdevkit.net/pkg/utilities/file"
	_string "parsdevkit.net/pkg/utilities/string"

	basic_environment_payload_structs "parsdevkit.net/modules/environment/basic_environment_payload/structs"
)

const (
	DEFAULT_WORKSPACE_PATH string = "environment"
	CURRENT_WORKSPACE_ID   string = "current_environment_id"
)

type EnvironmentService struct {
}

func NewEnvironmentService() basic_environment_contract.EnvironmentInterface {
	return &EnvironmentService{}
}

func (s EnvironmentService) List() ([]basic_environment_payload_structs.EnvironmentBaseStruct, error) {

	entityList, err := s.list()
	if err != nil {
		return nil, err
	}

	environmentList := make([]basic_environment_payload_structs.EnvironmentBaseStruct, 0)

	for _, entity := range entityList {
		environment := basic_environment_payload_structs.EnvironmentBaseStruct{
			Header: schemas.NewSchemaHeader(
				schemas.StructTypes.Environment,
				basic_environment_payload_structs.ENVIRONMENT_KIND,
				entity,
				schemas.Metadata{},
			),
			Specifications: applicationEnvironment.EnvironmentSpecification{
				EnvironmentIdentifier: environment.EnvironmentIdentifier{},
			},
		}

		environmentList = append(environmentList, environment)
	}

	return environmentList, nil
}

func (s EnvironmentService) list() ([]string, error) {
	var result = make([]string, 0)

	directory := application.GetDataLocation()

	desen := "^pars-?(.*?).db$"

	regexPattern, err := regexp.Compile(desen)
	if err != nil {
		return nil, err

	}

	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		matches := regexPattern.FindStringSubmatch(string(info.Name()))
		if len(matches) > 0 {
			envName := file.GetOnlyFileName(matches[1])
			if !_string.IsEmpty(envName) {
				result = append(result, envName)
			} else {
				result = append(result, "Default")
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}
