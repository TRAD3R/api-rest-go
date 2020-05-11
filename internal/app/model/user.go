package model

import "golang.org/x/crypto/bcrypt"

// User - сущность пользователя
type User struct {
	ID           int
	Email        string
	Password     string
	HashPassword string
}

// BeforeCreate - выполнение подготовительных работ перед
// сохранением пользователя в базу
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		hash, err := encryptString(u.Password)

		if err != nil {
			return err
		}

		u.HashPassword = hash
	}
	return nil
}

// encryptStirng - создание хэш строки
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
