package providers

func NodeJSExecute(path string, args ...string) error {

	var cmdPath string = "node"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("nodejs", "20.11.0"), "node.exe")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}
