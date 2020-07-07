#-*- coding:utf-8 -*-
import json
import os
import requests

CHAT_ID = os.environ['TELEGRAM_CHAT_ID']
TOKEN = os.environ['TELEGRAM_TOKEN']
BASE_URL = "https://api.telegram.org/bot{}".format(TOKEN)

def send_message(message):
    url = BASE_URL + "/sendMessage"
    strMessage = message
    if type(message) == list:
        strMessage = ''
        for m in message:
            strMessage += m
            strMessage += "\n"
    data = {"text": strMessage.encode("utf8"), "chat_id": CHAT_ID, "parse_mode":"Markdown"}
    requests.post(url, data)
    print("Send complete")

def read_file(filepath):
    data = ''
    with open(filepath) as f:
        data += f.read()

    send_message(data)
