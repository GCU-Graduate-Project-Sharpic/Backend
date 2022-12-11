package database

import (
	"database/sql"
	"fmt"
	"log"
	"mime/multipart"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"github.com/GCU-Sharpic/sharpic-server/types/album"
	"github.com/GCU-Sharpic/sharpic-server/types/image"
	"github.com/GCU-Sharpic/sharpic-server/types/user"
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
	signupData *user.User,
) error {
	encryptedPW, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}

	result, err := c.db.Exec(`INSERT INTO user_account (username, password, email) VALUES ($1, $2, $3);`, signupData.Username, string(encryptedPW), signupData.Email)
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err := result.RowsAffected(); err != nil {
		log.Println(err)
		return err
	}
	result, err = c.db.Exec(`INSERT INTO album (username, title) VALUES ($1, 'default');`, signupData.Username)
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err := result.RowsAffected(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *Client) FindUserByUsername(
	username string,
) (*user.User, error) {
	userData := user.User{}
	err := c.db.QueryRow(
		`SELECT * FROM user_account WHERE username=$1`,
		username,
	).Scan(
		&userData.Username,
		&userData.Password,
		&userData.Email,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no such user")
	} else if err != nil {
		log.Println(err)
		return nil, err
	}

	// if userData.Username == "" {
	// 	return nil, fmt.Errorf("no such user")
	// }

	return &userData, nil
}
func (c *Client) InsertNewAlbum(
	newAlbum *album.Album,
) error {
	result, err := c.db.Exec(`INSERT INTO album (username, title) VALUES ($1, $2);`, newAlbum.Username, newAlbum.Title)
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err := result.RowsAffected(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *Client) FindAlbumListByUsername(
	username string,
) ([]int, error) {
	rows, err := c.db.Query(`SELECT id FROM album WHERE username=$1`, username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ids := []int{}
	for rows.Next() {
		id := 0
		err := rows.Scan(&id)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func (c *Client) FindAlbumByID(
	albumId int,
) (*album.Album, error) {
	album := album.Album{}
	err := c.db.QueryRow(`SELECT username, title FROM album WHERE id=$1;`, albumId).Scan(&album.Username, &album.Title)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	album.ImageIds = []int{}
	rows, err := c.db.Query(`SELECT image_id FROM album_image WHERE album_id=$1;`, albumId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		id := 0
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		album.ImageIds = append(album.ImageIds, id)
	}
	return &album, nil
}

func (c *Client) FindImageByID(
	username string,
	imageId int,
) (*image.Image, error) {
	image := image.Image{}
	err := c.db.QueryRow(
		`SELECT image_name, image_file, size, added_date, up FROM image WHERE username=$1 AND id=$2;`,
		username,
		imageId,
	).Scan(
		&image.Filename,
		&image.File,
		&image.Size,
		&image.AddedDate,
		&image.UP,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &image, nil
}

func (c *Client) FindProcessedImageByID(
	username string,
	id int,
) (*image.Image, error) {
	image := image.Image{}
	err := c.db.QueryRow(
		`SELECT image_name, image_file, size, added_date, up FROM processed_image WHERE username=$1 AND id=$2;`,
		username,
		id,
	).Scan(
		&image.Filename,
		&image.File,
		&image.Size,
		&image.AddedDate,
		&image.UP,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &image, nil
}

func (c *Client) InsertImages(
	username string,
	albumId int,
	up int,
	headers []*multipart.FileHeader,
) error {
	for _, header := range headers {
		image, err := image.FromFileHeader(header, up)
		if err != nil {
			log.Println(err)
			return err
		}

		imageId := 0
		err = c.db.QueryRow(
			`INSERT INTO image (username, image_name, image_file, size, up) VALUES ($1, $2, $3, $4, $5) RETURNING id;`,
			username,
			image.Filename,
			image.File,
			image.Size,
			image.UP,
		).Scan(&imageId)
		if err != nil {
			log.Println(err)
			return err
		}

		// defaultId := 0
		// err = c.db.QueryRow(`SELECT id FROM album WHERE username=$1 AND title=$2;`, username, "default").Scan(&defaultId)
		// if err != nil {
		// 	log.Println(err)
		// 	return err
		// }

		result, err := c.db.Exec(
			`INSERT INTO album_image (album_id, image_id) VALUES ((SELECT id FROM album WHERE username=$1 AND title='default'), $2);`,
			username,
			imageId,
		)
		if err != nil {
			log.Println(err)
			return err
		}
		cnt, err := result.RowsAffected()
		if err != nil && cnt != 1 {
			log.Println(err)
			return err
		}

		if albumId != 0 {
			result, err = c.db.Exec(`INSERT INTO album_image (album_id, image_id) VALUES ($1, $2);`, albumId, imageId)
			if err != nil {
				log.Println(err)
				return err
			}
			cnt, err = result.RowsAffected()
			if err != nil && cnt != 1 {
				log.Println(err)
				return err
			}
		}
		log.Println(image.Filename + "uploaded")
	}

	return nil
}

func (c *Client) UpdateImageUp(
	username string,
	imageId int,
	up int,
) error {
	result, err := c.db.Exec(`UPDATE image SET up=$1 WHERE username=$2 AND id=$3;`, up, username, imageId)
	if err != nil {
		log.Println(err)
		return err
	}

	cnt, err := result.RowsAffected()
	if err != nil && cnt != 1 {
		log.Println(err)
		return err
	}

	result, err = c.db.Exec(`DELETE FROM processed_image WHERE username=$1 AND id=$2;`, username, imageId)
	if err != nil {
		log.Println(err)
		return err
	}

	cnt, err = result.RowsAffected()
	if err != nil && cnt != 1 {
		log.Println(err)
		return err
	}

	return nil
}
