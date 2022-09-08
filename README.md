# services-backend

The server part of the site with services for advertising and promotion of goods. At the moment there is a service for inviting and mailing in telegram. To do this, you will have to add telegram accounts on the site, which will perform these functions

![GO][go-version] ![Python][python-version]

---
## Installation

#### Requirements
* Golang 1.17  
* Python3.8+
* Linux, Windows or macOS

#### Installing
```
git clone https://github.com/Services-combine/services-backend.git
cd services-backend
```

#### Configure
To work, you must create a `.env` file in the main directory of the project and specify such variables as:
```
MONDO_DB_URL - link to mongodb database
SALT - a combination of characters to generate a password hash
SECRET_KEY - key for generating authentication tokens
FRONTEND_URL - the link from which the request will come from the frontend
FOLDER_PYTHON_SCRIPTS_VERIFY - the path to the folder where the script with account verification lies, in the project it is located along the path services-backend/python/account-telethon
FOLDER_ACCOUNTS - the path to the folder where the .session files of the added accounts will be located, in the project it is located along the path services-backend/accounts
FOLDER_CHANNELS - the path to the folder where the tokens of the applications and clients of the added channels will be located, in the project it is located along the path services-backend/channels
```

Also, in the `configs/config.yml` file, specify your mongodb login and the name of the database

---
## Usage
The port on which the service will be launched is specified in the file `configs/config.yml`

To start, run
```
go build -o services-backend cmd/app/main.go
./services-backend
```

---
## Additionally
A `services-backend.service` file was also created to run this bot on the server


[go-version]: https://img.shields.io/static/v1?label=GO&message=v1.17&color=blue
[python-version]: https://img.shields.io/static/v1?label=Python&message=v3.8&color=blue