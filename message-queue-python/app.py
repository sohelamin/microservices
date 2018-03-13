from flask import Flask, request, jsonify
from celery import Celery
import os
from flask_mail import Mail, Message

app = Flask(__name__)
app.config.from_pyfile('config.py')

celery = Celery(
    app.name,
    broker=app.config['CELERY_BROKER_URL'],
    backend=app.config['CELERY_RESULT_BACKEND']
)

mail = Mail(app)

@app.route('/')
def home():
    return jsonify({'message': 'There is no place like home.'})

@app.route('/emails', methods=['POST'])
def emails():
    # send to the queue
    for email in request.json['emails']:
        process_email.delay(email)

    return jsonify({'message': 'Queued'})

@celery.task
def process_email(email):
    with app.app_context():
        # send the email
        msg = Message(
            email['subject'],
            sender=app.config['MAIL_USERNAME'],
            recipients=[email['recipient']]
        )
        msg.body = email['message']
        mail.send(msg)
    return 'Sent'

if __name__ == "__main__":
    app.run(host=app.config['HOST'], debug=app.config['DEBUG'], port=app.config['PORT'])
