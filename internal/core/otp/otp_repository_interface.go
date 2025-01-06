package otp


type OtpRepositoryInterface interface {
	CreateOtp (data *OtpEntity) (OtpEntity, error)
	GetOtpByEmail (email string) (OtpEntity, error)
	UpdateOtpByEmail (email, otp string, expireAt int64) (OtpEntity, error)
}