package model_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/trad3r/rest-api-go/internal/app/model"
	"testing"
)

// TestUser_BeforeCreate ...
func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser()
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.HashPassword)
}
