package repository

import (
	iconModel "go-blog-api/internal/icon/models"
	"go-blog-api/internal/tag/handlers/requests"
	"go-blog-api/internal/tag/interfaces"
	"go-blog-api/internal/tag/models"
	userModel "go-blog-api/internal/user/models"
	"go-blog-api/pkg/pagination"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) interfaces.TagRepositoryInterface {
	return &TagRepository{db: db}
}

func (repo *TagRepository) CreateTag(input requests.TagCreateRequest) (models.Tag, error) {
	// Check if UserId exists
	var user userModel.User
	if err := repo.db.First(&user, input.UserId).Error; err != nil {
		return models.Tag{}, err
	}

	// Check if IconId exists
	var icon iconModel.Icon
	if err := repo.db.First(&icon, input.IconId).Error; err != nil {
		return models.Tag{}, err
	}

	// Check if ParentId exists
	var parentTag models.Tag
	if input.ParentId != nil {
		if err := repo.db.First(&parentTag, input.ParentId).Error; err != nil {
			return models.Tag{}, err
		}
		input.ParentId = &parentTag.ID
	}

	// Create the new tag
	tag := models.Tag{
		Name:     input.Name,
		ParentId: input.ParentId,
		UserId:   input.UserId,
		IconId:   input.IconId,
	}

	// Insert the tag into the database
	if err := repo.db.Create(&tag).Error; err != nil {
		return models.Tag{}, err
	}

	return tag, nil
}

func (repo *TagRepository) UpdateTag(id int, input requests.TagUpdateRequest) (models.Tag, error) {
	// First, find the existing tag by ID
	var tag models.Tag
	if err := repo.db.Preload("Parent").First(&tag, id).Error; err != nil {
		return models.Tag{}, err
	}

	// Check if the user exists
	var user userModel.User
	if err := repo.db.First(&user, input.UserId).Error; err != nil {
		return models.Tag{}, err
	}

	// Check if the Icon exists
	var icon iconModel.Icon
	if err := repo.db.First(&icon, input.IconId).Error; err != nil {
		return models.Tag{}, err
	}

	// Check if ParentId exists (if provided)
	var parentTag models.Tag
	if input.ParentId != nil {
		if err := repo.db.First(&parentTag, input.ParentId).Error; err != nil {
			return models.Tag{}, err
		}
	}

	// Set the fields for update (tag's fields)
	tag.Name = input.Name
	tag.ParentId = input.ParentId
	tag.UserId = input.UserId
	tag.IconId = input.IconId

	// Now, update the existing tag with the modified fields
	if err := repo.db.Model(&tag).Where("id = ?", tag.ID).Updates(tag).Error; err != nil {
		return models.Tag{}, err
	}

	// Return the updated tag
	return tag, nil
}

func (repo *TagRepository) FindWithPagination(ctx *gin.Context) ([]models.Tag, error) {
	var tags []models.Tag
	if err := repo.db.Scopes(pagination.Paginate(ctx)).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}
