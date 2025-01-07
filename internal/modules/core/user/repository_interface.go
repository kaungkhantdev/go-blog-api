package user

type UserRepositoryInterface interface {
	CreateUser(data *UserEntity) (UserEntity, error)
	FindByIdUser(id int) (UserEntity, error)
	FindByEmailUser(email string) (UserEntity, error)
	UpdateUser(id int, data *UserEntity) (UserEntity, error)
	FindByUserName(userName string) (UserEntity, error)
}
