package teststore

import (
	"github.com/trad3r/rest-api-go/internal/app/model"
	"github.com/trad3r/rest-api-go/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User - обертка для возможности работать с таблицей users
// только через наше подключение
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
