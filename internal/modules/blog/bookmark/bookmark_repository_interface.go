package bookmark

type BookmarkRepositoryInterface interface {
	CreateBookmark(data *BookmarkEntity) (BookmarkEntity, error)
	UpdateBookmark(id int, data *BookmarkEntity) (BookmarkEntity, error)
	FindOneById(id int, userId int) (BookmarkEntity, error)
}
