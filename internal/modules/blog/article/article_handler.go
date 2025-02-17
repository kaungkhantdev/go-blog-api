package article

import (
	"go-blog-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleService *ArticleService
}

func NewArticleHandler(articleService *ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: articleService}
}

func (handler *ArticleHandler) CreateArticle(context *gin.Context) {
	var input CreateArticleRequest
	if err := utils.BindAndValidate(context, &input); err != nil {
		return
	}

	userId, err := utils.GetUserIdFromGin(context)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.articleService.CreateArticle(userId, input)
	utils.HandleResponse(context, "Success", data, err)
}

func (handler *ArticleHandler) UpdateArticle(context *gin.Context) {
	id := context.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	var input UpdateArticleRequest
	if err := utils.BindAndValidate(context, &input); err != nil {
		return
	}

	userId, err := utils.GetUserIdFromGin(context)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.articleService.UpdateArticle(intId, userId, input)
	utils.HandleResponse(context, "Success", data, err)

}

func (handler *ArticleHandler) FindById(context *gin.Context) {
	id := context.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.articleService.FindOneById(intId)
	utils.HandleResponse(context, "Success", data, err)

}

func (handler *ArticleHandler) FindWithPagination(context *gin.Context) {
	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	pageSize, err := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	if err != nil {
		utils.ErrorResponse(context, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.articleService.FindWithPagination(page, pageSize)
	utils.HandleResponse(context, "Success", data, err)

}
