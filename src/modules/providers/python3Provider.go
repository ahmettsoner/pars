package providers

func Python3Execute(path string, args ...string) error {

	var cmdPath string = "python3"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("python", "3.12.1"), "python3.exe")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}
