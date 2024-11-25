package services

import (
	"errors"
	"log"
	"strings"
	"time"

	otpModel "go-blog-api/internal/otp/models"

	otp "go-blog-api/internal/otp/services"

	userModel "go-blog-api/internal/user/models"

	user "go-blog-api/internal/user/services"

	jwt "go-blog-api/pkg/jwt"
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

// Helper Methods

func (auth AuthService) defineExpire() int64 {
	return time.Now().Add( 1 * time.Minute).Unix()
}


func (auth AuthService) authBuildRes(user userModel.User, token string) (map[string]string, error) {
	return map[string]string{
		"email": user.Email,
		"user_name": user.UserName,
		"token": token,
	}, nil
}

func (auth AuthService) authCreateUser(email string) (userModel.User, error) {
	user := userModel.User{ 
		Email: email,
		UserName: strings.Split(email, "@")[0],
		VerifyAt: time.Now().Unix(),
	}

	newUser, err := auth.userService.CreateUser(&user)
	if err != nil {
		return userModel.User{}, errors.New("failed to create user")
	}

	return newUser, nil
}

func (auth AuthService) authCreateOtp(email, otp string) (string, error) {
	otpData := otpModel.Otp{
		Email: email,
		Otp: otp,
		ExpiresAt: auth.defineExpire(),
	}
	_, err := auth.otpService.CreateOtp(&otpData)
	if err != nil {
		return "", errors.New("otp create error")
	}	
	
	return "", nil
}

func (auth AuthService) authUpdateOtp(email, otp string) (string, error) {
	_, err := auth.otpService.UpdateOtpByEmail(email, otp, auth.defineExpire())
	if err != nil {
		return "", errors.New("otp update error")
	}

	return "", nil
}

func (auth AuthService) authSendOtpEmail(otp, email string) (string, error) {
	data := map[string]string{
		"OTP": otp,
	}

	// send otp via email
	err := auth.mailService.SendEmail(
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

func (auth AuthService) authPrepareUserData(data map[string]string) userModel.User {
	return userModel.User{
		Name:      data["name"],
		UserName:  data["user_name"],
		AvatarUrl: data["avatar_url"],
		Bio:       data["bio"],
	}
}

// Methods

func (auth AuthService) SignUp (data map[string]string) (map[string]string, error) {
	email, hasEmail := data["email"]
	resObj := map[string]string{}

	if !hasEmail {
		return resObj, errors.New("email is required")
	}

	// user exist in user table
	oldUser, _ := auth.userService.FindByEmailUser(email)
	if oldUser.Email == "" {
		return resObj, errors.New("email invalid")
	}

	user := auth.authPrepareUserData(data)
	if data["user_name"] != "" {
		userNameTaken, _ := auth.userService.FindByUserName(data["user_name"])
		if userNameTaken.UserName != "" {
			return resObj, errors.New("username is already taken")
		}
	}

	updateUser, err := auth.userService.UpdateUser(oldUser.ID, &user)
	if err != nil {
		return resObj, errors.New(err.Error())
	}

	token, err := jwt.GenerateJWT(updateUser.ID)
	if err != nil {
		return resObj, errors.New(err.Error())
	}

	return auth.authBuildRes(updateUser, token)
}

func (auth AuthService) SignIn(email string) (string, error) {
	// user exist in user table
	oldUser, _ := auth.userService.FindByEmailUser(email)

	if oldUser.Email == "" {
		return "", errors.New("user need to create")
	}

	return auth.GetOtpViaEmail(email)
}

func (auth AuthService) GetOtpViaEmail(email string) (string, error) {
	hasEmail, _ := auth.otpService.GetOtpByEmail(email)

	// generate otp
	otp, err := generateOtp.GenerateOtp(6)
	if err != nil { 
		return "", errors.New("otp generate error")
	}
	
	// check email it's ald exit or not
	if hasEmail.Email == "" {
		// create new email with otp
		auth.authCreateOtp(email, otp)		
	} else {
		// update otp
		auth.authUpdateOtp(email, otp)
	}

	return auth.authSendOtpEmail(otp, email)
}

func (auth AuthService) VerifyOtpViaEmail(data map[string]string) (map[string]string, error) {

	email, hasEmail := data["email"]
	otp, hasOtp		:= data["otp"]
	resObj := map[string]string{}

	if !hasEmail || !hasOtp {
		return resObj, errors.New("something is missing")
	}


	storedOtp, err := auth.otpService.GetOtpByEmail(email)
	if err != nil {
		return resObj, errors.New("error fetching OTP record")
	}

	// check email it's ald exit or not
	if storedOtp.Email == "" {
		return resObj, errors.New("email has not exit")
	}

	// check otp expires
	if storedOtp.ExpiresAt < time.Now().Unix() {
		return resObj, errors.New("otp has expired")
	}

	// check otp is valid or not
	if storedOtp.Otp != otp {
        return resObj, errors.New("invalid OTP")
    }

	// user exist in user table
	oldUser, _ := auth.userService.FindByEmailUser(email)
	if oldUser.Email == "" {
		newUser, err := auth.authCreateUser(email)
		if err != nil {
			return map[string]string{}, errors.New("failed to create user")
		}
	
		return auth.authBuildRes(newUser, "")
	} else {
		token, err := jwt.GenerateJWT(oldUser.ID)
		if err != nil {
			return resObj, errors.New(err.Error())
		}
	
		return auth.authBuildRes(oldUser, token)
	}

}
