package app

import (
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

	return route
}

func (s *Server) InitRoutes(router *gin.RouterGroup) (*gin.RouterGroup, *Dependencies) {
	deps, err := NewAppDependencies()

	if err != nil {
		return router, nil
	}

	return router, deps
}

func (s *Server) UserRoutes(router *gin.RouterGroup, deps *Dependencies) {
	userRouter := router.Group("/user")
	userRouter.GET("/:id", deps.UserHandler.FindOneByID)
}
