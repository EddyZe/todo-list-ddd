package authmanager

import (
	"context"
	"errors"
	"todo/internal/domain/user"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
)

type UserReader interface {
	GetByEmail(ctx context.Context, email user.Email) (*user.User, error)
}

type AuthManager struct {
	userReader UserReader
}

func New(userReader UserReader) *AuthManager {
	return &AuthManager{userReader: userReader}
}

func (a *AuthManager) Authenticate(ctx context.Context, email, password string) (*UserOutput, error) {
	emailVO, err := user.NewEmail(email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	u, err := a.userReader.GetByEmail(ctx, emailVO)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !u.ComparePassword(password) {
		return nil, ErrInvalidCredentials
	}

	return &UserOutput{
		Email: u.Email().String(),
		ID:    u.ID().String(),
	}, nil
}
