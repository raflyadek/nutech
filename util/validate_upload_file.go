package util

import (
	"errors"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

var allowedMimeTypes = map[string]bool{
	"image/png":  true,
	"image/jpeg": true,
}

var allowedExtensions = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
}

func ValidateImage(file *multipart.FileHeader) error {
	//validate extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		return errors.New("only PNG or JPEG images are allowed")
	}

	//validate MIME type
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		return err
	}

	contentType := http.DetectContentType(buffer)
	if !allowedMimeTypes[contentType] {
		return errors.New("Format Image tidak sesuai")
	}

	return nil
}