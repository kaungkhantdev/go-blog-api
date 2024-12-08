package handlers

import (
	"go-blog-api/internal/comment/handlers/requests"
	"go-blog-api/internal/comment/services"
	"go-blog-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service *services.CommentService
}

func NewCommentHandler(service *services.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

func (handler *CommentHandler) CreateComment(context *gin.Context) {
	var input requests.CreateCommentRequest
	if err := utils.BindAndValidate(context, &input); err != nil {
		return
	}

	userId, err := utils.GetUserIdFromGin(context)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.service.CreateComment(userId, input)
	utils.HandleResponse(context, "Success", data, err)
}

func (handler *CommentHandler) UpdateComment(context *gin.Context) {
	var input requests.UpdateCommentRequest
	if err := utils.BindAndValidate(context, &input); err != nil {
		return
	}

	id := context.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIdFromGin(context)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.service.UpdateComment(intId, userId, input)
	utils.HandleResponse(context, "Success", data, err)
}
