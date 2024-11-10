package services

import "go-blog-api/internal/user/interfaces"

type AuthService struct {
	repo *interfaces.UserRepositoryInterface
}

func NewAuthService(repo *interfaces.UserRepositoryInterface) *AuthService {
	return &AuthService{repo: repo}
}

func (auth AuthService) SignUp() {
	
}