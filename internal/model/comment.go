package model

import "time"

type Comment struct {
	ArticleId string    `json:"article_id"`
	UserId    string    `json:"user_id"`
	ParentId  string    `json:"parent_id"`
	Content   string    `json:"content"`
	DelFlag   string    `json:"del_flag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
