package reaction_type

type ReactionTypeService struct {
	repo ReactionTypeRepositoryInterface
}

func NewReactionTypeService(repo ReactionTypeRepositoryInterface) *ReactionTypeService {
	return &ReactionTypeService{repo: repo}
}

func (service *ReactionTypeService) CreateReactionType(data *ReactionTypeEntity) (ReactionTypeEntity, error) {
	return service.repo.CreateReactionType(data)
}

func (service *ReactionTypeService) UpdateReactionType(id int, data *ReactionTypeEntity) (ReactionTypeEntity, error) {
	return service.repo.UpdateReactionType(id, data)
}
