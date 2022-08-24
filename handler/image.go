package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetImage(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error"})
		c.Abort()
		return
	}
	image, err := h.dbClient.FindImageByID(id)
	c.Data(http.StatusOK, "image/png", image.File)
}

func (h *Handler) GetImageList(c *gin.Context) {
	cookie, err := c.Cookie("userId")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error"})
		c.Abort()
		return
	}

	list, err := h.dbClient.FindImageListByUsername(cookie)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": list})
}

func (h *Handler) PostImage(c *gin.Context) {
	cookie, err := c.Cookie("userId")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error"})
		c.Abort()
		return
	}

	form, _ := c.MultipartForm()
	files := form.File["images"]
	ids, err := h.dbClient.StoreImages(cookie, files)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("%d files uploaded!", len(ids))})
}
