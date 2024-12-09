package requests

type CreateReactionRequest struct {
	ArticleId      int `json:"article_id" validate:"required"`
	ReactionTypeId int `json:"reaction_type_id" validate:"required"`
}

type UpdateReactionRequest struct {
	CreateReactionRequest
}
