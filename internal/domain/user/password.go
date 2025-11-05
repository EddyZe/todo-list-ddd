package user

import (
	"todo/internal/pkg/passencode"
)

type Password string

func NewPassword(password string) (Password, error) {
	if password == "" {
		return "", ErrCreatingPassword
	}

	ok, hash := passencode.HashPassword(password)
	if !ok {
		return "", ErrCreatingPassword
	}

	return Password(hash), nil
}

func (p Password) String() string {
	return string(p)
}

func (p Password) Compare(password string) bool {
	return passencode.ComparePasswords(p.String(), password)
}
