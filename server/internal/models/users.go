package models

type Credentials struct {
	Email    string
	Password string
}

type User struct {
	ID       int    `json:"userid" db:"userid"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
