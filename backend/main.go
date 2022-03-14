package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matheusrf96/go-webserver/backend/src/ws"
)

func setupRoutes() {
	http.HandleFunc("/", ws.HomePage)
	http.HandleFunc("/ws", ws.WsEndpoint)
}

func main() {
	fmt.Println("Go websocket")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
