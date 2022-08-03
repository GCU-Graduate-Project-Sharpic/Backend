package main

import (
	"github.com/chromato99/go-react-test-web-app/userHandler"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("/Frontend", true)))

	userApi := router.Group("/user")
	{
		userApi.GET("/", userHandler.User)
		userApi.POST("signup", userHandler.Signup)
		userApi.POST("/login", userHandler.Login)
		userApi.POST("/logout", userHandler.Logout)
	}

	router.Run(":8005")
}
