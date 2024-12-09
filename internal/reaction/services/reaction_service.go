package services

import (
	articleInterface "go-blog-api/internal/article/interfaces"
	"go-blog-api/internal/reaction/handlers/requests"
	"go-blog-api/internal/reaction/interfaces"
	"go-blog-api/internal/reaction/models"
	reactionTypeInterface "go-blog-api/internal/reaction_type/interfaces"
)

type ReactionService struct {
	repo             interfaces.ReactionRepositoryInterface
	articleRepo      articleInterface.ArticleRepositoryInterfaces
	reactionTypeRepo reactionTypeInterface.ReactionTypeRepositoryInterface
}

func NewReactionService(
	repo interfaces.ReactionRepositoryInterface,
	articleRepo articleInterface.ArticleRepositoryInterfaces,
	reactionTypeRepo reactionTypeInterface.ReactionTypeRepositoryInterface,
) *ReactionService {
	return &ReactionService{repo: repo, articleRepo: articleRepo, reactionTypeRepo: reactionTypeRepo}
}

func (service *ReactionService) CreateReaction(userId int, input *requests.CreateReactionRequest) (models.Reaction, error) {

	if _, err := service.articleRepo.FindOneById(input.ArticleId); err != nil {
		return models.Reaction{}, err
	}

	if _, err := service.reactionTypeRepo.FindOneById(input.ArticleId); err != nil {
		return models.Reaction{}, err
	}

	reaction := &models.Reaction{
		UserId:         userId,
		ArticleId:      input.ArticleId,
		ReactionTypeId: input.ReactionTypeId,
	}

	return service.repo.CreateReaction(reaction)
}

func (service *ReactionService) UpdateReaction(id int, userId int, input *requests.UpdateReactionRequest) (models.Reaction, error) {
	if _, err := service.articleRepo.FindOneById(input.ArticleId); err != nil {
		return models.Reaction{}, err
	}

	if _, err := service.reactionTypeRepo.FindOneById(input.ArticleId); err != nil {
		return models.Reaction{}, err
	}

	reaction := &models.Reaction{
		UserId:         userId,
		ArticleId:      input.ArticleId,
		ReactionTypeId: input.ReactionTypeId,
	}

	return service.repo.UpdateReaction(id, reaction)
}
