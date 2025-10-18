package database

import (
	"errors"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func InitCloudinary() (*cloudinary.Cloudinary, error) {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, errors.New("Failed to initialize cloudinary client: " + err.Error())
	}
	log.Println("Cloudinary client initialized")
	return cld, nil
}
