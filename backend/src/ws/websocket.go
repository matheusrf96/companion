package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/matheusrf96/go-webserver/backend/src/db"
	"github.com/matheusrf96/go-webserver/backend/src/models"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		db, err := db.Connect()
		if err != nil {
			log.Println(err)
			return
		}
		defer db.Close()

		var access models.Access

		err = json.Unmarshal(data, &access)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(access)

		err = conn.WriteMessage(messageType, data)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.RemoteAddr, r.UserAgent())

	reader(ws)
}
