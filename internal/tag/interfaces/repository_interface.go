package interfaces

import (
	"go-blog-api/internal/tag/models"

	"github.com/gin-gonic/gin"
)

type TagRepositoryInterface interface {
	CreateTag(data *models.Tag) (models.Tag, error)
	UpdateTag(id int, data *models.Tag) (models.Tag, error)
	FindWithPagination(ctx *gin.Context) ([]models.Tag, error)
}
