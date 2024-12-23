package repository

import (
	"go-blog-api/internal/core/user/interfaces"
	"go-blog-api/internal/core/user/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(user *models.User) (models.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return models.User{}, err
	}

	return *user, nil
}

func (repo *UserRepository) UpdateUser(id int, data *models.User) (models.User, error) {

	var user models.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return models.User{}, err 
	}

	// Update the user fields
	if err := repo.db.Model(&user).Updates(data).Error; err != nil {
		return models.User{}, err 
	}

	return user, nil 
}

func (repo *UserRepository) FindByIdUser(id int) (models.User, error) {
	var user models.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo *UserRepository) FindByEmailUser(email string) (models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo *UserRepository) FindByUserName(userName string) (models.User, error) {
	var user models.User
	if err := repo.db.Where("user_name = ?", userName).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
