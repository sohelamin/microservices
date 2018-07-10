# Message Queue
> A message queue service written in Python using Celery & Redis.

### Installation
1. Go through the project directory then install the dependencies
    ```
    pip install -r requirements.txt
    ```
2. Set the configs on the `config.py` file
3. Run the Celery on a different tab
    ```
    celery worker -A app.celery --loglevel=info
    ```
4. If you want to use Flower to monitor the tasks then install and run (optional)
    ```
    celery -A app.celery flower
    ```

### Docker
You can deploy the service using docker
```
docker-compose up -d
```

### Usage
Make a POST request to execute any queues
```
curl -XPOST 'http://localhost:8080/emails/' -H 'Content-Type: application/json' -d '
{
    "emails": [
        {
            "message": "Hi, This message is came from a message queue.",
            "recipient": "sohel@sohelamin.com",
            "subject": "Greetings"
        },
        {
            "message": "Hi, This message is came from a message queue.",
            "recipient": "sohelamincse@gmail.com",
            "subject": "Greetings"
        }
    ]
}
'
```
