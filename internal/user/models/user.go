package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	UserName  string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"username"`
	AvatarUrl string    `gorm:"type:varchar(255)" json:"avatar_url"`
	Bio       string    `gorm:"type:text" json:"bio"`
	DelFlag   bool      `gorm:"type:tinyint(1);default:0" json:"del_flag"`
	VerifyAt  time.Time `json:"verify_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
