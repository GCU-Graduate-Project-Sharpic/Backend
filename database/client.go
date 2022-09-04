package database

import (
	"database/sql"
	"fmt"
	"log"
	"mime/multipart"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"github.com/GCU-Graduate-Project-Sharpic/Backend/types/image"
	"github.com/GCU-Graduate-Project-Sharpic/Backend/types/user"
)

type Client struct {
	config *Config
	db     *sql.DB
}

// Dial creates an instance of Client and dials the given postgresql.
func Dial(conf ...*Config) (*Client, error) {
	if len(conf) == 0 {
		conf = append(conf, NewConfig())
	} else if len(conf) > 1 {
		return nil, fmt.Errorf("too many arguments")
	}

	db, err := sql.Open("postgres", conf[0].PsqlConn())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Client{
		config: conf[0],
		db:     db,
	}, nil
}

func (c *Client) InsertNewUser(
	signupData user.User,
) error {
	encryptedPW, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return err
	}

	result, err := c.db.Exec(`INSERT INTO user_list (username, password, email) VALUES ($1, $2, $3);`, signupData.Username, string(encryptedPW), signupData.Email)
	if err != nil {
		return err
	}
	if _, err := result.RowsAffected(); err != nil {
		return err
	}

	return nil
}

func (c *Client) FindUserByUsername(
	username string,
) (*user.User, error) {
	rows, err := c.db.Query(`SELECT * FROM user_list WHERE username=$1`, username)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var userData user.User
	for rows.Next() {
		err := rows.Scan(&userData.Username, &userData.Password, &userData.Email)

		if err != nil {
			return nil, err
		}
	}

	return &userData, nil
}

func (c *Client) FindImageListByUsername(
	username string,
) ([]int, error) {
	rows, err := c.db.Query(`SELECT id FROM images WHERE username=$1`, username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var ids []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func (c *Client) FindImageByID(
	id int,
) (*image.Image, error) {
	rows, err := c.db.Query(`SELECT image_name, image_file, size, sr FROM images WHERE id=$1`, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	image := image.Image{}
	rows.Next()
	err = rows.Scan(&image.Filename, &image.File, &image.Size, &image.SR)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &image, nil
}

func (c *Client) StoreImages(
	username string,
	headers []*multipart.FileHeader,
) error {
	for _, header := range headers {
		image, err := image.FromFileHeader(header)
		if err != nil {
			return err
		}

		result, err := c.db.Exec(`INSERT INTO images (username, image_name, image_file, size, sr) VALUES ($1, $2, $3, $4, $5);`, username, image.Filename, image.File, image.Size, image.SR)
		if err != nil {
			return err
		}

		cnt, err := result.RowsAffected()
		if err != nil && cnt != 1 {
			return err
		}
		log.Println(image.Filename + "uploaded")
	}

	return nil
}
