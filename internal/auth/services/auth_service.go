package services

import (
	"errors"

	otp "go-blog-api/internal/otp/services"
	user "go-blog-api/internal/user/services"
)

type AuthService struct {
	otpService *otp.OtpService
	userService *user.UserService
}

func NewAuthService(
	otpService *otp.OtpService,
	userService *user.UserService,
) *AuthService {
	return &AuthService{
		otpService: otpService,
		userService: userService,
	}
}

func (auth AuthService) SignUp(data interface{}) {

}

func (auth AuthService) SignIn() {
	// TODO
	

}

func (auth AuthService) GetOtpViaEmail(email string) (string, error) {
	// TODO
	// get email
	// check email it's ald exit or not
	// generate otp
	// send otp via email

	hasEmail, _ := auth.otpService.GetOtpByEmail(email)

	if hasEmail.Email != "" {
		return "", errors.New("email has ald exit")
	}

	return "123", nil
}

func (auth AuthService) VerifyOtpViaEmail() {
	// TODO
	// get email
	// get otp
	// check otp is valid or not
	// if valid then create user, just email
	// add verify at in user's

}

