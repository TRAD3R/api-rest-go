package store_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/trad3r/rest-api-go/internal/app/model"
	"github.com/trad3r/rest-api-go/store"
	"testing"
)

// TestUserRepository_Create ...
func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser())

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

// TestUserRepository_FindByEmail ...
func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	// При поиске нусуществующего email возвращается ошибка
	_, err := s.User().FindByEmail("test@email.com")

	assert.Error(t, err)

	// Поиск существующего email
	u, err := s.User().Create(model.TestUser())

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
