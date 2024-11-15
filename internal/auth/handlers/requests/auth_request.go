package requests

type AuthSignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthSignUpRequest struct {
	Name      string `json:"name" validate:"required,min=1"`
	Email     string `json:"email" validate:"required,email"`
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