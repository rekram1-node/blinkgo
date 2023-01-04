package blink_test

import (
	"testing"

	"github.com/rekram1-node/blinkgo/blink"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	tests := []struct {
		name            string
		fakeEmail       string
		fakePass        string
		expectedAccount *blink.Account
	}{
		{
			name:      "NewAccount: Good Call",
			fakeEmail: "example@example.com",
			fakePass:  "fakepassword",
			expectedAccount: &blink.Account{
				Email:    "example@example.com",
				Password: "fakepassword",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			account := blink.NewAccount(test.fakeEmail, test.fakePass)
			assert.Equal(t, test.expectedAccount.Email, account.Email)
			assert.Equal(t, test.expectedAccount.Password, account.Password)
			assert.NotEqual(t, "", account.UniqueID)
		})
	}
}
