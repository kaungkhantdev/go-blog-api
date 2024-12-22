package repository

import (
	"go-blog-api/internal/blog/comment/interfaces"
	"go-blog-api/internal/blog/comment/models"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) interfaces.CommentRepositoryInterface {
	return &CommentRepository{db: db}
}

func (repo *CommentRepository) CreateComment(data *models.Comment) (models.Comment, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return models.Comment{}, err
	}
	return *data, nil
}

func (repo *CommentRepository) UpdateComment(id int, data *models.Comment) (models.Comment, error) {
	var comment models.Comment
	if err := repo.db.First(&comment, id).Error; err != nil {
		return models.Comment{}, err
	}

	if err := repo.db.Model(&comment).Updates(data).Error; err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

func (repo *CommentRepository) FindOneById(id int) (models.Comment, error) {
	var comment models.Comment
	if err := repo.db.First(&comment, id).Error; err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}
