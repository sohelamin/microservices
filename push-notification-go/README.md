# Push Notification
> A push notification service written in Golang.

### Pre-requisites
- HTTPS enabled site
- [Service Worker](https://developers.google.com/web/fundamentals/primers/service-workers/) (Client)
- [Web Push](https://github.com/sherclockholmes/webpush-go) (Server)

### Installation
1. Go through the project directory then install the dependencies
    ```
    go get github.com/sherclockholmes/webpush-go
    ```
2. Run the program
    ```
    go run main.go
    ```

### Docker
You can deploy the service using docker
```
docker-compose up -d
```

## Usage
1. Serve the client app with any http server
    ```
    cd client
    python -m SimpleHTTPServer 8888
    ```
2. Visit the client app and allow notification
    ![Push Notification](https://user-images.githubusercontent.com/1708683/42136238-64a2e2bc-7d79-11e8-95c4-80afdd9fa66b.png)

3. Send notification to the subscribed users
    Simple notification
    ```bash
    curl -XPOST -H 'Content-Type: application/json' http://localhost:8082/send -d '
    {
        "title": "Push Notification",
        "body": "Hey Dude!",
        "icon": "icon.png",
        "url": "https://www.appzcoder.com/"
    }
    '
    ```
    You will get the notification like this
    ![Push Notification](https://user-images.githubusercontent.com/1708683/42136289-50e94008-7d7a-11e8-9e60-7fe9484ae3ba.png)
