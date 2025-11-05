package user

import (
	"context"
	"todo/internal/domain/common"
)

type Writer interface {
	Save(ctx context.Context, usr *User) error
	DeleteByID(ctx context.Context, id common.ID) error
}

type Reader interface {
	GetByID(ctx context.Context, id common.ID) (*User, error)
	GetByEmail(ctx context.Context, email Email) (*User, error)
	GetAll(ctx context.Context, page, limit int) ([]*User, int, error)
}

type Repository interface {
	Writer
	Reader
}
