package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	// the second parameter is work factor, which determines the complexity of the hashing
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
