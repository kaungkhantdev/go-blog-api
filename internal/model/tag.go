package model

import "time"

type Tag struct {
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	IconUrl   string    `json:"icon_url"`
	ParentId  string    `json:"parent_id"`
	UserId    string    `json:"user_id"`
	DelFlag   string    `json:"del_flag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
