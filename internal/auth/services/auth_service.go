package services

import (
	"errors"
	"log"

	otp "go-blog-api/internal/otp/services"
	user "go-blog-api/internal/user/services"
	mail "go-blog-api/pkg/mail"

	generateOtp "go-blog-api/pkg/generate_otp"
)

type AuthService struct {
	otpService *otp.OtpService
	userService *user.UserService
	mailService *mail.EmailService
}

func NewAuthService(
	otpService *otp.OtpService,
	userService *user.UserService,
	mailService *mail.EmailService,
) *AuthService {
	return &AuthService{
		otpService: otpService,
		userService: userService,
		mailService: mailService,
	}
}

func (auth AuthService) SignUp(data interface{}) {

}

func (auth AuthService) SignIn() {
	// TODO
	
}

func (auth AuthService) GetOtpViaEmail(email string) (string, error) {
	hasEmail, _ := auth.otpService.GetOtpByEmail(email)

	// check email it's ald exit or not
	if hasEmail.Email != "" {
		return "", errors.New("email has ald exit")
	}

	// generate otp
	otp, err := generateOtp.GenerateOtp(6)
	if err != nil { 
		return "", errors.New("otp generate error")
	}

	data := map[string]string{
		"OTP": otp,
	}

	// send otp via email
	err = auth.mailService.SendEmail(
		[]string{email},
		"Testing",
		"email_template.html",
		data,
		nil,
		nil,
		nil,
	)

	if err != nil {
		log.Printf("Failed to send OTP via email: %v", err)
		return "", errors.New("failed to send OTP")
	}
	
	return "Otp code has just sent.", nil
}

func (auth AuthService) VerifyOtpViaEmail() {
	// TODO
	// get email
	// get otp
	// check otp is valid or not
	// if valid then create user, just email
	// add verify at in user's

}

