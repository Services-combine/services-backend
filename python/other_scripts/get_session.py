from telethon.sync import TelegramClient


API_ID = 0
API_HASH = ''
NAME_SESSION = 'vanya'

client = TelegramClient(NAME_SESSION, API_ID, API_HASH)
client.start()
