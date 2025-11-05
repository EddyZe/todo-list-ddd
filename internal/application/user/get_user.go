package user

import (
	"context"
	"errors"
	"todo/internal/domain/common"
	"todo/internal/domain/user"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidEmail = errors.New("invalid email")
)

func (s *UseCase) GetByID(ctx context.Context, id string) (*OutputUser, error) {
	return s.findUser(ctx, func(ctx context.Context) (*user.User, error) {
		ID, err := common.NewID(id)
		if err != nil {
			return nil, err
		}
		return s.userRepo.GetByID(ctx, ID)
	})
}

func (s *UseCase) GetByEmail(ctx context.Context, email string) (*OutputUser, error) {
	e, err := user.NewEmail(email)
	if err != nil {
		return nil, errors.Join(ErrInvalidEmail, err)
	}

	return s.findUser(ctx, func(ctx context.Context) (*user.User, error) {
		return s.userRepo.GetByEmail(ctx, e)
	})
}

func (s *UseCase) GetAll(ctx context.Context, input InputUserList) *OutputUserList {
	if input.Limit <= 0 {
		input.Limit = 20
	}

	usrs, totalPages, err := s.userRepo.GetAll(ctx, input.Page, input.Limit)
	if err != nil {
		return &OutputUserList{
			Users:       make([]*OutputUser, 0),
			TotalPages:  0,
			CurrentPage: input.Page,
		}
	}

	res := make([]*OutputUser, 0, input.Limit)
	for _, usr := range usrs {
		res = append(res, &OutputUser{
			ID:    usr.ID().String(),
			Email: usr.Email().String(),
		})
	}

	return &OutputUserList{
		Users:       res,
		CurrentPage: input.Page,
		TotalPages:  totalPages,
		Limit:       input.Limit,
	}
}

func (s *UseCase) findUser(ctx context.Context, finder func(ctx context.Context) (*user.User, error)) (*OutputUser, error) {
	u, err := finder(ctx)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) || errors.Is(err, user.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &OutputUser{
		ID:    u.ID().String(),
		Email: u.Email().String(),
	}, nil
}
