package services

import (
	"context"
	"mime/multipart"

	"github.com/Auxesia23/velarsyapi/internal/dto"
	"github.com/Auxesia23/velarsyapi/internal/repositories"
)

type ImageService interface {
	CreateImage(ctx context.Context, file *multipart.File, projectId *uint) (*dto.ImageResponse, error)
	Delete(ctx context.Context, id *uint) error
}

type imageService struct {
	imageRepository repositories.ImageRepository
}

func NewImageService(imageRepository repositories.ImageRepository) ImageService {
	return &imageService{imageRepository: imageRepository}
}

func (s *imageService) CreateImage(ctx context.Context, file *multipart.File, projectId *uint) (*dto.ImageResponse, error) {
	url, err := s.imageRepository.Upload(ctx, file)
	if err != nil {
		return nil, err
	}

	image, err := s.imageRepository.Create(ctx, projectId, url)
	if err != nil {
		return nil, err
	}

	imageResponse := &dto.ImageResponse{
		ID:  image.ID,
		URL: image.URL,
	}
	return imageResponse, nil
}

func (s *imageService) Delete(ctx context.Context, id *uint) error {
	return s.imageRepository.Delete(ctx, id)
}
