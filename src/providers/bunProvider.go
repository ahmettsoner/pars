package providers

func BUNExecute(version string, path string, args ...string) error {
	// platformVersion := ""

	// switch version {
	// default:
	// 	platformVersion = "20.11.0"
	// }

	var cmdPath string = "bun"
	// cmdPath = filepath.Join(application.GetBinaryLocation("nodejs", platformVersion), "bun.cmd")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}

func BUNExecuteWithOutput(version string, path string, args ...string) (string, error) {
	// platformVersion := ""

	// switch version {
	// default:
	// 	platformVersion = "20.11.0"
	// }

	var cmdPath string = "bun"
	// cmdPath = filepath.Join(application..GetBinaryLocation("nodejs", platformVersion), "bun.cmd")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return "", err
	}

	output, err := ExecuteWithOutput(path, cmdPath, args...)
	if err != nil {
		return "", err
	}
	return output, nil
}
