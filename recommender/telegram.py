#-*- coding:utf-8 -*-
import json
import os
import requests
import logging
from dotenv import load_dotenv
load_dotenv()

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
    data = {"text": strMessage.encode("utf8"), "chat_id": CHAT_ID}
    requests.post(url, data)

def read_file():
    data = ''
    with open('recommend.txt') as f:
        data += f.read()

    send_message(data)


read_file()
