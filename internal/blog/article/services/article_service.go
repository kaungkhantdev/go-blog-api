package services

import (
	"go-blog-api/internal/blog/article/handlers/requests"
	"go-blog-api/internal/blog/article/interfaces"
	"go-blog-api/internal/blog/article/models"
	"go-blog-api/pkg/pagination"

	tagInterfaces "go-blog-api/internal/blog/tag/interfaces"
	userInterfaces "go-blog-api/internal/core/user/interfaces"
)

type ArticleService struct {
	repo     interfaces.ArticleRepositoryInterfaces
	userRepo userInterfaces.UserRepositoryInterface
	tagRepo  tagInterfaces.TagRepositoryInterface
}

func NewArticleService(
	repo interfaces.ArticleRepositoryInterfaces,
	userRepo userInterfaces.UserRepositoryInterface,
	tagRepo tagInterfaces.TagRepositoryInterface,
) *ArticleService {
	return &ArticleService{repo: repo, userRepo: userRepo, tagRepo: tagRepo}
}

func (service *ArticleService) CreateArticle(userId int, input requests.CreateArticleRequest) (models.Article, error) {
	tags, err := service.tagRepo.FindByIdsTags(input.Tags)
	if err != nil {
		return models.Article{}, err
	}

	if err := service.validateUser(userId); err != nil {
		return models.Article{}, err
	}

	// Create Article
	article := models.Article{
		Title:   input.Title,
		Content: input.Content,
		Tag:     tags,
		UserId:  userId,
	}
	return service.repo.CreateArticle(&article)
}

func (service *ArticleService) UpdateArticle(id int, userId int, input requests.UpdateArticleRequest) (models.Article, error) {
	tags, err := service.tagRepo.FindByIdsTags(input.Tags)
	if err != nil {
		return models.Article{}, err
	}

	if err := service.validateUser(userId); err != nil {
		return models.Article{}, err
	}

	// Create Article
	article := models.Article{
		Title:   input.Title,
		Content: input.Content,
		UserId:  userId,
		Tag:     tags,
	}
	return service.repo.UpdateArticle(id, &article)
}

func (service *ArticleService) FindOneById(id int) (models.Article, error) {
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
