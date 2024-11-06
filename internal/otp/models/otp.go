package models

type Otp struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}
