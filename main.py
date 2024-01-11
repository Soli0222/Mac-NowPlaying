import subprocess
import pyperclip
import re

import os
from os.path import join, dirname
from dotenv import load_dotenv

import requests

def subprocessScript(script):
    p = subprocess.Popen(['osascript', '-'], stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    stdout, stderr = p.communicate(script.encode('utf-8'))
    if stderr:
        raise Exception(stderr.decode('utf-8').strip())
    if "missing value" in stdout.decode('utf-8').strip():
        return None
    return stdout.decode('utf-8')

def get_current_track_info():
    script_track = '''
        tell application "Swinsian"
            if it is running then
                if player state is playing then
                    set track_name to name of current track
                    return track_name
                end if
            end if
        end tell
    '''
    script_artist = '''
        tell application "Swinsian"
            if it is running then
                if player state is playing then
                    set artist_name to artist of current track
                    return artist_name
                end if
            end if
        end tell
    '''
    script_album = '''
        tell application "Swinsian"
            if it is running then
                if player state is playing then
                    set album_name to album of current track
                    return album_name
                end if
            end if
        end tell
    '''
    

    track_name = re.sub("\n", "", subprocessScript(script_track))
    artist_name = re.sub("\n", "", subprocessScript(script_artist))
    album_name = re.sub("\n", "", subprocessScript(script_album))

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