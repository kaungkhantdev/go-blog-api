package services

import (
	"go-blog-api/internal/blog/bookmark/handlers/requests"
	"go-blog-api/internal/blog/bookmark/interfaces"
	"go-blog-api/internal/blog/bookmark/models"

	articleInterface "go-blog-api/internal/blog/article/interfaces"
)

type BookmarkService struct {
	repo        interfaces.BookmarkRepositoryInterface
	articleRepo articleInterface.ArticleRepositoryInterfaces
}

func NewBookmarkService(
	repo interfaces.BookmarkRepositoryInterface,
	articleRepo articleInterface.ArticleRepositoryInterfaces,
) *BookmarkService {
	return &BookmarkService{repo: repo, articleRepo: articleRepo}
}

func (service *BookmarkService) CreateBookmark(userId int, input requests.CreateBookmarkRequest) (models.Bookmark, error) {
	article, err := service.articleRepo.FindOneById(input.ArticleId)
	if err != nil {
		return models.Bookmark{}, err
	}

	bookmark := &models.Bookmark{
		UserId:    userId,
		ArticleId: article.ID,
	}

	return service.repo.CreateBookmark(bookmark)
}

func (service *BookmarkService) UpdateBookmark(id int, userId int, input requests.UpdateBookmarkRequest) (models.Bookmark, error) {
	article, err := service.articleRepo.FindOneById(input.ArticleId)
	if err != nil {
		return models.Bookmark{}, err
	}

	bookmark := &models.Bookmark{
		UserId:    userId,
		ArticleId: article.ID,
	}

	return service.repo.UpdateBookmark(id, bookmark)
}

func (service *BookmarkService) FindOneById(id, userId int) (models.Bookmark, error) {
	return service.repo.FindOneById(id, userId)
}
