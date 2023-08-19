# Mac Now Playing Clipboard Copy Script

This script allows you to copy the currently playing song information on your Mac to the clipboard. It's designed to work with various music player applications. By default, it is set up to work with the "Swinsian" music player, but you can easily configure it to work with other supported applications as well.

## Supported Applications

The script is capable of retrieving the current track information from various music player applications. To configure the script for a specific music player, replace `<ApplicationName>` in the script below with the name of the desired music player application:

```applescript
tell application "<ApplicationName>"
    if it is running then
        if player state is playing then
            set track_name to name of current track
            set artist_name to artist of current track
            set album_name to album of current track
            return {track_name, artist_name, album_name}
        end if
    end if
end tell
```

Replace `<ApplicationName>` with the exact name of the music player you want to use, as it appears in the application's scripting dictionary.

Here are a few examples of how you might replace `<ApplicationName>` for different music player applications:

- For Swinsian:

  ```applescript
  tell application "Swinsian"
  ```

- For iTunes:

  ```applescript
  tell application "iTunes"
  ```

- For Spotify:

  ```applescript
  tell application "Spotify"
  ```

And so on for any other supported music player application you wish to use.

## How to Use

Follow these steps to use the script:

1. Clone the repository:

   ```sh
   git clone https://github.com/Soli0222/Mac-NowPlaying.git
   ```

2. Change into the cloned directory:

   ```sh
   cd Mac-NowPlaying
   ```

3. Create a virtual environment (optional but recommended):

   ```sh
   python -m venv venv
   ```

4. Activate the virtual environment:

   ```sh
   source venv/bin/activate
   ```

5. Install the required libraries:

   ```sh
   pip install -r requirements.txt
   ```

6. Edit the script to configure the target music player application:

   - Open `main.py` in a text editor.
   - Find the `<ApplicationName>` placeholder and replace it with the name of the music player you want to use.

7. Run the script:

   ```sh
   python main.py
   ```

If the configured music player is running and playing a track, the script will retrieve the track's name, artist, and album, and then copy this information to your clipboard.

## Configuration

To configure the script for a different music player application, open the `main.py` file and replace the `<ApplicationName>` placeholder with the actual name of the music player you want to use. Make sure to follow the same syntax and structure as provided in the script.

## Disclaimer

This script is provided as-is and may require adjustments to work with certain music player applications. The effectiveness of the script may depend on updates and changes made to the music player software.
