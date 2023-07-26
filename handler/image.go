package handler

import (
	"database/sql"
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

	username := c.Param("username")

	image, err := h.dbClient.FindImageByID(username, imageId)
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

	username := c.Param("username")

	image, err := h.dbClient.FindProcessedImageByID(username, imageId)
	if err != nil {
		log.Println(err)

		// img, _ := os.Open("./assets/processing.png")
		// imgData, _ := io.ReadAll(img)
		// c.Data(http.StatusOK, "image/png", imgData)
		c.JSON(http.StatusNotFound, gin.H{"status": "Not yet processed"})
		return
	}
	c.Data(http.StatusOK, "image/png", image.File)
}

func (h *Handler) GetImageInfo(c *gin.Context) {
	param := c.Param("imageId")
	imageId, err := strconv.Atoi(param)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	username := c.Param("username")

	image, err := h.dbClient.FindImageByID(username, imageId)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	status := false
	processedImage, err := h.dbClient.FindProcessedImageByID(username, imageId)
	if err == sql.ErrNoRows {
		log.Println("No processed image")
		status = false
	} else if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	} else if image.UP == processedImage.UP {
		status = true
	}

	c.JSON(http.StatusOK, gin.H{
		"fileName":   image.Filename,
		"size":       image.Size,
		"added_date": image.AddedDate,
		"up":         image.UP,
		"status":     status,
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

	username := c.Param("username")

	form, _ := c.MultipartForm()
	files := form.File["images"]

	err = h.dbClient.InsertImages(username, albumId, -1, files)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "images uploaded!"})
}

// TODO: image의 up 정보를 바꾸는 작업 필요
// up이 바뀌므로 이전에 작업된 processed image는 삭제
func (h *Handler) PatchImageUp(c *gin.Context) {
	param := c.Param("imageId")
	imageId, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	param = c.Param("newUp")
	newUp, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	username := c.Param("username")

	err = h.dbClient.UpdateImageUp(username, imageId, newUp)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "up changed"})
}
