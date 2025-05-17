package user

import "golang.org/x/crypto/bcrypt"

// Проверка пароля
func checkPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return NewAuthorizationError("Check password!", err)
}
