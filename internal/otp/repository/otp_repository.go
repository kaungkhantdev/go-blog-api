package repository

import (
	"go-blog-api/internal/otp/models"
	"go-blog-api/internal/otp/interfaces"

	"gorm.io/gorm"
)

type OtpRepository struct {
	db *gorm.DB
}

func NewOtpRepository(db *gorm.DB) interfaces.OtpRepositoryInterface {
	return &OtpRepository{db: db}
}

func (repo *OtpRepository) CreateOtp(data *models.Otp) (models.Otp, error) {
	
	if err := repo.db.Create(data).Error; err != nil { 
		return models.Otp{}, err
	}
	return *data, nil
}

func (repo *OtpRepository) GetOtpByEmail(email string) (models.Otp, error) {

	var otpData models.Otp
	if err := repo.db.Where("email = ?", email).First(&otpData).Error; err != nil {
		return models.Otp{}, err
	}
	return otpData, nil
}

func (repo *OtpRepository) UpdateOtpByEmail(email, otp string, expireAt int64) (models.Otp, error) {
	var otpData models.Otp
	if err := repo.db.Where("email = ?", email).First(&otpData).Error; err != nil {
		return models.Otp{}, err
	}
	otpData.Otp = otp
	otpData.ExpiresAt = expireAt
	if err := repo.db.Save(&otpData).Error; err != nil {
        return models.Otp{}, err
    }

	return otpData, nil
}