import os

# App
DEBUG = False
HOST = '0.0.0.0'
PORT = 5000

# Celery
CELERY_BROKER_URL = os.environ.get('CELERY_BROKER_URL') or 'redis://localhost:6379/0'
CELERY_RESULT_BACKEND = os.environ.get('CELERY_RESULT_BACKEND') or 'redis://localhost:6379/0'

# SMTP
MAIL_SERVER = 'smtp.gmail.com'
MAIL_PORT = 465
MAIL_USERNAME = '<YOUR-EMAIL-ADDRESS>'
MAIL_PASSWORD = '<YOUR-EMAIL-PASSWORD>'
MAIL_USE_TLS = False
MAIL_USE_SSL = True
