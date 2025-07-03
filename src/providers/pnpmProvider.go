package providers

func PNPMExecute(version string, path string, args ...string) error {
	// platformVersion := ""

	// switch version {
	// default:
	// 	platformVersion = "20.11.0"
	// }

	var cmdPath string = "pnpm"
	// cmdPath = filepath.Join(application.GetBinaryLocation("nodejs", platformVersion), "pnpm.cmd")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}

func PNPMExecuteWithOutput(version string, path string, args ...string) (string, error) {
	// platformVersion := ""

	// switch version {
	// default:
	// 	platformVersion = "20.11.0"
	// }

	var cmdPath string = "pnpm"
	// cmdPath = filepath.Join(application.GetBinaryLocation("nodejs", platformVersion), "pnpm.cmd")

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
