package interfaces

import "go-blog-api/internal/comment/models"

type CommentRepositoryInterface interface {
	CreateComment(data *models.Comment) (models.Comment, error)
	UpdateComment(id int, data *models.Comment) (models.Comment, error)
	FindOneById(id int) (models.Comment, error)
}
