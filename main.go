package main

import (
	"github.com/GCU-Graduate-Project-Sharpic/Backend/handler"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userHandler := handler.New()

	router.Use(static.Serve("/", static.LocalFile("/Frontend", true)))

	router.POST("/login", userHandler.PostLogin)
	router.POST("/signup", userHandler.PostSignup)

	router.Use(userHandler.SessionAuth)

	router.POST("logout", userHandler.PostLogout)

	userApi := router.Group("/user")
	{
		userApi.GET("/", userHandler.GetUserData)
	}

	router.Run(":8005")
}
