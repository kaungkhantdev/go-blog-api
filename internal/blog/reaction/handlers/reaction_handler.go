package handlers

import (
	"go-blog-api/internal/blog/reaction/handlers/requests"
	"go-blog-api/internal/blog/reaction/services"
	"go-blog-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReactionHandler struct {
	service *services.ReactionService
}

func NewReactionHandler(service *services.ReactionService) *ReactionHandler {
	return &ReactionHandler{service: service}
}

func (handler *ReactionHandler) CreateReaction(context *gin.Context) {
	var input requests.CreateReactionRequest
	if err := utils.BindAndValidate(context, &input); err != nil {
		return
	}

	userId, err := utils.GetUserIdFromGin(context)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.service.CreateReaction(userId, &input)
	utils.HandleResponse(context, "Success", data, err)
}

func (handler *ReactionHandler) UpdateReaction(context *gin.Context) {
	var input requests.UpdateReactionRequest
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

	data, err := handler.service.UpdateReaction(intId, userId, &input)
	utils.HandleResponse(context, "Success", data, err)
}
