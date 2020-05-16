package model_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/trad3r/rest-api-go/internal/app/model"
	"testing"
)

// TestUser_Validate ...
func TestUser_Validate(t *testing.T) {
	// Табличные тесты
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User { //nolint:gocritic
				return model.TestUser()
			},
			isValid: true,
		},
		{
			name: "with hashpassword",
			u: func() *model.User {
				u := model.TestUser()
				u.HashPassword = "hashpassword"
				u.Password = ""

				return u
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User { //nolint:gocritic
				u := model.TestUser()
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User { //nolint:gocritic
				u := model.TestUser()
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User { //nolint:gocritic
				u := model.TestUser()
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User { //nolint:gocritic
				u := model.TestUser()
				u.Password = "short"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.isValid {
				assert.Error(t, tc.u().Validate())
			} else {
				assert.NoError(t, tc.u().Validate())
			}
		})
	}
}

// TestUser_BeforeCreate ...
func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser()
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.HashPassword)
}
