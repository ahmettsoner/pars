package cmd

import (
	"fmt"
	"os"

	"parsdevkit.net/application"
	"parsdevkit.net/pkg/logs"
	_string "parsdevkit.net/pkg/utilities/string"

	cmdApply "parsdevkit.net/cmd/apply"
	cmdBrowse "parsdevkit.net/cmd/browse"
	cmdDestroy "parsdevkit.net/cmd/destroy"

	// cmdBuild "parsdevkit.net/cmd/build"
	cmdClean "parsdevkit.net/cmd/clean"
	cmdEnvironment "parsdevkit.net/cmd/environment"
	cmdGroup "parsdevkit.net/cmd/group"
	cmdInfo "parsdevkit.net/cmd/info"

	// cmdInstall "parsdevkit.net/cmd/install"
	cmdProject "parsdevkit.net/cmd/project"
	cmdResource "parsdevkit.net/cmd/resource"

	// cmdTask "parsdevkit.net/cmd/task"
	cmdTemplate "parsdevkit.net/cmd/template"
	cmdTest "parsdevkit.net/cmd/test"

	cmdInit "parsdevkit.net/cmd/init"
	cmdOpen "parsdevkit.net/cmd/open"
	cmdRun "parsdevkit.net/cmd/run"

	// cmdRelease "parsdevkit.net/cmd/release"
	// cmdRemote "parsdevkit.net/cmd/remote"
	cmdWorkspace "parsdevkit.net/cmd/workspace"

	// cmdGit "parsdevkit.net/cmd/git"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type RootOptions struct {
	cfgFile          string
	environment      string
	logLevelEnumFlag logs.LogLevelEnumFlag
}

var commandOptions RootOptions

var RootCmd = &cobra.Command{
	Use:           "pars [type] [command] [options] [flags]",
	Short:         "Smart Software Development Process Automation",
	Long:          `Smart Software Development Process Automation`,
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !_string.IsEmpty(commandOptions.environment) {
			application.SetEnvironment(commandOptions.environment)
		}

		application.SetLogLevel(commandOptions.logLevelEnumFlag.Value)

		if commandOptions.logLevelEnumFlag.Value != logs.LogLevels.Silence {
			if logrusLogLevel, err := log.ParseLevel(string(commandOptions.logLevelEnumFlag.Value)); err != nil {
				fmt.Println(err)
				// file, err := os.OpenFile(filepath.Join(application.GetLogLocation(), "app.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
				// if err != nil {
				// 	log.Fatal(err)
				// }
				// defer file.Close()
				// log.SetOutput(file)
			} else {
				log.SetLevel(logrusLogLevel)
			}
		}
	},
}

func Run() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	application.PrepareLocations()

	if !_string.IsEmpty(application.GetEnvironment()) {
		fmt.Printf("\nRunning on '%v' environment\n", application.GetEnvironment())
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&commandOptions.cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")
	RootCmd.PersistentFlags().StringVarP(&commandOptions.environment, "env", "e", "", "Environment (dev, prod, test, ...)")

	logLevelValues := logs.LogLevelToArray()
	commandOptions.logLevelEnumFlag.Value = logs.LogLevels.Error
	RootCmd.PersistentFlags().VarP(&commandOptions.logLevelEnumFlag, "log-level", "", fmt.Sprintf("Select log level %v", logLevelValues))

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubCommands()
}

func addSubCommands() {
	RootCmd.AddCommand(cmdApply.ApplyCmd)
	RootCmd.AddCommand(cmdDestroy.DestroyCmd)
	RootCmd.AddCommand(cmdInfo.InfoCmd)
	RootCmd.AddCommand(cmdInit.InitCmd)
	RootCmd.AddCommand(cmdGroup.GroupCmd)
	// RootCmd.AddCommand(cmdTask.TaskCmd)
	RootCmd.AddCommand(cmdProject.ProjectCmd)
	RootCmd.AddCommand(cmdResource.ResourceCmd)
	RootCmd.AddCommand(cmdTemplate.TemplateCmd)
	RootCmd.AddCommand(cmdEnvironment.EnvironmentCmd)
	// RootCmd.AddCommand(cmdGenerate.GenerateCmd)
	RootCmd.AddCommand(cmdClean.CleanCmd)
	// RootCmd.AddCommand(cmdInstall.InstallCmd)
	// RootCmd.AddCommand(cmdBuild.BuildCmd)
	RootCmd.AddCommand(cmdBrowse.BrowseCmd)
	RootCmd.AddCommand(cmdTest.TestCmd)
	// RootCmd.AddCommand(cmdRelease.ReleaseCmd)
	// RootCmd.AddCommand(cmdRemote.RemoteCmd)
	RootCmd.AddCommand(cmdRun.RunCmd)
	RootCmd.AddCommand(cmdOpen.OpenCmd)
	// RootCmd.AddCommand(cmdWork.WorkCmd)
	// RootCmd.AddCommand(cmdGit.GitCmd)
	// RootCmd.AddCommand(cmdContainerize.ContainerizeCmd)
	// RootCmd.AddCommand(cmdDistribute.DistributeCmd)
	RootCmd.AddCommand(cmdWorkspace.WorkspaceCmd)
	RootCmd.AddCommand(cmdWorkspace.WorkspaceListShorthandsCmd)

}

func init() {
	RegisterServices()
}

func initConfig() {
	if !_string.IsEmpty(commandOptions.cfgFile) {
		viper.SetConfigFile(commandOptions.cfgFile)
	} else {
		configDir := application.GetConfigLocation()

		viper.AddConfigPath(configDir)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
