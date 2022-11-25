package main

import (
	"os"

	"github.com/GCU-Sharpic/sharpic-server/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	handler := handler.New(os.Getenv("DOMAIN"))

	router.POST("/signup", handler.PostSignup)
	router.POST("/login", handler.PostLogin)

	router.Use(handler.SessionAuth)

	router.POST("/logout", handler.PostLogout)

	userApi := router.Group("/user")
	{
		userApi.GET("/", handler.GetUserData)
	}

	albumApi := router.Group("/album")
	{
		albumApi.GET("/list", handler.GetAlbumList)
		albumApi.GET("/:albumId", handler.GetAlbum)
		albumApi.POST("/new", handler.PostNewAlbum)
		// albumApi.POST("/remove/:albumId", handler.PostRemoveAlbum)
	}

	imageApi := router.Group("/image")
	{
		imageApi.GET("/:imageId", handler.GetImage)
		imageApi.GET("/processed/:imageId", handler.GetProcessedImage)
		imageApi.GET("/info/:imageId", handler.GetImageInfo)
		imageApi.POST("/new/:albumId/:up", handler.PostNewImage)
		// imageApi.POST("/remove/:imageId", handler.PostRemoveImage)
	}

	router.Run(":8005")
}
