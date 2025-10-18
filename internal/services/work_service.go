package services

import (
	"context"
	"mime/multipart"

	"github.com/Auxesia23/velarsyapi/internal/dto"
	"github.com/Auxesia23/velarsyapi/internal/repositories"
	"github.com/Auxesia23/velarsyapi/internal/utils"
)

type WorkService interface {
	CreateWork(ctx context.Context, title *string, file *multipart.File) (*dto.WorkResponse, error)
	GetAllWork(ctx context.Context) (*[]dto.WorkResponse, error)
	GetOneWork(ctx context.Context, slug *string) (*dto.WorkDetailResponse, error)
	UpdateWork(ctx context.Context, title *string, file *multipart.File, id *uint) (*dto.WorkResponse, error)
	DeleteWork(ctx context.Context, id *uint) error
}

type workService struct {
	workRepository    repositories.WorkRepository
	imageRepository   repositories.ImageRepository
	projectRepository repositories.ProjectRepository
}

func NewWorkService(
	workRepository repositories.WorkRepository,
	imageRepository repositories.ImageRepository,
	projectRepository repositories.ProjectRepository,
) WorkService {
	return &workService{
		workRepository,
		imageRepository,
		projectRepository,
	}
}

func (s *workService) CreateWork(ctx context.Context, title *string, file *multipart.File) (*dto.WorkResponse, error) {
	url, err := s.imageRepository.Upload(ctx, file)
	if err != nil {
		return nil, err
	}

	slug := utils.ToSlug(*title)

	work, err := s.workRepository.Create(ctx, title, url, &slug)
	if err != nil {
		return nil, err
	}

	response := &dto.WorkResponse{
		ID:    work.ID,
		Slug:  work.Slug,
		Title: work.Title,
		Image: work.Image,
	}
	return response, nil
}

func (s *workService) GetAllWork(ctx context.Context) (*[]dto.WorkResponse, error) {
	works, err := s.workRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	response := []dto.WorkResponse{}
	for _, work := range *works {
		response = append(response, dto.WorkResponse{
			ID:    work.ID,
			Slug:  work.Slug,
			Title: work.Title,
			Image: work.Image,
		})
	}
	return &response, nil
}

func (s *workService) GetOneWork(ctx context.Context, slug *string) (*dto.WorkDetailResponse, error) {
	work, err := s.workRepository.GetOne(ctx, slug)
	if err != nil {
		return nil, err
	}

	projects, err := s.projectRepository.GetByWorkID(ctx, &work.ID)
	if err != nil {
		return nil, err
	}

	var projectResponses []dto.ProjectResponse
	for _, project := range *projects {
		projectResponses = append(projectResponses, dto.ProjectResponse{
			ID:        project.ID,
			Slug:      project.Slug,
			Name:      project.Name,
			Thumbnail: project.Thumbnail,
		})
	}

	workResponse := dto.WorkDetailResponse{
		ID:       work.ID,
		Slug:     work.Slug,
		Title:    work.Title,
		Image:    work.Image,
		Projects: projectResponses,
	}

	return &workResponse, nil
}

func (s *workService) UpdateWork(ctx context.Context, title *string, file *multipart.File, id *uint) (*dto.WorkResponse, error) {
	url, err := s.imageRepository.Upload(ctx, file)
	if err != nil {
		return nil, err
	}
	newSlug := utils.ToSlug(*title)
	updatedWork, err := s.workRepository.Update(ctx, title, url, &newSlug, id)
	if err != nil {
		return nil, err
	}
	response := &dto.WorkResponse{
		ID:    *id,
		Slug:  updatedWork.Slug,
		Title: updatedWork.Title,
		Image: updatedWork.Image,
	}
	return response, nil
}

func (s *workService) DeleteWork(ctx context.Context, id *uint) error {
	err := s.workRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
