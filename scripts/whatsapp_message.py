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
import webbrowser as web
from urllib.parse import quote
from filelock import FileLock, Timeout

def send_whatsapp(phone, msg, txt="py", x_cord=958, y_cord=968):
    lock_file = "whatsapp.lock"
    lock = FileLock(lock_file, timeout=1000)

    try:
        with lock:
            try:
                web.open(f"https://web.whatsapp.com/send?phone={phone}&text={quote(msg)}")
                time.sleep(20)

                pyautogui.click(x_cord, y_cord)
                time.sleep(2)

                pyautogui.press('enter')
                time.sleep(4)

                pyautogui.hotkey('ctrl+w')
                time.sleep(1)
                pyautogui.hotkey('enter')

                print("-Message sent..!!", f"{phone} - {msg.replace('\n', ' ')[:10]}\n")
                with open(f"scripts/done_{txt}.txt", "a") as x:
                    x.write(f"{phone} - {msg.replace('\n', ' ')[:10]} - {time.strftime('%Y-%m-%d %H:%M:%S')}\n")
            except Exception as e:
                with open(f"scripts/error_{txt}.txt", "a") as error_file:
                    error_file.write(f"{phone} - {msg.replace('\n', ' ')[:10]} - {time.strftime('%Y-%m-%d %H:%M:%S')}\n")
                print(f"Failed to send message to {phone}: {e}")
    except Timeout:
        print("Another instance of whatsapp is running. Waiting for it to finish...")
        lock.acquire()
        send_whatsapp(phone, msg, txt, x_cord, y_cord)
    finally:
        if lock.is_locked:
            lock.release()

if __name__ == "__main__":
    phone = sys.argv[1]
    message = sys.argv[2]
    txt_file_name = sys.argv[3]
    x_cord = os.getenv('x_cord')
    y_cord = os.getenv('y_cord')
    send_whatsapp(phone, message, txt_file_name, x_cord, y_cord)
