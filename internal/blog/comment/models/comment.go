package models

import (
	articleModel "go-blog-api/internal/blog/article/models"
	userModel "go-blog-api/internal/core/user/models"
	"time"
)

type Comment struct {
	ID        int                   `gorm:"primaryKey;autoIncrement" json:"id"`
	ParentId  *int                  `gorm:"type:int" json:"parent_id"`
	Parent    *Comment              `gorm:"foreignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parent"`
	Content   string                `json:"content"`
	ArticleId int                   `json:"article_id"`
	Article   *articleModel.Article `gorm:"foreignKey:ArticleId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"article"`
	UserId    int                   `json:"user_id"`
	User      *userModel.User       `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	DelFlag   bool                  `gorm:"type:tinyint(1);default:0" json:"del_flag"`
	CreatedAt time.Time             `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time             `gorm:"autoUpdateTime" json:"updated_at"`
}
