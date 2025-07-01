package providers

func GVimExecute(path string, args ...string) error {

	var cmdPath string = "gvim"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("gvim", "9.1.0"), "vim.exe")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}
