package user

import (
	"context"
	"errors"
	"todo/internal/domain/user"
)

var (
	ErrCreatingUser = errors.New("error creating user")
	ErrUserIsExists = errors.New("user already is exists")
)

func (s *UseCase) CreateUser(ctx context.Context, u InputUser) (*OutputUser, error) {
	email, err := user.NewEmail(u.Email)
	if err != nil {
		return nil, errors.Join(ErrCreatingUser, err)
	}

	password, err := user.NewPassword(u.Password)
	if err != nil {
		return nil, errors.Join(ErrCreatingUser, err)
	}

	newUser := user.NewUser(email, password)

	if err := s.userRepo.Save(ctx, newUser); err != nil {
		if errors.Is(err, user.ErrUserAlreadyExists) {
			return nil, ErrUserIsExists
		}
		return nil, errors.Join(ErrCreatingUser, err)
	}

	output := OutputUser{
		ID:    newUser.ID().String(),
		Email: email.String(),
	}

	return &output, nil
}
