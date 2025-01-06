package user

type UserService struct {
	repo UserRepositoryInterface
}

func NewUserService(repo UserRepositoryInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (service *UserService) CreateUser(data *UserEntity) (UserEntity, error) {
	newUser, err := service.repo.CreateUser(data)
	if err != nil {
		return UserEntity{}, err
	}
	return newUser, nil
}

func (service *UserService) FindOneById(id int) (UserEntity, error) {
	user, err := service.repo.FindByIdUser(id)
	if err != nil {
		return UserEntity{}, err
	}
	return user, nil
}

func (service *UserService) FindByEmailUser(email string) (UserEntity, error) {
	user, err := service.repo.FindByEmailUser(email)
	if err != nil {
		return UserEntity{}, err
	}
	return user, nil
}

func (service *UserService) FindByUserName(userName string) (UserEntity, error) {
	user, err := service.repo.FindByUserName(userName)
	if err != nil {
		return UserEntity{}, err
	}
	return user, nil
}

func (service *UserService) UpdateUser(id int, data *UserEntity) (UserEntity, error) {
	updateUser, err := service.repo.UpdateUser(id, data)
	if err != nil {
		return UserEntity{}, err
	}
	return updateUser, nil
}
