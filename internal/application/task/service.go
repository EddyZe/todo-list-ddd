package task

import (
	"context"
	"todo/internal/application/user"
	"todo/internal/domain/task"
)

type UserReader interface {
	GetByID(ctx context.Context, id string) (*user.OutputUser, error)
}

type UseCase struct {
	repo       task.Repository
	userReader UserReader
}

func NewUseCase(repo task.Repository, userReader UserReader) *UseCase {
	return &UseCase{
		repo:       repo,
		userReader: userReader,
	}
}
