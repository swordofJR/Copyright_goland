package utils

import (
	"time"
)

func TimeAfter(d time.Duration) int64 {
	return time.Now().Add(d).Unix()
}

func GetGlobalJwtSecret() []byte {
	return []byte(GetEnv("JWT_SECRET"))
}

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      TimeAfter(time.Hour * 24),
	})

	signedToken, err := token.SignedString(GetGlobalJwtSecret())
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return GetGlobalJwtSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}
