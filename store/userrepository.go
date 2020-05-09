package store

import "github.com/trad3r/rest-api-go/internal/app/model"

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create - добавление нового пользователя в БД
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, hash_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.HashPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

// FindByEmail - поиск пользователя по email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, hash_password FROM users WHERE email = $1",
		email,
	).Scan(
		u.ID,
		u.Email,
		u.HashPassword,
	); err != nil {
		return nil, err
	}

	return u, nil
}
