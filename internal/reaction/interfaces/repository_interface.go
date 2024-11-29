package interfaces

import "go-blog-api/internal/reaction/models"

type ReactionRepositoryInterface interface {
	CreateReaction(data *models.Reaction) (models.Reaction, error)
	UpdateReaction(id int, data *models.Reaction) (models.Reaction, error)
}
