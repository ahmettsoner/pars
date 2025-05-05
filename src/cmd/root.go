package cmd

import (
	"fmt"
	"os"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core"

	cmdBrowse "parsdevkit.net/cmd/browse"
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

	// cmdWork "parsdevkit.net/cmd/work"

	// cmdContainerize "parsdevkit.net/cmd/containerize"
	// cmdDistribute "parsdevkit.net/cmd/distribute"
	// cmdGenerate "parsdevkit.net/cmd/generate"
	cmdEdit "parsdevkit.net/cmd/edit"
	cmdExecute "parsdevkit.net/cmd/execute"
	cmdInit "parsdevkit.net/cmd/init"
	cmdOpen "parsdevkit.net/cmd/open"

	// cmdRelease "parsdevkit.net/cmd/release"
	// cmdRemote "parsdevkit.net/cmd/remote"
	cmdWorkspace "parsdevkit.net/cmd/workspace"

	// cmdGit "parsdevkit.net/cmd/git"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile          string
	environment      string
	logLevelEnumFlag core.LogLevelEnumFlag
)

var RootCmd = &cobra.Command{
	Use:   "pars [type] [command] [options] [flags]",
	Short: "Smart Software Development Process Automation",
	Long:  `Smart Software Development Process Automation`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !utils.IsEmpty(environment) {
			utils.SetEnvironment(environment)
		}

		utils.SetLogLevel(logLevelEnumFlag.Value)

		if logLevelEnumFlag.Value != core.LogLevels.Silence {
			if logrusLogLevel, err := log.ParseLevel(string(logLevelEnumFlag.Value)); err != nil {
				fmt.Println(err)
				// file, err := os.OpenFile(filepath.Join(utils.GetLogLocation(), "app.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	utils.PrepareLocations()

	if !utils.IsEmpty(utils.GetEnvironment()) {
		fmt.Printf("\nRunning on '%v' environment\n", utils.GetEnvironment())
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")
	RootCmd.PersistentFlags().StringVarP(&environment, "env", "e", "", "Environment (dev, prod, test, ...)")

	logLevelValues := core.LogLevelToArray()
	logLevelEnumFlag.Value = core.LogLevels.Error
	RootCmd.PersistentFlags().VarP(&logLevelEnumFlag, "log-level", "", fmt.Sprintf("Select log level %v", logLevelValues))

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubCommands()
}

func addSubCommands() {
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
	RootCmd.AddCommand(cmdExecute.ExecuteCmd)
	RootCmd.AddCommand(cmdOpen.OpenCmd)
	RootCmd.AddCommand(cmdEdit.EditCmd)
	// RootCmd.AddCommand(cmdWork.WorkCmd)
	// RootCmd.AddCommand(cmdGit.GitCmd)
	// RootCmd.AddCommand(cmdContainerize.ContainerizeCmd)
	// RootCmd.AddCommand(cmdDistribute.DistributeCmd)
	RootCmd.AddCommand(cmdWorkspace.WorkspaceCmd)
	RootCmd.AddCommand(cmdWorkspace.WorkspaceListShorthandsCmd)

}

func initConfig() {
	if !utils.IsEmpty(cfgFile) {
		viper.SetConfigFile(cfgFile)
	} else {
		configDir := utils.GetConfigLocation()

		viper.AddConfigPath(configDir)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
