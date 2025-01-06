package comment

import "go-blog-api/internal/blog/article"

type CommentService struct {
	repo        CommentRepositoryInterface
	articleRepo article.ArticleRepositoryInterfaces
}

func NewCommentService(repo CommentRepositoryInterface, articleRepo article.ArticleRepositoryInterfaces) *CommentService {
	return &CommentService{repo: repo, articleRepo: articleRepo}
}

func (service *CommentService) CreateComment(userId int, input CreateCommentRequest) (CommentEntity, error) {
	if _, err := service.articleRepo.FindOneById(input.ArticleId); err != nil {
		return CommentEntity{}, err
	}

	if input.ParentId != 0 {
		if _, err := service.repo.FindOneById(input.ParentId); err != nil {
			return CommentEntity{}, err
		}
	}

	comment := &CommentEntity{
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

func (service *CommentService) UpdateComment(id, userId int, input UpdateCommentRequest) (CommentEntity, error) {
	if _, err := service.articleRepo.FindOneById(input.ArticleId); err != nil {
		return CommentEntity{}, err
	}

	if input.ParentId != 0 {
		if _, err := service.repo.FindOneById(input.ParentId); err != nil {
			return CommentEntity{}, err
		}
	}

	comment := &CommentEntity{
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

func (service *CommentService) FindOneById(id int) (CommentEntity, error) {
	return service.repo.FindOneById(id)
}
