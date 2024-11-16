/*
Copyright © 2024 Soli
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Soli0222/Mac-NowPlaying/modules"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// nowplayingCmd represents the nowplaying command
var nowplayingCmd = &cobra.Command{
	Use:   "nowplaying",
	Short: "Displays the current track information",
	Long: `The nowplaying command retrieves and displays the current track information
from your music player. It also copies the track information to the clipboard for easy sharing.

Usage examples:

  # Display the current track information
  nowplaying`,
	Run: func(cmd *cobra.Command, args []string) {

		postText, err := modules.GetCurrentTrackInfo()
		if err != nil {
			log.Fatalf("Get Track Info is Error: %v", err)
		}
		fmt.Println(postText)

		err = clipboard.WriteAll(postText)
		if err != nil {
			log.Fatalf("Error copying to clipboard: %v", err)
		} else {
			fmt.Println("\nSuccess copying to clipboard")
		}
	},
}

func init() {
	rootCmd.AddCommand(nowplayingCmd)
}
