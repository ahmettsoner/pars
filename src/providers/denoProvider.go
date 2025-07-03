package providers

func DenoExecute(version string, path string, args ...string) error {
	// platformVersion := ""

	// switch version {
	// default:
	// 	platformVersion = "20.11.0"
	// }

	var cmdPath string = "deno"
	// cmdPath = filepath.Join(application.GetBinaryLocation("nodejs", platformVersion), "deno.cmd")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}

func DenoExecuteWithOutput(version string, path string, args ...string) (string, error) {
	// platformVersion := ""

	// switch version {
	// default:
	// 	platformVersion = "20.11.0"
	// }

	var cmdPath string = "deno"
	// cmdPath = filepath.Join(application.GetBinaryLocation("nodejs", platformVersion), "deno.cmd")

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
