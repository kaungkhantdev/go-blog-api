package app

import (
	"go-blog-api/database"
	"go-blog-api/pkg/validator"

	otpRepo 	"go-blog-api/internal/otp/repository"
	otpService "go-blog-api/internal/otp/services"
	
	userRepo 	"go-blog-api/internal/user/repository"
	userService "go-blog-api/internal/user/services"
	userHandler "go-blog-api/internal/user/handlers"

	authService	"go-blog-api/internal/auth/services"
	authHandler "go-blog-api/internal/auth/handlers"
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

	var OtpRepo = otpRepo.NewOtpRepository(DB)
	OtpService := otpService.NewOtpService(OtpRepo)
	
	var UserRepo = userRepo.NewUserRepository(DB)
	UserService := userService.NewUserService(UserRepo)
	UserHandler := userHandler.NewUserHandler(UserService)
	
	// auth
	AuthService := authService.NewAuthService(OtpService, UserService)
	AuthHandler := authHandler.NewAuthHandler(AuthService)

	return &Dependencies{
		UserHandler: UserHandler,
		AuthHandler: AuthHandler,
	}, err
}
