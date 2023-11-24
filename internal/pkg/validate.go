package pkg

import "regexp"

func ValidatePassword(password string) bool {
	containNums, _ := regexp.Match(`[0123456789]`, []byte(password))
	containUpper, _ := regexp.Match(`[A-Z][a-z]`, []byte(password))
	containSymbols, _ := regexp.Match(`[!@#$%^&*_-]`, []byte(password))

	if (len(password) > 8 && len(password) < 20) && containNums && containUpper && containSymbols {
		return true
	}
	return false
}
