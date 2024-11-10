package requests

type AuthSignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthSignUpRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	UserName  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
	Bio       string `json:"bio"`
	OtpId     string `json:"otp_id"`
}
