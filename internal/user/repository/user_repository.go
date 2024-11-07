package repository

import (
	"go-blog-api/internal/user/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB;
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) CreateUser(user *models.User) (models.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return models.User{}, err
	}

	return *user, nil
}

func (repo *UserRepository) GetUserByID(id int) (*models.User, error) {
    var user models.User
    if err := repo.db.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}