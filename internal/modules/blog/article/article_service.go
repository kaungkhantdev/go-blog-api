package article

import (
	"go-blog-api/internal/modules/blog/tag"
	"go-blog-api/internal/modules/core/user"
	"go-blog-api/pkg/pagination"
)

type ArticleService struct {
	repo     ArticleRepositoryInterfaces
	userRepo user.UserRepositoryInterface
	tagRepo  tag.TagRepositoryInterface
}

func NewArticleService(
	repo ArticleRepositoryInterfaces,
	userRepo user.UserRepositoryInterface,
	tagRepo tag.TagRepositoryInterface,
) *ArticleService {
	return &ArticleService{repo: repo, userRepo: userRepo, tagRepo: tagRepo}
}

func (service *ArticleService) CreateArticle(userId int, input CreateArticleRequest) (ArticleEntity, error) {
	tags, err := service.tagRepo.FindByIdsTags(input.Tags)
	if err != nil {
		return ArticleEntity{}, err
	}

	if err := service.validateUser(userId); err != nil {
		return ArticleEntity{}, err
	}

	// Create Article
	article := ArticleEntity{
		Title:   input.Title,
		Content: input.Content,
		Tag:     tags,
		UserId:  userId,
	}
	return service.repo.CreateArticle(&article)
}

func (service *ArticleService) UpdateArticle(id int, userId int, input UpdateArticleRequest) (ArticleEntity, error) {
	tags, err := service.tagRepo.FindByIdsTags(input.Tags)
	if err != nil {
		return ArticleEntity{}, err
	}

	if err := service.validateUser(userId); err != nil {
		return ArticleEntity{}, err
	}

	// Create Article
	article := ArticleEntity{
		Title:   input.Title,
		Content: input.Content,
		UserId:  userId,
		Tag:     tags,
	}
	return service.repo.UpdateArticle(id, &article)
}

func (service *ArticleService) FindOneById(id int) (ArticleEntity, error) {
	return service.repo.FindOneById(id)
}

func (service *ArticleService) FindWithPagination(page, pageSize int) (*pagination.PaginatedResponse, error) {
	return service.repo.FindWithPagination(page, pageSize)
}

// private
func (service *ArticleService) validateUser(userId int) error {
	if _, err := service.userRepo.FindByIdUser(userId); err != nil {
		return err
	}

	return nil
}
