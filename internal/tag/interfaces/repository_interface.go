package interfaces

import (
	"go-blog-api/internal/tag/handlers/requests"
	"go-blog-api/internal/tag/models"

	"github.com/gin-gonic/gin"
)

type TagRepositoryInterface interface {
	CreateTag(input requests.TagCreateRequest) (models.Tag, error)
	UpdateTag(id int, data requests.TagUpdateRequest) (models.Tag, error)
	FindWithPagination(ctx *gin.Context) ([]models.Tag, error)
}
