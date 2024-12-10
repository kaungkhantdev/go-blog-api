package handlers

import (
	"go-blog-api/internal/tag/handlers/requests"
	"go-blog-api/internal/tag/services"
	"go-blog-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	tagService *services.TagService
}

func NewTagHandler(tagService *services.TagService) *TagHandler {
	return &TagHandler{
		tagService: tagService,
	}
}

// Methods

func (handler *TagHandler) CreateTag(context *gin.Context) {
	var input requests.TagCreateRequest
	if utils.BindAndValidate(context, &input) != nil {
		return
	}

	userId, err := utils.GetUserIdFromGin(context)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.tagService.CreateTag(input, userId)
	utils.HandleResponse(context, "Success", data, err)
}

func (handler *TagHandler) UpdateTag(context *gin.Context) {
	id := context.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	var input requests.TagUpdateRequest
	if utils.BindAndValidate(context, &input) != nil {
		return
	}

	userId, err := utils.GetUserIdFromGin(context)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.tagService.UpdateTag(intId, input, userId)
	utils.HandleResponse(context, "Success", data, err)

}

func (handler *TagHandler) FindWithPagination(context *gin.Context) {
	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	pageSize, err := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.tagService.FindWithPagination(page, pageSize)
	utils.HandleResponse(context, "Success", data, err)

}

func (handler *TagHandler) FindById(context *gin.Context) {
	id := context.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	tag, err := handler.tagService.FindByIdTag(intId)
	utils.HandleResponse(context, "Success", tag, err)
}
