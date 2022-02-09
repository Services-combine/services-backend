import time
from telethon.sync import TelegramClient
from telethon.tl.functions.channels import InviteToChannelRequest
from database import Database
from config import *


ERROR_PRIVACY = "The user's privacy settings"
ERROR_MANY_REQUEST = "Too many requests"
ERROR_NO_ADMIN = "You can't write in this chat"


class Inviting:   

    def __init__(self):
        try:
            self.db = Database() 
            logger.info("Success connect to MySQL")
        except Exception as error:
            logger.error(error)


    def inviting_users(self, account, usernames, chat):
        try:
            client = TelegramClient(f"{FOLDER_ACCOUNTS}/{account['phone']}.session", account['api_id'], account['api_hash'])
            client.connect()
            id_channel = client.get_entity(chat)
            id_channel = id_channel.to_dict()['id']
    
            for user in usernames:
                try:
                    logger.info(f"Start added @{user}")
                    client(InviteToChannelRequest(channel=id_channel, users=[user]))
                    logger.info(f"Success added @{user}")

                except Exception as error:
                    if ERROR_PRIVACY in str(error):
                        logger.error(f"PRIVATE USER {user}")
                        if user != usernames[-1]:
                            time.sleep(account['interval'])
                        continue
                    elif ERROR_MANY_REQUEST in str(error):
                        logger.error(f"TOO MANY REQUEST {account['name']}")
                        break
                    elif ERROR_NO_ADMIN in str(error):
                        logger.error(f"NO ADMIN ACCOUNT IN CHAT {account['name']}")
                        break
                    else:
                        logger.error(error)

                if user != usernames[-1]:
                    time.sleep(account['interval'])

            client.disconnect()
        except Exception as error:
            logger.error(error)


    def mailing_users(self, account, usernames, message):
        try:
            client = TelegramClient(f"{FOLDER_ACCOUNTS}/{account['phone']}.session", account['api_id'], account['api_hash'])
            client.connect()
    
            for user in usernames:
                try:
                    logger.info(f"Start send message @{user}")
                    client.send_message(f'@{user}', message)
                    logger.info(f"Success send message @{user}")
                    
                except Exception as error:
                    if ERROR_PRIVACY in str(error):
                        logger.error(f"PRIVATE USER {user}")
                        if user != usernames[-1]:
                            time.sleep(account['interval'])
                        continue
                    elif ERROR_MANY_REQUEST in str(error):
                        logger.error(f"TOO MANY REQUEST {account['name']}")
                        break
                    else:
                        logger.error(error)

                if user != usernames[-1]:
                    time.sleep(account['interval'])

            client.disconnect()
        except Exception as error:
            logger.error(error)


    def mailing_group(self, account, groups, message):
        try:
            client = TelegramClient(f"{FOLDER_ACCOUNTS}/{account['phone']}.session", account['api_id'], account['api_hash'])
            client.connect()
    
            for group in groups:
                try:
                    logger.info(f"Start send message {group}")
                    client.send_message(entity=group, message=message)
                    logger.info(f"Success send message {group}")
                    
                except Exception as error:
                    if ERROR_MANY_REQUEST in str(error):
                        logger.error(f"TOO MANY REQUEST {account['name']}")
                        break
                    else:
                        logger.error(error)

                if group != groups[-1]:
                    time.sleep(account['interval'])

            client.disconnect()
        except Exception as error:
            logger.error(error)


    def preparation_inviting(self, folder):
        try:
            usernames = folder["usernames"]
            chat = folder["chat"]
            accounts = self.db.get_accounts_folder(folder["_id"])
            settings = self.db.get_settings()
            
            for account in accounts:
                if account["interval"] != None and account["interval"] != 0 and account["launch"] and account["verify"] and account["status_block"] == 'clean':
                    try:
                        if usernames != []:
                            logger.info(f'Start inviting {account["name"]}')

                            if len(usernames) <= settings["countInviting"]:
                                self.inviting_users(account, usernames, chat)
                                usernames = []
                            else:
                                self.inviting_users(account, usernames[:settings["countInviting"]], chat)
                                usernames = usernames[settings["countInviting"]:]
                        else:
                            break
                    except Exception as error:
                        logger.error(error)
                        break

                logger.info(f"Change status launch account {account['name']}")
                result_change = self.db.suspend_account(account['_id'])
                if result_change == 0:
                    continue
            
            logger.info(f"Change status launch folder {folder['name']}")
            result_suspend = self.db.suspend_folder(folder['_id'], "inviting")
            if result_suspend == 0:
                logger.error(result_suspend)

            result_add_remainig_usernames = self.db.add_remaining_usernames(folder['_id'], usernames)
            if result_add_remainig_usernames == 0:
                logger.error(result_add_remainig_usernames)

        except Exception as error:
            logger.error(error)
            print(f"[ERROR] {error}")


    def preparation_mailing_usernames(self, folder):
        try:
            usernames = folder["usernames"]
            message = folder["message"]
            accounts = self.db.get_accounts_folder(folder["_id"])
            settings = self.db.get_settings()
            
            for account in accounts:
                if account["interval"] != None and account["interval"] != 0 and account["launch"] and account["verify"] and account["status_block"] == 'clean':
                    try:
                        if usernames != []:
                            logger.info(f"Start mailing {account['name']}")

                            if len(usernames) <= settings['countMailing']:
                                self.mailing_users(account, usernames, message)
                                usernames = []
                            else:
                                self.mailing_users(account, usernames[:settings['countMailing']], message)
                                usernames = usernames[settings['countMailing']:]
                        else:
                            break
                    except Exception as error:
                        logger.error(error)
                        break

                logger.info(f"Change status launch account {account['name']}")
                result_change = self.db.suspend_account(account['_id'])

                if result_change == 0:
                    continue

            result_add_remainig_usernames = self.db.add_remaining_usernames(folder['_id'], usernames)
            if result_add_remainig_usernames == 0:
                logger.error(result_add_remainig_usernames)

            logger.info(f"Change status launch folder {folder['name']}")
            result_suspend = self.db.suspend_folder(folder['_id'], 'mailing_usernames')
            if result_suspend == 0:
                logger.error(result_suspend)

        except Exception as error:
            logger.error(error)
            print(f'[ERROR] {error}')


    def preparation_mailing_groups(self, folder):
        try:
            groups = folder['groups']
            message = folder["message"]
            accounts = self.db.get_accounts_folder(folder["_id"])
            settings = self.db.get_settings()
                
            for account in accounts:
                if account["interval"] != None and account["interval"] != 0 and account["launch"] and account["verify"] and account["status_block"] == 'clean':
                    try:
                        logger.info(f"Start mailing {account['name']}")
                        self.mailing_group(account, groups, message)
                    except Exception as error:
                        logger.error(error)
                        break

                logger.info(f"Change status launch account {account['name']}")
                result_change = self.db.suspend_account(account['_id'])

                if result_change == 0:
                    continue

            logger.info(f"Change status launch folder {folder['name']}")
            result_suspend = self.db.suspend_folder(folder['_id'], 'mailing_groups')
            if result_suspend == 0:
                logger.error(result_suspend)

        except Exception as error:
            logger.error(error)
            print(f'[ERROR] {error}')


    def start_check(self):
        try:
            while True:
                folder = self.db.get_folder_launch_inviting()
                if folder != None:
                    self.preparation_inviting(folder)

                folder = self.db.get_folder_launch_mailing_usernames()
                if folder != None:
                    self.preparation_mailing_usernames(folder)

                folder = self.db.get_folder_launch_mailing_groups()
                if folder != None:
                    self.preparation_mailing_groups(folder)

                time.sleep(15)
        except Exception as error:
            logger.error(error)


inviting = Inviting()
inviting.start_check()