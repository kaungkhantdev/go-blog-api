package services

import (
	"go-blog-api/internal/bookmark/interfaces"
	"go-blog-api/internal/bookmark/models"
)

type BookmarkService struct {
	repo interfaces.BookmarkRepositoryInterface
}

func NewBookmarkService(repo interfaces.BookmarkRepositoryInterface) *BookmarkService {
	return &BookmarkService{repo: repo}
}

func (service *BookmarkService) CreateBookmark(data *models.Bookmark) (models.Bookmark, error) {
	return service.repo.CreateBookmark(data)
}

func (service *BookmarkService) UpdateBookmark(id int, data *models.Bookmark) (models.Bookmark, error) {
	return service.repo.UpdateBookmark(id, data)
}
