package blink

import "github.com/twinj/uuid"

type Account struct {
	Email     string
	Password  string
	AuthToken string
	Tier      string
	UniqueID  string

	ID       int
	ClientID int
}

// Creates an account object that can be used to login
func NewAccount(email, pass string) *Account {
	u := uuid.NewV4() // used for refresh authentication

	return &Account{
		Email:    email,
		Password: pass,
		UniqueID: u.String(),
	}
}
