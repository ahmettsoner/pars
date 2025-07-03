package providers

func VSCodeExecute(path string, args ...string) error {

	var cmdPath string = "code"
	// cmdPath = filepath.Join(application.GetBinaryLocation("vscode", "1.86.0"), "Code.exe")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}
