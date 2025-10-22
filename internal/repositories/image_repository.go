package repositories

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Auxesia23/velarsyapi/internal/models"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gorm.io/gorm"
)

type ImageRepository interface {
	Upload(ctx context.Context, file *os.File) (*string, error)
	Create(ctx context.Context, projectId *uint, url *string) (*models.Image, error)
	GetByProjectID(ctx context.Context, projectId *uint) (*[]models.Image, error)
	Delete(ctx context.Context, id *uint) error
}

type imageRepository struct {
	cld *cloudinary.Cloudinary
	db  *gorm.DB
}

func NewImageRepository(cld *cloudinary.Cloudinary, db *gorm.DB) ImageRepository {
	return &imageRepository{cld, db}
}

func (r *imageRepository) Upload(ctx context.Context, file *os.File) (*string, error) {
	resp, err := r.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: fmt.Sprintf("upload-%d", time.Now().UnixMilli()),
	})
	if err != nil {
		return nil, errors.New("Upload error")
	}

	return &resp.SecureURL, nil
}

func (r *imageRepository) Create(ctx context.Context, projectId *uint, url *string) (*models.Image, error) {
	newImage := &models.Image{
		URL:       *url,
		ProjectID: *projectId,
	}
	if err := r.db.WithContext(ctx).Create(&newImage).Error; err != nil {
		return nil, errors.New("Image create error")
	}

	return newImage, nil
}

func (r *imageRepository) GetByProjectID(ctx context.Context, projectId *uint) (*[]models.Image, error) {
	var images []models.Image
	if err := r.db.WithContext(ctx).Find(&images, "project_id = ?", projectId).Error; err != nil {
		return nil, errors.New("Image get error")
	}
	return &images, nil
}

func (r *imageRepository) Delete(ctx context.Context, id *uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.Image{}, id).Error; err != nil {
		return errors.New("Image delete error")
	}
	return nil
}
