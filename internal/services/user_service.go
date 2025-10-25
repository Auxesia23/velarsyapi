package services

import (
	"context"

	"github.com/Auxesia23/velarsyapi/internal/dto"
	"github.com/Auxesia23/velarsyapi/internal/repositories"
	"github.com/Auxesia23/velarsyapi/internal/utils"
)

type UserService interface {
	CreateUser(ctx context.Context, user *dto.UserRequest) (*dto.UserResponse, error)
	GetAllUser(ctx context.Context) ([]*dto.UserResponse, error)
	UpdateUser(ctx context.Context, id uint, user *dto.UserRequest) (*dto.UserResponse, error)
	DeleteUser(ctx context.Context, id uint) error
	LoginUser(ctx context.Context, user *dto.UserRequest) (string, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *dto.UserRequest) (*dto.UserResponse, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	createdUser, err := s.userRepository.Create(ctx, user.Username, hashedPassword)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponse{
		ID:        createdUser.ID,
		Username:  createdUser.Username,
		Password:  createdUser.Password,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}

	return response, nil
}

func (s *userService) GetAllUser(ctx context.Context) ([]*dto.UserResponse, error) {
	users, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*dto.UserResponse
	for _, user := range users {
		response := &dto.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *userService) UpdateUser(ctx context.Context, id uint, user *dto.UserRequest) (*dto.UserResponse, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	updatedUser, err := s.userRepository.Update(ctx, id, user.Username, hashedPassword)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponse{
		ID:        id,
		Username:  updatedUser.Username,
		Password:  updatedUser.Password,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return response, nil
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	err := s.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) LoginUser(ctx context.Context, user *dto.UserRequest) (string, error) {
	loginUser, err := s.userRepository.Login(ctx, user.Username)
	if err != nil {
		return "", err
	}

	if err := utils.CheckPassword(user.Password, loginUser.Password); err != nil {
		return "", err
	}

	token := utils.GenerateToken(loginUser)

	return token, nil
}
