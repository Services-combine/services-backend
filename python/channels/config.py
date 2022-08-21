import os
import logging
from dotenv import load_dotenv


dotenv_path = '../../.env'
load_dotenv(dotenv_path)

FOLDER_CHANNELS = os.getenv("FOLDER_CHANNELS")
PYTHON_SCRIPTS = os.getenv("FOLDER_PYTHON_SCRIPTS_CHANNELS")

if not os.path.exists(f'{PYTHON_SCRIPTS}logs'):
	os.mkdir(f'{PYTHON_SCRIPTS}logs')

logging.basicConfig(filename=f"{PYTHON_SCRIPTS}logs/errors.log", format = u'[%(levelname)s][%(asctime)s] %(funcName)s:%(lineno)s: %(message)s', level='INFO')
logger = logging.getLogger()
