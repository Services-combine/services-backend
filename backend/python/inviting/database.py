from telnetlib import EC
import pymongo
from config import *


class Database:

    def __init__(self):
        client = pymongo.MongoClient(URL_DB)
        self.db = client.services
        self.folders = self.db.folders
        self.accounts = self.db.accounts
        self.settings = self.db.settings

    
    def get_folder_launch_inviting(self):
        try:
            folder = self.folders.find_one({"inviting": True})
            return folder
        except Exception as error:
            logger.error(error)
            return 0


    def get_folder_launch_mailing_usernames(self):
        try:
            folder = self.folders.find_one({"mailing_usernames": True})
            return folder
        except Exception as error:
            logger.error(error)
            return 0

    
    def get_folder_launch_mailing_groups(self):
        try:
            folder = self.folders.find_one({"mailing_groups": True})
            return folder
        except Exception as error:
            logger.error(error)
            return 0


    def get_accounts_folder(self, folderID):
        try:
            accounts = self.accounts.find({"folder": folderID})
            return accounts
        except Exception as error:
            logger.error(error)
            return 0


    def get_settings(self):
        try:
            settings = self.settings.find_one()
            return settings
        except Exception as error:
            logger.error(error)
            return 0

    
    def suspend_account(self, id):
        try:
            self.accounts.update_one({"_id": id}, {"$set": {"launch": False}})
            return 1
        except Exception as error:
            logger.error(error)
            return 0


    def suspend_folder(self, id, mode):
        try:
            self.folders.update_one({"_id": id}, {"$set": {mode: False}})
            return 1
        except Exception as error:
            logger.error(error)
            return 0


    def add_remaining_usernames(self, id, usernames):
        try:
            self.folders.update_one({"_id": id}, {"$set": {"usernames": usernames}})
            return 1
        except Exception as error:
            logger.error(error)
            return 0