package pagination

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

type PaginationLinks struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type PaginationMeta struct {
	Total       int64           `json:"total"`
	Count       int             `json:"count"`
	PerPage     int             `json:"per_page"`
	CurrentPage int             `json:"current_page"`
	TotalPages  int             `json:"total_pages"`
	Links       PaginationLinks `json:"links"`
}

type PaginatedResponse struct {
	Data interface{}    `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
func GetPaginatedItems[T any](db *gorm.DB, model T, page, pageSize int) (*PaginatedResponse, error) {
	var items []T
	var totalCount int64

	// Apply pagination
	dbModel := db.Model(&model)
	dbModel = Paginate(page, pageSize)(dbModel)

	// Fetch the items with pagination
	if err := dbModel.Find(&items).Error; err != nil {
		return nil, err
	}

	// Get the total count of records
	db.Model(&model).Count(&totalCount)

	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	// Prepare pagination metadata
	paginationMeta := PaginationMeta{
		Total:       totalCount,
		Count:       len(items),
		PerPage:     pageSize,
		CurrentPage: page,
		TotalPages:  totalPages,
		Links: PaginationLinks{
			Next: fmt.Sprintf("?page=%d&page_size=%d", page+1, pageSize),
			Prev: fmt.Sprintf("?page=%d&page_size=%d", page-1, pageSize),
		},
	}

	// Prepare the final response
	response := &PaginatedResponse{
		Data: items,
		Meta: paginationMeta,
	}

	return response, nil
}
