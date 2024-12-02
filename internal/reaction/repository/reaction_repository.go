package repository

import (
	"go-blog-api/internal/reaction/interfaces"
	"go-blog-api/internal/reaction/models"

	"gorm.io/gorm"
)

type ReactionRepository struct {
	db *gorm.DB
}

func NewReactionRepository(db *gorm.DB) interfaces.ReactionRepositoryInterface {
	return &ReactionRepository{db: db}
}

func (repo *ReactionRepository) CreateReaction(data *models.Reaction) (models.Reaction, error) {
	if err := repo.db.Create(&data).Error; err != nil {
		return models.Reaction{}, err
	}

	return *data, nil
}

func (repo *ReactionRepository) UpdateReaction(id int, data *models.Reaction) (models.Reaction, error) {
	var reaction models.Reaction
	if err := repo.db.First(&reaction).Error; err != nil {
		return models.Reaction{}, err
	}

	if err := repo.db.Model(&reaction).Updates(data).Error; err != nil {
		return models.Reaction{}, err
	}
	return reaction, nil
}
