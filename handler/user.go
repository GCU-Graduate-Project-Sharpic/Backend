package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GCU-Sharpic/sharpic-server/types/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) PostSignup(c *gin.Context) {
	signupData, err := user.NewShouldBindJSON(c)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.dbClient.InsertNewUser(signupData); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "signup success"})
}

func (h *Handler) PostLogin(c *gin.Context) {
	loginData, err := user.NewShouldBindJSON(c)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userData, err := h.dbClient.FindUserByUsername(loginData.Username)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// compare password with sotred password
	bcryptErr := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(loginData.Password))

	if loginData.Username != userData.Username || bcryptErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.SetCookie("username", userData.Username, 3600, "/", h.domain, false, true)
	c.JSON(http.StatusOK, gin.H{"status": "login success"})
}

func (h *Handler) PostLogout(c *gin.Context) {
	c.SetCookie("username", "", -1, "/", h.domain, false, true)

	c.JSON(http.StatusOK, gin.H{"status": "logout success"})
}

func (h *Handler) GetUserData(c *gin.Context) {
	username, err := c.Cookie("username")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	userData, err := h.dbClient.FindUserByUsername(username)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(userData)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	c.JSON(http.StatusOK, string(data))
}
