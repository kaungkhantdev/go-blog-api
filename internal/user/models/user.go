package models

import "time"

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UserName  string    `json:"username"`
	AvatarUrl string    `json:"avatar_url"`
	Bio       string    `json:"bio"`
	DelFlag   string    `json:"del_flag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	OtpId     string    `json:"otp_id"`
}
