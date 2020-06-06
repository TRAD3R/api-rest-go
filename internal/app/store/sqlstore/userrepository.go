package sqlstore

import (
	"database/sql"
	"github.com/trad3r/rest-api-go/internal/app/model"
	"github.com/trad3r/rest-api-go/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create - добавление нового пользователя в БД
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, hash_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.HashPassword,
	).Scan(&u.ID)
}

// FindByEmail - поиск пользователя по email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, hash_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.HashPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
