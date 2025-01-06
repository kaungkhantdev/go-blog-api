package reaction

import (
	"go-blog-api/internal/blog/article"
	"go-blog-api/internal/blog/reaction_type"
)

type ReactionService struct {
	repo             ReactionRepositoryInterface
	articleRepo      article.ArticleRepositoryInterfaces
	reactionTypeRepo reaction_type.ReactionTypeRepositoryInterface
}

func NewReactionService(
	repo ReactionRepositoryInterface,
	articleRepo article.ArticleRepositoryInterfaces,
	reactionTypeRepo reaction_type.ReactionTypeRepositoryInterface,
) *ReactionService {
	return &ReactionService{repo: repo, articleRepo: articleRepo, reactionTypeRepo: reactionTypeRepo}
}

func (service *ReactionService) CreateReaction(userId int, input *CreateReactionRequest) (ReactionEntity, error) {

	if _, err := service.articleRepo.FindOneById(input.ArticleId); err != nil {
		return ReactionEntity{}, err
	}

	if _, err := service.reactionTypeRepo.FindOneById(input.ArticleId); err != nil {
		return ReactionEntity{}, err
	}

	reaction := &ReactionEntity{
		UserId:         userId,
		ArticleId:      input.ArticleId,
		ReactionTypeId: input.ReactionTypeId,
	}

	return service.repo.CreateReaction(reaction)
}

func (service *ReactionService) UpdateReaction(id int, userId int, input *UpdateReactionRequest) (ReactionEntity, error) {
	if _, err := service.articleRepo.FindOneById(input.ArticleId); err != nil {
		return ReactionEntity{}, err
	}

	if _, err := service.reactionTypeRepo.FindOneById(input.ArticleId); err != nil {
		return ReactionEntity{}, err
	}

	reaction := &ReactionEntity{
		UserId:         userId,
		ArticleId:      input.ArticleId,
		ReactionTypeId: input.ReactionTypeId,
	}

	return service.repo.UpdateReaction(id, reaction)
}
