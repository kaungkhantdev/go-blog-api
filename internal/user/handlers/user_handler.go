package handlers

import (
	"fmt"
	"go-blog-api/internal/user/services"
	"go-blog-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) FindOneByID(context *gin.Context) {
	id := context.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
	
	user, err := handler.service.FindOneById(int(intId))
	
	if err != nil {
        fmt.Println("Error:", err)
        return
    }

	utils.SuccessResponse(context, user, "user data", http.StatusOK)
}
