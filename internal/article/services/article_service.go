package services

import (
	"go-blog-api/internal/article/interfaces"
	"go-blog-api/internal/article/models"

	userInterfaces "go-blog-api/internal/user/interfaces"
)

type ArticleService struct {
	repo     interfaces.ArticleRepositoryInterfaces
	userRepo userInterfaces.UserRepositoryInterface
}

func NewArticleService(
	repo interfaces.ArticleRepositoryInterfaces,
	userRepo userInterfaces.UserRepositoryInterface,
) *ArticleService {
	return &ArticleService{repo: repo, userRepo: userRepo}
}

func (service *ArticleService) CreateArticle(data *models.Article) (models.Article, error) {
	if err := service.validateUser(data.UserId); err != nil {
		return models.Article{}, err
	}
	return service.repo.CreateArticle(data)
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

// private
func (service *ArticleService) validateUser(userId int) error {
	if _, err := service.userRepo.FindByIdUser(userId); err != nil {
		return err
	}

	return nil
}
