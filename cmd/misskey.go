/*
Copyright © 2024 Soli
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Soli0222/Mac-NowPlaying/modules"
	"github.com/spf13/cobra"
)

// misskeyCmd represents the misskey command
var misskeyCmd = &cobra.Command{
	Use:   "misskey",
	Short: "Post the currently playing track to Misskey.",
	Long: `The 'misskey' command allows you to share your currently playing track directly to Misskey.
You need to configure the command by setting the environment variables MISSKEY_HOST and MISSKEY_TOKEN.

Usage examples:

  # Post the currently playing track to Misskey
  export MISSKEY_HOST=your_misskey_instance
  export MISSKEY_TOKEN=your_access_token
  misskey`,
	Run: func(cmd *cobra.Command, args []string) {

		postText, err := modules.GetCurrentTrackInfo()
		if err != nil {
			log.Fatalf("Get Track Info is Error: %v", err)
		}
		fmt.Println(postText)

		host := os.Getenv("MISSKEY_HOST")
		if host == "" {
			log.Fatalf("Missing Misskey Host. Please set MISSKEY_HOST env.")
		}

		token := os.Getenv("MISSKEY_TOKEN")
		if token == "" {
			log.Fatalf("Missing Misskey Token. Please set MISSKEY_TOKEN env.")
		}

		url := fmt.Sprintf("https://%s/api/notes/create", host)
		payload := map[string]string{
			"i":    token,
			"text": postText,
		}
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			log.Fatalf("Error marshalling JSON: %v", err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
		if err != nil {
			log.Fatalf("Error posting to Misskey: %v", err)
		}
		defer resp.Body.Close()
	},
}

func init() {
	rootCmd.AddCommand(misskeyCmd)
}
