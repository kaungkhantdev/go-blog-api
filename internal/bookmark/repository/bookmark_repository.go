package repository

import (
	"go-blog-api/internal/bookmark/interfaces"
	"go-blog-api/internal/bookmark/models"

	"gorm.io/gorm"
)

type BookmarkRepository struct {
	db *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) interfaces.BookmarkRepositoryInterface {
	return &BookmarkRepository{db: db}
}

func (repo *BookmarkRepository) CreateBookmark(data *models.Bookmark) (models.Bookmark, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return models.Bookmark{}, err
	}

	return *data, nil
}

func (repo *BookmarkRepository) UpdateBookmark(id int, data *models.Bookmark) (models.Bookmark, error) {
	var bookmark models.Bookmark
	if err := repo.db.First(&bookmark, id).Error; err != nil {
		return models.Bookmark{}, err
	}

	if err := repo.db.Model(&bookmark).Updates(data).Error; err != nil {
		return models.Bookmark{}, err
	}
	return bookmark, nil
}
