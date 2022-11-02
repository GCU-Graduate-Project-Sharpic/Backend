package handler

import (
	"log"
	"net/http"

	"github.com/GCU-Graduate-Project-Sharpic/Backend/database"
	"github.com/gin-gonic/gin"
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

func (h *Handler) SessionAuth(c *gin.Context) {

	// TODO: Using sessions library

	// session := sessions.Default(c)
	// user := session.Get("user")
	// if user == nil {
	// 	log.Println("User not logged in")
	// 	c.Redirect(http.StatusFound, "/login")
	// 	c.Abort()
	// 	return
	// }

	cookie, err := c.Cookie("userId")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := h.dbClient.FindUserByUsername(cookie)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if user == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}
