package main

import (
	"parsdevkit.net/cmd"
	"parsdevkit.net/core/utils"
)

var version string

func main() {

	cmd.Execute()
	// logLevel := utils.GetLogLevel()

	// if logLevel != core.LogLevels.None {
	// 	if logrusLogLevel, err := log.ParseLevel(string(logLevel)); err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		log.SetLevel(logrusLogLevel)
	// 	}
	// } else {
	// 	fmt.Println("Zaten istenmiyor!!!")
	// }

	utils.SetVersion(version)
}
