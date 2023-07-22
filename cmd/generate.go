package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/marcoantoni0/go-cli/templates"
	"github.com/marcoantoni0/go-cli/util"
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:     "generate [name]",
	Aliases: []string{"g"},
	Short:   "Generate an project to REST API using the flags",
	Long: `Generate an project with right pattern that I personal use for my projects.\n
	Can you use the flags to customize.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		urlFlag, _ := cmd.Flags().GetString("url")
		projectName := args[0]
		fmt.Printf("Generating %s...\n", projectName)

		// Generate all folders
		err := generateFolders(projectName)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		// Generate all folders
		err = generateFiles(projectName)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		err = generateMod(projectName, urlFlag)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	},
}

func generateFolders(projectName string) error {
	// Creating the root project folder
	err := util.CreateFolder(projectName)
	if err != nil {
		return err
	}

	// Create all the folders for pattern
	folders := []string{"controllers", "services", "repositories", "routes"}
	for _, folder := range folders {
		err := util.CreateFolderWithParent(projectName+"/app", folder)
		if err != nil {
			return err
		} else {
			fmt.Printf("Folder %s created successfully.\n", folder)
		}
	}

	return nil
}

func generateFiles(projectName string) error {
	data := struct {
		Name string
	}{
		Name: "Marco",
	}
	_, err := util.CreateGoFileByTemplate(projectName, "main", templates.MainTemplate, data)
	if err != nil {
		return err
	}
	fmt.Printf("main.go file is create successfully.")

	return nil
}

func generateMod(projectName, url string) error {
	var projectFullName string
	if url != "" {
		projectFullName = url + projectName
	} else {
		projectFullName = projectName
	}

	err := os.Chdir(projectName)
	if err != nil {
		return err
	}
	cmd := exec.Command("go", "mod", "init", projectFullName)

	// Set the command's standard output and standard error to our program's standard output.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command.
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("Go module initialized successfully.")
	return nil
}
