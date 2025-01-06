package otp


type OtpService struct {
	repo OtpRepositoryInterface
}

func NewOtpService(repo OtpRepositoryInterface) *OtpService {
	return &OtpService{repo: repo}
}

func (service *OtpService) CreateOtp(data *OtpEntity) (OtpEntity, error) {
	newOtp, err := service.repo.CreateOtp(data)
	if err != nil {
		return OtpEntity{}, err
	}

	return newOtp, nil;
}

func (service *OtpService) GetOtpByEmail(email string) (OtpEntity, error) {
	otp, err := service.repo.GetOtpByEmail(email)

	if err != nil {
		return OtpEntity{}, err
	}

	return otp, nil
}

func (service *OtpService) UpdateOtpByEmail(email, otp string, expireAt int64) (OtpEntity, error) {
	updatedOtp, err := service.repo.UpdateOtpByEmail(email, otp, expireAt)

	if err != nil {
		return OtpEntity{}, err
	}

	return updatedOtp, nil;
}