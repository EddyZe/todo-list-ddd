package task

import (
	"context"
	"todo/internal/domain/common"
)

type Writer interface {
	Save(ctx context.Context, task *Task) error
	DeleteByID(ctx context.Context, id common.ID) error
}

type Reader interface {
	GetByID(ctx context.Context, id common.ID) (*Task, error)
	GetByAuthorID(ctx context.Context, id common.ID) ([]*Task, error)
	GetByAssigneeID(ctx context.Context, id common.ID) ([]*Task, error)
}

type Repository interface {
	Writer
	Reader
}
