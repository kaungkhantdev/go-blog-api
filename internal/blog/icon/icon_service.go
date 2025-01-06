package icon

type IconService struct {
	repo IconRepositoryInterface
}

func NewIconService(repo *IconRepositoryInterface) *IconService {
	return &IconService{repo: *repo}
}

func (service *IconService) CreateIcon(data *IconEntity) (IconEntity, error) {
	newIcon, err := service.repo.CreateIcon(data)
	if err != nil {
		return IconEntity{}, err
	}

	return newIcon, nil
}

func (service *IconService) UpdateIcon(id int, data *IconEntity) (IconEntity, error) {
	updateIcon, err := service.repo.UpdateIcon(id, data)
	if err != nil {
		return IconEntity{}, err
	}

	return updateIcon, nil
}

func (service *IconService) FindByName(name string) (IconEntity, error) {
	icon, err := service.repo.FindByName(name)
	if err != nil {
		return IconEntity{}, err
	}

	return icon, nil
}

func (service *IconService) FindbyIdIcon(id int) (IconEntity, error) {
	icon, err := service.repo.FindByIdIcon(id)
	if err != nil {
		return IconEntity{}, err
	}

	return icon, nil
}
