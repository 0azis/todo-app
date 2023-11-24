package models

import (
	"fmt"
	"todo/internal/pkg"
	"todo/internal/store"
)

type IsUserValid struct {
	Email    string `json:"isEmail"`
	Password string `json:"isPassword"`
}

func (is IsUserValid) ValidCredentials(credentials Credentials) IsUserValid {
	var validateUser IsUserValid
	var resultUser User

	db, err := store.NewConnection()
	if err != nil {
		return validateUser
	}
	defer db.Close()

	err = db.Get(&resultUser, fmt.Sprintf("select * from users where email = '%s'", credentials.Email))

	if resultUser.Email == credentials.Email {
		validateUser.Email = "Пользователь с такой почтой уже сущесвует"
	}
	if !pkg.ValidatePassword(credentials.Password) {
		validateUser.Password = "Пароль не соответсвует требованиям"
	}

	return validateUser
}
