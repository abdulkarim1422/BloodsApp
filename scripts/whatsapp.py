import sys
import keyboard as k
import pyautogui
import time
import json
import webbrowser as web
from urllib.parse import quote
from PIL import Image

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
            time.sleep(10)
            pyautogui.click(x_cord, y_cord)
            time.sleep(2)
            k.press_and_release('enter')
            time.sleep(4)
            k.press_and_release('ctrl+w')
            time.sleep(1)
            k.press_and_release('enter')
            counter += 1
            print(counter, "-Message sent..!!", f"{phone_number} - {name} - {blood} / {patient_name}\n")
            with open("done.txt", "a") as x:
                x.write(f"{phone_number} - {name} - {blood} / {patient_name}\n")
        except Exception as e:
            with open("error.txt", "a") as error_file:
                error_file.write(f"{phone_number} - {name} - {blood} / {patient_name}\n")
            print(f"Failed to send message to {phone_number}: {e}")
    print("Done!")


# data_file_json = "body.json"
# message_file_text = "message.txt"
# send_whatsapp(data_file_json, message_file_text)

if __name__ == "__main__":
    data_file_json = sys.argv[1]
    message_file_text = sys.argv[2]
    x_cord = int(sys.argv[3])
    y_cord = int(sys.argv[4])
    send_whatsapp(data_file_json, message_file_text, x_cord, y_cord)
