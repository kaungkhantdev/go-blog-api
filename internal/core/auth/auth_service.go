package auth

import (
	"errors"
	"log"
	"strings"
	"time"

	"go-blog-api/internal/core/otp"
	"go-blog-api/internal/core/user"

	"go-blog-api/pkg/jwt"
	jwtService "go-blog-api/pkg/jwt"
	mailService "go-blog-api/pkg/mail"

	generateOtp "go-blog-api/pkg/generate_otp"
)

type AuthService struct {
	otpService  *otp.OtpService
	userService *user.UserService
	mailService *mailService.EmailService
}

func NewAuthService(
	otpService *otp.OtpService,
	userService *user.UserService,
	mailService *mailService.EmailService,
) *AuthService {
	return &AuthService{
		otpService:  otpService,
		userService: userService,
		mailService: mailService,
	}
}

// Helper Methods

func (auth AuthService) defineExpire() int64 {
	return time.Now().Add(1 * time.Minute).Unix()
}

func (auth AuthService) authBuildRes(user user.UserEntity, showToken bool, token map[string]string) (map[string]string, error) {
	response := map[string]string{
		"email":     user.Email,
		"user_name": user.UserName,
	}

	if showToken {
		response["refresh_token"] = token["refresh_token"]
		response["access_token"] = token["access_token"]
	}

	return response, nil
}

func (auth AuthService) authCreateUser(email string) (user.UserEntity, error) {
	newUser := user.UserEntity{
		Email:    email,
		UserName: strings.Split(email, "@")[0],
		VerifyAt: time.Now().Unix(),
	}

	createdUser, err := auth.userService.CreateUser(&newUser)
	if err != nil {
		return user.UserEntity{}, errors.New("failed to create user")
	}

	return createdUser, nil
}

func (auth AuthService) authCreateOtp(email, otpInput string) (string, error) {
	otpData := otp.OtpEntity{
		Email:     email,
		Otp:       otpInput,
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

func (auth AuthService) authPrepareUserData(data map[string]string) user.UserEntity {
	return user.UserEntity{
		Name:     data["name"],
		UserName: data["user_name"],
		Avatar:   data["avatar"],
		Bio:      data["bio"],
	}
}

func (auth AuthService) tokens(userId int) (map[string]string, error) {
	accessToken, err := jwtService.GenerateJWT(userId, jwtService.GetJWTSecret(), jwtService.GetJWTExpireMinutes())
	refreshToken, err := jwtService.GenerateJWT(userId, jwtService.GetJWTAccessTokenSecret(), jwtService.GetJWTAccessTokenExpireMinutes())
	if err != nil {
		return map[string]string{}, errors.New(err.Error())
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

// Methods

func (auth AuthService) SignUp(data map[string]string) (map[string]string, error) {
	email, hasEmail := data["email"]
	resObj := map[string]string{}

	if !hasEmail {
		return resObj, errors.New("email is required")
	}

	// has email in otp
	optUser, _ := auth.otpService.GetOtpByEmail(email)
	if optUser.Email == "" {
		return resObj, errors.New("please, verify email first.")
	}

	// user exist in user table
	oldUser, _ := auth.userService.FindByEmailUser(email)
	if oldUser.Email != "" {
		return resObj, errors.New("email already taken")
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

	tokens, err := auth.tokens(updateUser.ID)
	if err != nil {
		return resObj, errors.New(err.Error())
	}

	return auth.authBuildRes(updateUser, true, tokens)
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
	otp, hasOtp := data["otp"]
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

		return auth.authBuildRes(newUser, false, resObj)
	} else {
		tokens, err := auth.tokens(oldUser.ID)
		if err != nil {
			return resObj, errors.New(err.Error())
		}

		return auth.authBuildRes(oldUser, true, tokens)
	}

}

func (auth AuthService) VerifyRefreshToken(token string) (map[string]string, error) {
	claims, err := jwt.VerifyJWT(token, jwt.GetJWTAccessTokenSecret())
	if err != nil {
		return map[string]string{}, errors.New(err.Error())
	}

	tokens, err := auth.tokens(claims.UserId)
	if err != nil {
		return map[string]string{}, errors.New(err.Error())
	}

	return tokens, nil
}
