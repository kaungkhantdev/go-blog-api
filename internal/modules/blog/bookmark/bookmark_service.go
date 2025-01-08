package bookmark

import "go-blog-api/internal/modules/blog/article"

type BookmarkService struct {
	repo        BookmarkRepositoryInterface
	articleRepo article.ArticleRepositoryInterfaces
}

func NewBookmarkService(
	repo BookmarkRepositoryInterface,
	articleRepo article.ArticleRepositoryInterfaces,
) *BookmarkService {
	return &BookmarkService{repo: repo, articleRepo: articleRepo}
}

func (service *BookmarkService) CreateBookmark(userId int, input CreateBookmarkRequest) (BookmarkEntity, error) {
	article, err := service.articleRepo.FindOneById(input.ArticleId)
	if err != nil {
		return BookmarkEntity{}, err
	}

	bookmark := &BookmarkEntity{
		UserId:    userId,
		ArticleId: article.ID,
	}

	return service.repo.CreateBookmark(bookmark)
}

func (service *BookmarkService) UpdateBookmark(id int, userId int, input UpdateBookmarkRequest) (BookmarkEntity, error) {
	article, err := service.articleRepo.FindOneById(input.ArticleId)
	if err != nil {
		return BookmarkEntity{}, err
	}

	bookmark := &BookmarkEntity{
		UserId:    userId,
		ArticleId: article.ID,
	}

	return service.repo.UpdateBookmark(id, bookmark)
}

func (service *BookmarkService) FindOneById(id, userId int) (BookmarkEntity, error) {
	return service.repo.FindOneById(id, userId)
}
