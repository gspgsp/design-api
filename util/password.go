package util

import "golang.org/x/crypto/bcrypt"

/**
加密字符串
 */
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

/**
验证字符串
 */
func PasswordCheck(encodePassword, rawPassword []byte) error {
	err := bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(rawPassword))
	return err
}
