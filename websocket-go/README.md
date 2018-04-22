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
// Connecting to websocket server
var mySocket = new WebSocket("ws://localhost:8080/ws");

// Sending message
mySocket.onopen = function (event) {
  var msg = {
    email: "sohelamincse@gmail.com",
    message: "Hello"
  };

  mySocket.send(JSON.stringify(msg));
};

// Recieving message
mySocket.onmessage = function (event) {
  console.log(event.data);
}
'
```
