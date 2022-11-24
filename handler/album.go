package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/GCU-Sharpic/sharpic-server/types/album"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAlbumList(c *gin.Context) {
	cookie, err := c.Cookie("username")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	list, err := h.dbClient.FindAlbumListByUsername(cookie)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": list})
}

func (h *Handler) GetAlbum(c *gin.Context) {
	param := c.Param("albumId")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	album, err := h.dbClient.FindAlbumByID(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	encodedAlbum, err := album.JSON()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		c.Abort()
		return
	}

	// json 반환 되는지 확인 필요. 안되면 gin.H{"response": encodedAlbum} 시도하기
	c.JSON(http.StatusOK, string(encodedAlbum))
}

func (h *Handler) PostNewAlbum(c *gin.Context) {
	newAlbum, err := album.NewShouldBindJSON(c)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.dbClient.InsertNewAlbum(newAlbum); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "insert success"})
}
