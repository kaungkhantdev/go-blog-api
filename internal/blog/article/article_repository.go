package article

import (
	"go-blog-api/internal/blog/tag"
	"go-blog-api/pkg/pagination"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepositoryInterfaces {
	return &ArticleRepository{db: db}
}

func (repo *ArticleRepository) CreateArticle(data *ArticleEntity) (ArticleEntity, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return ArticleEntity{}, err
	}
	return *data, nil
}

func (repo *ArticleRepository) UpdateArticle(id int, data *ArticleEntity) (ArticleEntity, error) {
	var article ArticleEntity
	if err := repo.db.Preload("Tag").First(&article, id).Error; err != nil {
		return ArticleEntity{}, err
	}
	if err := repo.db.Model(&article).Updates(data).Error; err != nil {
		return ArticleEntity{}, err
	}

	if data.Tag == nil {
		data.Tag = []tag.TagEntity{}
	}
	repo.db.Model(&article).Association("Tag").Replace(data.Tag)
	return article, nil
}

func (repo *ArticleRepository) FindOneById(id int) (ArticleEntity, error) {
	var article ArticleEntity
	if err := repo.db.First(&article, id).Error; err != nil {
		return ArticleEntity{}, err
	}
	return article, nil
}

func (repo *ArticleRepository) FindWithPagination(page, pageSize int) (*pagination.PaginatedResponse, error) {
	return pagination.GetPaginatedItems(repo.db, ArticleEntity{}, page, pageSize)
}
