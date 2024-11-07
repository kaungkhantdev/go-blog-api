package models

import "time"

type ArticleTag struct {
	ArticleId string    `json:"article_id"`
	TagId     string    `json:"tag_id"`
	DelFlag   string    `json:"del_flag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
