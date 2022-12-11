package handler

import (
	"log"

	"github.com/GCU-Sharpic/sharpic-server/database"
)

type Handler struct {
	dbClient *database.Client
}

func New() *Handler {
	dbClient, err := database.Dial()
	if err != nil {
		log.Println(err)
		return nil
	}

	return &Handler{
		dbClient: dbClient,
	}
}
