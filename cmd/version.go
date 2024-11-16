/*
Copyright © 2024 Soli
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the application.",
	Long: `The 'version' command prints the version number of the application.
This is useful for checking which version of the application you are currently using.

Usage examples:

  # Print the version number
  version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
