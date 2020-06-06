package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/trad3r/rest-api-go/internal/app/store"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User - обертка для возможности работать с таблицей users
// только через наше подключение
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
