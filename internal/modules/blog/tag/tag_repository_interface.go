package tag

import (
	"go-blog-api/pkg/pagination"
)

type TagRepositoryInterface interface {
	CreateTag(input TagCreateRequest, userId int) (TagEntity, error)
	FindByIdTag(id int) (TagEntity, error)
	UpdateTag(id int, data TagUpdateRequest) (TagEntity, error)
	FindWithPagination(page, pageSize int) (*pagination.PaginatedResponse, error)
	FindByIdsTags(tagIds []int) ([]TagEntity, error)
}
