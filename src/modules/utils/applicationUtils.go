package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/core"
)

var (
	environment string
	version     string
	stage       string
	logLevel    core.LogLevel
)

type StageType string

var StageTypes = struct {
	None   StageType
	Dev    StageType
	Test   StageType
	Stabil StageType
}{
	None:   "none",
	Dev:    "dev",
	Test:   "test",
	Stabil: "stabil",
}

func getDefaultPlatformApplicationDir() string {
	snap := os.Getenv("SNAP")

	switch runtime.GOOS {
	case "windows":
		programFiles := os.Getenv("PROGRAMFILES")
		return filepath.Join(programFiles, "Pars/bin")
	case "darwin":
		if !IsEmpty(snap) {
			return filepath.Join(snap, "bin")
		}
		return "/usr/bin"
	case "linux":
		if !IsEmpty(snap) {
			return filepath.Join(snap, "bin")
		}
		return "/usr/bin"
	case "freebsd":
		return "/usr/bin"
	case "openbsd":
		return "/usr/bin"
	case "netbsd":
		return "/usr/bin"
	default:
		return "/usr/bin"
	}
}

func getDefaultPlatformLibraryDir() string {
	snap := os.Getenv("SNAP")
	snapCommon := os.Getenv("SNAP_USER_COMMON")

	switch runtime.GOOS {
	case "windows":
		programFiles := os.Getenv("PROGRAMFILES")
		return filepath.Join(programFiles, "Pars/lib")
	case "darwin":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "lib")
		}
		return "/usr/lib/pars"
	case "linux":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "lib")
		}
		return "/usr/lib/pars"
	case "freebsd":
		return "/usr/lib/pars"
	case "openbsd":
		return "/usr/lib/pars"
	case "netbsd":
		return "/usr/lib/pars"
	default:
		return "/usr/lib/pars"
	}
}

func getDefaultPlatformPluginDir() string {
	snap := os.Getenv("SNAP")
	snapCommon := os.Getenv("SNAP_USER_COMMON")

	switch runtime.GOOS {
	case "windows":
		programFiles := os.Getenv("PROGRAMFILES")
		return filepath.Join(programFiles, "Pars/plugins")
	case "darwin":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "plugins")
		}
		return "/usr/share/pars/plugins"
	case "linux":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "plugins")
		}
		return "/usr/share/pars/plugins"
	case "freebsd":
		return "/usr/share/pars/plugins"
	case "openbsd":
		return "/usr/share/pars/plugins"
	case "netbsd":
		return "/usr/share/pars/plugins"
	default:
		return "/usr/share/pars/plugins"
	}
}

func getDefaultPlatformDocumentDir() string {
	snap := os.Getenv("SNAP")
	snapCommon := os.Getenv("SNAP_USER_COMMON")

	switch runtime.GOOS {
	case "windows":
		programFiles := os.Getenv("PROGRAMFILES")
		return filepath.Join(programFiles, "Pars/doc")
	case "darwin":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "doc")
		}
		return "/usr/share/doc/pars"
	case "linux":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "doc")
		}
		return "/usr/share/doc/pars"
	case "freebsd":
		return "/usr/share/doc/pars"
	case "openbsd":
		return "/usr/share/doc/pars"
	case "netbsd":
		return "/usr/share/doc/pars"
	default:
		return "/usr/share/doc/pars"
	}
}

func getDefaultPlatformConfigDir() string {
	snap := os.Getenv("SNAP")
	snapCommon := os.Getenv("SNAP_USER_COMMON")

	switch runtime.GOOS {
	case "windows":
		programData := os.Getenv("PROGRAMDATA")
		return filepath.Join(programData, "Pars/config")
	case "darwin":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "config")
		}
		return "/etc/pars"
	case "linux":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "config")
		}
		return "/etc/pars"
	case "freebsd":
		return "/etc/pars"
	case "openbsd":
		return "/etc/pars"
	case "netbsd":
		return "/etc/pars"
	default:
		return "/etc/pars"
	}
}

