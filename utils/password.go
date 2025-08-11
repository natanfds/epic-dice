package utils

import "golang.org/x/crypto/bcrypt"

type Password struct {
}

func (p *Password) Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (p *Password) Validate(password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(password),
		[]byte(password),
	)
}
