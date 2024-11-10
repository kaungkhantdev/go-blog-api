package models

import "time"

type Otp struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Otp       string    `gorm:"type:varchar(50)" json:"otp_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
