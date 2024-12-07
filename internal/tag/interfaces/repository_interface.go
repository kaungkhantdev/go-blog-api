package interfaces

import (
	"go-blog-api/internal/tag/handlers/requests"
	"go-blog-api/internal/tag/models"
	"go-blog-api/pkg/pagination"
)

type TagRepositoryInterface interface {
	CreateTag(input requests.TagCreateRequest) (models.Tag, error)
	FindByIdTag(id int) (models.Tag, error)
	UpdateTag(id int, data requests.TagUpdateRequest) (models.Tag, error)
	FindWithPagination(page, pageSize int) (*pagination.PaginatedResponse, error)
	FindByIdsTags(tagIds []int) ([]models.Tag, error)
}
