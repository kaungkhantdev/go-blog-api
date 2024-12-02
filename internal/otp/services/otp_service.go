package services

import (
	"go-blog-api/internal/otp/models"
	"go-blog-api/internal/otp/interfaces"
)

type OtpService struct {
	repo interfaces.OtpRepositoryInterface
}

func NewOtpService(repo interfaces.OtpRepositoryInterface) *OtpService {
	return &OtpService{repo: repo}
}

func (service *OtpService) CreateOtp(data *models.Otp) (models.Otp, error) {
	newOtp, err := service.repo.CreateOtp(data)
	if err != nil {
		return models.Otp{}, err
	}

	return newOtp, nil;
}

func (service *OtpService) GetOtpByEmail(email string) (models.Otp, error) {
	otp, err := service.repo.GetOtpByEmail(email)

	if err != nil {
		return models.Otp{}, err
	}

	return otp, nil
}

func (service *OtpService) UpdateOtpByEmail(email, otp string, expireAt int64) (models.Otp, error) {
	updatedOtp, err := service.repo.UpdateOtpByEmail(email, otp, expireAt)

	if err != nil {
		return models.Otp{}, err
	}

	return updatedOtp, nil;
}