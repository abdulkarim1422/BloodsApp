from dotenv import load_dotenv
import os
load_dotenv()

try:
    os.environ['DISPLAY'] = os.getenv('DISPLAY') # echo $DISPLAY # ps aux | grep Xorg
except:
    print("DISPLAY not set, using default Display")
    pass

import sys
import pyautogui
import time
import pandas as pd
import webbrowser as web
from urllib.parse import quote
from PIL import Image
# import keyboard as k # Windows

def send_whatsapp (phone, msg, txt="py", x_cord=958, y_cord=968):
    try:
        web.open(f"https://web.whatsapp.com/send?phone={phone}&text={quote(msg)}")
        time.sleep (20)

        pyautogui.click(x_cord, y_cord)
        time.sleep (2)

        pyautogui.press('enter')
        time.sleep (4)

        pyautogui.hotkey('ctrl+w')
        time.sleep (1)
        pyautogui.hotkey('enter')
        
        print ("-Message sent..!!", f"{phone} - {msg.replace('\n', ' ')[:10]}\n")
        with open(f"scripts/done_{txt}.txt", "a") as x:
            x.write(f"{phone} - {msg.replace('\n', ' ')[:10]} - {time.strftime('%Y-%m-%d %H:%M:%S')}\n")
    except Exception as e:
        with open(f"scripts/error_{txt}.txt", "a") as error_file:
            error_file.write(f"{phone} - {msg.replace('\n', ' ')[:10]} - {time.strftime('%Y-%m-%d %H:%M:%S')}\n")
        print(f"Failed to send message to {phone}: {e}")
    print ("Done!")


if __name__ == "__main__":
    phone = sys.argv[1]
    message = sys.argv[2]
    txt_file_name = sys.argv[3]
    x_cord = os.getenv('x_cord')
    y_cord = os.getenv('y_cord')
    send_whatsapp(phone, message, txt_file_name, x_cord, y_cord)
