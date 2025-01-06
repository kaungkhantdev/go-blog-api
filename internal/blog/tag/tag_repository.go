package tag

import (
	"errors"
	"go-blog-api/pkg/pagination"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepositoryInterface {
	return &TagRepository{db: db}
}

func (repo *TagRepository) FindByIdTag(id int) (TagEntity, error) {
	var tag TagEntity
	if err := repo.db.First(&tag, id).Error; err != nil {
		return TagEntity{}, err
	}
	return tag, nil
}

func (repo *TagRepository) CreateTag(input TagCreateRequest, userId int) (TagEntity, error) {

	// Create the new tag
	tag := TagEntity{
		Name:     input.Name,
		ParentId: input.ParentId,
		UserId:   userId,
		IconId:   input.IconId,
	}

	// Insert the tag into the database
	if err := repo.db.Create(&tag).Error; err != nil {
		return TagEntity{}, err
	}

	return tag, nil
}

func (repo *TagRepository) UpdateTag(id int, input TagUpdateRequest) (TagEntity, error) {
	// First, find the existing tag by ID
	var tag TagEntity
	if err := repo.db.Preload("Parent").First(&tag, id).Error; err != nil {
		return TagEntity{}, err
	}

	// Set the fields for update (tag's fields)
	tag.Name = input.Name
	tag.ParentId = input.ParentId
	tag.IconId = input.IconId

	// Now, update the existing tag with the modified fields
	if err := repo.db.Model(&tag).Where("id = ?", tag.ID).Updates(tag).Error; err != nil {
		return TagEntity{}, err
	}

	// Return the updated tag
	return tag, nil
}

func (repo *TagRepository) FindWithPagination(page, pageSize int) (*pagination.PaginatedResponse, error) {
	return pagination.GetPaginatedItems(repo.db, TagEntity{}, page, pageSize)
}

func (repo *TagRepository) FindByIdsTags(tagIds []int) ([]TagEntity, error) {
	if len(tagIds) == 0 { // Check for empty slice
		return []TagEntity{}, errors.New("tags cannot be empty")
	}

	var tags []TagEntity // Use a slice for multiple records
	if err := repo.db.Find(&tags, tagIds).Error; err != nil {
		return []TagEntity{}, err
	}

	return tags, nil
}
