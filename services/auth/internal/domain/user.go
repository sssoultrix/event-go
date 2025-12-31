package domain

import "time"

type User struct {
	ID        string
	Email     string
	CreatedAt time.Time
}

type CreateUserParams struct {
	Email           string
	Password        string
	PasswordConfirm string
}

type LoginParams struct {
	Email    string
	Password string
}
