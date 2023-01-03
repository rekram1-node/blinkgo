package client_test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/rekram1-node/blinkgo/internal/client"
	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Default: Good Call",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var c *resty.Client
			defaultClient := client.Default()
			assert.NotNil(t, defaultClient)
			assert.IsType(t, c, defaultClient)
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name          string
		mockAuthToken string
	}{
		{
			name:          "New: Good Call",
			mockAuthToken: "fake-token",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var c *resty.Client
			newClient := client.New(test.mockAuthToken)
			assert.NotNil(t, newClient)
			assert.IsType(t, c, newClient)
		})
	}
}
