package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Auth(c *gin.Context) {
	tokenString, err := c.Cookie("token")
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

	c.AddParam("username", username)
	c.Next()
}
