package user

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	ErrInvalidEmail     = errors.New("invalidate email")
	ErrCreatingPassword = errors.New("error creating password")
)

type Email string

func NewEmail(email string) (Email, error) {
	if email == "" {
		return "", fmt.Errorf("%w: email не должен быть пустым", ErrInvalidEmail)
	}

	if !emailRegex.MatchString(email) {
		return "", fmt.Errorf("%w: невалидный email %s. Email должен быть в формате: mail@mail.com", ErrInvalidEmail, email)
	}

	return Email(email), nil
}

func (e Email) String() string {
	return string(e)
}
