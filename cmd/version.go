package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Aliases: []string{"v"},
	Short: "Print the version number of Go CLI.",
	Long:  `This command prints the version number of Go CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go-cli version 1.0.0")
	},
}
