package interfaces

import "go-blog-api/internal/blog/reaction_type/models"

type ReactionTypeRepositoryInterface interface {
	CreateReactionType(data *models.ReactionType) (models.ReactionType, error)
	UpdateReactionType(id int, data *models.ReactionType) (models.ReactionType, error)
	FindOneById(id int) (models.ReactionType, error)
}
