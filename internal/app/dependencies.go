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

	articleHandler "go-blog-api/internal/article/handlers"
	articleRepo "go-blog-api/internal/article/repository"
	articleService "go-blog-api/internal/article/services"

	iconRepo "go-blog-api/internal/icon/repository"

	bookmarkHandler "go-blog-api/internal/bookmark/handlers"
	bookmarkRepo "go-blog-api/internal/bookmark/repository"
	bookmarkService "go-blog-api/internal/bookmark/services"

	reactionHandler "go-blog-api/internal/reaction/handlers"
	reactionRepo "go-blog-api/internal/reaction/repository"
	reactionService "go-blog-api/internal/reaction/services"

	reactionTypeRepo "go-blog-api/internal/reaction_type/repository"

	commentHandler "go-blog-api/internal/comment/handlers"
	commentRepo "go-blog-api/internal/comment/repository"
	commentService "go-blog-api/internal/comment/services"
)

type Dependencies struct {
	UserHandler     *userHandler.UserHandler
	AuthHandler     *authHandler.AuthHandler
	TagHandler      *tagHandler.TagHandler
	ArticleHandler  *articleHandler.ArticleHandler
	BookmarkHandler *bookmarkHandler.BookmarkHandler
	ReactionHandler *reactionHandler.ReactionHandler
	CommentHandler  *commentHandler.CommentHandler
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

	// icon
	IconRepo := iconRepo.NewIconRepository(DB)

	// tag
	TagRepo := tagRepo.NewTagRepository(DB)
	TagService := tagService.NewTagService(TagRepo, UserRepo, IconRepo)
	TagHandler := tagHandler.NewTagHandler(TagService)

	// article
	ArticleRepo := articleRepo.NewArticleRepository(DB)
	ArticleService := articleService.NewArticleService(ArticleRepo, UserRepo, TagRepo)
	ArticleHandler := articleHandler.NewArticleHandler(ArticleService)

	// bookmark
	BookmarkRepo := bookmarkRepo.NewBookmarkRepository(DB)
	BookmarkService := bookmarkService.NewBookmarkService(BookmarkRepo, ArticleRepo)
	BookmarkHandler := bookmarkHandler.NewBookmarkHandler(BookmarkService)

	// reaction type
	ReactionTypeRepo := reactionTypeRepo.NewReactionTypeRepository(DB)

	// reaction
	ReactionRepo := reactionRepo.NewReactionRepository(DB)
	ReactionService := reactionService.NewReactionService(ReactionRepo, ArticleRepo, ReactionTypeRepo)
	ReactionHandler := reactionHandler.NewReactionHandler(ReactionService)

	// comment
	CommentRepo := commentRepo.NewCommentReposity(DB)
	CommentService := commentService.NewCommentService(CommentRepo, ArticleRepo)
	CommentHandler := commentHandler.NewCommentHandler(CommentService)

	return &Dependencies{
		UserHandler:     UserHandler,
		AuthHandler:     AuthHandler,
		TagHandler:      TagHandler,
		ArticleHandler:  ArticleHandler,
		BookmarkHandler: BookmarkHandler,
		ReactionHandler: ReactionHandler,
		CommentHandler:  CommentHandler,
	}, err
}
