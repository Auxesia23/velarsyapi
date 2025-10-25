package services

import (
	"context"

	"github.com/Auxesia23/velarsyapi/internal/dto"
	"github.com/Auxesia23/velarsyapi/internal/repositories"
)

type ServiceService interface {
	CreateService(ctx context.Context, service *dto.ServiceRequest) (*dto.ServiceResponse, error)
	GetAllServices(ctx context.Context) (*[]dto.ServiceResponse, error)
	UpdateService(ctx context.Context, service *dto.ServiceRequest, id uint) (*dto.ServiceResponse, error)
	DeleteService(ctx context.Context, id uint) error
}

type serviceService struct {
	serviceRepo repositories.ServiceRepository
}

func NewServiceService(serviceRepo repositories.ServiceRepository) ServiceService {
	return &serviceService{
		serviceRepo: serviceRepo,
	}
}

func (s *serviceService) CreateService(ctx context.Context, service *dto.ServiceRequest) (*dto.ServiceResponse, error) {
	createdService, err := s.serviceRepo.Create(ctx, service)
	if err != nil {
		return nil, err
	}

	response := &dto.ServiceResponse{
		ID:          createdService.ID,
		Title:       createdService.Title,
		Description: createdService.Description,
		Icon:        createdService.Icon,
	}

	return response, nil
}

func (s *serviceService) GetAllServices(ctx context.Context) (*[]dto.ServiceResponse, error) {
	services, err := s.serviceRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	response := []dto.ServiceResponse{}
	for _, service := range *services {
		response = append(response, dto.ServiceResponse{
			ID:          service.ID,
			Title:       service.Title,
			Description: service.Description,
			Icon:        service.Icon,
		})
	}

	return &response, nil
}

func (s *serviceService) UpdateService(ctx context.Context, service *dto.ServiceRequest, id uint) (*dto.ServiceResponse, error) {
	updatedService, err := s.serviceRepo.Update(ctx, service, id)
	if err != nil {
		return nil, err
	}
	response := &dto.ServiceResponse{
		ID:          id,
		Title:       updatedService.Title,
		Description: updatedService.Description,
		Icon:        updatedService.Icon,
	}
	return response, nil
}

func (s *serviceService) DeleteService(ctx context.Context, id uint) error {
	if err := s.serviceRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
