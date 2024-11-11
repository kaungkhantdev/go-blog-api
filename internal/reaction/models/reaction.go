package models

import "time"

type Reaction struct {
	ArticleId      string    `json:"article_id"`
	UserId         string    `json:"user_id"`
	ReactionTypeId string    `json:"reaction_type_id"`
	DelFlag        string    `json:"del_flag"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
