package repositories

import (
	"context"
	"errors"

	"github.com/Auxesia23/velarsyapi/internal/dto"
	"github.com/Auxesia23/velarsyapi/internal/models"
	"gorm.io/gorm"
)

type ServiceRepository interface {
	Create(ctx context.Context, service *dto.ServiceRequest) (*models.Service, error)
	GetAll(ctx context.Context) (*[]models.Service, error)
	Update(ctx context.Context, service *dto.ServiceRequest, id uint) (*models.Service, error)
	Delete(ctx context.Context, id uint) error
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepository{db: db}
}

func (r *serviceRepository) Create(ctx context.Context, service *dto.ServiceRequest) (*models.Service, error) {
	newService := &models.Service{
		Title:       service.Title,
		Description: service.Description,
		Icon:        service.Icon,
	}

	if err := r.db.WithContext(ctx).Create(&newService).Error; err != nil {
		return nil, errors.New("Service creation error")
	}

	return newService, nil
}

func (r *serviceRepository) GetAll(ctx context.Context) (*[]models.Service, error) {
	var services []models.Service
	if err := r.db.WithContext(ctx).Find(&services).Error; err != nil {
		return nil, errors.New("Service get error")
	}
	return &services, nil
}

func (r *serviceRepository) Update(ctx context.Context, service *dto.ServiceRequest, id uint) (*models.Service, error) {
	updatedService := &models.Service{
		Title:       service.Title,
		Description: service.Description,
		Icon:        service.Icon,
	}

	if err := r.db.WithContext(ctx).Model(&models.Service{}).Where("id = ?", id).Updates(updatedService).Error; err != nil {
		return nil, errors.New("Service update error")
	}
	return updatedService, nil
}

func (r *serviceRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.Service{}, id).Error; err != nil {
		return errors.New("Service delete error")
	}
	return nil
}
