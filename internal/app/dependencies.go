package app

import (
	"go-blog-api/database"
	"go-blog-api/internal/user/handlers"
	"go-blog-api/internal/user/repository"
	"go-blog-api/internal/user/services"
)

type Dependencies struct {
	UserHandler *handlers.UserHandler
}

func NewAppDependencies() (*Dependencies, error) {

	// database connection
	DB, err := database.Connect()
	if err != nil {
		return nil, err
	}

	var UserRepo = repository.NewUserRepository(DB)
	UserService := services.NewUserService(UserRepo)
	UserHandler := handlers.NewUserHandler(UserService)

	return &Dependencies{
		UserHandler: UserHandler,
	}, err
}
