package interfaces

import "go-blog-api/internal/article/models"

type ArticleRepositoryInterfaces interface {
	CreateArticle(data *models.Article) (models.Article, error)
	UpdateArticle(id int, data *models.Article) (models.Article, error)
	FindOneById(id int) (models.Article, error)
}
