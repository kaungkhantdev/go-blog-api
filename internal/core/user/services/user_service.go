package services

import (
	"go-blog-api/internal/core/user/models"
	"go-blog-api/internal/core/user/interfaces"
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

func (service *UserService) FindByEmailUser(email string) (models.User, error) {
	user, err := service.repo.FindByEmailUser(email);
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (service *UserService) FindByUserName(userName string) (models.User, error) {
	user, err := service.repo.FindByUserName(userName);
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}



func (service *UserService) UpdateUser(id int, data *models.User) (models.User, error) {
	updateUser, err := service.repo.UpdateUser(id, data)
	if err != nil {
		return models.User{}, err
	}
	return updateUser, nil
}