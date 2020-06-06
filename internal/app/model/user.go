package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

// User - сущность пользователя
type User struct {
	ID           int
	Email        string
	HashPassword string
	Password     string
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

// Validate - валидация данных пользователя
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.HashPassword == "")), validation.Length(6, 100)),
	)
}

// encryptStirng - создание хэш строки
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
