package handlers

import (
	"go-blog-api/internal/auth/handlers/requests"
	"go-blog-api/internal/auth/services"
	"go-blog-api/pkg/utils"
	"go-blog-api/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (handler *AuthHandler) GetOtpViaEmail(context *gin.Context) {
	var input requests.AuthOtpRequest

	if err := context.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.ValidateStruct(&input); err != nil {
		utils.ErrorResponse(context, "Email field is missing.", http.StatusBadRequest)
		return
	}

	data, err := handler.authService.GetOtpViaEmail(input.Email)


	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SuccessResponse(context, "", data, http.StatusOK)
}


func (handler *AuthHandler) VerifyOtpViaEmail(context *gin.Context) {

	var input requests.AuthVerifyOtpRequest

	if err := context.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(context, "Invalid inputs", http.StatusBadRequest)
		return
	}

	if err := validator.ValidateStruct(&input); err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	inputData := map[string]string{
		"otp": input.Otp,
		"email": input.Email,
	}

	data, err := handler.authService.VerifyOtpViaEmail(inputData)

	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}


	utils.SuccessResponse(context, data, "success", http.StatusOK)
}


func (handler *AuthHandler) SignUp(context *gin.Context) {
	var inputs requests.AuthSignUpRequest

	if err := context.ShouldBindJSON(&inputs); err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.ValidateStruct(&inputs); err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	inputData :=  map[string]string{
		"email": inputs.Email,
		"name": inputs.Name,
		"user_name": inputs.UserName,
		"avatar_url": inputs.AvatarUrl,
		"bio": inputs.Bio,
	}

	data, err := handler.authService.SignUp(inputData)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SuccessResponse(context, data, "success", http.StatusOK)
}


func (handler *AuthHandler) SignIn(context *gin.Context) {

	var inputs requests.AuthOtpRequest

	if err := context.ShouldBindJSON(&inputs); err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.ValidateStruct(&inputs); err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.authService.SignIn(inputs.Email)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SuccessResponse(context, data, "success", http.StatusOK)
}
