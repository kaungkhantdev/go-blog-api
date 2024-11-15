package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId int	`json:"user_id"`
	jwt.RegisteredClaims
}

func GetJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func GetJWTExpiration() time.Duration {
	expirationMinutes := os.Getenv("JWT_EXPIRATION_MINUTES")
	expirationTime, err := time.ParseDuration(expirationMinutes + "m")

	if err != nil {
		fmt.Println("Error parsing JWT expiration time")
		return 15 * time.Minute // default to 15 minutes
	}

	return expirationTime
}

func GenerateJWT(userId int) (string, error) {
	expirationTime := time.Now().Add(GetJWTExpiration())

	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetJWTSecret())
}

func VerifyJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid signature")
		}
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	
	return claims, nil
}