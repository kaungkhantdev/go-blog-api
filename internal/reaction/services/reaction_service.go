package services

import (
	"go-blog-api/internal/reaction/interfaces"
	"go-blog-api/internal/reaction/models"
)

type ReactionService struct {
	repo interfaces.ReactionRepositoryInterface
}

func NewReactionService(repo interfaces.ReactionRepositoryInterface) *ReactionService {
	return &ReactionService{repo: repo}
}

func (service *ReactionService) CreateReaction(data *models.Reaction) (models.Reaction, error) {
	return service.repo.CreateReaction(data)
}

func (service *ReactionService) UpdateReaction(id int, data *models.Reaction) (models.Reaction, error) {
	return service.repo.UpdateReaction(id, data)
}
