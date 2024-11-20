package requests


type AuthSignUpRequest struct {
	Name      string `json:"name" validate:"required,min=1"`
	Email     string `json:"email" validate:"required,email"`
	UserName  string `json:"name" validate:"required,min=1"`
	AvatarUrl string `json:"avatar_url"`
	Bio       string `json:"bio"`
}

type AuthOtpRequest struct {
	Email    string `json:"email" validate:"required,email"`
}

type AuthVerifyOtpRequest struct {
	Email   string `json:"email" validate:"required,email"`
	Otp     string `json:"otp" validate:"required"`
}