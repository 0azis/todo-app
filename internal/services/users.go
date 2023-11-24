package services

import (
	"fmt"
	"todo/internal/models"
	"todo/internal/store"
)

type UsersServices struct{}

func (us UsersServices) InsertOne(user models.User) (int, error) {
	var userID int
	db, err := store.NewConnection()
	if err != nil {
		return userID, err
	}
	defer db.Close()

	err = db.Get(&userID, fmt.Sprintf("insert into public.users (email, password) values ('%s', '%s') RETURNING userid", user.Email, user.Password))
	return userID, err
}

func (us UsersServices) GetByEmail(email string) (models.User, error) {
	var resultUser models.User
	db, err := store.NewConnection()
	if err != nil {
		return resultUser, err
	}
	defer db.Close()

	err = db.Get(&resultUser, fmt.Sprintf("select * from users where email = '%s'", email))
	return resultUser, err
}
