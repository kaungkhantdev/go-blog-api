package tag

import (
	"go-blog-api/internal/modules/blog/icon"
	"go-blog-api/internal/modules/core/user"
	"time"
)

type TagEntity struct {
	ID        int              `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string           `gorm:"type:varchar(255);uniqueIndex;not null" json:"name"`
	IconId    int              `json:"icon_id"`
	Icon      *icon.IconEntity `gorm:"foreignKey:IconId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"icon"`
	ParentId  *int             `gorm:"type:int" json:"parent_id"`
	Parent    *TagEntity       `gorm:"foreignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parent"`
	UserId    int              `json:"user_id"`
	User      *user.UserEntity `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	DelFlag   bool             `gorm:"type:tinyint(1);default:0" json:"del_flag"`
	CreatedAt time.Time        `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time        `gorm:"autoUpdateTime" json:"updated_at"`
}
