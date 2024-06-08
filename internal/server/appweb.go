package server

import (
	"log"
	"golang.org/x/net/websocket"
)

func HandleMessages(){
	for {
		msg := <- broadcast
		for client := range clients {
			err := websocket.JSON.Send(client, msg)
			if err != nil {
				log.Print("Err sensing message: ")
				client.Close()
				delete(clients, client)
			}
		}
	}
}