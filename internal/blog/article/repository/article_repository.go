package repository

import (
	"go-blog-api/internal/blog/article/interfaces"
	"go-blog-api/internal/blog/article/models"
	"go-blog-api/pkg/pagination"

	tagModel "go-blog-api/internal/blog/tag/models"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) interfaces.ArticleRepositoryInterfaces {
	return &ArticleRepository{db: db}
}

func (repo *ArticleRepository) CreateArticle(data *models.Article) (models.Article, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return models.Article{}, err
	}
	return *data, nil
}

func (repo *ArticleRepository) UpdateArticle(id int, data *models.Article) (models.Article, error) {
	var article models.Article
	if err := repo.db.Preload("Tag").First(&article, id).Error; err != nil {
		return models.Article{}, err
	}
	if err := repo.db.Model(&article).Updates(data).Error; err != nil {
		return models.Article{}, err
	}

	if data.Tag == nil {
		data.Tag = []tagModel.Tag{}
	}
	repo.db.Model(&article).Association("Tag").Replace(data.Tag)
	return article, nil
}

func (repo *ArticleRepository) FindOneById(id int) (models.Article, error) {
	var article models.Article
	if err := repo.db.First(&article, id).Error; err != nil {
		return models.Article{}, err
	}
	return article, nil
}

func (repo *ArticleRepository) FindWithPagination(page, pageSize int) (*pagination.PaginatedResponse, error) {
	return pagination.GetPaginatedItems(repo.db, models.Article{}, page, pageSize)
}
