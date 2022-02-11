import os
import pymongo
from collections import Counter


PATH = "/home/q/p/projects/services/backend/accounts"
URL_DB = ""

client = pymongo.MongoClient(URL_DB)
db = client.services
accounts_coll = db.accounts

def get_numbers():
	list_phone = accounts_coll.find({}, {"_id": 0, "phone": 1})
	return list_phone


def main():
    list_phone = [i["phone"] for i in get_numbers()]
    list_files = [i.split(".")[0] for i in os.listdir(PATH)]
    print(len(list_phone), len(list_files))

    counter = Counter(list_phone)

    '''for i in list_files:
        if i not in list_phone:
            os.remove(f"{PATH}/{i}.session")
            print(f"Remove {i}.session")'''


    for i in list_phone:
        if i not in list_files:
            print(i)


if __name__ == "__main__":
    main()