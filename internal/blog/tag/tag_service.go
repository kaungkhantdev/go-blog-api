package tag

import (
	"go-blog-api/internal/blog/icon"
	"go-blog-api/internal/core/user"
	"go-blog-api/pkg/pagination"
)

type TagService struct {
	repo     TagRepositoryInterface
	userRepo user.UserRepositoryInterface
	iconRepo icon.IconRepositoryInterface
}

// instance
func NewTagService(
	repo TagRepositoryInterface,
	userRepo user.UserRepositoryInterface,
	iconRepo icon.IconRepositoryInterface,
) *TagService {
	return &TagService{
		repo:     repo,
		userRepo: userRepo,
		iconRepo: iconRepo,
	}
}

// public methods
func (service *TagService) CreateTag(data TagCreateRequest, userId int) (TagEntity, error) {
	if err := service.validateDependencies(userId, data.IconId, data.ParentId); err != nil {
		return TagEntity{}, err
	}

	return service.repo.CreateTag(data, userId)
}

func (service *TagService) UpdateTag(id int, data TagUpdateRequest, userId int) (TagEntity, error) {

	if err := service.validateDependencies(userId, data.IconId, data.ParentId); err != nil {
		return TagEntity{}, err
	}

	return service.repo.UpdateTag(id, data)
}

func (service *TagService) FindWithPagination(page, pageSize int) (*pagination.PaginatedResponse, error) {
	return service.repo.FindWithPagination(page, pageSize)
}

func (service *TagService) FindByIdTag(id int) (TagEntity, error) {
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
