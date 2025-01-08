package reaction

import (
	"gorm.io/gorm"
)

type ReactionRepository struct {
	db *gorm.DB
}

func NewReactionRepository(db *gorm.DB) ReactionRepositoryInterface {
	return &ReactionRepository{db: db}
}

func (repo *ReactionRepository) CreateReaction(data *ReactionEntity) (ReactionEntity, error) {
	if err := repo.db.Create(&data).Error; err != nil {
		return ReactionEntity{}, err
	}

	return *data, nil
}

func (repo *ReactionRepository) UpdateReaction(id int, data *ReactionEntity) (ReactionEntity, error) {
	var reaction ReactionEntity
	if err := repo.db.First(&reaction).Error; err != nil {
		return ReactionEntity{}, err
	}

	if err := repo.db.Model(&reaction).Updates(data).Error; err != nil {
		return ReactionEntity{}, err
	}
	return reaction, nil
}
