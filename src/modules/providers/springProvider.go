package providers

func SpringExecute(path string, args ...string) error {

	var cmdPath string = "spring"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("spring", "0.8.1"), "jars", "spring-cli-0.8.1.jar")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}
