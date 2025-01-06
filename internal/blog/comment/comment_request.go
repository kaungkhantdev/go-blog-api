package comment

type CreateCommentRequest struct {
	ParentId  int    `json:"parent_id"`
	Content   string `json:"content" validate:"required"`
	ArticleId int    `json:"article_id" validate:"required"`
}

type UpdateCommentRequest struct {
	CreateCommentRequest
}
