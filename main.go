package main

import (
	"fmt"
	"os"

	"github.com/marcoantoni0/go-cli/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "My awesome CLI application.",
	Long:  `A longer description that spans multiple lines and tells users what this application is about.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This function is executed when the root command is called.
		cmd.Help()
	},
}

func main() {
	Execute()
}

func init() {
	cmd.NewCmd.Flags().StringP("url", "u", "", "Your github url")

	rootCmd.AddCommand(cmd.NewCmd)
	rootCmd.AddCommand(cmd.VersionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
