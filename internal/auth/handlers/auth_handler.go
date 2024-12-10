package handlers

import (
	"go-blog-api/internal/auth/handlers/requests"
	"go-blog-api/internal/auth/services"
	"go-blog-api/pkg/utils"

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

// Methods

func (handler *AuthHandler) GetOtpViaEmail(context *gin.Context) {
	var input requests.AuthOtpRequest

	if utils.BindAndValidate(context, &input) != nil {
		return
	}

	data, err := handler.authService.GetOtpViaEmail(input.Email)
	utils.HandleResponse(context, "Success", data, err)
}

func (handler *AuthHandler) VerifyOtpViaEmail(context *gin.Context) {

	var input requests.AuthVerifyOtpRequest

	if utils.BindAndValidate(context, &input) != nil {
		return
	}

	inputData := map[string]string{
		"otp":   input.Otp,
		"email": input.Email,
	}

	data, err := handler.authService.VerifyOtpViaEmail(inputData)
	utils.HandleResponse(context, "Success", data, err)
}

func (handler *AuthHandler) SignUp(context *gin.Context) {
	var inputs requests.AuthSignUpRequest

	if utils.BindAndValidate(context, &inputs) != nil {
		return
	}

	inputData := map[string]string{
		"email":     inputs.Email,
		"name":      inputs.Name,
		"user_name": inputs.UserName,
		"avatar":    inputs.Avatar,
		"bio":       inputs.Bio,
	}

	data, err := handler.authService.SignUp(inputData)
	utils.HandleResponse(context, "Success", data, err)
}

func (handler *AuthHandler) SignIn(context *gin.Context) {

	var inputs requests.AuthOtpRequest

	if utils.BindAndValidate(context, &inputs) != nil {
		return
	}

	data, err := handler.authService.SignIn(inputs.Email)
	utils.HandleResponse(context, "Success", data, err)
}

func (handler *AuthHandler) VerifyRefreshToken(context *gin.Context) {

	var inputs requests.VerifyRefreshTokenRequest

	if utils.BindAndValidate(context, &inputs) != nil {
		return
	}

	data, err := handler.authService.VerifyRefreshToken(inputs.RefreshToken)
	utils.HandleResponse(context, "Success", data, err)
}
