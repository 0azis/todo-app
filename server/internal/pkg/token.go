package pkg

import (
	"github.com/golang-jwt/jwt/v5"
)

// KEY Слово-секрет, нужен для расшифровки токена
var KEY = []byte("secret")

// CreateAccessToken Метод создания access токена
func CreateAccessToken(userId int) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// Создаем payload структуру
		"userId": userId,
	}).SignedString(KEY)
	return token, err
}

// GetIdentity Расшифровываем токен и получаем из него данные (identity)
func GetIdentity(token string) (int, error) {
	identity, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return KEY, nil
	})

	if err != nil {
		return 0, err
	}

	payload := identity.Claims.(jwt.MapClaims)
	userId := int(payload["userId"].(float64))

	return userId, nil

}
