package services

import (
	"go-blog-api/internal/reaction_type/interfaces"
	"go-blog-api/internal/reaction_type/models"
)

type ReactionTypeService struct {
	repo interfaces.ReactionTypeRepositoryInterface
}

func NewReactionTypeService(repo interfaces.ReactionTypeRepositoryInterface) *ReactionTypeService {
	return &ReactionTypeService{repo: repo}
}

func (service *ReactionTypeService) CreateReactionType(data *models.ReactionType) (models.ReactionType, error) {
	return service.repo.CreateReactionType(data)
}

func (service *ReactionTypeService) UpdateReactionType(id int, data *models.ReactionType) (models.ReactionType, error) {
	return service.repo.UpdateReactionType(id, data)
}
