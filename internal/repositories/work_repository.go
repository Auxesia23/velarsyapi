package repositories

import (
	"context"
	"errors"

	"github.com/Auxesia23/velarsyapi/internal/models"
	"gorm.io/gorm"
)

type WorkRepository interface {
	Create(ctx context.Context, title, url, slug *string) (*models.OurWork, error)
	GetAll(ctx context.Context) (*[]models.OurWork, error)
	GetOne(ctx context.Context, slug *string) (*models.OurWork, error)
	Update(ctx context.Context, title, url, slug *string, id *uint) (*models.OurWork, error)
	Delete(ctx context.Context, id *uint) error
}

type workRepository struct {
	db *gorm.DB
}

func NewWorkRepository(db *gorm.DB) WorkRepository {
	return &workRepository{db: db}
}

func (r *workRepository) Create(ctx context.Context, title, url, slug *string) (*models.OurWork, error) {
	ourWork := &models.OurWork{
		Title: *title,
		Image: *url,
		Slug:  *slug,
	}

	if err := r.db.WithContext(ctx).Create(ourWork).Error; err != nil {
		return nil, errors.New("Work create error")
	}

	return ourWork, nil
}

func (r *workRepository) GetAll(ctx context.Context) (*[]models.OurWork, error) {
	var works []models.OurWork
	if err := r.db.WithContext(ctx).Find(&works).Error; err != nil {
		return nil, errors.New("Work get error")
	}
	return &works, nil
}

func (r *workRepository) GetOne(ctx context.Context, slug *string) (*models.OurWork, error) {
	var work models.OurWork
	if err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&work).Error; err != nil {
		return nil, errors.New("Work get error")
	}
	return &work, nil
}

func (r *workRepository) Update(ctx context.Context, title, url, slug *string, id *uint) (*models.OurWork, error) {
	updatedWork := &models.OurWork{
		Title: *title,
		Image: *url,
		Slug:  *slug,
	}
	if err := r.db.WithContext(ctx).Model(&models.OurWork{}).Where("id = ?", id).Updates(updatedWork).Error; err != nil {
		return nil, errors.New("Work update error")
	}
	return updatedWork, nil
}

func (r *workRepository) Delete(ctx context.Context, id *uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.OurWork{}, id).Error; err != nil {
		return errors.New("Work delete error")
	}
	return nil
}
