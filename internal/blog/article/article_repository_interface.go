package article

import (
	"go-blog-api/pkg/pagination"
)

type ArticleRepositoryInterfaces interface {
	CreateArticle(data *ArticleEntity) (ArticleEntity, error)
	UpdateArticle(id int, data *ArticleEntity) (ArticleEntity, error)
	FindOneById(id int) (ArticleEntity, error)
	FindWithPagination(page, pageSize int) (*pagination.PaginatedResponse, error)
}
