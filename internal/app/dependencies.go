package app

import (
	"go-blog-api/database"
	authHandler "go-blog-api/internal/auth/handlers"
	userHandler "go-blog-api/internal/user/handlers"
	"go-blog-api/internal/user/repository"
	"go-blog-api/internal/user/services"
	"go-blog-api/pkg/validator"
)

type Dependencies struct {
	UserHandler *userHandler.UserHandler
	AuthHandler *authHandler.AuthHandler
}

func NewAppDependencies() (*Dependencies, error) {

	// database connection
	DB, err := database.Connect()
	if err != nil {
		return nil, err
	}

	// init validator
	validator.InitValidator()

	// auth
	AuthHandler := authHandler.NewAuthHandler()

	var UserRepo = repository.NewUserRepository(DB)
	UserService := services.NewUserService(UserRepo)
	UserHandler := userHandler.NewUserHandler(UserService)

	return &Dependencies{
		UserHandler: UserHandler,
		AuthHandler: AuthHandler,
	}, err
}
