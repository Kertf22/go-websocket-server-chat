package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan []byte)            // broadcast channel

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	// get the client IP address
	ip := r.RemoteAddr + ":"
	fmt.Println("Client IP address: ", ip)
	if err != nil {
		fmt.Println(err)
		return
	}

	// register client
	clients[conn] = true

	// sendNumberOfClients(conn)
	broadcastMessage(conn, ip)
}

func sendNumberOfClients(conn *websocket.Conn) {
	for {

		// Use conn.ReadMessage to read messages from the client
		messageType, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// client into a string
		numberOfClientsString := strconv.Itoa(len(clients))
		// convert string to byte array
		numberOfClientsByte := []byte(numberOfClientsString)
		// Use conn.WriteMessage to send messages to the client
		if err := conn.WriteMessage(messageType, numberOfClientsByte); err != nil {

			fmt.Println(err)
			return
		}
	}
}

func broadcastMessage(conn *websocket.Conn, ip string) {
	for {
		// Use conn.ReadMessage to read messages from the client
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Change the message to add the client IP address in the front
		message = append([]byte(ip), message...)

		// // grab the next message from the broadcast channel
		// message := <-broadcast

		// send it out to every client that is currently connected
		for client := range clients {

			// if the client is the one that sent the message, don't send it back to them
			if client == conn {
				continue
			}

			if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
				fmt.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
