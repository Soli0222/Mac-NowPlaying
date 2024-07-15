package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"encoding/json"
	"log"
	"net/http"

	"github.com/atotto/clipboard"
	"github.com/joho/godotenv"
)

func execAppleScript(script string) (string, error) {
	cmd := exec.Command("osascript", "-e", script)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func getCurrentTrackInfo() (string, string, string, error) {
	scriptTrack := `
		tell application "Swinsian"
			if it is running then
				if player state is playing then
					set track_name to name of current track
					return track_name
				end if
			end if
		end tell
	`
	scriptArtist := `
		tell application "Swinsian"
			if it is running then
				if player state is playing then
					set artist_name to artist of current track
					return artist_name
				end if
			end if
		end tell
	`
	scriptAlbum := `
		tell application "Swinsian"
			if it is running then
				if player state is playing then
					set album_name to album of current track
					return album_name
				end if
			end if
		end tell
	`

	trackName, err := execAppleScript(scriptTrack)
	if err != nil {
		return "", "", "", err
	}
	artistName, err := execAppleScript(scriptArtist)
	if err != nil {
		return "", "", "", err
	}
	albumName, err := execAppleScript(scriptAlbum)
	if err != nil {
		return "", "", "", err
	}

	trackName = regexp.MustCompile(`\n`).ReplaceAllString(trackName, "")
	artistName = regexp.MustCompile(`\n`).ReplaceAllString(artistName, "")
	albumName = regexp.MustCompile(`\n`).ReplaceAllString(albumName, "")

	return trackName, artistName, albumName, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	trackName, artistName, albumName, err := getCurrentTrackInfo()
	if err != nil {
		log.Fatalf("Error getting track info: %v", err)
	}
	text := fmt.Sprintf("%s / %s\n%s\n#NowPlaying #PsrPlaying", trackName, artistName, albumName)

	if os.Getenv("SHARE_MISSKEY") == "True" {
		host := os.Getenv("MISSKEY_HOST")
		token := os.Getenv("MISSKEY_TOKEN")

		url := fmt.Sprintf("https://%s/api/notes/create", host)
		payload := map[string]string{
			"i":    token,
			"text": text,
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
	} else {
		err := clipboard.WriteAll(text)
		if err != nil {
			log.Fatalf("Error copying to clipboard: %v", err)
		}
	}
}
