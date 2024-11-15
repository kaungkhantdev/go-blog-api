package interfaces

import "go-blog-api/internal/user/models"

type UserRepositoryInterface interface {
	CreateUser(data *models.User) (models.User, error)
	FindByIdUser(id int) (models.User, error)
	//FindAllUser() ([]models.User, error)
	//UpdateUser(id int, data *models.User) (models.User, error)
}
