package utils

import "golang.org/x/crypto/bcrypt"

type password struct {
}

func (p *password) Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (p *password) Validate(password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(password),
		[]byte(password),
	)
}

var Password password = password{}
