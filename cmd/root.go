/*
Copyright © 2024 Soli
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mnp",
	Short: "A tool to share your currently playing track information.",
	Long: `The 'mnp' tool allows you to share your currently playing track information in various ways.
You can display the track information, copy it to the clipboard, or post it directly to social media platforms like Twitter and Misskey.

By setting the environment variable MNP_APP_NAME, you can retrieve information from any application. 
If the environment variable MNP_APP_NAME is not set, the default application name "Swinsian" will be used.

Available commands:

  # Display the current track information and copy it to the clipboard
  nowplaying

  # Post the currently playing track to Twitter
  tweet

  # Post the currently playing track to Misskey
  export MISSKEY_HOST=your_misskey_instance
  export MISSKEY_TOKEN=your_access_token
  misskey`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
