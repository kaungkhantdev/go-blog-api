package handlers

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/internal/auth/handlers/requests"
	"net/http"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (handler *AuthHandler) SignUp(context *gin.Context) {
	var inputs requests.AuthSignUpRequest

	if err := context.ShouldBindJSON(&inputs); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inputs"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": inputs})
}
