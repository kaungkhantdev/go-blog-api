package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	UserName  string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	AvatarUrl string    `gorm:"type:varchar(255)" json:"avatar_url"`
	Bio       string    `gorm:"type:text" json:"bio"`
	DelFlag   string    `gorm:"type:char(1);default:'N'" json:"del_flag"` // Assuming 'N' means not deleted
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	OtpId     string    `gorm:"type:varchar(50)" json:"otp_id"`
}
