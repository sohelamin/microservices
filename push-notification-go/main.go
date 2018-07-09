package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	webpush "github.com/sherclockholmes/webpush-go"
)

const (
	vapidPrivateKey = "ql8hU1yB-T48dBGzI9vapTav4IX8n-WqNXURKHhMFLs"
)

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Icon  string `json:"icon"`
	Url   string `json:"url"`
}

var subscriptions = make([]*webpush.Subscription, 0)

func subscribe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	var subscription *webpush.Subscription
	jsonErr := json.NewDecoder(r.Body).Decode(&subscription)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	subscriptions = append(subscriptions, subscription)
	fmt.Fprintf(w, "Subscribed")
}

func sendNotification(w http.ResponseWriter, r *http.Request) {
	var notification Notification
	jsonErr := json.NewDecoder(r.Body).Decode(&notification)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	data, _ := json.Marshal(notification)

	// Send Notification
	for _, subscription := range subscriptions {
		_, sendErr := webpush.SendNotification([]byte(string(data)), subscription, &webpush.Options{
			Subscriber:      "mailto:<sohelamincse@gmail.com>",
			VAPIDPrivateKey: vapidPrivateKey,
		})
		if sendErr != nil {
			log.Fatal(sendErr)
		}
	}

	fmt.Fprintf(w, "Sent")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "There is no place like home.")
	})

	http.HandleFunc("/subscribe", subscribe)
	http.HandleFunc("/send", sendNotification)

	log.Println("http server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
