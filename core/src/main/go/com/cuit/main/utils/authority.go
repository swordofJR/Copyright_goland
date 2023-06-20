package utils

import (
	"fmt"
)

func GetAuthoritiesByUsername(username string) ([]string, error) {
	userRepo := GetGlobalUserRepository()
	user, err := userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	authorityRepo := GetGlobalAuthorityRepository()
	authorities, err := authorityRepo.FindByUserId(user.ID)
	if err != nil {
		return nil, err
	}

	var authorityValues []string
	for _, authority := range authorities {
		authorityValues = append(authorityValues, authority.Value)
	}

	return authorityValues, nil
}

func GetAuthoritiesByToken(tokenString string) ([]string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetEnv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	authorities, err := GetAuthoritiesByUsername(username)
	if err != nil {
		return nil, err
	}

	return authorities, nil
}

func HasAuthority(authorities []string, authority string) bool {
	for _, a := range authorities {
		if a == authority {
			return true
		}
	}
	return false
}
