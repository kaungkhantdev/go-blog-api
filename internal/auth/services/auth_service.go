package services

import (
	"errors"

	otp "go-blog-api/internal/otp/services"
	user "go-blog-api/internal/user/services"
	"go-blog-api/pkg/generate_otp"
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

	hasEmail, _ := auth.otpService.GetOtpByEmail(email)

	// check email it's ald exit or not
	if hasEmail.Email != "" {
		return "", errors.New("email has ald exit")
	}

	// generate otp
	otp, err := generate_otp.GenerateOtp(6)
	if err != nil { 
		return "", errors.New("otp generate error")
	}

	// send otp via email

	return otp, nil
}

func (auth AuthService) VerifyOtpViaEmail() {
	// TODO
	// get email
	// get otp
	// check otp is valid or not
	// if valid then create user, just email
	// add verify at in user's

}

