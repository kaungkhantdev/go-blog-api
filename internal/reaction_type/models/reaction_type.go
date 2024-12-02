package models

import (
	iconModel "go-blog-api/internal/icon/models"
	"time"
)

type ReactionEnum string

const (
	Love   ReactionEnum = "love"
	Haha   ReactionEnum = "haha"
	Like   ReactionEnum = "like"
	Unlike ReactionEnum = "unlike"
)

type ReactionType struct {
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Type      ReactionEnum   `gorm:"type:enum('love','haha','like','unlike');not null" json:"type"`
	IconId    int            `json:"icon_id"`
	Icon      iconModel.Icon `gorm:"foreignKey:IconId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"icon"`
	DelFlag   bool           `gorm:"type:tinyint(1);default:0" json:"del_flag"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}
