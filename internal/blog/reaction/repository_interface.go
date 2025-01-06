package reaction

type ReactionRepositoryInterface interface {
	CreateReaction(data *ReactionEntity) (ReactionEntity, error)
	UpdateReaction(id int, data *ReactionEntity) (ReactionEntity, error)
}
