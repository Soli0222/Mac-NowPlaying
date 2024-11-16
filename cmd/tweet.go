/*
Copyright © 2024 Soli
*/
package cmd

import (
	"fmt"
	"log"
	"net/url"

	"github.com/Soli0222/Mac-NowPlaying/modules"
	"github.com/spf13/cobra"
)

// tweetCmd represents the tweet command
var tweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "Post the currently playing track to Twitter.",
	Long: `The 'tweet' command allows you to share your currently playing track directly to Twitter.
It opens a web browser with a pre-filled tweet containing the track information.

Usage examples:

  # Post the currently playing track to Twitter
  tweet`,
	Run: func(cmd *cobra.Command, args []string) {

		postText, err := modules.GetCurrentTrackInfo()
		if err != nil {
			log.Fatalf("Get Track Info is Error: %v", err)
		}
		fmt.Println(postText)

		modules.OpenBrowser("https://x.com/intent/tweet?text=" + url.QueryEscape(postText))
	},
}

func init() {
	rootCmd.AddCommand(tweetCmd)
}
