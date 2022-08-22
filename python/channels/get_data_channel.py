import argparse
from config import *
from creds import *

def get_arguments():
    parser = argparse.ArgumentParser()
    parser.add_argument('-C', '--channelId', type=str, help='Channel id')
    parser.add_argument('-A', '--apiKey', type=str, help='Api key')

    options = parser.parse_args()
    return options


def get_data_channel(channelId, apiKey):
    try:
        data = get_service_simple(apiKey, "youtube", "v3").channels().list(id=channelId, part='snippet, statistics').execute()

        title = data['items'][0]['snippet']['title']
        description = data['items'][0]['snippet']['description']
        photo = data['items'][0]['snippet']['thumbnails']['default']['url']
        viewCount = data['items'][0]['statistics']['viewCount']
        subscriberCount = data['items'][0]['statistics']['subscriberCount']
        videoCount = data['items'][0]['statistics']['videoCount']

        data_string = "{" + \
                    "\"title\":\"" + title + \
                    "\",\"description\":\"" + description + \
                    "\",\"photo\":\"" + photo + \
                    "\",\"viewCount\":" + viewCount + \
                    ",\"subscriberCount\":" + subscriberCount + \
                    ",\"videoCount\":" + videoCount + \
                    "}"
        print(data_string)
    except Exception as error:
        logger.error(f"[{channelId}] {error}")
        print(f"ERROR: {error}")


def main():
    options = get_arguments()
    '''get_service_creds(
        app_token = f"{FOLDER_CHANNELS}app_token_{options.channelId}.json",
        user_token = f"{FOLDER_CHANNELS}user_token_{options.channelId}.json",
        service = 'youtube', 
        version = 'v3'
    )'''
    get_data_channel(options.channelId, options.apiKey)


if __name__ == "__main__":
    main()