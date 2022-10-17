package main

import (
	"os"

	"github.com/GCU-Graduate-Project-Sharpic/Backend/handler"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	handler := handler.New(os.Getenv("DOMAIN"))

	router.GET("/", handler.SessionAuth, static.Serve("/", static.LocalFile("/Frontend", true)))
	router.Use(static.Serve("/", static.LocalFile("/Frontend", true)))

	router.GET("/login", static.Serve("/login", static.LocalFile("/Frontend", true)))
	router.GET("/signup", static.Serve("/signup", static.LocalFile("/Frontend", true)))

	router.POST("/login", handler.PostLogin)
	router.POST("/signup", handler.PostSignup)

	router.Use(handler.SessionAuth)

	router.POST("/logout", handler.PostLogout)

	userApi := router.Group("/user")
	{
		userApi.GET("/", handler.GetUserData)
	}

	imageApi := router.Group("/image")
	{
		imageApi.GET("/:id", handler.GetImage)
		imageApi.GET("/list", handler.GetImageList)
		imageApi.POST("/", handler.PostImage)
	}

	router.Run(":8005")
}
