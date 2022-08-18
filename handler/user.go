package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GCU-Graduate-Project-Sharpic/Backend/database"
	"github.com/GCU-Graduate-Project-Sharpic/Backend/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	dbClient *database.Client
}

func New() *UserHandler {
	dbClient, err := database.Dial()
	if err != nil {
		log.Println(err)
		return nil
	}
	return &UserHandler{
		dbClient: dbClient,
	}
}

func (h *UserHandler) PostSignup(c *gin.Context) {
	signupData, err := user.NewShouldBindJSON(c)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.dbClient.InsertNewUser(*signupData); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "signup success"})
}

func (h *UserHandler) PostLogin(c *gin.Context) {
	loginData, err := user.NewShouldBindJSON(c)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userData, err := h.dbClient.FindUserByUsername(loginData.Username)

	// compare password with sotred password
	bcryptErr := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(loginData.Password))

	if loginData.Username != userData.Username || bcryptErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.SetCookie("userId", userData.Username, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "login success"})
	return
}

func (h *UserHandler) PostLogout(c *gin.Context) {
	c.SetCookie("userId", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "logout success"})
}

func (h *UserHandler) GetUserData(c *gin.Context) {
	username, err := c.Cookie("userId")
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

func (h *UserHandler) SessionAuth(c *gin.Context) {
	// session := sessions.Default(c)
	// user := session.Get("user")
	// if user == nil {
	// 	log.Println("User not logged in")
	// 	c.Redirect(http.StatusMovedPermanently, "/login")
	// 	c.Abort()
	// 	return
	// }

	cookie, err := c.Cookie("userId")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		c.Abort()
		return
	}

	user, err := h.dbClient.FindUserByUsername(cookie)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		c.Abort()
		return
	}

	if user == nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}
