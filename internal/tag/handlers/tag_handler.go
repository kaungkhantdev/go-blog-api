package handlers

import (
	"go-blog-api/internal/tag/handlers/requests"
	"go-blog-api/internal/tag/models"
	"go-blog-api/internal/tag/services"
	"go-blog-api/pkg/utils"
	"go-blog-api/pkg/validator"
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

// Helper Methods

func (handler *TagHandler) bindAndValidate(context *gin.Context, input interface{}) error {
	if err := context.ShouldBindJSON(input); err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return err
	}

	if err := validator.ValidateStruct(input); err != nil {
		utils.ErrorResponse(context, "Email field is missing.", http.StatusBadRequest)
		return err
	}

	return nil

}

func (handler *TagHandler) handleResponse(context *gin.Context, message string, data interface{}, err error) {
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SuccessResponse(context, data, message, http.StatusOK)
}

// Methods

func (handler *TagHandler) CreateTag(context *gin.Context) {
	var input requests.TagCreateRequest
	if handler.bindAndValidate(context, &input) != nil {
		return
	}

	tag := &models.Tag{
		Name:     input.Name,
		UserId:   input.UserId,
		ParentId: &input.ParentId,
		IconId:   input.IconId,
	}

	data, err := handler.tagService.CreateTag(tag)
	handler.handleResponse(context, "Success", data, err)
}

func (handler *TagHandler) UpdateTag(context *gin.Context) {
	id := context.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	var input requests.TagUpdateRequest
	if handler.bindAndValidate(context, input) != nil {
		return
	}

	tag := &models.Tag{
		Name:     input.Name,
		UserId:   input.UserId,
		ParentId: &input.ParentId,
		IconId:   input.IconId,
	}

	data, err := handler.tagService.UpdateTag(intId, tag)
	handler.handleResponse(context, "Success", data, err)

}

func (handler *TagHandler) FindWithPagination(context *gin.Context) {

	data, err := handler.tagService.FindWithPagination(context)
	handler.handleResponse(context, "Success", data, err)

}
