package app

import (
	"go-blog-api/database"
	"go-blog-api/pkg/validator"

	mail "go-blog-api/pkg/mail"

	otpRepo "go-blog-api/internal/otp/repository"
	otpService "go-blog-api/internal/otp/services"

	userHandler "go-blog-api/internal/user/handlers"
	userRepo "go-blog-api/internal/user/repository"
	userService "go-blog-api/internal/user/services"

	authHandler "go-blog-api/internal/auth/handlers"
	authService "go-blog-api/internal/auth/services"

	tagHandler "go-blog-api/internal/tag/handlers"
	tagRepo "go-blog-api/internal/tag/repository"
	tagService "go-blog-api/internal/tag/services"
)

type Dependencies struct {
	UserHandler *userHandler.UserHandler
	AuthHandler *authHandler.AuthHandler
	TagHandler *tagHandler.TagHandler
}

func NewAppDependencies() (*Dependencies, error) {

	// database connection
	DB, err := database.Connect()
	if err != nil {
		return nil, err
	}


	// Create an instance of EmailService
	emailConfig := mail.NewEmailConfig()
	emailService := mail.NewEmailService(emailConfig)


	// init validator
	validator.InitValidator()

	OtpRepo := otpRepo.NewOtpRepository(DB)
	OtpService := otpService.NewOtpService(OtpRepo)

	UserRepo := userRepo.NewUserRepository(DB)
	UserService := userService.NewUserService(UserRepo)
	UserHandler := userHandler.NewUserHandler(UserService)

	// auth
	AuthService := authService.NewAuthService(OtpService, UserService, emailService)
	AuthHandler := authHandler.NewAuthHandler(AuthService)

	// tag
	TagRepo := tagRepo.NewTagRepository(DB)
	TagService := tagService.NewTagService(TagRepo)
	TagHandler := tagHandler.NewTagHandler(TagService)

	return &Dependencies{
		UserHandler: UserHandler,
		AuthHandler: AuthHandler,
		TagHandler: TagHandler,
	}, err
}
