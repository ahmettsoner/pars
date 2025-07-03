package providers

func NPXExecute(version string, path string, args ...string) error {
	// platformVersion := ""

	// switch version {
	// default:
	// 	platformVersion = "20.11.0"
	// }

	var cmdPath string = "npx"
	// cmdPath = filepath.Join(application.GetBinaryLocation("nodejs", platformVersion), "npx.cmd")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}

func NPMXExecuteWithOutput(version string, path string, args ...string) (string, error) {
	// platformVersion := ""

	// switch version {
	// default:
	// 	platformVersion = "20.11.0"
	// }

	var cmdPath string = "npx"
	// cmdPath = filepath.Join(application.GetBinaryLocation("nodejs", platformVersion), "npx.cmd")

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
