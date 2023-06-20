package utils

import (
	"regexp"
)

func IsValidUsername(username string) bool {
	if len(username) < 4 || len(username) > 16 {
		return false
	}
	match, _ := regexp.MatchString("^[A-Za-z0-9]+$", username)
	if !match {
		return false
	}
	return true
}

func IsValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 16 {
		return false
	}
	match, _ := regexp.MatchString("^[A-Za-z0-9]+$", password)
	if !match {
		return false
	}
	return true
}

func IsValidEmail(email string) bool {
	match, _ := regexp.MatchString("^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}$", email)
	if !match {
		return false
	}
	return true
}

func IsValidPhone(phone string) bool {
	match, _ := regexp.MatchString("^[0-9]+$", phone)
	if !match {
		return false
	}
	return true
}
