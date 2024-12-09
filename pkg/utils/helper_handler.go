package utils

import (
	"go-blog-api/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindAndValidate(context *gin.Context, input interface{}) error {
	if err := context.ShouldBindJSON(input); err != nil {
		ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return err
	}

	if err := validator.ValidateStruct(input); err != nil {
		ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return err
	}

	return nil

}

func HandleResponse(context *gin.Context, message string, data interface{}, err error) {
	if err != nil {
		ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessResponse(context, data, message, http.StatusOK)
}
