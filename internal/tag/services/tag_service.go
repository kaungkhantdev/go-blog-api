package services

import (
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

func (service *TagService) CreateTag(data *models.Tag) (models.Tag, error) {
	return service.repo.CreateTag(data)
}

func (service *TagService) UpdateTag(id int, data *models.Tag) (models.Tag, error) {
	return service.repo.UpdateTag(id, data)
}

func (service *TagService) FindWithPagination(ctx *gin.Context) ([]models.Tag, error) {
	return service.repo.FindWithPagination(ctx)
}
