package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	cookie, err := c.Cookie("username")
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
