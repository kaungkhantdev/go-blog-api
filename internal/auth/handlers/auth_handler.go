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

// Helper Methods

func (handler *AuthHandler) bindAndValidate(context *gin.Context, input interface{}) error {
	if err := context.ShouldBindJSON(input); err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return err
	}

	if err := validator.ValidateStruct(input); err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return err
	}

	return nil

}

func (handler *AuthHandler) handleResponse(context *gin.Context, message string, data interface{}, err error) {
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SuccessResponse(context, data, message, http.StatusOK)
}

// Methods

func (handler *AuthHandler) GetOtpViaEmail(context *gin.Context) {
	var input requests.AuthOtpRequest

	if handler.bindAndValidate(context, &input) != nil {
		return
	}

	data, err := handler.authService.GetOtpViaEmail(input.Email)
	handler.handleResponse(context, "Success", data, err)
}

func (handler *AuthHandler) VerifyOtpViaEmail(context *gin.Context) {

	var input requests.AuthVerifyOtpRequest

	if handler.bindAndValidate(context, &input) != nil {
		return
	}

	inputData := map[string]string{
		"otp":   input.Otp,
		"email": input.Email,
	}

	data, err := handler.authService.VerifyOtpViaEmail(inputData)
	handler.handleResponse(context, "Success", data, err)
}

func (handler *AuthHandler) SignUp(context *gin.Context) {
	var inputs requests.AuthSignUpRequest

	if handler.bindAndValidate(context, &inputs) != nil {
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
	handler.handleResponse(context, "Success", data, err)
}

func (handler *AuthHandler) SignIn(context *gin.Context) {

	var inputs requests.AuthOtpRequest

	if handler.bindAndValidate(context, &inputs) != nil {
		return
	}

	data, err := handler.authService.SignIn(inputs.Email)
	handler.handleResponse(context, "Success", data, err)
}

func (handler *AuthHandler) VerifyRefreshToken(context *gin.Context) {

	var inputs requests.VerifyRefreshTokenRequest

	if handler.bindAndValidate(context, &inputs) != nil {
		return
	}

	data, err := handler.authService.VerifyRefreshToken(inputs.RefreshToken)
	handler.handleResponse(context, "Success", data, err)
}
