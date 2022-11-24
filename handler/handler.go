package handler

import (
	"log"

	"github.com/GCU-Sharpic/sharpic-server/database"
)

type Handler struct {
	dbClient *database.Client
	domain   string
}

func New(domain string) *Handler {
	dbClient, err := database.Dial()
	if err != nil {
		log.Println(err)
		return nil
	}

	if domain == "" {
		domain = "localhost"
	}

	return &Handler{
		dbClient: dbClient,
		domain:   domain,
	}
}
