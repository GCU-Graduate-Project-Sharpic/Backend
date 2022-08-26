package image

import (
	"io/ioutil"
	"mime/multipart"
)

type Image struct {
	Filename string
	File     []byte
	Size     int64
	SR       bool
}

func FromFileHeader(header *multipart.FileHeader) (*Image, error) {
	file, err := header.Open()
	if err != nil {
		return nil, err
	}
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &Image{
		Filename: header.Filename,
		File:     fileData,
		Size:     header.Size,
		SR:       false,
	}, nil
}
