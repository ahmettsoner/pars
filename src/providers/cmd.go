package providers

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

func ExecuteWithOutput(path string, command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = path

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	return stdout.String(), nil
}
func Execute(path string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = path

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
func ExecuteCombined(path string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = path
	// currentDir, err := os.Getwd()
	// if err != nil {
	// 	return err
	// }

	// err = os.Chdir(path)
	// if err != nil {
	// 	return err
	// }

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	// err = os.Chdir(currentDir)
	// if err != nil {
	// 	return err
	// }

	fmt.Println(output)

	return nil
}
func ExecuteQuick(command string, args ...string) error {
	fmt.Println(args)
	cmd := exec.Command(command, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
