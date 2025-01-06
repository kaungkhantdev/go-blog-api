package bookmark

import (
	"gorm.io/gorm"
)

type BookmarkRepository struct {
	db *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) BookmarkRepositoryInterface {
	return &BookmarkRepository{db: db}
}

func (repo *BookmarkRepository) CreateBookmark(data *BookmarkEntity) (BookmarkEntity, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return BookmarkEntity{}, err
	}

	return *data, nil
}

func (repo *BookmarkRepository) UpdateBookmark(id int, data *BookmarkEntity) (BookmarkEntity, error) {
	var bookmark BookmarkEntity
	if err := repo.db.First(&bookmark, id).Error; err != nil {
		return BookmarkEntity{}, err
	}

	if err := repo.db.Model(&bookmark).Updates(data).Error; err != nil {
		return BookmarkEntity{}, err
	}
	return bookmark, nil
}

func (repo *BookmarkRepository) FindOneById(id int, userId int) (BookmarkEntity, error) {
	var bookmark BookmarkEntity
	err := repo.db.
		Preload("Article").
		Preload("User").
		Where("user_id = ? AND id = ?", userId, id).
		First(&bookmark).
		Error

	if err != nil {
		return BookmarkEntity{}, err
	}

	return bookmark, nil
}
