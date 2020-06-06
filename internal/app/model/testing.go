package model

// TestUser - возвращает валидного юзера для тестов
func TestUser() *User {
	return &User{
		Email:    "test@example.com",
		Password: "password",
	}
}
