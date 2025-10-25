package repositories

import (
	"context"
	"errors"

	"github.com/Auxesia23/velarsyapi/internal/dto"
	"github.com/Auxesia23/velarsyapi/internal/models"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(ctx context.Context, project *dto.ProjectRequest, thumbnail, slug string, workId uint) (*models.Project, error)
	GetAll(ctx context.Context) ([]*models.Project, error)
	GetOne(ctx context.Context, slug string) (*models.Project, error)
	GetByWorkID(ctx context.Context, workId uint) (*[]models.Project, error)
	Update(ctx context.Context, project *dto.ProjectRequest, thumbnail, slug string, id uint) (*models.Project, error)
	Delete(ctx context.Context, projectId uint) error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{
		db: db,
	}
}

func (r *projectRepository) Create(ctx context.Context, project *dto.ProjectRequest, thumbnail, slug string, workId uint) (*models.Project, error) {
	newProject := &models.Project{
		Slug:            slug,
		Name:            project.Name,
		Thumbnail:       thumbnail,
		AboutBrand:      project.AboutBrand,
		DesignExecution: project.DesignExecution,
		OurWorkID:       workId,
	}

	if err := r.db.WithContext(ctx).Create(newProject).Error; err != nil {
		return nil, errors.New("Project create error")
	}

	return newProject, nil
}

func (r *projectRepository) GetAll(ctx context.Context) ([]*models.Project, error) {
	var projects []*models.Project
	if err := r.db.WithContext(ctx).Find(&projects).Error; err != nil {
		return nil, errors.New("Project find error")
	}
	return projects, nil
}

func (r *projectRepository) GetByWorkID(ctx context.Context, workId uint) (*[]models.Project, error) {
	var projects []models.Project
	if err := r.db.WithContext(ctx).Find(&projects, "our_work_id = ?", workId).Error; err != nil {
		return nil, errors.New("Project find error")
	}
	return &projects, nil
}

func (r *projectRepository) GetOne(ctx context.Context, slug string) (*models.Project, error) {
	var project models.Project
	if err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&project).Error; err != nil {
		return nil, errors.New("Project find error")
	}
	return &project, nil
}

func (r *projectRepository) Update(ctx context.Context, project *dto.ProjectRequest, thumbnail, slug string, id uint) (*models.Project, error) {
	updatedProject := &models.Project{
		Slug:            slug,
		Name:            project.Name,
		Thumbnail:       thumbnail,
		AboutBrand:      project.AboutBrand,
		DesignExecution: project.DesignExecution,
	}
	if err := r.db.WithContext(ctx).Model(&models.Project{}).Where("id = ?", id).Updates(updatedProject).Error; err != nil {
		return nil, errors.New("Project update error")
	}
	return updatedProject, nil
}

func (r *projectRepository) Delete(ctx context.Context, projectId uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.Project{}, projectId).Error; err != nil {
		return errors.New("Project delete error")
	}
	return nil
}
