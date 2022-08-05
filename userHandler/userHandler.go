package userHandler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type loginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type user struct {
	Username string `json:"username" binding:"resuired"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), 5432, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

func User(c *gin.Context) {
	username := checkLogin(c)

	if username != nil {
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

func Signup(c *gin.Context) {
	var inputSignupData user

	if err := c.ShouldBindJSON(&inputSignupData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert to password excluding \n character at the end
	encryptedPW, err := bcrypt.GenerateFromPassword([]byte(inputSignupData.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	id := 0
	err = db.QueryRow(`INSERT INTO user_list (username, password, email) VALUES ($1, $2, $3);`, inputSignupData.Username, string(encryptedPW), inputSignupData.Email).Scan(&id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "signup fail", "error": err})
	}

	c.JSON(http.StatusOK, gin.H{"status": "signup success", "id": id})
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

		// compare password with sotred password
		bcryptErr := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(inputLoginData.Password))

		if inputLoginData.Username != userData.Username || bcryptErr != nil {
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
