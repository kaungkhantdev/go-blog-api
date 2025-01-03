package models

import (
	iconModel "go-blog-api/internal/blog/icon/models"
	userModel "go-blog-api/internal/core/user/models"
	"time"
)

type Tag struct {
	ID        int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string          `gorm:"type:varchar(255);uniqueIndex;not null" json:"name"`
	IconId    int             `json:"icon_id"`
	Icon      *iconModel.Icon `gorm:"foreignKey:IconId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"icon"`
	ParentId  *int            `gorm:"type:int" json:"parent_id"`
	Parent    *Tag            `gorm:"foreignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parent"`
	UserId    int             `json:"user_id"`
	User      *userModel.User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	DelFlag   bool            `gorm:"type:tinyint(1);default:0" json:"del_flag"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}
