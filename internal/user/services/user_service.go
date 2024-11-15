package services

import (
	"go-blog-api/internal/user/models"
	"go-blog-api/internal/user/interfaces"
)

type UserService struct {
	repo interfaces.UserRepositoryInterface
}

func NewUserService(repo interfaces.UserRepositoryInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (service *UserService) CreateUser(data *models.User) (models.User, error) {
	newUser, err := service.repo.CreateUser(data)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func (service *UserService) FindOneById(id int) (models.User, error) {
	user, err := service.repo.FindByIdUser(id);
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}