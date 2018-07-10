package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Message   string `json:"message"`
	Recipient string `json:"recipient"`
}

type Subscriber struct {
	Email  string
	Socket *websocket.Conn
}

var subscribers = make([]Subscriber, 0)
var messageChannel = make(chan Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	email := r.FormValue("email")
	subscriber := Subscriber{Email: email, Socket: ws}
	subscribers = append(subscribers, subscriber)
	fmt.Println(subscribers)

	for {
		var msg Message

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			// Delete the last subscriber just appended
			subscribers = subscribers[:len(subscribers)-1]
			break
		}

		messageChannel <- msg
	}
}

func handleMessages() {
	for {
		msg := <-messageChannel

		for subscriberIndex, subscriber := range subscribers {
			// Send only to the recipient
			if subscriber.Email == msg.Recipient {
				err := subscriber.Socket.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					subscriber.Socket.Close()
					// Delete the subscriber
					subscribers = append(subscribers[:subscriberIndex], subscribers[subscriberIndex+1:]...)
				}
			}
		}
	}
}

func handleBroadcast(w http.ResponseWriter, r *http.Request) {
	msg := "Announcement for all"

	for subscriberIndex, subscriber := range subscribers {
		err := subscriber.Socket.WriteJSON(msg)
		if err != nil {
			log.Printf("error: %v", err)
			subscriber.Socket.Close()
			// Delete the subscriber
			subscribers = append(subscribers[:subscriberIndex], subscribers[subscriberIndex+1:]...)
		}
	}

	fmt.Fprintf(w, "Broadcasted")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "There is no place like home.")
	})

	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/broadcast", handleBroadcast)

	go handleMessages()

	log.Println("http server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
