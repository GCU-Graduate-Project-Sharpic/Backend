package handler

import (
	"log"
	"net/http"
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
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}
	c.Data(http.StatusOK, "image/png", image.File)
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

	c.JSON(http.StatusOK, gin.H{"status": "files uploaded!"})
}
