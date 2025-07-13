package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"parsdevkit.net/application"
	"parsdevkit.net/cmd"
)

var version string

func main() {

	if len(os.Args) > 1 && (os.Args[1] == "-i" || os.Args[1] == "--interactive") {
		runInteractiveMode()
	} else {
		if err := cmd.RootCmd.Execute(); err != nil {
			fmt.Println("\n❌ Error:", err)
			os.Exit(1)
		} else {
			fmt.Println("\n✅ Success: Operation completed.")
		}
	}
	// logLevel := application.GetLogLevel()

	// if logLevel != core.LogLevels.None {
	// 	if logrusLogLevel, err := log.ParseLevel(string(logLevel)); err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		log.SetLevel(logrusLogLevel)
	// 	}
	// } else {
	// 	fmt.Println("Zaten istenmiyor!!!")
	// }

	application.SetVersion(version)
}

func runInteractiveMode() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🔁 Interactive Pars CLI Mode. Type 'exit' to quit.")

	for {
		fmt.Print("pars >> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "exit" || line == "quit" {
			fmt.Println("👋 Bye!")
			break
		}

		if line == "" {
			continue
		}

		args := strings.Split(line, " ")
		cmd.RootCmd.SetArgs(args)

		if err := cmd.RootCmd.Execute(); err != nil {
			fmt.Println("\n❌ Error:", err)
		} else {
			fmt.Println("\n✅ Success: Operation completed.")
		}
	}
}
