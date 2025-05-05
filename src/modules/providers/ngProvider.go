package providers

func NGExecute(version string, path string, args ...string) error {
	platformVersion := ""

	switch version {
	case "V15":
		platformVersion = "15.2.10"
	case "V16":
		platformVersion = "16.2.12"
	case "V17":
		platformVersion = "17.1.3"
	}

	args = append([]string{"--package", "@angular/cli@" + platformVersion, "ng"}, args...)
	err := NPXExecute("", path, args...)
	if err != nil {
		return err
	}

	return nil
}

func NGNodeExecute(path string, args ...string) error {

	var cmdPath string = "ng"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("nodejs", "20.11.0"), "ng.cmd")

	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}
