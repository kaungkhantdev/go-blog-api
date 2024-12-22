package models

import (
	tagModel "go-blog-api/internal/blog/tag/models"
	userModel "go-blog-api/internal/core/user/models"
	"time"
)

type Article struct {
	ID        int             `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId    int             `json:"user_id"`
	User      *userModel.User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Title     string          `gorm:"type:varchar(255)" json:"title"`
	Content   string          `gorm:"type:longtext" json:"content"`
	Tag       []tagModel.Tag  `gorm:"many2many:article_tags;" json:"tags"`
	IsFree    bool            `gorm:"type:tinyint(1);default:0" json:"is_free"`
	DelFlag   bool            `gorm:"type:tinyint(1);default:0" json:"del_flag"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}
