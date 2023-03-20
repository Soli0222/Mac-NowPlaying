import subprocess
import pyperclip

def get_current_track_info():
    script = '''
        tell application "Swinsian"
            if it is running then
                if player state is playing then
                    set track_name to name of current track
                    set artist_name to artist of current track
                    set album_name to album of current track
                    return {track_name, artist_name, album_name}
                end if
            end if
        end tell
    '''
    p = subprocess.Popen(['osascript', '-'], stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    stdout, stderr = p.communicate(script.encode('utf-8'))
    if stderr:
        raise Exception(stderr.decode('utf-8').strip())
    if "missing value" in stdout.decode('utf-8').strip():
        return None
    
    # 中括弧を除去する代わりに、split()関数で余分な空白文字を削除する
    track_name, artist_name, album_name = [s.strip() for s in stdout.decode('utf-8').split(',', 2)]

    return track_name, artist_name, album_name

if __name__ == "__main__":
    track_name, artist_name, album_name = get_current_track_info()
    pyperclip.copy(track_name + " / " + artist_name + "\n#NowPlaying")