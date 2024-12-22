package interfaces

import (
	"go-blog-api/internal/blog/article/models"
	"go-blog-api/pkg/pagination"
)

type ArticleRepositoryInterfaces interface {
	CreateArticle(data *models.Article) (models.Article, error)
	UpdateArticle(id int, data *models.Article) (models.Article, error)
	FindOneById(id int) (models.Article, error)
	FindWithPagination(page, pageSize int) (*pagination.PaginatedResponse, error)
}
