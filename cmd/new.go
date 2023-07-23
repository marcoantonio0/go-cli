package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/marcoantoni0/go-cli/templates"
	"github.com/marcoantoni0/go-cli/util"
	"github.com/spf13/cobra"
)

var projectName string

var NewCmd = &cobra.Command{
	Use:     "new [name]",
	Aliases: []string{"n"},
	Short:   "Create an project to REST API using the flags",
	Long: `Create an project with right pattern that I personal use for my projects.\n
	Can you use the flags to customize.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		urlFlag, _ := cmd.Flags().GetString("url")
		projectNameArg := args[0]
		// Validate if has github url
		if urlFlag != "" {
			projectName = urlFlag + projectNameArg
		} else {
			projectName = projectNameArg
		}
		fmt.Printf("✅ Creating %s...\n", projectName)

		// Generate all folders
		err := generateFolders()
		if err != nil {
			fmt.Println("❌ Error: ", err)
		}
		// Generate all folders
		err = generateFiles()
		if err != nil {
			fmt.Println("❌ Error: ", err)
		}
		// Navigate to project folder
		err = os.Chdir(projectName)
		if err != nil {
			fmt.Println("❌ Error: ", err)
		}
		// Create go.mod
		err = generateMod(urlFlag)
		if err != nil {
			fmt.Println("❌ Error: ", err)
		}
		// Install all initial dependencies
		err = installDepedencies()
		if err != nil {
			fmt.Println("❌ Error: ", err)
		}

		fmt.Printf("🎉 The %s project has been created.\n", projectName)
	},
}

func generateFolders() error {
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
			fmt.Printf("✅ Folder %s created successfully.\n", folder)
		}
	}

	return nil
}

func generateFiles() error {
	data := struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}
	_, err := util.CreateGoFileByTemplate(projectName, "main", templates.MainTemplate, data)
	if err != nil {
		return err
	}
	fmt.Printf("✅ main.go file is create successfully.")

	return nil
}

func generateMod(projecName string) error {
	cmd := exec.Command("go", "mod", "init", projectName)

	// Set the command's standard output and standard error to our program's standard output.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command.
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("✅ Go module initialized successfully.")
	return nil
}

func installDepedencies() error {
	fmt.Println("⏩ Installing all packages dependencies...")
	packages := []string{
		"github.com/gin-gonic/gin",
	}
	for _, packageName := range packages {
		err := util.InstallPackage(packageName)
		if err != nil {
			return err
		}
	}
	fmt.Println("✅ All packages has been installed successfully.")
	return nil
}
