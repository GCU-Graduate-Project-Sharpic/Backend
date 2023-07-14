package handler

import (
	"log"
	"os"
	"strconv"

	"github.com/GCU-Sharpic/sharpic-server/database"
	"github.com/GCU-Sharpic/sharpic-server/types/token"
)

type Handler struct {
	dbClient     *database.Client
	tokenManager *token.Token
}

func New() *Handler {
	dbClient, err := database.Dial()
	if err != nil {
		log.Println(err)
		return nil
	}

	tokenLifespan, err := strconv.Atoi(os.Getenv("JWT_TOKEN_LIFESPAN"))
	if err != nil {
		log.Println(err)
		return nil
	}

	tokenManager := &token.Token{
		SecretKey:     os.Getenv("JWT_SECRET"),
		TokenLifespan: tokenLifespan,
	}

	return &Handler{
		dbClient:     dbClient,
		tokenManager: tokenManager,
	}
}
