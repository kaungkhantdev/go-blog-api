package reaction_type

type ReactionTypeRepositoryInterface interface {
	CreateReactionType(data *ReactionTypeEntity) (ReactionTypeEntity, error)
	UpdateReactionType(id int, data *ReactionTypeEntity) (ReactionTypeEntity, error)
	FindOneById(id int) (ReactionTypeEntity, error)
}
