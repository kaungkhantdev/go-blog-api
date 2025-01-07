package otp

import (
	"gorm.io/gorm"
)

type OtpRepository struct {
	db *gorm.DB
}

func NewOtpRepository(db *gorm.DB) OtpRepositoryInterface {
	return &OtpRepository{db: db}
}

func (repo *OtpRepository) CreateOtp(data *OtpEntity) (OtpEntity, error) {

	if err := repo.db.Create(data).Error; err != nil {
		return OtpEntity{}, err
	}
	return *data, nil
}

func (repo *OtpRepository) GetOtpByEmail(email string) (OtpEntity, error) {

	var otpData OtpEntity
	if err := repo.db.Where("email = ?", email).First(&otpData).Error; err != nil {
		return OtpEntity{}, err
	}
	return otpData, nil
}

func (repo *OtpRepository) UpdateOtpByEmail(email, otp string, expireAt int64) (OtpEntity, error) {
	var otpData OtpEntity
	if err := repo.db.Where("email = ?", email).First(&otpData).Error; err != nil {
		return OtpEntity{}, err
	}
	otpData.Otp = otp
	otpData.ExpiresAt = expireAt
	if err := repo.db.Save(&otpData).Error; err != nil {
		return OtpEntity{}, err
	}

	return otpData, nil
}
