package user

import (
	"todo/internal/domain/user"
)

type UseCase struct {
	userRepo user.Repository
}

func NewUseCase(repo user.Repository) *UseCase {
	return &UseCase{
		userRepo: repo,
	}
}
