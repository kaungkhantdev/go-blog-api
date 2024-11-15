package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status		int			`json:"status"`
	Message		string		`json:"message"`
	Data		interface{}	`json:"data,omitempty"`
}

func SuccessResponse(context *gin.Context, data interface{}, message string, status int) {
	context.JSON(status, Response{
		Status: status,
		Message: message,
		Data: data,
	})
}

func ErrorResponse(context *gin.Context, message string, status int) {
	context.JSON(status, Response{
		Status: status,
		Message: message,
	})
}