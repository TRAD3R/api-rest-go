package teststore_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/trad3r/rest-api-go/internal/app/model"
	"github.com/trad3r/rest-api-go/internal/app/store"
	"github.com/trad3r/rest-api-go/internal/app/store/teststore"
	"testing"
)

// TestUserRepository_Create ...
func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	assert.NoError(t, s.User().Create(model.TestUser()))
}

// TestUserRepository_FindByEmail ...
func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "test@gmail.com"
	// При поиске нусуществующего email возвращается ошибка
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	// Поиск существующего email
	u := model.TestUser()
	u.Email = email
	err = s.User().Create(u)
	assert.NoError(t, err)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
