package repository

import (
	"errors"
	"go-blog-api/internal/blog/tag/handlers/requests"
	"go-blog-api/internal/blog/tag/interfaces"
	"go-blog-api/internal/blog/tag/models"
	"go-blog-api/pkg/pagination"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) interfaces.TagRepositoryInterface {
	return &TagRepository{db: db}
}

func (repo *TagRepository) FindByIdTag(id int) (models.Tag, error) {
	var tag models.Tag
	if err := repo.db.First(&tag, id).Error; err != nil {
		return models.Tag{}, err
	}
	return tag, nil
}

func (repo *TagRepository) CreateTag(input requests.TagCreateRequest, userId int) (models.Tag, error) {

	// Create the new tag
	tag := models.Tag{
		Name:     input.Name,
		ParentId: input.ParentId,
		UserId:   userId,
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

	// Set the fields for update (tag's fields)
	tag.Name = input.Name
	tag.ParentId = input.ParentId
	tag.IconId = input.IconId

	// Now, update the existing tag with the modified fields
	if err := repo.db.Model(&tag).Where("id = ?", tag.ID).Updates(tag).Error; err != nil {
		return models.Tag{}, err
	}

	// Return the updated tag
	return tag, nil
}

func (repo *TagRepository) FindWithPagination(page, pageSize int) (*pagination.PaginatedResponse, error) {
	return pagination.GetPaginatedItems(repo.db, models.Tag{}, page, pageSize)
}

func (repo *TagRepository) FindByIdsTags(tagIds []int) ([]models.Tag, error) {
	if len(tagIds) == 0 { // Check for empty slice
		return []models.Tag{}, errors.New("tags cannot be empty")
	}

	var tags []models.Tag // Use a slice for multiple records
	if err := repo.db.Find(&tags, tagIds).Error; err != nil {
		return []models.Tag{}, err
	}

	return tags, nil
}
