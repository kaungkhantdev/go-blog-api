package services

import (
	iconRepoInterface "go-blog-api/internal/icon/interfaces"
	"go-blog-api/internal/tag/handlers/requests"
	"go-blog-api/internal/tag/interfaces"
	"go-blog-api/internal/tag/models"
	userRepoInterface "go-blog-api/internal/user/interfaces"
	"go-blog-api/pkg/pagination"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagService struct {
	repo     interfaces.TagRepositoryInterface
	userRepo userRepoInterface.UserRepositoryInterface
	iconRepo iconRepoInterface.IconRepositoryInterface
}

// instance
func NewTagService(
	repo interfaces.TagRepositoryInterface,
	userRepo userRepoInterface.UserRepositoryInterface,
	iconRepo iconRepoInterface.IconRepositoryInterface,
) *TagService {
	return &TagService{
		repo:     repo,
		userRepo: userRepo,
		iconRepo: iconRepo,
	}
}

// public methods
func (service *TagService) CreateTag(data requests.TagCreateRequest) (models.Tag, error) {
	if err := service.validateDependencies(data.UserId, data.IconId, data.ParentId); err != nil {
		return models.Tag{}, err
	}

	return service.repo.CreateTag(data)
}

func (service *TagService) UpdateTag(id int, data requests.TagUpdateRequest) (models.Tag, error) {

	if err := service.validateDependencies(data.UserId, data.IconId, data.ParentId); err != nil {
		return models.Tag{}, err
	}

	return service.repo.UpdateTag(id, data)
}

func (service *TagService) FindWithPagination(ctx *gin.Context) (*pagination.PaginatedResponse, error) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	return service.repo.FindWithPagination(page, pageSize)
}

func (service *TagService) FindByIdTag(id int) (models.Tag, error) {
	return service.repo.FindByIdTag(id)
}

// private methods
func (service *TagService) validateDependencies(userId, iconId int, parentId *int) error {
	if _, err := service.userRepo.FindByIdUser(userId); err != nil {
		return err
	}

	// Check if IconId exists
	if _, err := service.iconRepo.FindByIdIcon(iconId); err != nil {
		return err
	}

	// Check if ParentId exists
	if parentId != nil {
		// Call FindByIdTag and capture both the tag and error
		if _, err := service.repo.FindByIdTag(*parentId); err != nil {
			return err
		}
	}

	return nil
}
