import re
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

def check_block(phone, hash, id):
    try:
        status_text = asyncio.run(get_status_block(phone, hash, id))

        if status_text == 0:
            status = "ERROR"
        elif "UTC" in status_text:
            if "автоматически сняты" in status_text:
                status = re.split(r'автоматически сняты', status_text)[1].strip()
            else:
                status = re.split(r'automatically released on', status_text)[1].strip()

            status = re.split(r'UTC', status)[0]
        else:
            status = "clean"

        print(status, end="")
    except Exception as error:
        logger.error(f"[{phone}] {error}")
        print("ERROR")

async def get_status_block(phone, hash, id):
    try:
        BOT = "@SpamBot"

        path_to_file = FOLDER_ACCOUNTS + f"{phone}.session"
        client = TelegramClient(path_to_file, id, hash)
        await client.connect()
        await client.send_message(entity=BOT, message="/start")

        async for message in client.iter_messages(BOT):
            answer = message.text
            break

        await client.disconnect()
        return answer
    except Exception as error:
        logger.error(f"[{phone}] {error}")
        return 0

def main():
    options = get_arguments()
    check_block(options.phone.strip(), options.hash.strip(), options.id)


if __name__ == "__main__":
    main()