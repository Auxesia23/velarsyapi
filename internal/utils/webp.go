package utils

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/chai2010/webp"
)

// ToWebp convert image into a webp format.
// It receive two parameters, data is the data to be converted and
// dest is a destination for the new converted data to be writen to.
// If there is an error, it will return error, otherwise it return nil
func ToWebp(data io.Reader, dest io.Writer) error {
	img, _, err := image.Decode(data)
	if err != nil {
		return err
	}

	return webp.Encode(dest, img, &webp.Options{Lossless: false, Quality: 80})
}
