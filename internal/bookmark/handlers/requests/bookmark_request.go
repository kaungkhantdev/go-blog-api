package requests

type CreateBookmarkRequest struct {
	ArticleId int `json:"article_id" validate:"required"`
}

type UpdateBookmarkRequest struct {
	CreateBookmarkRequest
}
