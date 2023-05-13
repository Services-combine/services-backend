import argparse
import asyncio
from telethon.sync import TelegramClient
from telethon.tl.functions.channels import JoinChannelRequest
from config import *


def get_arguments():
    parser = argparse.ArgumentParser()
    parser.add_argument('-P', '--phone', type=str, help='Phone')
    parser.add_argument('-H', '--hash', type=str, help='Api hash')
    parser.add_argument('-I', '--id', type=int, help='Api id')
    parser.add_argument('-G', '--group', type=str, help='Group')

    options = parser.parse_args()
    return options

async def join_channel(phone, hash, id, group):
    try:
        path_to_file = FOLDER_ACCOUNTS + f"{phone}.session"
        client = TelegramClient(path_to_file, id, hash)
        await client.connect()
        await client(JoinChannelRequest(group))
        await client.disconnect()
        print("OK")
    except Exception as error:
        logger.error(f"[{phone}] {error}")
        print("ERROR")

def main():
    options = get_arguments()
    asyncio.run(join_channel(options.phone.strip(), options.hash.strip(), options.id, options.group))


if __name__ == "__main__":
    main()