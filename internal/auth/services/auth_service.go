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

func (auth AuthService) SignUp (data map[string]string) (string, error) {
	// Todo
	email, hasEmail := data["email"]
	name, hasName := data["name"]
	userName, hasUserName := data["user_name"]
	avatarUrl, hasAvatarUrl := data["avatar_url"]
	bio, hasBio := data["bio"]

	if !hasEmail {
		return "", errors.New("email is required")
	}

	// user exist in user table
	oldUser, _ := auth.userService.FindByEmailUser(email)

	if oldUser.Email == "" {
		return "", errors.New("email invalid")
	}

	user := userModel.User{}

	if hasName {
		user.Name = name
	}

	if hasUserName {
		hasUserNameInDb, _ := auth.userService.FindByUserName(userName)
		if hasUserNameInDb.UserName != "" { 
			return "", errors.New("username is already taken")
		}
		user.UserName = userName;
	}

	if hasAvatarUrl {
		user.AvatarUrl = avatarUrl
	}

	if hasBio {
		user.Bio = bio
	}

	updateUser, err := auth.userService.UpdateUser(oldUser.ID, &user)

	if err != nil {
		return "", errors.New(err.Error())
	}

	token, err := jwt.GenerateJWT(updateUser.ID)

	if err != nil {
		return "", errors.New(err.Error())
	}

	return token, nil
}

func (auth AuthService) SignIn(email string) (string, error) {
	hasEmail, _ := auth.otpService.GetOtpByEmail(email)

	// check email it's ald exit or not
	if hasEmail.Email == "" {
		return "", errors.New("email invalid")
	}

	// generate otp
	otp, err := generateOtp.GenerateOtp(6)
	if err != nil { 
		return "", errors.New("otp generate error")
	}

	// create new email with otp
	expireAt := time.Now().Add( 1 * time.Minute).Unix();

	_, err = auth.otpService.UpdateOtpByEmail(email, otp, expireAt)
	if err != nil {
		return "", errors.New("otp create error")
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

	// create new email with otp
	otpData := otpModel.Otp{
		Email: email,
		Otp: otp,
		ExpiresAt: time.Now().Add( 1 * time.Minute).Unix(),
	}
	_, err = auth.otpService.CreateOtp(&otpData)
	if err != nil {
		return "", errors.New("otp create error")
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

func (auth AuthService) VerifyOtpViaEmail(data map[string]string) (map[string]string, error) {

	email, hasEmail := data["email"]
	otp, hasOtp		:= data["otp"]

	resObj := map[string]string{};

	if !hasEmail || !hasOtp {
		return resObj, errors.New("email or otp is missing")
	}

	// TODO
	// get email
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

	if oldUser.Email != "" {
		return resObj, errors.New("user already exist")
	}

	// if valid then create user, just email
	userName := strings.Split(email, "@")
	user := userModel.User{ 
		Email: email,
		UserName: userName[0],
		VerifyAt: time.Now().Unix(),
	}

	newUser, err := auth.userService.CreateUser(&user)

	if err != nil {
		return map[string]string{}, errors.New("failed to create user")
	}

	result :=  map[string]string{
		"email": newUser.Email,
		"user_name": newUser.UserName,
	}


	return result, nil

}


func (auth AuthService) VerifyOtpViaEmailSignIn(data map[string]string) (string, error) {

	email, hasEmail := data["email"]
	otp, hasOtp		:= data["otp"]

	if !hasEmail || !hasOtp {
		return "", errors.New("email or otp is missing")
	}

	// TODO
	// get email
	storedOtp, err := auth.otpService.GetOtpByEmail(email)

	if err != nil {
		return "", errors.New("error fetching OTP record")
	}

	// check email it's ald exit or not
	if storedOtp.Email == "" {
		return "", errors.New("email has not exit")
	}

	// check otp expires
	if storedOtp.ExpiresAt < time.Now().Unix() {
		return "", errors.New("otp has expired")
	}

	// check otp is valid or not
	if storedOtp.Otp != otp {
        return "", errors.New("invalid OTP")
    }

	// user exist in user table
	oldUser, _ := auth.userService.FindByEmailUser(email)

	if oldUser.Email != "" {
		return "", errors.New("user already exist")
	}

	if err != nil {
		return "", errors.New("failed to create user")
	}

	token, err := jwt.GenerateJWT(oldUser.ID)

	if err != nil {
		return "", errors.New(err.Error())
	}

	return token, nil

}

