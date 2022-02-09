import argparse
import asyncio
from telethon.sync import TelegramClient
from config import logger


def get_arguments():
    parser = argparse.ArgumentParser()
    parser.add_argument('-P', '--phone', type=str, help='Phone')
    parser.add_argument('-H', '--hash', type=str, help='Api hash')
    parser.add_argument('-I', '--id', type=int, help='Api id')
    parser.add_argument('-C', '--code', type=int, help='Code telegram')
    parser.add_argument('-G', '--pch', type=str, help='Phone code hash')

    options = parser.parse_args()
    return options

def verify_account(phone, hash, id, code, phone_code_hash):
    try:
        loop = asyncio.new_event_loop()
        asyncio.set_event_loop(loop)
        
        client = TelegramClient(f"accounts/{phone}", id, hash)
        client.connect()
        client.sign_in(phone, code, phone_code_hash=phone_code_hash)
        client.disconnect()
        print("SUCCESS")
    except Exception as error:
        logger.error(f"[{phone}] {error}")
        print("ERROR")

def main():
    options = get_arguments()
    verify_account(options.phone.strip(), options.hash.strip(), options.id, options.code, options.pch.strip())


if __name__ == "__main__":
    main()