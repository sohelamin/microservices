FROM golang:latest

COPY . /code
WORKDIR /code

RUN go get github.com/gorilla/websocket
RUN go build -o main .

CMD ["/code/main"]
