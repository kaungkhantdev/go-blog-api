package services

import (
	"go-blog-api/internal/article/interfaces"
	"go-blog-api/internal/article/models"
)

type ArticleService struct {
	repo interfaces.ArticleRepositoryInterfaces
}

func NewArticleService(repo interfaces.ArticleRepositoryInterfaces) *ArticleService {
	return &ArticleService{repo: repo}
}

func (service *ArticleService) CreateArticle(data *models.Article) (models.Article, error) {
	return service.repo.CreateArticle(data)
}

func (service *ArticleService) UpdateArticle(id int, data *models.Article) (models.Article, error) {
	return service.repo.UpdateArticle(id, data)
}

func (service *ArticleService) FindOneById(id int) (models.Article, error) {
	return service.repo.FindOneById(id)
}
