package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusrf96/go-webserver/backend/src/config"
	"github.com/matheusrf96/go-webserver/backend/src/ws"
)

func setupRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Ok") })
	r.HandleFunc("/ws", ws.WsEndpoint)

	return r
}

func main() {
	config.Load()

	log.Printf("Go Websocket running at the port :%d", config.Port)

	r := mux.NewRouter()
	r = setupRoutes(r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
