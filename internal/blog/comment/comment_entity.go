package comment

import (
	"go-blog-api/internal/blog/article"
	"go-blog-api/internal/core/user"
	"time"
)

type CommentEntity struct {
	ID        int                    `gorm:"primaryKey;autoIncrement" json:"id"`
	ParentId  *int                   `gorm:"type:int" json:"parent_id"`
	Parent    *CommentEntity         `gorm:"foreignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parent"`
	Content   string                 `json:"content"`
	ArticleId int                    `json:"article_id"`
	Article   *article.ArticleEntity `gorm:"foreignKey:ArticleId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"article"`
	UserId    int                    `json:"user_id"`
	User      *user.UserEntity       `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	DelFlag   bool                   `gorm:"type:tinyint(1);default:0" json:"del_flag"`
	CreatedAt time.Time              `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time              `gorm:"autoUpdateTime" json:"updated_at"`
}
