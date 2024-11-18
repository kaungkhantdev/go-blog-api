package app

import (
	"go-blog-api/internal/middleware"
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
	s.UserRoutes(router, deps)

	s.AuthRoutes(router, deps)

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
	authRoute.POST("/sign-up", middleware.AuthMiddleware(),  deps.AuthHandler.SignUp)
	authRoute.POST("/sign-in", middleware.AuthMiddleware(), deps.AuthHandler.SignIn)
}

func (s *Server) UserRoutes(router *gin.RouterGroup, deps *Dependencies) {
	userRouter := router.Group("/user")
	userRouter.GET("/:id", deps.UserHandler.FindOneByID)
}
