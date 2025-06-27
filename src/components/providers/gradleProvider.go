package providers

func GradleExecute(path string, args ...string) error {

	var cmdPath string = "gradle"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("gradle", "7.5"), "bin", "gradle.cmd")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}
