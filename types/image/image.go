package image

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"mime/multipart"
	"time"
)

type Image struct {
	Filename  string
	Hash      string
	File      []byte
	Size      int64
	AddedDate time.Time
	UP        int
}

func FromFileHeader(header *multipart.FileHeader, up int) (*Image, error) {
	file, err := header.Open()
	if err != nil {
		return nil, err
	}
	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Transform image file to sha256 hash with crypto/sha256
	hasher := sha256.New()
	hasher.Write(fileData)
	hash := hasher.Sum(nil)
	hashString := hex.EncodeToString(hash)

	return &Image{
		Filename: header.Filename,
		Hash:     hashString,
		File:     fileData,
		Size:     header.Size,
		UP:       up,
	}, nil
}
