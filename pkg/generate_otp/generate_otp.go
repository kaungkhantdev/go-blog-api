package generate_otp

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateOtp(length int) (string, error) {

	if (length <= 0) {
		return "", fmt.Errorf("length must be greater than 0")
	}

	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	newOtp := make([]byte, length)
	for i := range newOtp {
		max := big.NewInt(int64(len(charSet)))

		randomIndex, err := rand.Int(rand.Reader, max)

		if err != nil {
			return "", fmt.Errorf("failed to generate secure random number: %v", err)
		}

		newOtp[i] = charSet[randomIndex.Int64()]
	}

	return string(newOtp), nil
	
}