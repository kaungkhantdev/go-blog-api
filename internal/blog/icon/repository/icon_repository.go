package repository

import (
	"go-blog-api/internal/blog/icon/interfaces"
	"go-blog-api/internal/blog/icon/models"

	"gorm.io/gorm"
)

type IconRepository struct {
	db *gorm.DB
}

func NewIconRepository(db *gorm.DB) interfaces.IconRepositoryInterface {
	return &IconRepository{db: db}
}

func (repo *IconRepository) CreateIcon(data *models.Icon) (models.Icon, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return models.Icon{}, err
	}

	return *data, nil
}

func (repo *IconRepository) UpdateIcon(id int, data *models.Icon) (models.Icon, error) {
	var icon models.Icon
	if err := repo.db.First(&icon, id).Error; err != nil {
		return models.Icon{}, err
	}

	if err := repo.db.Model(&icon).Updates(data).Error; err != nil {
		return models.Icon{}, err
	}
	return icon, nil
}

func (repo *IconRepository) FindByName(name string) (models.Icon, error) {
	var icon models.Icon
	if err := repo.db.Where("name = ?", name).First(&icon).Error; err != nil {
		return models.Icon{}, err
	}

	return icon, nil
}

func (repo *IconRepository) FindByIdIcon(id int) (models.Icon, error) {
	var icon models.Icon
	if err := repo.db.First(&icon, id).Error; err != nil {
		return models.Icon{}, err
	}
	return icon, nil
}
