package user

import "time"

type User struct {
	ID         string
	Email      string
	Password   string
	SignedUpAt time.Time
	Role       string
}

func New(email, password string) User {
	return User{
		ID:         time.Now().String(),
		Role:       "buyer",
		Email:      email,
		Password:   password,
		SignedUpAt: time.Now(),
	}
}
