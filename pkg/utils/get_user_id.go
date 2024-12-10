package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromGin(context *gin.Context) (int, error) {
	userId, exists := context.Get("userId")
	if !exists {
		return 0, errors.New("User ID not found in context")
	}

	// Assert userId to string
	intId, ok := userId.(int)
	if !ok {
		return 0, errors.New("Invalid user ID format")
	}
	return intId, nil
}
