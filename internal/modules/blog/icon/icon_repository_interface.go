package icon

type IconRepositoryInterface interface {
	CreateIcon(data *IconEntity) (IconEntity, error)
	UpdateIcon(id int, data *IconEntity) (IconEntity, error)
	FindByIdIcon(id int) (IconEntity, error)
	FindByName(name string) (IconEntity, error)
}
