package image

import (
	"io"
	"mime/multipart"
)

type Image struct {
	Filename string
	File     []byte
	Size     int64
	UP       int
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

	return &Image{
		Filename: header.Filename,
		File:     fileData,
		Size:     header.Size,
		UP:       up,
	}, nil
}
