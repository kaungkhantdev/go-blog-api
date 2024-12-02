package interfaces

import (
	"go-blog-api/internal/tag/models"
)

type TagRepositoryInterface interface {
	CreateTag(data *models.Tag) (models.Tag, error)
	UpdateTag(id int, data *models.Tag) (models.Tag, error)
}
