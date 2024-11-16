# MacNowPlaying (mnp)

This tool allows you to share your currently playing track information in various ways. You can display the track information, copy it to the clipboard, or post it directly to social media platforms like Twitter and Misskey.

By setting the environment variable `MNP_APP_NAME`, you can retrieve information from any application. If the environment variable `MNP_APP_NAME` is not set, the default application name "Swinsian" will be used.

## Installation

You can run the program using the pre-built binary without needing to build the source code.

1. Download the latest binary for your OS from the [releases page](https://github.com/Soli0222/Mac-NowPlaying/releases).
2. Extract the downloaded binary and make it executable (if necessary).

   ```bash
   chmod +x mnp
   ```

3. Move it to a directory in your PATH or use it directly as a command.

## Usage

### Display the Current Track Information

The `nowplaying` command retrieves and displays the current track information from your music player. It also copies the track information to the clipboard for easy sharing.

```sh
# Display the current track information
nowplaying
```

### Post the Currently Playing Track to Twitter

The `tweet` command allows you to share your currently playing track directly to Twitter. It opens a web browser with a pre-filled tweet containing the track information.

```sh
# Post the currently playing track to Twitter
tweet
```

### Post the Currently Playing Track to Misskey

The `misskey` command allows you to share your currently playing track directly to Misskey. You need to configure the command by setting the environment variables `MISSKEY_HOST` and `MISSKEY_TOKEN`.

```sh
# Post the currently playing track to Misskey
export MISSKEY_HOST=your.misskey.server.tld
export MISSKEY_TOKEN=your_access_token
misskey
```

## Environment Variables

- `MNP_APP_NAME`: The name of the application to retrieve track information from. If not set, the default application name "Swinsian" will be used.
- `MISSKEY_HOST`: The host of your Misskey instance.
- `MISSKEY_TOKEN`: Your access token for Misskey.

## Example

When you run the `nowplaying` command and the retrieved track information is as follows:

- Track Name: 灼熱にて純情(wii-wii-woo)
- Artist: 星街すいせい
- Album Name: Specter

The following text will be copied to your clipboard:

```plaintext
灼熱にて純情(wii-wii-woo) / 星街すいせい
Specter
#NowPlaying #PsrPlaying
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
