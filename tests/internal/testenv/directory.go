package test

import (
	"parsdevkit.net/application"

	"os"
	"path/filepath"
	"time"
)

func CreateTempTestDirectory(testName string) (string, error) {
	testDir := filepath.Join(application.GetTempsLocation(), "test", testName)
	err := os.MkdirAll(testDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	format := "2006-01-02T15-04-05.000"
	tempDir, err := os.MkdirTemp(testDir, time.Now().Format(format))
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(filepath.Join(tempDir, "test-files"), os.ModePerm)
	if err != nil {
		return "", err
	}

	return tempDir, nil
}
