package services

import (
	"go-blog-api/internal/article/handlers/requests"
	"go-blog-api/internal/article/interfaces"
	"go-blog-api/internal/article/models"
	"go-blog-api/pkg/pagination"

	tagInterfaces "go-blog-api/internal/tag/interfaces"
	userInterfaces "go-blog-api/internal/user/interfaces"
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

func (service *ArticleService) CreateArticle(input requests.CreateArticleRequest) (interface{}, error) {
	return service.tagRepo.FindByIdsTags(input.Tags)
	// if err != nil {
	// 	return nil, err
	// }

	// if err := service.validateUser(input.UserId); err != nil {
	// 	return models.Article{}, err
	// }
	// return tags, nil

	// tag1 := Tag{Name: "Go"}
	// tag2 := Tag{Name: "Programming"}

	// // Create Article
	// article := Article{
	// 	Title:   "Learn Go",
	// 	Content: "Go is an awesome language!",
	// 	Tag:     []Tag{tag1, tag2}, // Associate tags
	// }
	// return service.repo.CreateArticle(data)
}

func (service *ArticleService) UpdateArticle(id int, data *models.Article) (models.Article, error) {
	if err := service.validateUser(data.UserId); err != nil {
		return models.Article{}, err
	}
	return service.repo.UpdateArticle(id, data)
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
