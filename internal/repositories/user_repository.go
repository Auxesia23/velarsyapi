package repositories

import (
	"context"
	"errors"

	"github.com/Auxesia23/velarsyapi/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, username, hashedPassword *string) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, id *uint, username, hashedPassword *string) (*models.User, error)
	Delete(ctx context.Context, id *uint) error
	Login(ctx context.Context, username *string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, username, hashedPassword *string) (*models.User, error) {
	newUser := models.User{
		Username: *username,
		Password: *hashedPassword,
	}

	err := r.db.WithContext(ctx).Create(&newUser).Error
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User

	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Update(ctx context.Context, id *uint, username, hashedPassword *string) (*models.User, error) {
	var userToUpdate models.User
	findResult := r.db.WithContext(ctx).First(&userToUpdate, id)

	if findResult.Error != nil {
		if findResult.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("User not found")
		}
		return nil, findResult.Error
	}

	// 2. Tentukan data baru untuk update
	updateData := models.User{
		Username: *username,
		Password: *hashedPassword,
	}

	result := r.db.WithContext(ctx).Where("id = ?", id).Updates(updateData)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("Failed to update user data")
	}

	var updatedUser models.User
	err := r.db.WithContext(ctx).First(&updatedUser, id).Error
	if err != nil {
		return nil, errors.New("Failed to retrieve updated user data")
	}

	return &updatedUser, nil
}
func (r *userRepository) Delete(ctx context.Context, id *uint) error {
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Login(ctx context.Context, username *string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
