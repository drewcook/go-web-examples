package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Define the websocket handler for gorilla/websocket
// Set 1MB buffer sizes
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func echo(w http.ResponseWriter, r *http.Request) {
	// Establish connection to web socket
	conn, _ := upgrader.Upgrade(w, r, nil)

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Do something with the incoming message
		// In this case, log to console the message and who sent it
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write the message back to the browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func serve(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "websockets.html")
}

func main() {
	http.HandleFunc("/", serve)
	http.HandleFunc("/echo", echo)
	http.ListenAndServe(":80", nil)
}
