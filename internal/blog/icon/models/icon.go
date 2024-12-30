package models

import "time"

type Icon struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Icon      string    `gorm:"type:varchar(255)" json:"icon"`
	Url       string    `gorm:"type:text" json:"url"`
	DelFlag   bool      `gorm:"type:tinyint(1);default:0" json:"del_flag"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
