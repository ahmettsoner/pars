package providers

func PythonExecute(path string, args ...string) error {

	var cmdPath string = "python"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("python", "3.12.1"), "python.exe")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}
