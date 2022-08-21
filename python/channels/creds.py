import os
from google.auth.transport.requests import Request
from google.oauth2.credentials import Credentials
from google_auth_oauthlib.flow import InstalledAppFlow
from googleapiclient.discovery import build
from config import *

SCOPES = [
    'https://www.googleapis.com/auth/youtube.force-ssl',
    'https://www.googleapis.com/auth/userinfo.profile',
    'https://www.googleapis.com/auth/youtube.upload',
]

def get_creds_saved(app_token, user_token):
    creds = None

    if os.path.exists(user_token):
        creds = Credentials.from_authorized_user_file(user_token, SCOPES)
        creds.refresh(Request())

    if not creds or not creds.valid:
        flow = InstalledAppFlow.from_client_secrets_file(app_token, SCOPES)
        creds = flow.run_local_server(port=0)

    with open(user_token, 'w') as token:
        token.write(creds.to_json())

    return creds


def get_service_creds(app_token, user_token, service = 'youtube', version = 'v3'):
    creds = get_creds_saved(app_token, user_token)
    service = build(service, version, credentials=creds)
    return service


def get_service_simple(api_key, service = 'youtube', version = 'v3'):
    return  build(service, version, developerKey=api_key)