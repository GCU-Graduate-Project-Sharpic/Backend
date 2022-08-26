package user

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func NewShouldBindJSON(c *gin.Context) (*User, error) {
	var new User

	if err := c.ShouldBindJSON(&new); err != nil {
		return nil, err
	}
	return &new, nil
}
