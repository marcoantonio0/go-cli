package util

import (
	"os"
	"os/exec"
)

func InstallPackage(packageName string) error {
	cmd := exec.Command("go", "get", "-u", packageName)

	// Set the command's standard output and standard error to our program's standard output.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command.
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
