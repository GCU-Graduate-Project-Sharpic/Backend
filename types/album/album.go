package album

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Album struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	ImageIds []int  `json:"imageIds"`
}

func NewShouldBindJSON(c *gin.Context) (*Album, error) {
	new := Album{}
	if err := c.ShouldBindJSON(&new); err != nil {
		return nil, err
	}
	return &new, nil
}

func (a *Album) JSON() ([]byte, error) {
	encodedAlbum, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		return nil, err
	}
	return encodedAlbum, nil
}
