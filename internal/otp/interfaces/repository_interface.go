package interfaces

import "go-blog-api/internal/otp/models"

type OtpRepositoryInterface interface {
	CreateOtp (data *models.Otp) (models.Otp, error)
	GetOtpByEmail (email string) (models.Otp, error)
	UpdateOtpByEmail (email, otp string) (models.Otp, error)
}