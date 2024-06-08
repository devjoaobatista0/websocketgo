package server

import (
	"websocket/internal/models"
	"log"
	"golang.org/x/net/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)

func HandleConnections(ws *websocket.Conn) {
	defer ws.Close()
	clients[ws] = true

	for {
		var msg models.Message
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			log.Printf("Err reaging message: %v", err)
			delete(clients, ws)
			break


		}
		broadcast <- msg
	}
}
