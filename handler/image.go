package handler

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetImage(c *gin.Context) {
	param := c.Param("imageId")
	imageId, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}
	cookie, err := c.Cookie("username")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	image, err := h.dbClient.FindImageByID(cookie, imageId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"status": "error"})
		c.Abort()
		return
	}
	c.Data(http.StatusOK, "image/png", image.File)
}

func (h *Handler) GetProcessedImage(c *gin.Context) {
	param := c.Param("imageId")
	imageId, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}
	cookie, err := c.Cookie("username")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	image, err := h.dbClient.FindProcessedImageByID(cookie, imageId)
	if err != nil {
		log.Println(err)

		img, _ := os.Open("./assets/processing.png")
		imgData, _ := io.ReadAll(img)
		c.Data(http.StatusOK, "image/png", imgData)
		return
	}
	c.Data(http.StatusOK, "image/png", image.File)
}

func (h *Handler) GetImageInfo(c *gin.Context) {
	param := c.Param("imageId")
	imageId, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}
	cookie, err := c.Cookie("username")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	image, err := h.dbClient.FindImageByID(cookie, imageId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"fileName":   image.Filename,
		"size":       image.Size,
		"added_date": image.AddedDate,
		"up":         image.UP,
	})
}

func (h *Handler) PostNewImage(c *gin.Context) {
	param := c.Param("albumId")
	albumId, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}
	param = c.Param("up")
	up, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	cookie, err := c.Cookie("username")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	form, _ := c.MultipartForm()
	files := form.File["images"]

	err = h.dbClient.InsertImages(cookie, albumId, up, files)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "images uploaded!"})
}
