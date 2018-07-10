# Websocket
> A websocket service written in Golang.

### Installation
1. Go through the project directory then install the dependencies
    ```
    go get github.com/gorilla/websocket
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

### Usage
Put the codes into your frontend app
```js
// Client 1
// Connecting to websocket server
var email = "sohelamincse@gmail.com";
var mySocket = new WebSocket("ws://localhost:8081/ws?email=" + email);
// Recieving message
mySocket.onmessage = function (event) {
  console.log(event.data);
}

// Client 2
// Connecting to websocket server
var email = "sohel@sohelamin.com";
var mySocket = new WebSocket("ws://localhost:8081/ws?email=" + email);
// Recieving message
mySocket.onmessage = function (event) {
  console.log(event.data);
}

// Sending message to client 2
var msg = {
  message: "Say hi to client 2",
  recipient: "sohel@sohelamin.com"
};

mySocket.send(JSON.stringify(msg));

// Sending message to client 1
var msg = {
  message: "Say hi to client 1",
  recipient: "sohelamincse@gmail.com"
};

mySocket.send(JSON.stringify(msg));
```

Broadcast a message to all clients/subscribers
```
Hit http://localhost:8081/broadcast
```
