package reaction

import (
	"go-blog-api/internal/blog/article"
	"go-blog-api/internal/blog/reaction_type"
	"go-blog-api/internal/core/user"
	"time"
)

type ReactionEntity struct {
	ID             int                               `gorm:"primaryKey;autoIncrement" json:"id"`
	ArticleId      int                               `json:"article_id"`
	Article        *article.ArticleEntity            `gorm:"foreignKey:ArticleId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"article"`
	UserId         int                               `json:"user_id"`
	User           *user.UserEntity                  `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	ReactionTypeId int                               `json:"reaction_type_id"`
	ReactionType   *reaction_type.ReactionTypeEntity `gorm:"foreignKey:ReactionTypeId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reaction_type"`
	DelFlag        bool                              `gorm:"type:tinyint(1);default:0" json:"del_flag"`
	CreatedAt      time.Time                         `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time                         `gorm:"autoUpdateTime" json:"updated_at"`
}
