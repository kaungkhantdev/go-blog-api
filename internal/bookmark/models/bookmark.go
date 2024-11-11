package models

import "time"

type Bookmark struct {
	UserId    string    `json:"user_id"`
	ArticleId string    `json:"article_id"`
	DelFlag   string    `json:"del_flag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
