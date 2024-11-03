package model

import "time"

type ReactionType struct {
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	AvatarUrl string    `json:"avatar_url"`
	DelFlag   string    `json:"del_flag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
