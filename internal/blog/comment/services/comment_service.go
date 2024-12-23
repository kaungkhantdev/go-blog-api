package services

import (
	"go-blog-api/internal/blog/comment/handlers/requests"
	"go-blog-api/internal/blog/comment/interfaces"
	"go-blog-api/internal/blog/comment/models"

	articleInterface "go-blog-api/internal/blog/article/interfaces"
)

type CommentService struct {
	repo        interfaces.CommentRepositoryInterface
	articleRepo articleInterface.ArticleRepositoryInterfaces
}

func NewCommentService(repo interfaces.CommentRepositoryInterface, articleRepo articleInterface.ArticleRepositoryInterfaces) *CommentService {
	return &CommentService{repo: repo, articleRepo: articleRepo}
}

func (service *CommentService) CreateComment(userId int, input requests.CreateCommentRequest) (models.Comment, error) {
	if _, err := service.articleRepo.FindOneById(input.ArticleId); err != nil {
		return models.Comment{}, err
	}

	if input.ParentId != 0 {
		if _, err := service.repo.FindOneById(input.ParentId); err != nil {
			return models.Comment{}, err
		}
	}

	comment := &models.Comment{
		UserId:    userId,
		ArticleId: input.ArticleId,
		ParentId:  nil, // Set to nil for top-level comments
		Content:   input.Content,
	}

	if input.ParentId != 0 {
		comment.ParentId = &input.ParentId
	}

	return service.repo.CreateComment(comment)
}

func (service *CommentService) UpdateComment(id, userId int, input requests.UpdateCommentRequest) (models.Comment, error) {
	if _, err := service.articleRepo.FindOneById(input.ArticleId); err != nil {
		return models.Comment{}, err
	}

	if input.ParentId != 0 {
		if _, err := service.repo.FindOneById(input.ParentId); err != nil {
			return models.Comment{}, err
		}
	}

	comment := &models.Comment{
		UserId:    userId,
		ArticleId: input.ArticleId,
		ParentId:  nil, // Set to nil for top-level comments
		Content:   input.Content,
	}

	if input.ParentId != 0 {
		comment.ParentId = &input.ParentId
	}

	return service.repo.UpdateComment(id, comment)
}

func (service *CommentService) FindOneById(id int) (models.Comment, error) {
	return service.repo.FindOneById(id)
}
