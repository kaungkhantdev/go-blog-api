package article

type CreateArticleRequest struct {
	Title   string `json:"title" validate:"required,min=5"`
	Content string `json:"content" validate:"required,min=5"`
	Tags    []int  `json:"tags" validate:"required"`
}

type UpdateArticleRequest struct {
	CreateArticleRequest
}
