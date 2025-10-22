package utils

import (
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"os"

	"github.com/chai2010/webp"
)

func ToWebp(file *multipart.File) (string, error) {
	img, _, err := image.Decode(*file)
	if err != nil {
		return "", err
	}

	output, err := os.CreateTemp("", "image-*.webp")
	if err != nil {
		return "", errors.New("Failed to create temporary file")
	}
	defer output.Close()

	if err := webp.Encode(output, img, &webp.Options{Lossless: false, Quality: 80}); err != nil {
		return "", errors.New("Failed to encode image")
	}
	return output.Name(), nil
}
