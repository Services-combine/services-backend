import os
import logging

if not os.path.exists('python/logs'):
	os.mkdir('python/logs')

logging.basicConfig(filename="python/logs/errors.log", format = u'[%(levelname)s][%(asctime)s] %(funcName)s:%(lineno)s: %(message)s', level='INFO')
logger = logging.getLogger()
