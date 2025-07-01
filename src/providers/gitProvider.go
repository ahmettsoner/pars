package providers

func GitExecute(path string, args ...string) error {

	var cmdPath string = "git"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("git", "2"), "cmd", "git.exe")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}
