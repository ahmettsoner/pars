package services

import (
	"os"
	"path/filepath"
	"regexp"

	"parsdevkit.net/core/utils"
)

type EnvironmentService struct {
}

func NewEnvironmentService() *EnvironmentService {
	return &EnvironmentService{}
}

func (s *EnvironmentService) List() ([]string, error) {
	var result = make([]string, 0)

	directory := utils.GetDataLocation()

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
			envName := utils.GetOnlyFileName(matches[1])
			if !utils.IsEmpty(envName) {
				result = append(result, envName)
				// } else {
				// 	result = append(result, "pars")
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}