func getDefaultPlatformDataDir() string {
	snap := os.Getenv("SNAP")
	snapCommon := os.Getenv("SNAP_USER_COMMON")

	switch runtime.GOOS {
	case "windows":
		programData := os.Getenv("PROGRAMDATA")
		return filepath.Join(programData, "Pars/data")
	case "darwin":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "data")
		}
		return "/var/lib/pars/data"
	case "linux":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "data")
		}
		return "/var/lib/pars/data"
	case "freebsd":
		return "/var/lib/pars/data"
	case "openbsd":
		return "/var/lib/pars/data"
	case "netbsd":
		return "/var/lib/pars/data"
	default:
		return "/var/lib/pars/data"
	}
}

func getDefaultPlatformLogDir() string {
	snap := os.Getenv("SNAP")
	snapCommon := os.Getenv("SNAP_USER_COMMON")

	switch runtime.GOOS {
	case "windows":
		programData := os.Getenv("PROGRAMDATA")
		return filepath.Join(programData, "Pars/logs")
	case "darwin":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "logs")
		}
		return "/var/log/pars"
	case "linux":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "logs")
		}
		return "/var/log/pars"
	case "freebsd":
		return "/var/log/pars"
	case "openbsd":
		return "/var/log/pars"
	case "netbsd":
		return "/var/log/pars"
	default:
		return "/var/log/pars"
	}
}

func getDefaultPlatformCacheDir() string {
	snap := os.Getenv("SNAP")
	snapCommon := os.Getenv("SNAP_USER_COMMON")

	switch runtime.GOOS {
	case "windows":
		programData := os.Getenv("PROGRAMDATA")
		return filepath.Join(programData, "Pars/cache")
	case "darwin":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "cache")
		}
		return "/var/cache/pars"
	case "linux":
		if !IsEmpty(snap) {
			return filepath.Join(snapCommon, "cache")
		}
		return "/var/cache/pars"
	case "freebsd":
		return "/var/cache/pars"
	case "openbsd":
		return "/var/cache/pars"
	case "netbsd":
		return "/var/cache/pars"
	default:
		return "/var/cache/pars"
	}
}

func getDefaultPlatformTempDir() string {
	return os.TempDir()
}

func SetStage(stage string) {
	stage = stage
}
func GetStage() string {
	if IsEmpty(stage) {
		return string(StageTypes.None)
	}
	return stage
}

func SetVersion(version string) {
	version = version
}
func GetVersion() string {
	if IsEmpty(version) {
		ver, _ := getWorkingVersion(filepath.Join(GetCodeBaseLocation(), "VERSION"))
		return ver
	}
	return version
}

func SetEnvironment(env string) {
	environment = env
}
func GetEnvironment() string {
	return environment
}

func GetPlatform() string {
	snap := os.Getenv("SNAP")

	if !IsEmpty(snap) {
		return "snap"
	} else if isRunningInDocker() {
		return "docker"
	}

	return "native"
}

func SetLogLevel(LogLevel core.LogLevel) {
	logLevel = LogLevel
}
func GetLogLevel() core.LogLevel {
	return logLevel
}

func GetCodeBaseLocation() string {
	if GetStage() == string(StageTypes.None) {
		_, file, _, ok := runtime.Caller(0)
		if !ok {
			fmt.Println("Error getting caller")
		}
		project_root, err := BaseDirFromFilePath(file, "src/modules/utils/applicationUtils.go")
		if err != nil {
			fmt.Println("Error:", err)
		}
		return project_root
	} else if GetStage() == string(StageTypes.Dev) {
		projectRoot := os.Getenv("PARS_PROJECT_ROOT")

		return projectRoot
	}

	return ""
}

func GetExecutableLocation() string {

	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Get the directory containing the executable
	exeDir := filepath.Dir(exePath)
	return exeDir
}

func PrepareLocations() error {
	err := PrepareConfigLocation()
	if err != nil {
		return err
	}
	err = PrepareLogLocation()
	if err != nil {
		return err
	}

	err = PrepareDocumentLocation()
	if err != nil {
		return err
	}

	err = PrepareDataLocation()
	if err != nil {
		return err
	}

	err = PrepareCacheLocation()
	if err != nil {
		return err
	}

	err = PrepareBinariesLocation()
	if err != nil {
		return err
	}

	err = PreparePluginsLocation()
	if err != nil {
		return err
	}

	return nil
}

