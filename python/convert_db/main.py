import os
import pymongo
import mysql.connector
from dotenv import load_dotenv
from bson import ObjectId


dotenv_path = '/home/q/p/python/projects/.env'
load_dotenv(dotenv_path)

IP_DB = os.getenv("INVITING_IP")
USER_DB = os.getenv("USER_DB")
PASSWORD_DB = os.getenv("INVITING_PASSWORD")
DATABASE = os.getenv("DATABASE")

TABLE_ACCOUNTS = 'accounts'

URL_DB = ""

db_mysql = mysql.connector.connect(host=IP_DB,
                                user=USER_DB,
                                password=PASSWORD_DB,
                                database=DATABASE)
sql_mysql = db_mysql.cursor(buffered=True)

client = pymongo.MongoClient(URL_DB)
db = client.services
accounts_coll = db.accounts


def get_accounts_folder(hash):
    sql_mysql.execute(f"SELECT * FROM {TABLE_ACCOUNTS} WHERE hash_folder='{hash}'")
    accounts = sql_mysql.fetchall()

    return accounts


def transport_accounts(accounts, hash):
    for account in accounts:
        data_account = {
            "name": account[1], 
            "phone": account[2], 
            "folder": ObjectId(hash), 
            "api_id": account[3], 
            "api_hash": account[4],
            "verify": account[7],
            "launch": account[8],
            "interval": account[10],
            "status_block": account[11],
            "random_hash": account[6],
            "phone_code_hash": account[5]

        }
        accounts_coll.insert_one(data_account)
        print(f"[INFO] Создан {account[1]}")


def main():
    accounts = get_accounts_folder("")
    transport_accounts(accounts, "")


if __name__ == "__main__":
    main()