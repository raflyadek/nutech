package util

import (
	"io"
	"mime/multipart"
	"os"
)

func SaveUploadFile(file *multipart.FileHeader, path string) error {
	//open the upload file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	//destination
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	//copy the uploaded file to the destination
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}