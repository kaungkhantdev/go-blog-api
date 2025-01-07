package comment

import (
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepositoryInterface {
	return &CommentRepository{db: db}
}

func (repo *CommentRepository) CreateComment(data *CommentEntity) (CommentEntity, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return CommentEntity{}, err
	}
	return *data, nil
}

func (repo *CommentRepository) UpdateComment(id int, data *CommentEntity) (CommentEntity, error) {
	var comment CommentEntity
	if err := repo.db.First(&comment, id).Error; err != nil {
		return CommentEntity{}, err
	}

	if err := repo.db.Model(&comment).Updates(data).Error; err != nil {
		return CommentEntity{}, err
	}
	return comment, nil
}

func (repo *CommentRepository) FindOneById(id int) (CommentEntity, error) {
	var comment CommentEntity
	if err := repo.db.First(&comment, id).Error; err != nil {
		return CommentEntity{}, err
	}

	return comment, nil
}
