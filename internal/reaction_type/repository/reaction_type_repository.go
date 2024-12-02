package repository

import (
	"go-blog-api/internal/reaction_type/interfaces"
	"go-blog-api/internal/reaction_type/models"

	"gorm.io/gorm"
)

type ReactionTypeRepository struct {
	db *gorm.DB
}

func NewReactionTypeRepository(db *gorm.DB) interfaces.ReactionTypeRepositoryInterface {
	return &ReactionTypeRepository{db: db}
}

func (repo *ReactionTypeRepository) CreateReactionType(data *models.ReactionType) (models.ReactionType, error) {
	if err := repo.db.Create(&data).Error; err != nil {
		return models.ReactionType{}, err
	}

	return *data, nil
}

func (repo *ReactionTypeRepository) UpdateReactionType(id int, data *models.ReactionType) (models.ReactionType, error) {
	var reactionType models.ReactionType
	if err := repo.db.First(&reactionType).Error; err != nil {
		return models.ReactionType{}, err
	}

	if err := repo.db.Model(&reactionType).Updates(data).Error; err != nil {
		return models.ReactionType{}, err
	}
	return reactionType, nil
}
