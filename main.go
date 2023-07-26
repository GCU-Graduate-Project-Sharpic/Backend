package main

import (
	"github.com/GCU-Sharpic/sharpic-server/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	handler := handler.New()

	api := router.Group("/api")
	{
		api.POST("/signup", handler.PostSignup)
		api.POST("/login", handler.PostLogin)

		api.Use(handler.Auth)

		api.POST("/logout", handler.PostLogout)

		userApi := api.Group("/user")
		{
			userApi.GET("", handler.GetUserData)
		}

		albumApi := api.Group("/album")
		{
			albumApi.GET("/list", handler.GetAlbumList)
			albumApi.GET("/:albumId", handler.GetAlbum)
			albumApi.POST("/new", handler.PostNewAlbum)
			// albumApi.POST("/remove/:albumId", handler.PostRemoveAlbum)
		}

		imageApi := api.Group("/image")
		{
			imageApi.GET("/:imageId", handler.GetImage)
			imageApi.GET("/processed/:imageId", handler.GetProcessedImage)
			imageApi.GET("/info/:imageId", handler.GetImageInfo)
			imageApi.POST("/new/:albumId", handler.PostNewImage)
			imageApi.PATCH("/up/:imageId/:newUp", handler.PatchImageUp)
			// imageApi.POST("/remove/:imageId", handler.PostRemoveImage)
		}
	}

	router.Run(":8005")
}
