package interfaces

import "go-blog-api/internal/blog/icon/models"

type IconRepositoryInterface interface {
	CreateIcon(data *models.Icon) (models.Icon, error)
	UpdateIcon(id int, data *models.Icon) (models.Icon, error)
	FindByIdIcon(id int) (models.Icon, error)
	FindByName(name string) (models.Icon, error)
}
