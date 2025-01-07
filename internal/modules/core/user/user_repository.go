package user

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(user *UserEntity) (UserEntity, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return UserEntity{}, err
	}

	return *user, nil
}

func (repo *UserRepository) UpdateUser(id int, data *UserEntity) (UserEntity, error) {

	var user UserEntity
	if err := repo.db.First(&user, id).Error; err != nil {
		return UserEntity{}, err
	}

	// Update the user fields
	if err := repo.db.Model(&user).Updates(data).Error; err != nil {
		return UserEntity{}, err
	}

	return user, nil
}

func (repo *UserRepository) FindByIdUser(id int) (UserEntity, error) {
	var user UserEntity
	if err := repo.db.First(&user, id).Error; err != nil {
		return UserEntity{}, err
	}
	return user, nil
}

func (repo *UserRepository) FindByEmailUser(email string) (UserEntity, error) {
	var user UserEntity
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return UserEntity{}, err
	}
	return user, nil
}

func (repo *UserRepository) FindByUserName(userName string) (UserEntity, error) {
	var user UserEntity
	if err := repo.db.Where("user_name = ?", userName).First(&user).Error; err != nil {
		return UserEntity{}, err
	}
	return user, nil
}
