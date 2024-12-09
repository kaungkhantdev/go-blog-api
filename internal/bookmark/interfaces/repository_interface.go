package interfaces

import "go-blog-api/internal/bookmark/models"

type BookmarkRepositoryInterface interface {
	CreateBookmark(data *models.Bookmark) (models.Bookmark, error)
	UpdateBookmark(id int, data *models.Bookmark) (models.Bookmark, error)
	FindOneById(id int, userId int) (models.Bookmark, error)
}
