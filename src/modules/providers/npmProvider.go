package providers

func NPMExecute(path string, args ...string) error {

	var cmdPath string = "npm"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("nodejs", "20.11.0"), "npm.cmd")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}

func NPMExecuteWithOutput(path string, args ...string) (string, error) {

	var cmdPath string = "npm"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("nodejs", "20.11.0"), "npm.cmd")

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
