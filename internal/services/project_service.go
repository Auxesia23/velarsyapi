package services

import (
	"context"
	"mime/multipart"

	"github.com/Auxesia23/velarsyapi/internal/dto"
	"github.com/Auxesia23/velarsyapi/internal/repositories"
	"github.com/Auxesia23/velarsyapi/internal/utils"
)

type ProjectService interface {
	CreateProject(ctx context.Context, project *dto.ProjectRequest, image *multipart.File, workId *uint) (*dto.ProjectResponse, error)
	GetSingleProject(ctx context.Context, slug *string) (*dto.ProjectDetailResponse, error)
	UpdateProject(ctx context.Context, project *dto.ProjectRequest, image *multipart.File, id *uint) (*dto.ProjectDetailResponse, error)
	DeleteProject(ctx context.Context, id *uint) error
}

type projectService struct {
	projectRepository repositories.ProjectRepository
	imageRepository   repositories.ImageRepository
}

func NewProjectService(projectRepository repositories.ProjectRepository, imageRepository repositories.ImageRepository) ProjectService {
	return &projectService{projectRepository, imageRepository}
}

func (s *projectService) CreateProject(ctx context.Context, project *dto.ProjectRequest, image *multipart.File, workId *uint) (*dto.ProjectResponse, error) {
	thumbnailUrl, err := s.imageRepository.Upload(ctx, image)
	if err != nil {
		return nil, err
	}

	slug := utils.ToSlug(project.Name)

	createdProject, err := s.projectRepository.Create(ctx, project, thumbnailUrl, &slug, workId)
	if err != nil {
		return nil, err
	}

	response := &dto.ProjectResponse{
		ID:        createdProject.ID,
		Slug:      createdProject.Slug,
		Name:      createdProject.Name,
		Thumbnail: createdProject.Thumbnail,
	}

	return response, nil
}

func (s *projectService) GetSingleProject(ctx context.Context, slug *string) (*dto.ProjectDetailResponse, error) {
	project, err := s.projectRepository.GetOne(ctx, slug)
	if err != nil {
		return nil, err
	}

	images, err := s.imageRepository.GetByProjectID(ctx, &project.ID)
	if err != nil {
		return nil, err
	}

	var imageResponse []dto.ImageResponse
	for _, image := range *images {
		imageResponse = append(imageResponse, dto.ImageResponse{
			ID:  image.ID,
			URL: image.URL,
		})
	}

	response := &dto.ProjectDetailResponse{
		ID:              project.ID,
		Slug:            project.Slug,
		Name:            project.Name,
		Thumbnail:       project.Thumbnail,
		AboutBrand:      project.AboutBrand,
		DesignExecution: project.DesignExecution,
		Images:          imageResponse,
	}
	return response, nil
}

func (s *projectService) UpdateProject(ctx context.Context, project *dto.ProjectRequest, image *multipart.File, id *uint) (*dto.ProjectDetailResponse, error) {
	imageUrl, err := s.imageRepository.Upload(ctx, image)
	if err != nil {
		return nil, err
	}

	slug := utils.ToSlug(project.Name)

	updatedProject, err := s.projectRepository.Update(ctx, project, imageUrl, &slug, id)
	if err != nil {
		return nil, err
	}
	response := &dto.ProjectDetailResponse{
		ID:              *id,
		Slug:            updatedProject.Slug,
		Name:            updatedProject.Name,
		Thumbnail:       updatedProject.Thumbnail,
		AboutBrand:      updatedProject.AboutBrand,
		DesignExecution: updatedProject.DesignExecution,
	}
	return response, nil
}

func (s *projectService) DeleteProject(ctx context.Context, id *uint) error {
	err := s.projectRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
