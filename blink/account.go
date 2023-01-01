package blink

type Account struct {
	Email     string
	Password  string
	AuthToken string
	Tier      string

	ID       int
	ClientID int

	SyncModules *[]SyncModule
	Networks    *[]Network
}

func NewAccount(email, pass string) *Account {
	return &Account{
		Email:    email,
		Password: pass,
	}
}
