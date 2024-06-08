package server

import (
	"log"
	"net/http"
	
	"golang.org/x/net/websocket"
)

func StartServer() {
	dir := "C:/Users/jvara/Documents/PROJETOS/Websocket/public"
	
    
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)
	http.Handle("/ws", websocket.Handler(HandleConnections))

	go HandleMessages()

	log.Println("Server started on port :8081")
	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		log.Fatal("Err to started the server", err)
	}
}