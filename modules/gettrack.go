package modules

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func GetCurrentTrackInfo() (string, error) {

	appName := os.Getenv("MNP_APP_NAME")
	if appName == "" {
		appName = "Swinsian"
	}

	scriptTrack := fmt.Sprintf(`
        tell application "%s"
            if it is running then
                if player state is playing then
                    set track_name to name of current track
                    return track_name
                end if
            end if
        end tell
    `, appName)

	scriptArtist := fmt.Sprintf(`
        tell application "%s"
            if it is running then
                if player state is playing then
                    set artist_name to artist of current track
                    return artist_name
                end if
            end if
        end tell
    `, appName)

	scriptAlbum := fmt.Sprintf(`
        tell application "%s"
            if it is running then
                if player state is playing then
                    set album_name to album of current track
                    return album_name
                end if
            end if
        end tell
    `, appName)

	trackName, err := execAppleScript(scriptTrack)
	if err != nil {
		return "", err
	}
	artistName, err := execAppleScript(scriptArtist)
	if err != nil {
		return "", err
	}
	albumName, err := execAppleScript(scriptAlbum)
	if err != nil {
		return "", err
	}

	trackName = regexp.MustCompile(`\n`).ReplaceAllString(trackName, "")
	artistName = regexp.MustCompile(`\n`).ReplaceAllString(artistName, "")
	albumName = regexp.MustCompile(`\n`).ReplaceAllString(albumName, "")

	if trackName == "" && artistName == "" && albumName == "" {
		log.Fatalf("All item is empty")
	}

	postText := fmt.Sprintf("%s / %s\n%s\n#NowPlaying #PsrPlaying", trackName, artistName, albumName)

	return postText, nil
}

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
