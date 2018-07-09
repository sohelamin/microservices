FROM golang:latest

COPY . /code
WORKDIR /code

RUN go get github.com/sherclockholmes/webpush-go
RUN go build -o main .

CMD ["/code/main"]
