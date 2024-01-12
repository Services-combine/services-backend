import os
import argparse
import asyncio
from telethon.sync import TelegramClient
from config import *


def get_arguments():
    parser = argparse.ArgumentParser()
    parser.add_argument('-P', '--phone', type=str, help='Phone')
    parser.add_argument('-H', '--hash', type=str, help='Api hash')
    parser.add_argument('-I', '--id', type=int, help='Api id')

    options = parser.parse_args()
    return options

def send_code_account(phone, hash, id):
    try:
        loop = asyncio.new_event_loop()
        asyncio.set_event_loop(loop)

        path_to_file = f"{FOLDER_ACCOUNTS}{phone}.session"
        client = TelegramClient(path_to_file, id, hash)
        client.connect()
        client.send_code_request(phone)

        phone_code_hash = client.send_code_request(phone).phone_code_hash
        client.disconnect()

        print(phone_code_hash, end="")
    except Exception as error:
        logger.error(f"[{phone}] {error}")
        print("ERROR")

def main():
    options = get_arguments()
    send_code_account(options.phone.strip(), options.hash.strip(), options.id)


if __name__ == "__main__":
    main()