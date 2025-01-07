package comment

type CommentRepositoryInterface interface {
	CreateComment(data *CommentEntity) (CommentEntity, error)
	UpdateComment(id int, data *CommentEntity) (CommentEntity, error)
	FindOneById(id int) (CommentEntity, error)
}
