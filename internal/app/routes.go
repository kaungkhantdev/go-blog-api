package app

import (
	authMiddleware "go-blog-api/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {

	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "hello, welcome"})
	})

	routerGroup := route.Group("/api/v1")
	router, deps := s.InitRoutes(routerGroup)

	s.AuthRoutes(router, deps)
	s.UserRoutes(router, deps)
	s.TagRoutes(router, deps)
	s.ArticleRoutes(router, deps)
	s.BookmarkRoutes(router, deps)

	return route
}

func (s *Server) InitRoutes(router *gin.RouterGroup) (*gin.RouterGroup, *Dependencies) {
	deps, err := NewAppDependencies()

	if err != nil {
		return router, nil
	}

	return router, deps
}

func (s *Server) AuthRoutes(router *gin.RouterGroup, deps *Dependencies) {
	authRoute := router.Group("/auth")
	authRoute.POST("/get-otp", deps.AuthHandler.GetOtpViaEmail)
	authRoute.POST("/verify-otp", deps.AuthHandler.VerifyOtpViaEmail)
	authRoute.POST("/sign-up", deps.AuthHandler.SignUp)
	authRoute.POST("/sign-in", deps.AuthHandler.SignIn)
	authRoute.POST("/verify-refresh-token", deps.AuthHandler.VerifyRefreshToken)
}

func (s *Server) UserRoutes(router *gin.RouterGroup, deps *Dependencies) {
	userRouter := router.Group("/users")
	userRouter.GET("/:id", authMiddleware.AuthMiddleware(), deps.UserHandler.FindOneByID)
}

func (s *Server) TagRoutes(router *gin.RouterGroup, deps *Dependencies) {
	tagRouter := router.Group("/tags")
	tagRouter.POST("", authMiddleware.AuthMiddleware(), deps.TagHandler.CreateTag)
	tagRouter.PUT("/:id", authMiddleware.AuthMiddleware(), deps.TagHandler.UpdateTag)

	tagRouter.GET("", deps.TagHandler.FindWithPagination)
	tagRouter.GET("/:id", deps.TagHandler.FindById)
}

func (s *Server) ArticleRoutes(router *gin.RouterGroup, deps *Dependencies) {
	articleRouter := router.Group("/articles")
	articleRouter.POST("", authMiddleware.AuthMiddleware(), deps.ArticleHandler.CreateArticle)
	articleRouter.PUT("/:id", authMiddleware.AuthMiddleware(), deps.ArticleHandler.UpdateArticle)

	articleRouter.GET("", deps.ArticleHandler.FindWithPagination)
	articleRouter.GET("/:id", deps.ArticleHandler.FindById)
}

func (s *Server) BookmarkRoutes(router *gin.RouterGroup, deps *Dependencies) {
	bookmarkRouter := router.Group("/bookmarks")
	bookmarkRouter.POST("", authMiddleware.AuthMiddleware(), deps.BookmarkHandler.CreateBookmark)
	bookmarkRouter.PUT("/:id", authMiddleware.AuthMiddleware(), deps.BookmarkHandler.UpdateBookmark)
	bookmarkRouter.GET("/:id", authMiddleware.AuthMiddleware(), deps.BookmarkHandler.FindOneById)
}