func GetConfigLocation() string {
	if GetStage() == string(StageTypes.None) {
		return filepath.Join(GetCodeBaseLocation(), "bin")
	} else if GetStage() == string(StageTypes.Dev) {
		return filepath.Join(GetCodeBaseLocation(), "config")
	} else if GetStage() == string(StageTypes.Test) {
		return filepath.Join(GetExecutableLocation(), "config")
	}
	return getDefaultPlatformConfigDir()
}
func PrepareConfigLocation() error {
	path := GetConfigLocation()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func GetLogLocation() string {
	if GetStage() == string(StageTypes.None) {
		return filepath.Join(GetCodeBaseLocation(), "log")
	} else if GetStage() == string(StageTypes.Dev) {
		return filepath.Join(GetCodeBaseLocation(), "log")
	} else if GetStage() == string(StageTypes.Test) {
		return filepath.Join(GetExecutableLocation(), "log")
	}
	return getDefaultPlatformLogDir()
}
func PrepareLogLocation() error {
	path := GetLogLocation()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func GetDocumentLocation() string {
	if GetStage() == string(StageTypes.None) {
		return filepath.Join(GetCodeBaseLocation(), "docs")
	} else if GetStage() == string(StageTypes.Dev) {
		return filepath.Join(GetCodeBaseLocation(), "docs")
	} else if GetStage() == string(StageTypes.Test) {
		return filepath.Join(GetExecutableLocation(), "docs")
	}
	return getDefaultPlatformDocumentDir()
}
func PrepareDocumentLocation() error {
	path := GetDocumentLocation()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func GetDataLocation() string {
	if GetStage() == string(StageTypes.None) {
		return filepath.Join(GetCodeBaseLocation(), "data")
	} else if GetStage() == string(StageTypes.Dev) {
		return filepath.Join(GetCodeBaseLocation(), "data")
	} else if GetStage() == string(StageTypes.Test) {
		return filepath.Join(GetExecutableLocation(), "data")
	}

	dataDir := getDefaultPlatformDataDir()

	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err := os.MkdirAll(dataDir, 0755)
		if err != nil {
			log.Fatalf("Error creating data directory: %v", err)
		}
	}

	err := os.Chmod(dataDir, 0755)
	if err != nil {
		log.Fatalf("Error setting permissions for data directory: %v", err)
	}
	return dataDir
}
func PrepareDataLocation() error {
	path := GetDataLocation()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
func GetDBLocation(environment string) string {
	dbName := "pars.db"

	if !IsEmpty(environment) {
		dbName = fmt.Sprintf("pars-%v.db", environment)
	}

	path := filepath.Join(GetDataLocation(), dbName)
	return path
}

func GetCacheLocation() string {
	if GetStage() == string(StageTypes.None) {
		return filepath.Join(GetCodeBaseLocation(), "cache")
	} else if GetStage() == string(StageTypes.Dev) {
		return filepath.Join(GetCodeBaseLocation(), "cache")
	} else if GetStage() == string(StageTypes.Test) {
		return filepath.Join(GetExecutableLocation(), "cache")
	}
	return getDefaultPlatformCacheDir()
}
func PrepareCacheLocation() error {
	path := GetCacheLocation()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
func GetBinariesLocation() string {
	if GetStage() == string(StageTypes.None) {
		return filepath.Join(GetCodeBaseLocation(), "bin")
	} else if GetStage() == string(StageTypes.Dev) {
		return filepath.Join(GetCodeBaseLocation(), "bin")
	} else if GetStage() == string(StageTypes.Test) {
		return filepath.Join(GetExecutableLocation(), "bin")
	}
	return getDefaultPlatformLibraryDir()
}
func PrepareBinariesLocation() error {
	path := GetBinariesLocation()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
func GetBinaryLocation(name, version string) string {
	return filepath.Join(GetBinariesLocation(), name, version)
}

func GetPluginsLocation() string {
	if GetStage() == string(StageTypes.None) {
		return filepath.Join(GetCodeBaseLocation(), "plugins")
	} else if GetStage() == string(StageTypes.Dev) {
		return filepath.Join(GetCodeBaseLocation(), "plugins")
	} else if GetStage() == string(StageTypes.Test) {
		return filepath.Join(GetExecutableLocation(), "plugins")
	}
	return getDefaultPlatformPluginDir()
}
func PreparePluginsLocation() error {
	path := GetPluginsLocation()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func GetPluginLocation(pluginName string) string {
	return filepath.Join(GetPluginsLocation(), pluginName)
}

func GetManagerTemplatesLocation() string {
	return filepath.Join(GetExecutableLocation(), "templates")
}

func GetSourceLocation() string {
	return filepath.Join(GetCodeBaseLocation(), "src")
}

// TODO: Gözden geçirilerek değerlendirilecek
func GetTempsLocation() string {
	if GetStage() == string(StageTypes.None) {
		return filepath.Join(GetCodeBaseLocation(), "temp")
	} else if GetStage() == string(StageTypes.Dev) {
		return filepath.Join(GetCodeBaseLocation(), "temp")
	} else if GetStage() == string(StageTypes.Test) {
		return filepath.Join(GetExecutableLocation(), "temp")
	}
	return getDefaultPlatformTempDir()
}

// TODO: Gözden geçirilerek değerlendirilecek
func GetTestsLocation() string {
	if GetStage() == string(StageTypes.None) {
		return filepath.Join(GetCodeBaseLocation(), "tests")
	} else if GetStage() == string(StageTypes.Dev) {
		return filepath.Join(GetCodeBaseLocation(), "tests")
	}
	return ""
}
func GetTestFileLocation(file string) string {
	return filepath.Join(GetTestsLocation(), file)
}

// TODO: Gözden geçirilerek değerlendirilecek
func GenerateTestArea() string {
	parsTestDir := GetTestsLocation()
	file, err := os.Getwd()
	if err != nil {
		logrus.Fatalf("Error occured when getting test file location: %v", err)
	}

	if strings.ToLower(filepath.VolumeName(parsTestDir)) == strings.ToLower(filepath.VolumeName(file)) {
		drive := filepath.VolumeName(parsTestDir)

		testArea := strings.TrimPrefix(filepath.Join(strings.ToUpper(drive), file[len(drive):]), parsTestDir)
		testArea = strings.TrimPrefix(testArea, "/")
		testArea = strings.TrimPrefix(testArea, "\\")

		return filepath.Join(testArea)

	} else {
		logrus.Fatalf("Test Drivers are not same")
	}
	return ""
}

func GetTestFileFromCurrentLocation(file string) string {
	return GetTestFileLocation(filepath.Join(GenerateTestArea(), file))
}

func GetUnitTestsLocation() string {
	return filepath.Join(GetTestsLocation(), "unit")
}

func GetUnitTestFileLocation(file string) string {
	return filepath.Join(GetUnitTestsLocation(), file)
}

func GetE2ETestsLocation() string {
	return filepath.Join(GetTestsLocation(), "e2e")
}
func GetE2ETestFileLocation(file string) string {
	return filepath.Join(GetE2ETestsLocation(), file)
}

func GetIntegrationTestsLocation() string {
	return filepath.Join(GetTestsLocation(), "integration")
}
func GetIntegrationTestFileLocation(file string) string {
	return filepath.Join(GetIntegrationTestsLocation(), file)
}

func GetScenarioTestsLocation() string {
	return filepath.Join(GetTestsLocation(), "scenario")
}
func GetScenarioTestFileLocation(file string) string {
	return filepath.Join(GetScenarioTestsLocation(), file)
}

func GetBenchmarkTestsLocation() string {
	return filepath.Join(GetTestsLocation(), "benchmark")
}
func GetBenchmarkTestFileLocation(file string) string {
	return filepath.Join(GetBenchmarkTestsLocation(), file)
}

func isRunningInDocker() bool {
	file, err := os.Open("/proc/1/cgroup")
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "docker") || strings.Contains(line, "kubepods") {
			return true
		}
	}
	return false
}

func getWorkingVersion(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "WORKING_VERSION=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return parts[1], nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("WORKING_VERSION not found")
}
