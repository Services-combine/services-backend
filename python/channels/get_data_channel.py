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
        print(data['items'][0])
    except Exception as error:
        logger.error(f"[{channelId}] {error}")
        print(f"ERROR: {error}")


def main():
    options = get_arguments()
    get_data_channel(options.channelId, options.apiKey)


if __name__ == "__main__":
    main()