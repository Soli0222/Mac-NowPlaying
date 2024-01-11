import subprocess
import pyperclip

import os
from os.path import join, dirname
from dotenv import load_dotenv

import requests

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
    
    track_name, artist_name, album_name = [s.strip() for s in stdout.decode('utf-8').split(',', 2)]

    return track_name, artist_name, album_name

if __name__ == "__main__":
    load_dotenv(verbose=True)
    dotenv_path = join(dirname(__file__), '.env')
    load_dotenv(dotenv_path)

    track_name, artist_name, album_name = get_current_track_info()
    text = track_name + " / " + artist_name + "\n" + album_name+ "\n#NowPlaying #PsrPlaying"
    
    if os.environ.get("SHARE_MISSKEY") == "True":
        HOST = os.environ.get("MISSKEY_HOST")
        TOKEN = os.environ.get("MISSKEY_TOKEN")

        url = "https://"+HOST+"/api/notes/create"

        headers = {
            "Content-Type": "application/json",
        }

        data = {
            "i": TOKEN,
            "text": text,
        }

        response = requests.post(url, headers=headers, json=data)

    else:
        pyperclip.copy(text)