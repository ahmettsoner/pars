package providers

import (
	"github.com/sirupsen/logrus"
)

func GoExecute(path string, args ...string) error {

	var cmdPath string = "go"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("go", "1.21.6"), "bin", "go.exe")

	logrus.Debugf("executing go command: %v on path: %v", args, path)

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}

	return nil
}
