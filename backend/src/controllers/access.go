package controllers

import (
	"encoding/json"
	"log"

	"github.com/matheusrf96/go-webserver/backend/src/db"
	"github.com/matheusrf96/go-webserver/backend/src/models"
)

func HandleAccess(data []byte) {
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

	// repo := repositories.NewAccessRepository(db)
	// err = repo.CreateDayTable()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	log.Println(access)

	return
}
