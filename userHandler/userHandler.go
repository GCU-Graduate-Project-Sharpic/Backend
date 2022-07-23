package userHandler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type loginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type user struct {
	Username string
	Password string
	Email    string
}

func User(c *gin.Context) {
	username := checkLogin(c)

	if username != nil {
		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), 5432, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

		db, err := sql.Open("postgres", psqlconn)
		if err != nil {
			fmt.Println(err)
		}

		defer db.Close()

		rows, err := db.Query(`SELECT * FROM user_list WHERE username=$1`, username)

		if err != nil {
			fmt.Println(err)
		}

		var userData user
		for rows.Next() {
			err := rows.Scan(&userData.Username, &userData.Password, &userData.Email)

			if err != nil {
				fmt.Println(err)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"username": userData.Username,
			"email":    userData.Email,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"username": "Not Login",
			"email":    "",
		})
	}

}

func Login(c *gin.Context) {
	var inputLoginData loginData

	if err := c.ShouldBindJSON(&inputLoginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := checkLogin(c)
	if username != nil {
		c.JSON(http.StatusOK, gin.H{"status": "you are already logged in"})
		return
	} else {
		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), 5432, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

		db, err := sql.Open("postgres", psqlconn)
		if err != nil {
			fmt.Println(err)
		}

		defer db.Close()

		rows, err := db.Query(`SELECT * FROM user_list WHERE username=$1`, inputLoginData.Username)

		if err != nil {
			fmt.Println(err)
		}

		var userData user
		for rows.Next() {
			err := rows.Scan(&userData.Username, &userData.Password, &userData.Email)

			if err != nil {
				fmt.Println(err)
			}
		}

		if inputLoginData.Username != userData.Username || inputLoginData.Password != userData.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.SetCookie("userId", userData.Username, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		return
	}
}

func Logout(c *gin.Context) {
	c.SetCookie("userId", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "you are logged out"})
}

func checkLogin(c *gin.Context) *string {
	cookie, err := c.Cookie("userId")

	if err != nil {
		return nil
	} else {
		return &cookie
	}
}
