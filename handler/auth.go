package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Auth(c *gin.Context) {

	// TODO: Using sessions library

	// session := sessions.Default(c)
	// user := session.Get("user")
	// if user == nil {
	// 	log.Println("User not logged in")
	// 	c.Redirect(http.StatusFound, "/login")
	// 	c.Abort()
	// 	return
	// }

	tokenString, err := c.Cookie("token")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = h.tokenManager.TokenValid(tokenString)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	username, err := h.tokenManager.ExtractTokenUsername(tokenString)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := h.dbClient.FindUserByUsername(username)
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
