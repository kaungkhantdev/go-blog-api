package reaction_type

import (
	"gorm.io/gorm"
)

type ReactionTypeRepository struct {
	db *gorm.DB
}

func NewReactionTypeRepository(db *gorm.DB) ReactionTypeRepositoryInterface {
	return &ReactionTypeRepository{db: db}
}

func (repo *ReactionTypeRepository) CreateReactionType(data *ReactionTypeEntity) (ReactionTypeEntity, error) {
	if err := repo.db.Create(&data).Error; err != nil {
		return ReactionTypeEntity{}, err
	}

	return *data, nil
}

func (repo *ReactionTypeRepository) UpdateReactionType(id int, data *ReactionTypeEntity) (ReactionTypeEntity, error) {
	var reactionType ReactionTypeEntity
	if err := repo.db.First(&reactionType).Error; err != nil {
		return ReactionTypeEntity{}, err
	}

	if err := repo.db.Model(&reactionType).Updates(data).Error; err != nil {
		return ReactionTypeEntity{}, err
	}
	return reactionType, nil
}

func (repo *ReactionTypeRepository) FindOneById(id int) (ReactionTypeEntity, error) {
	var reactionType ReactionTypeEntity
	if err := repo.db.First(&reactionType, id).Error; err != nil {
		return ReactionTypeEntity{}, err
	}

	return reactionType, nil
}
