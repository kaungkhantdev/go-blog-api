package interfaces

import "go-blog-api/internal/core/user/models"

type UserRepositoryInterface interface {
	CreateUser(data *models.User) (models.User, error)
	FindByIdUser(id int) (models.User, error)
	FindByEmailUser(email string) (models.User, error)
	UpdateUser(id int, data *models.User) (models.User, error)
	FindByUserName(userName string) (models.User, error)
}
