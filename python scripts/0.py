import keyboard as k
import pyautogui
import time
import pandas as pd
import webbrowser as web
from urllib.parse import quote
from PIL import Image
# import win32clipboard
# import io


def send_whatsapp (data_file_excel,message_file_text,x_cord=958,y_cord=968):
    df = pd.read_excel(data_file_excel, dtype={str: str})
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

excel_path=r"D:\work\24-Q4\python tests\contacts.xlsx"
msg_path=r"D:\work\24-Q4\python tests\message.txt"

send_whatsapp(excel_path,msg_path)
