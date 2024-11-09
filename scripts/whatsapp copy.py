import sys
import keyboard as k
import pyautogui
import time
import pandas as pd
import webbrowser as web
from urllib.parse import quote
from PIL import Image
# import win32clipboard
# import io

def send_whatsapp (data_file_csv,message_file_text,x_cord=958,y_cord=968):
    df = pd.read_csv(data_file_csv, dtype={str: str})
    name = df['Name'].values
    contact = df['Number'].values
    user = df['username'].values
    pass0 = df['pass'].values
    files = message_file_text

    with open(files, encoding='utf-8') as f:
        file_data = f.read()
    zipped = zip(name, contact, user, pass0)
    counter = 0
    for (a,b,c,d) in zipped:
        try:
            msg=file_data.format (a,c,d)
            web.open(f"https://web.whatsapp.com/send?phone={b}&text={quote(msg)}")
            # if counter == 0:
            #     time.sleep(60)
            time.sleep (10)
            #adjust duration if required
            pyautogui.click(x_cord, y_cord)
            time.sleep (2)
            k.press_and_release('enter')
            time.sleep (4)
            k.press_and_release('ctrl+w')
            time.sleep(1)
            k.press_and_release('enter')
            counter+=1
            print (counter, "-Message sent..!!")
            with open("done.txt", "a") as x:
                x.write(f"{b}\n")
        except Exception as e:
            with open("error.txt", "a") as error_file:
                error_file.write(f"{b}\n")
            print(f"Failed to send message to {b}: {e}")
    print ("Done!")

# csv_path=r"contacts.csv"
# msg_path=r"message.txt"

# send_whatsapp(csv_path, msg_path)


if __name__ == "__main__":
    data_file_csv = sys.argv[1]
    message_file_text = sys.argv[2]
    x_cord = int(sys.argv[3])
    y_cord = int(sys.argv[4])
    send_whatsapp(data_file_csv, message_file_text, x_cord, y_cord)

