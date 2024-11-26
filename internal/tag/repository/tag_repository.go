package repository

import (
	"go-blog-api/internal/tag/interfaces"
	"go-blog-api/internal/tag/models"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) interfaces.TagRepositoryInterface {
	return &TagRepository{db: db}
}

func (repo *TagRepository) CreateTag(data *models.Tag) (models.Tag, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return models.Tag{}, err
	}
	return *data, nil
}

func (repo *TagRepository) UpdateTag(id int, data *models.Tag) (models.Tag, error) {
	var tag models.Tag
	if err := repo.db.First(&tag, id).Error; err != nil {
		return models.Tag{}, nil
	}

	if err := repo.db.Model(&tag).Updates(data).Error; err != nil {
		return models.Tag{}, nil
	}
	return tag, nil
}
