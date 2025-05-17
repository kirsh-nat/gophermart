package user

import "golang.org/x/crypto/bcrypt"

// Хеширование пароля
func hashPassword(password string) (string, error) {
	// Генерация соли и хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
