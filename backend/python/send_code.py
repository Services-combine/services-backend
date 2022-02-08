import argparse
import asyncio
from telethon.sync import TelegramClient


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
        client = TelegramClient(f"accounts/{phone}", id, hash)
        client.connect()
        client.send_code_request(phone)

        phone_code_hash = client.send_code_request(phone).phone_code_hash
        print(phone_code_hash, end="")
        client.disconnect()
    except Exception as error:
        print(f"[ERROR] {error}")

def main():
    options = get_arguments()
    send_code_account(options.phone.strip(), options.hash.strip(), options.id)


if __name__ == "__main__":
    main()