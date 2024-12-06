package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(ctx.Query("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
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

// {
//   "data": [
//     {
//       "id": 1,
//       "name": "Item 1",
//       "description": "Description for Item 1"
//     },
//     {
//       "id": 2,
//       "name": "Item 2",
//       "description": "Description for Item 2"
//     },
//     {
//       "id": 3,
//       "name": "Item 3",
//       "description": "Description for Item 3"
//     }
//   ],
//   "meta": {
//     "pagination": {
//       "total": 50,
//       "count": 3,
//       "per_page": 3,
//       "current_page": 1,
//       "total_pages": 17,
//       "links": {
//         "next": "/api/items?page=2&page_size=3",
//         "prev": ""
//       }
//     }
//   }
// }
