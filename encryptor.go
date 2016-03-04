package devisecrypto

import (
	"bytes"

	"golang.org/x/crypto/bcrypt"
)

// Digest digests a password and  stretches amount and a pepper which can be blank.
func Digest(password string, stretches int, pepper string) (string, error) {
	var passwordBuffer bytes.Buffer
	passwordBuffer.WriteString(password)

	if pepper != "" {
		passwordBuffer.WriteString(pepper)
	}

	digestedPassword, err := bcrypt.GenerateFromPassword(passwordBuffer.Bytes(), stretches)

	if err != nil {
		return "", err
	}

	return string(digestedPassword), nil
}

// Compare  compares a password with a salt against a hashed password.
func Compare(password string, pepper string, hashedPassword string) bool {
	if hashedPassword == "" {
		return false
	}

	var passwordBuffer bytes.Buffer
	passwordBuffer.WriteString(password)

	if pepper != "" {
		passwordBuffer.WriteString(pepper)
	}

	val := bcrypt.CompareHashAndPassword([]byte(hashedPassword), passwordBuffer.Bytes())

	return (val == nil)
}
