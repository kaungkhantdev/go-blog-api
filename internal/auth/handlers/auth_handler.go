package handlers

import (
	"go-blog-api/internal/auth/handlers/requests"
	"go-blog-api/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (handler *AuthHandler) GetOtp(context *gin.Context) {
	var input requests.AuthEmailRequest

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inputs"})
	}

	if err := validator.ValidateStruct(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Email field is missing."})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"data": input})

}


func (handler *AuthHandler) SignUp(context *gin.Context) {
	var inputs requests.AuthSignUpRequest

	if err := context.ShouldBindJSON(&inputs); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inputs"})
		return
	}

	if err := validator.ValidateStruct(&inputs); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Some field is missing."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": inputs})
}
