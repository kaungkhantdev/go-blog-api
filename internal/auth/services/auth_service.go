package services

import (
	"go-blog-api/internal/user/interfaces"
)

type AuthService struct {
	repo *interfaces.UserRepositoryInterface
}

func NewAuthService(repo *interfaces.UserRepositoryInterface) *AuthService {
	return &AuthService{repo: repo}
}

func (auth AuthService) SignUp(data interface{}) {

}

func (auth AuthService) SignIn() {
	// TODO
	

}

func (auth AuthService) GetOtpViaEmail() {
	// TODO
	// get email
	// check email it's ald exit or not
	// generate otp
	// send otp via email

}

func (auth AuthService) VerifyOtpViaEmail() {
	// TODO
	// get email
	// get otp
	// check otp is valid or not
	// if valid then create user, just email
	// add verify at in user's

}

