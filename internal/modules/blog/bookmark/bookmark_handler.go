package bookmark

import (
	"go-blog-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookmarkHandler struct {
	service *BookmarkService
}

func NewBookmarkHandler(service *BookmarkService) *BookmarkHandler {
	return &BookmarkHandler{service: service}
}

func (handler *BookmarkHandler) CreateBookmark(context *gin.Context) {
	var input CreateBookmarkRequest
	if err := utils.BindAndValidate(context, &input); err != nil {
		return
	}

	userId, err := utils.GetUserIdFromGin(context)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.service.CreateBookmark(userId, input)
	utils.HandleResponse(context, "Success", data, err)
}

func (handler *BookmarkHandler) UpdateBookmark(context *gin.Context) {
	id := context.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	var input UpdateBookmarkRequest
	if err := utils.BindAndValidate(context, &input); err != nil {
		return
	}

	userId, err := utils.GetUserIdFromGin(context)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.service.UpdateBookmark(intId, userId, input)
	utils.HandleResponse(context, "Success", data, err)
}

func (handler *BookmarkHandler) FindOneById(context *gin.Context) {
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

	data, err := handler.service.FindOneById(intId, userId)
	utils.HandleResponse(context, "Success", data, err)
}
