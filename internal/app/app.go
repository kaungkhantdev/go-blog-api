package app

import (
	"go-blog-api/database"
	"go-blog-api/internal/blog/article"
	"go-blog-api/internal/blog/bookmark"
	"go-blog-api/internal/blog/comment"
	"go-blog-api/internal/blog/icon"
	"go-blog-api/internal/blog/reaction"
	"go-blog-api/internal/blog/reaction_type"
	"go-blog-api/internal/blog/tag"
	"go-blog-api/internal/core/auth"
	"go-blog-api/internal/core/otp"
	"go-blog-api/internal/core/user"
	"go-blog-api/pkg/validator"

	mail "go-blog-api/pkg/mail"
)

type Dependencies struct {
	UserHandler     *user.UserHandler
	AuthHandler     *auth.AuthHandler
	TagHandler      *tag.TagHandler
	ArticleHandler  *article.ArticleHandler
	BookmarkHandler *bookmark.BookmarkHandler
	ReactionHandler *reaction.ReactionHandler
	CommentHandler  *comment.CommentHandler
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

	OtpRepo := otp.NewOtpRepository(DB)
	OtpService := otp.NewOtpService(OtpRepo)

	UserRepo := user.NewUserRepository(DB)
	UserService := user.NewUserService(UserRepo)
	UserHandler := user.NewUserHandler(UserService)

	// auth
	AuthService := auth.NewAuthService(OtpService, UserService, emailService)
	AuthHandler := auth.NewAuthHandler(AuthService)

	// icon
	IconRepo := icon.NewIconRepository(DB)

	// tag
	TagRepo := tag.NewTagRepository(DB)
	TagService := tag.NewTagService(TagRepo, UserRepo, IconRepo)
	TagHandler := tag.NewTagHandler(TagService)

	// article
	ArticleRepo := article.NewArticleRepository(DB)
	ArticleService := article.NewArticleService(ArticleRepo, UserRepo, TagRepo)
	ArticleHandler := article.NewArticleHandler(ArticleService)

	// bookmark
	BookmarkRepo := bookmark.NewBookmarkRepository(DB)
	BookmarkService := bookmark.NewBookmarkService(BookmarkRepo, ArticleRepo)
	BookmarkHandler := bookmark.NewBookmarkHandler(BookmarkService)

	// reaction type
	ReactionTypeRepo := reaction_type.NewReactionTypeRepository(DB)

	// reaction
	ReactionRepo := reaction.NewReactionRepository(DB)
	ReactionService := reaction.NewReactionService(ReactionRepo, ArticleRepo, ReactionTypeRepo)
	ReactionHandler := reaction.NewReactionHandler(ReactionService)

	// comment
	CommentRepo := comment.NewCommentRepository(DB)
	CommentService := comment.NewCommentService(CommentRepo, ArticleRepo)
	CommentHandler := comment.NewCommentHandler(CommentService)

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
