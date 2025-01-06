package article

import (
	"go-blog-api/internal/blog/tag"
	"go-blog-api/internal/core/user"
	"time"
)

type ArticleEntity struct {
	ID        int              `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId    int              `json:"user_id"`
	User      *user.UserEntity `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Title     string           `gorm:"type:varchar(255)" json:"title"`
	Content   string           `gorm:"type:longtext" json:"content"`
	Tag       []tag.TagEntity  `gorm:"many2many:article_tags;" json:"tags"`
	IsFree    bool             `gorm:"type:tinyint(1);default:0" json:"is_free"`
	DelFlag   bool             `gorm:"type:tinyint(1);default:0" json:"del_flag"`
	CreatedAt time.Time        `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time        `gorm:"autoUpdateTime" json:"updated_at"`
}
