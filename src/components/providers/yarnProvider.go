package providers

func YarnExecute(path string, args ...string) error {

	var cmdPath string = "yarn"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("nodejs", "20.11.0"), "yarn.cmd")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}

func YarnExecuteWithOutput(path string, args ...string) (string, error) {

	var cmdPath string = "yarn"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("nodejs", "20.11.0"), "yarn.cmd")

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
