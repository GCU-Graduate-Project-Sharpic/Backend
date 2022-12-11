package album

import (
	"github.com/gin-gonic/gin"
)

type Album struct {
	Id       int    `json:"id"`
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
