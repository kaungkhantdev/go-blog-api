package services

import (
	"go-blog-api/internal/tag/handlers/requests"
	"go-blog-api/internal/tag/interfaces"
	"go-blog-api/internal/tag/models"

	"github.com/gin-gonic/gin"
)

type TagService struct {
	repo interfaces.TagRepositoryInterface
}

func NewTagService(repo interfaces.TagRepositoryInterface) *TagService {
	return &TagService{repo: repo}
}

func (service *TagService) CreateTag(data requests.TagCreateRequest) (models.Tag, error) {
	return service.repo.CreateTag(data)
}

func (service *TagService) UpdateTag(id int, data requests.TagUpdateRequest) (models.Tag, error) {
	return service.repo.UpdateTag(id, data)
}

func (service *TagService) FindWithPagination(ctx *gin.Context) ([]models.Tag, error) {
	return service.repo.FindWithPagination(ctx)
}
