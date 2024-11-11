package models

import "time"

type Article struct {
	UserId    string    `json:"user_id"`
	IsFree    string    `json:"is_free"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	DelFlag   string    `json:"del_flag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
