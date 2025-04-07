package providers

func DotnetExecute(version string, path string, args ...string) error {
	// platformVersion := ""

	// switch version {
	// case "Net5":
	// 	platformVersion = "5.0.408"
	// case "Net6":
	// 	platformVersion = "6.0.418"
	// case "Net7":
	// 	platformVersion = "7.0.405"
	// case "Net8":
	// 	platformVersion = "8.0.101"
	// default:
	// 	platformVersion = "8.0.101"
	// }

	var cmdPath string = "dotnet"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("dotnet", platformVersion), "dotnet.exe")
	err := Execute(path, cmdPath, args...)
	if err != nil {
		return err
	}
	return nil
}

func DotnetExecuteWithOutput(version string, path string, args ...string) (string, error) {
	// platformVersion := ""

	// switch version {
	// case "Net5":
	// 	platformVersion = "5.0.408"
	// case "Net6":
	// 	platformVersion = "6.0.418"
	// case "Net7":
	// 	platformVersion = "7.0.405"
	// case "Net8":
	// 	platformVersion = "8.0.101"
	// default:
	// 	platformVersion = "8.0.101"
	// }

	var cmdPath string = "dotnet"
	// cmdPath = filepath.Join(utils.GetBinaryLocation("dotnet", platformVersion), "dotnet.exe")

	output, err := ExecuteWithOutput(path, cmdPath, args...)
	if err != nil {
		return "", err
	}
	return output, nil
}
