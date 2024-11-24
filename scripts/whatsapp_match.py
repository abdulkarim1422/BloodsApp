from dotenv import load_dotenv
load_dotenv()

import os
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
import json
# import keyboard as k # Windows

def send_whatsapp(data_file_json, message_file_text, x_cord=958, y_cord=968):
    with open(data_file_json, 'r', encoding='utf-8') as f:
        data = json.load(f)
    
    donors = data['donors']
    files = message_file_text
    with open(files, encoding='utf-8') as f:
        file_data = f.read()
    
    counter = 0
    for donor in donors:
        try:
            name = donor['first_name'] + " " + donor['last_name']
            blood = donor['bloodType']
            phone_number = donor['phone_number']
            patient_name = data['patient_first_name'] + " " + data['patient_last_name']
            message = data['message']
            msg = file_data.format(name, blood, message)
            web.open(f"https://web.whatsapp.com/send?phone={phone_number}&text={quote(msg)}")
            if counter == 0:
                time.sleep(20)
            time.sleep(10)

            pyautogui.click(x_cord, y_cord)
            time.sleep(2)

            pyautogui.click('enter')
            time.sleep(4)

            pyautogui.click('ctrl+w')
            time.sleep(1)
            pyautogui.click('enter')

            counter += 1
            print(counter, "-Message sent..!!", f"{phone_number} - {name} - {blood} / {patient_name}\n")
            with open("scripts/donors_sent.txt", "a") as x:
                x.write(f"{phone_number} - {name} - {blood} / patient: {patient_name} - {time.strftime('%Y-%m-%d %H:%M:%S')}\n")
        except Exception as e:
            with open("scripts/donors_sent_error.txt", "a") as error_file:
                error_file.write(f"{phone_number} - {name} - {blood} / patient: {patient_name} - {time.strftime('%Y-%m-%d %H:%M:%S')}\n")
            print(f"Failed to send message to {phone_number}: {e}")
    print("Done!")

if __name__ == "__main__":
    data_file_json = sys.argv[1]
    message_file_text = sys.argv[2]
    try:
        x_cord: = int(sys.argv[3])
    except:
        pass
    try:
        y_cord = int(sys.argv[4])
    except:
        pass
    send_whatsapp(data_file_json, message_file_text, x_cord, y_cord)